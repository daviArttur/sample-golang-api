package query

var FindUserByEmail = `SELECT id, email, password FROM users WHERE email = $1;`
var FindUserById = `SELECT id, email, password FROM users WHERE id = $1;`
var CreateUser = `INSERT INTO users (email, password) VALUES  ($1, $2);`
