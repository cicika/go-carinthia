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