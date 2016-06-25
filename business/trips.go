package business

import (
  "log"
  "github.com/cicika/go-carinthia/db"
  _ "github.com/cicika/go-carinthia/model"
)

func StartTrip(paramMap map[string]string){
  params := []string{paramMap["StartedAt"]}

  if !alreadyStarted(paramMap["UserId"]) {
    db.Execute(1, params, db.TripQ("StartTrip"))
  }  
}

func EndTrip(paramMap map[string]string){
  params := []string{paramMap["EndedAt"], paramMap["UserId"]}

  db.Execute(1, params, db.TripQ("EndTrip"))
}

func AddTripSegment(paramMap map[string]string){
  params := []string{paramMap["UserId"], paramMap["BeaconIdentifier"], 
                     string(tripByUserId(paramMap["UserId"])), paramMap["CheckInType"], 
                     "", paramMap["Latitude"], paramMap["Longitude"], paramMap["CreatedAt"]}

  db.Execute(len(params), params, db.TripQ("InsertSegment"))
}

func AddPassenger(paramMap map[string]string){}

func RemovePassenger(paramMap map[string]string){}

func alreadyStarted(userId string)(bool) {
  params := []string{userId}
  var id, user int64
  var res map[int64]int64
  rows := db.Execute(1, params, db.TripQ("StartedTripByUserId"))
  
  for rows.Next() {
    err := rows.Scan(&id, &userId)
      if err != nil {
        log.Fatal(err)
      }
    res[id] = user
  }
  return len(res) != 0
}

func tripByUserId(userId string)(int64) {
  params := []string{userId}
  var id int64
  var res []int64
  rows := db.Execute(1, params, db.TripQ("StartedTripByUserId"))
  
  for rows.Next() {
    err := rows.Scan(&id)
      if err != nil {
        log.Fatal(err)
      }
    res = append(res, id)
  }
  return res[0]
}
