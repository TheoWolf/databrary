package models

import dbrary._

/** The possible levels of permission used for site access, user delegation, and study permissions.
  * Must match the corresponding postgres "permission" type. */
object Permission extends PGEnum("permission") {
  val NONE, VIEW, DOWNLOAD, CONTRIBUTE, ADMIN = Value
  // aliases or equivalent permissions (do not use val here)
  def OWN = ADMIN
  def EDIT = CONTRIBUTE
  def DATA = DOWNLOAD
  def COMMENT = VIEW
}

/** The possible levels of participant consent governing [Classification.IDENTIFIED] data.
  * Must match the corresponding postgres "consent" type, except for the NONE value which represents NULL (missing) as this is a common possibility.
  * Should thus often be constructed as `consent.getOrElse(Consent.NONE)` and used as `util.maybe(consent, Consent.NONE)`. */
object Consent extends PGEnum("consent") {
  val NONE, PRIVATE, SHARED, EXCERPTS, PUBLIC = Value
}

/** The possible types of data sensitivity according to the presence of identifying user data.
  * Must match the corresponding postgres "consent" type. */
object Classification extends PGEnum("classification") {
  val IDENTIFIED, EXCERPT, DEIDENTIFIED, ANALYSIS, PRODUCT, MATERIAL = Value
}
