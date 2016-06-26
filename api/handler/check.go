package handler

import (
  _ "fmt"
  "github.com/julienschmidt/httprouter"
  _ "github.com/cicika/go-carinthia/business"
  "github.com/cicika/go-carinthia/model"
  "encoding/json"
  "net/http"
  "time"
)

func PassengerCheck(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  lastCheckInAt := time.Now().Unix() - 300
  beacon := "EC:F2:FE:70:29:41"
  paxCheck := model.PassengerCheck{lastCheckInAt, beacon}
  json, _ := json.Marshal(paxCheck)
  response := model.HttpResponse{200, string(json[:])}  
  
  RespondWith(w, response)
}
