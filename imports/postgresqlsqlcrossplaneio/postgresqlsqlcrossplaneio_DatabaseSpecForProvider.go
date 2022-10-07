// postgresqlsqlcrossplaneio
package postgresqlsqlcrossplaneio


// DatabaseParameters are the configurable fields of a Database.
type DatabaseSpecForProvider struct {
	// If false then no one can connect to this database.
	//
	// The default is true, allowing connections (except as restricted by other mechanisms, such as GRANT/REVOKE CONNECT).
	AllowConnections *bool `field:"optional" json:"allowConnections" yaml:"allowConnections"`
	// How many concurrent connections can be made to this database.
	//
	// -1 (the default) means no limit.
	ConnectionLimit *float64 `field:"optional" json:"connectionLimit" yaml:"connectionLimit"`
	// Character set encoding to use in the new database.
	//
	// Specify a string constant (e.g., 'SQL_ASCII'), or an integer encoding number, or DEFAULT to use the default encoding (namely, the encoding of the template database). The character sets supported by the PostgreSQL server are described in Section 23.3.1. See below for additional restrictions.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// If true, then this database can be cloned by any user with CREATEDB privileges;
	//
	// if false (the default), then only superusers or the owner of the database can clone it.
	IsTemplate *bool `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// Collation order (LC_COLLATE) to use in the new database.
	//
	// This affects the sort order applied to strings, e.g. in queries with ORDER BY, as well as the order used in indexes on text columns. The default is to use the collation order of the template database. See below for additional restrictions.
	LcCollate *string `field:"optional" json:"lcCollate" yaml:"lcCollate"`
	// Character classification (LC_CTYPE) to use in the new database.
	//
	// This affects the categorization of characters, e.g. lower, upper and digit. The default is to use the character classification of the template database. See below for additional restrictions.
	LcCType *string `field:"optional" json:"lcCType" yaml:"lcCType"`
	// The role name of the user who will own the new database, or DEFAULT to use the default (namely, the user executing the command).
	//
	// To create a database owned by another role, you must be a direct or indirect member of that role, or be a superuser.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The name of the tablespace that will be associated with the new database, or DEFAULT to use the template database's tablespace.
	//
	// This tablespace will be the default tablespace used for objects created in this database. See CREATE TABLESPACE for more information.
	Tablespace *string `field:"optional" json:"tablespace" yaml:"tablespace"`
	// The name of the template from which to create the new database, or DEFAULT to use the default template (template1).
	Template *string `field:"optional" json:"template" yaml:"template"`
}

