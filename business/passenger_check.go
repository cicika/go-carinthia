package business

import (
  "log"
  "github.com/cicika/go-carinthia/db"
  "github.com/cicika/go-carinthia/model"
)

func PassengerCheck(paramMap map[string]string) model.PassengerCheck {
  rows := db.Execute(2, []string{paramMap["UserId"], paramMap["BeaconIdentifier"]}, db.ReportQ("PassengerCheck"))
  var userId int64
  var beaconIdentifier string
  var res [](model.PassengerCheck)

  for rows.Next() {
    err := rows.Scan(&userId, &beaconIdentifier)
    if err != nil {
      log.Fatal(err)
    }
    res = append(res, model.PassengerCheck{userId, beaconIdentifier})
  }

  return res[0]
}