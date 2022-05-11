package sql

const Users = `CREATE TABLE IF NOT EXISTS users(
	id BIGSERIAL PRIMARY KEY,
	username VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	password VARCHAR(100) NOT NULL,
	image VARCHAR(255) NOT NULL,
	create_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL
);`

const R_users = `DROP TABLE IF EXISTS users;`

const Func_create_validate = `CREATE OR REPLACE FUNCTION createValidate(n VARCHAR, e VARCHAR, out nameC int, out emailC int)
language plpgsql
as 
$$
BEGIN
	SELECT count(*) INTO nameC FROM users WHERE username = $1;
	SELECT count(*) INTO emailC FROM users WHERE email = $2;
END;
$$;`