package sql

const GetOne = `SELECT id, username, email, password, image, create_at, update_at FROM users WHERE id = $1`

const UpdateUser = `UPDATE users SET username = $2, email = $3, image = $4, update_at = NOW() WHERE id = $1`
