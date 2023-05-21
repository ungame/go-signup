package repository

const (
	createAuthenticationUsersQuery = `
INSERT INTO authentication_users
	(id, email, username, phone, password, created_at, updated_at)
VALUES
    (?,?,?,?,?,?,?)
`

	updateAuthenticationUsersQuery = `
UPDATE authentication_users 
SET 
    email = ?,
    username = ?,
    phone = ?,
    password = ?,
    updated_at = ?
WHERE
    id = ?                      
`
	getAuthenticationUserQuery = `SELECT * FROM authentication_users WHERE ? = ?`

	getAllAuthenticationUsersQuery = `SELECT * FROM authentication_users LIMIT ? OFFSET ?`

	deleteAuthenticationUsersQuery = `DELETE FROM authentication_users WHERE id = ?`
)
