package query

const (
  InsertNewTrip = "INSERT INTO trips (user_id, started_at) VALUES ($1, $2);"
  SetTripEndTime = "UPDATE trips SET ended_at=$1 WHERE user_id=$2;"
  SetTripTotalPrice = "UPDATE trips set total_price=$1;"
  StartedTripByUserId = "SELECT * FROM trips WHERE user_id=$1 and ended_at is null;"
  InsertSegment = "INSERT INTO trip_segments (user_id, beacon_identifier, trip_id, check_in_type, transport_station_id, latitude, longitude, passed_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8);"
)
