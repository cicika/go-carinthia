package query

const (
  InsertNewTrip = "INSERT INTO trips SET (user_id, started_at) VALUES ($1, $2);"
  SetTripEndTime = "UPDATE trips SET ended_at=$1"
  SetTripTotalPrice = "UPDATE trips set total_price=$1"
)