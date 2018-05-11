package userdecl

// TODO
//  Iterate through the current list of users. For each user:
//   - if unmanaged, skip
//   - if managed, check if user should exist
//     - if not, drop it
//
//  Iterate through the desired list of users. For each user:
//   - if unmanaged, error out
//   - if exists
//     - check login/password
//     - check role membership
//     - check direct perms
//   - if doesn't exist, create it
//     - set role membership
//     - set direct perms

const (
	// NB: pg_* should probably be filtered out
	queryRoles = `SELECT rolname, rolcanlogin FROM pg_roles`
)
