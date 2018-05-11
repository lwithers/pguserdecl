/*
Package userdecl can apply a declaratively-defined set of users and permissions
to a PostgresQL database.
*/
package userdecl

// Decl is a declaration of users and permissions.
type Decl struct {
	// Databases holds details of which databases within a PostgresQL
	// cluster will have their permissions managed.
	Databases map[string]DatabaseDecl

	// Roles defines the set of PostgresQL roles (users) to manage.
	Roles struct {
		// Unmanaged is a blacklist of users that will never be altered
		// by this tool. It has the highest precedence within this
		// section.
		Unmanaged []string

		// ManagedPatterns tells the tool that it has ownership over a
		// set of names. The patterns themselves are globs. For example,
		// "staff_*" would tell the tool that it should manage any user
		// whose username is prefixed with "staff_" — and thus, if a
		// user "staff_jsmith" is removed from the Managed list below,
		// the next invocation of the tool knows it is supposed to now
		// remove "staff_jsmith" as a user.
		ManagedPatterns []string `yaml:"managed_patterns"`

		// Managed is a list of users to create/manage.
		Managed []RoleDecl
	}
}

// DatabaseDecl is the set of data stored for each named database.
type DatabaseDecl struct {
	// Owner is the name of the user which owns the database. This user will
	// have full read/write access to the schema and the data. The user must
	// be declared as a managed user.
	Owner string

	// Schemas is a list of zero or more schemas to be created. There is
	// always an implicit “public” schema. Ownership for schemas follows
	// ownership of the database itself. Schemas which are not mentioned
	// will not be removed; this field exists only for automating
	// creation and ownership.
	Schemas []string
}

// RoleDecl is the set of data stored for each named user.
type RoleDecl struct {
	// Password for the user. If absent, the role will not be granted login.
	// This is typically used for "group" roles, which hold permissions that
	// may be inherited by other roles.
	Password string

	// MemberOf is the list of role groups this user is a member of. This
	// mechanism allows permissions to be inherited from other roles.
	MemberOf []string `yaml:"member_of"`

	// Select is a list of names of each table a user who is a member of
	// this role group should be able to SELECT data from.
	Select []string `yaml:"SELECT"`

	// Insert names tables into which data may be INSERTed.
	Insert []string `yaml:"INSERT"`

	// TODO: sequences too
}
