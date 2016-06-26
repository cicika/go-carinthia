package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/business"
  "github.com/cicika/go-carinthia/model"
  "encoding/json"
  "net/http"
)

func PassengerCheck(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
   paramMap := map[string]string {
    "UserId": params.ByName("user_id"),
    "BeaconIdentifier": params.ByName("beacon_identifier"),
  }

  paxCheck := business.PassengerCheck(paramMap)
  json, _ := json.Marshal(paxCheck)
  response := model.HttpResponse{200, string(json[:])}  
  
  RespondWith(w, response)
}