# pguserdecl

Declarative user management for PostgresQL.

Generally, the schema of a database should be managed using SQL. As changes to
the schema are made, these should be represented as a series of SQL scripts (one
in the forward path, one in the rollback path). And in fact the tool at
https://github.com/mattes/migrate/ provides a nice, simple way to execute such
script series.

However, this sort of scheme does not sit very well with user management, which
is a management function that is not synchronised with schema updates. And the
SQL scripts should be under developer SCM control — which means that any
passwords in scripts will be exposed to the world.

This tool envisages pulling out user management into a separate function, with
users declared in a consistent way and the tool taking care of ensuring this
application on the system.

## Example YAML file

The user declarations are placed into a YAML file as follows:

```yaml
pguserdecl:
  # list of databases to be managed by this tool — other databases will be
  # left untouched
  databases:
    my_stuff:
      owner: "my_stuff_owner"      # user declared below
      schemas:                     # optional; members will be created if needed
        - other_schema

  # list of roles (users) to be managed by this tool
  roles:
    # any users listed here are never altered by this tool; this section has the
    # highest precedence in case of conflict
    unmanaged:
      - admin

    # inform the tool it explicitly manages any users whose name matches one of
    # these glob patterns. Such users will be deleted if they are not listed in
    # the "managed" section, below.
    managed_patterns:
      - staff_*
      - my_stuff*
      - role*

    # list of current users
    managed:
      # this could be handed to an operator with read-only access
      staff_lwithers:
        password: "xyz"   # if present, user can login; if absent, user cannot
        member_of:
          - role1

      my_stuff_owner:
        password: "abc123"

      # this could be handed to an application with read/write capability
      my_stuff_access:
        password: "abc123"
        member_of:
          - role1
          - role2

      # these roles are not used directly for login, but are instead groups
      # which may be inherited by other roles
      role1:
        SELECT:
          - my_stuff.*        # only matches the public schema
          - my_stuff.*.*      # matches all schema including public

      role2:
        INSERT:
          - my_stuff.*
          - my_stuff.other_schema.*
        UPDATE:
          - my_stuff.table1
          - my_stuff.other_schema.table2
```

## Assumptions of the tool

The tool assumes it will be run with the permissions of a database superuser.

Databases will be owned by users that can login. It is suggested that these
credentials are used only for management / migration tools (for example
`mattes/migrate`), and that specific users are created for application-level
access.

Schemas are used to logically partition data that is accessed by a single,
cooperating set of applications. Schemas are not directly used for
access control (though it is simple to arrange such).
