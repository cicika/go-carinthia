package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  "net/http"
)

func WeeklyReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}
func MonthlyReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func CurrentSpendings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  responseString := "{\"Amount\": 10.54, \"Currency\": \"EUR\"}"
  response := model.HttpResponse{200, responseString}
  RespondWith(w, response)
}
