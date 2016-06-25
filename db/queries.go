package db

import "github.com/cicika/go-carinthia/db/query"

func UserQ(q string)string {
  mapping := map[string]string {
    "CreateUser": query.CreateUser,
    "UpdateAuthToken": query.UpdateAuthToken,
    "CheckLoginDetails" : query.CheckLoginDetails,
    "CheckToken" : query.CheckToken,
  }
  return mapping[q]
}

func TripQ(q string)string {
  mapping := map[string]string {
    "InsertNewTrip": query.InsertNewTrip,
    "SetTripEndTime": query.SetTripEndTime,
    "SetTripTotalPrice": query.SetTripTotalPrice,
    "StartedTripByUserId": query.StartedTripByUserId,
    "InsertSegment": query.InsertSegment,
  }
  return mapping[q]
}
