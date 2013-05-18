package models

import play.api.Play.current
import play.api.db.slick
import slick.DB
import slick.Config.driver.simple._
import collection.mutable.HashMap

final case class Entity(id : Int, var name : String) extends TableRow {
  override def hashCode = id
  override def equals(e : Any) = e match {
    case Entity(i, _) => i == id
    case _ => false
  }

  def commit = DB.withSession { implicit session =>
    Entity.byId(id).map(_.update_*) update (name)
  }

  def account = Account.getId(id)
  private val _access = CachedVal[Permission.Value](Authorize.access_check(id))
  def access : Permission.Value = _access
  def authorizeParents(all : Boolean = false) = Authorize.getParents(id, all)
  def authorizeChildren(all : Boolean = false) = Authorize.getChildren(id, all)
}

private object EntityCache extends HashMap[Int, Entity]

object Entity extends Table[Entity]("entity") {
  def id = column[Int]("id", O.PrimaryKey, O.AutoInc)
  def name = column[String]("name", O.DBType("text"))

  def * = id ~ name <> (Entity.apply _, Entity.unapply _)
  private def update_* = name
  private[this] def insert_* = update_*

  private def byId(i : Int) = Query(this).where(_.id === i)

  def cache(e : Entity, a : Permission.Value = null) : Entity = {
    e._access() = a
    EntityCache.put(e.id, e)
    e
  }
  def get(i : Int) : Entity =
    EntityCache.getOrElseUpdate(i,
      DB.withSession { implicit session =>
        byId(i).firstOption.orNull
      })
  def create(n : String) : Entity = {
    val i = DB.withSession { implicit session =>
      insert_* returning id insert n
    }
    val e = Entity(i, n)
    EntityCache.update(i, e)
    e
  }

  def byName(n : String) = {
    // should clearly be improved and/or indexed
    val w = "%" + n.split("\\s+").filter(!_.isEmpty).mkString("%") + "%"
    for {
      (e, a) <- Entity leftJoin Account on (_.id === _.id)
      if a.username === n || DBFunctions.ilike(e.name, w)
    } yield e
  }

  final val NOBODY : Int = -1
  final val ROOT : Int = 0
}
