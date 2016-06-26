package query

const (
  CreateUser = 
    "INSERT INTO users (email, password, name, age, auth_token, type, beacon_identifier) VALUES ($1, $2, $3, $4, $5, $6, &7) RETURNING id"
  UpdateAuthToken = "UPDATE users SET auth_token=$1"
  CheckLoginDetails = "SELECT * FROM users WHERE email=$1 AND password=$2"
  CheckToken = "SELECT * FROM users WHERE auth_token=$1"
  AddPaymentMethod = "INSERT INTO payment_methods (user_id, card_number, card_holder, expiration, cvv) VALUES($1, $2, $3, $4, $5)"
)
