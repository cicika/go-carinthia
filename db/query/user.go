package query

const (
  CreateUser = 
    "INSERT INTO users SET (email, password, name, auth_token, type, beacon_identifier) VALUES ($1, $2, $3, $4, $5, $6)"
  UpdateAuthToken = "UPDATE users SET auth_token=$1"
  CheckLoginDetails = "SELECT * FROM users WHERE email=$1 AND password=$2"
  CheckToken = "SELECT * FROM users WHERE auth_token=$1"
)
