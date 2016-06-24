package query

const (

  CreateUserTable = "CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, email text UNIQUE, name text, identifier text);"
)