package controllers

import scala.concurrent.Future
import play.api._
import          Play.current
import          mvc._
import          data._
import          i18n.Messages
import          libs.Files.TemporaryFile
import          libs.iteratee.Enumerator
import          libs.concurrent.Execution.Implicits.defaultContext
import macros._
import macros.async._
import dbrary._
import site._
import models._

private[controllers] sealed class AssetController extends ObjectController[Asset] {
  protected def action(i : models.Asset.Id, p : Permission.Value) =
    RequestObject.check(models.Asset.get(i)(_), p)

  protected def Action(i : models.Asset.Id, p : Permission.Value) =
    SiteAction andThen action(i, p)

  final val defaultThumbSize : Int = 480

  private[controllers] def assetResult(asset : BackedAsset, size : Option[Int] = None, saveAs : Option[String] = None)(implicit request : SiteRequest[_]) : Future[Result] = {
    val tag = asset.etag
    /* Assuming assets are immutable, any if-modified-since header is good enough */
    if (HTTP.notModified(tag, new Timestamp(0)))
      async(NotModified)
    else for {
      data <- store.Asset.read(asset, size)
      date = store.Asset.timestamp(asset)
    } yield {
      val size = data.size
      val range = if (request.headers.get(IF_RANGE).forall(HTTP.unquote(_).equals(tag)))
          request.headers.get(RANGE).flatMap(HTTP.parseRange(_, size))
        else
          None
      val subdata = range.fold(data)((data.range _).tupled)
      val headers = Seq[Option[(String, String)]](
        Some(CONTENT_LENGTH -> subdata.size.toString),
        range.map(r => CONTENT_RANGE -> ("bytes " + (if (r._1 >= size) "*" else r._1.toString + "-" + r._2.toString) + "/" + size.toString)),
        Some(CONTENT_TYPE -> asset.format.mimetype),
        saveAs.map(name => CONTENT_DISPOSITION -> ("attachment; filename=" + HTTP.quote(name + asset.format.extension.fold("")("." + _)))),
        Some(LAST_MODIFIED -> HTTP.date(new Timestamp(date))),
        Some(ETAG -> HTTP.quote(tag)),
        Some(CACHE_CONTROL -> "max-age=31556926, private") /* this needn't be private for public data */
      ).flatten
        Result(
          header = ResponseHeader(range.fold(OK)(r => if (r._1 >= size) REQUESTED_RANGE_NOT_SATISFIABLE else PARTIAL_CONTENT),
            Map(headers : _*)),
          body = subdata)
      }
  }

  private def set(asset : Asset, form : AssetController.AssetForm)(implicit request : SiteRequest[_]) =
    for {
      container <- form.container.get.mapAsync(Container.get(_).map(_ getOrElse
	form.container.withError("object.invalid", "container")._throw))
      _ <- asset.change(name = form.name.get, classification = form.classification.get)
      sa <- container.mapAsync(asset.link(_, form.position.get))
      _ <- form.excerpt.get.foreachAsync(Excerpt.set(asset, Range.full, _))
    } yield (sa.fold(result(asset))(SlotAssetController.result _))

  def update(o : models.Asset.Id) =
    Action(o, Permission.EDIT).async { implicit request =>
      val form = new AssetController.ChangeForm()._bind
      set(request.obj, form)
    }

  private def create(form : AssetController.AssetUploadForm) : Future[Asset] = {
    implicit val request = form.request
    val adm = request.access.isAdmin
    val ifmt = form.format.get.filter(_ => adm).flatMap(AssetFormat.get(_))
    val (file, fmt, fname) =
      form.file.get.fold[(TemporaryFile, AssetFormat, Option[String])] {
	/* local file handling, for admin only: */
	val file = store.Stage.file(form.localfile.get.filter(_ => adm) getOrElse
	  form.file.withError("error.required")._throw)
	val name = file.getName
	if (!file.isFile)
	  form.localfile.withError("File not found")._throw
	val ffmt = ifmt orElse AssetFormat.getFilename(name) getOrElse
	  form.format.withError("file.format.unknown", "unknown")._throw
	(store.TemporaryFileLinkOrCopy(file), ffmt, Some(name))
      } { file =>
	val ffmt = ifmt orElse AssetFormat.getFilePart(file) getOrElse
	  form.file.withError("file.format.unknown", file.contentType.getOrElse("unknown"))._throw
	(file.ref, ffmt, Maybe(file.filename).opt)
      }
    for {
      asset <- fmt match {
	case fmt : TimeseriesFormat if adm && form.timeseries.get =>
	  val probe = media.AV.probe(file.file)
	  models.Asset.create(form.volume, fmt, form.classification.get.getOrElse(Classification(0)), probe.duration, fname, file)
	case _ =>
	  if (fmt.isTranscodable) try {
	    media.AV.probe(file.file)
	  } catch { case e : media.AV.Error =>
	    form.file.withError("file.invalid", e.getMessage)._throw
	  }
	  models.Asset.create(form.volume, fmt, form.classification.get.getOrElse(Classification(0)), fname, file)
      }
      _ = if (fmt.isTranscodable && !form.timeseries.get)
	store.Transcode.start(asset)
    } yield (asset)
  }

  def upload(v : models.Volume.Id) =
    VolumeHtml.Action(v, Permission.CONTRIBUTE).async { implicit request =>
      val form = new AssetController.UploadForm()._bind
      create(form).flatMap(set(_, form))
    }

  def replace(o : models.Asset.Id) =
    Action(o, Permission.CONTRIBUTE).async { implicit request =>
      val form = new AssetController.ReplaceForm()._bind
      for {
	done <- request.obj.isSuperseded
	_ = if (done) form.withGlobalError("file.superseded")._throw
	asset <- create(form)
	_ <- asset.supersede(request.obj)
	r <- set(asset, form)
      } yield (r)
    }

  def remove(a : models.Asset.Id) = Action(a, Permission.EDIT).async { implicit request =>
    for {
      _ <- request.obj.unlink
    } yield (result(request.obj))
  }
}

object AssetController extends AssetController {
  sealed abstract class AssetForm(action : Call)(implicit val request : RequestObject.Site[_])
    extends HtmlForm[AssetForm](action,
      views.html.asset.edit(_)) {
    def actionName : String
    def formName : String = actionName + " Asset"
    def volume : Volume

    val name = Field(OptionMapping(Mappings.maybeText))
    val classification = Field(OptionMapping(Mappings.enum(Classification)))
    val container = Field(Forms.optional(Forms.of[Container.Id]))
    val position = Field(Forms.optional(Forms.of[Offset]))
    val excerpt = Field(OptionMapping(Forms.optional(Mappings.enum(Classification))))
  }

  sealed trait AssetUpdateForm extends AssetForm {
    def asset : Asset
    def volume = asset.volume
    name.fill(Some(asset.name))
    classification.fill(Some(asset.classification))
    container.fill(None)
    position.fill(None)
  }

  final class ChangeForm(implicit request : Request[_])
    extends AssetForm(routes.AssetHtml.update(request.obj.id))
    with AssetUpdateForm {
    def actionName = "Update"
    override def formName = "Edit Asset"
    def asset = request.obj
  }

  sealed trait AssetUploadForm extends AssetForm {
    val format = Field(Forms.optional(Forms.of[AssetFormat.Id]))
    val timeseries = Field(Forms.boolean)
    val localfile = Field(Forms.optional(Forms.nonEmptyText))
    val file = OptionalFile()
  }

  final class UploadForm(implicit request : VolumeController.Request[_])
    extends AssetForm(routes.AssetHtml.upload(request.obj.id))
    with AssetUploadForm {
    def actionName = "Upload"
    def volume = request.obj
    classification.fill(Some(Classification.IDENTIFIED))
  }

  final class ReplaceForm(implicit request : Request[_])
    extends AssetForm(routes.AssetHtml.replace(request.obj.id))
    with AssetUpdateForm with AssetUploadForm {
    def actionName = "Replace"
    def asset = request.obj
  }
}

object AssetHtml extends AssetController with HtmlController {
  import AssetController._

  def view(o : models.Asset.Id) = Action(o, Permission.VIEW).async { implicit request =>
    request.obj.slot.map(_.fold[Result](
      throw NotFoundException /* TODO */)(
      sa => Redirect(sa.inContainer.pageURL)))
  }

  def edit(o : models.Asset.Id) =
    Action(o, Permission.EDIT).async { implicit request =>
      request.obj.slot.flatMap { sa =>
	val form = new ChangeForm()
	form.excerpt.fill(Some(sa.flatMap(_.excerpt.map(_.classification))))
	form.Ok
      }
    }

  def replaceView(o : models.Asset.Id) =
    Action(o, Permission.EDIT).async { implicit request =>
      new ReplaceForm().Ok
    }

  def create(v : models.Volume.Id, c : Option[Container.Id], pos : Option[Offset]) =
    VolumeHtml.Action(v, Permission.CONTRIBUTE).async { implicit request =>
      val form = new UploadForm()
      form.container.fill(c)
      form.position.fill(pos)
      form.Ok
    }

  def transcode(a : models.Asset.Id, stop : Boolean = false) =
    (SiteAction.rootAccess(Permission.ADMIN) andThen action(a, Permission.EDIT)) { implicit request =>
      if (stop) store.Transcode.stop(request.obj.id)
      else      store.Transcode.start(request.obj)
      Ok("transcoding")
    }

  def formats = SiteAction.Unlocked {implicit request =>
		  Ok(views.html.asset.formats(
		  AssetFormat.getAll.toSeq.groupBy(_.mimeSubTypes._1)))
  }
}

object AssetApi extends AssetController with ApiController {
  def get(i : models.Asset.Id) = Action(i, Permission.VIEW).async { implicit request =>
    request.obj.json(request.apiOptions).map(Ok(_))
  }

  class TranscodedForm(aid : Asset.Id) extends {
      val auth = play.api.libs.Crypto.sign(aid.toString)
    } with StructForm(routes.AssetApi.transcoded(aid, auth)) {
    val pid = Field(Forms.number)
    val res = Field(Forms.number)
    val log = Field(Forms.text)
  }

  /** Called from remote transcoding process only. */
  def transcoded(i : models.Asset.Id, auth : String) =
    play.api.mvc.Action { implicit request =>
      val form = new TranscodedForm(i)._bind
      if (!auth.equals(form.auth) || form.hasErrors)
	BadRequest("")
      else {
	store.Transcode.collect(i, form.pid.get, form.res.get, form.log.get)
	Ok("")
      }
    }
}
