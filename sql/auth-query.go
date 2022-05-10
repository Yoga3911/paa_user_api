package sql

const CreateUser = `INSERT INTO users (username, email, password, create_at, update_at)
					VALUES ($1, $2, $3, NOW(), NOW())`

const VerifyCredential = `SELECT id, username, email, password, create_at, update_at FROM users WHERE email = $1`

const GetLastId = `SELECT COUNT(*) FROM users`

const GetByEmail = `SELECT COUNT(*) FROM users WHERE email = $1`

const GetByName = `SELECT COUNT(*) FROM users WHERE username = $1`

const RegisterVal = `SELECT * FROM createValidate($1, $2)`