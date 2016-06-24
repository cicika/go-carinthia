package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  "net/http"
)

func TripStart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func TripSegment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func TripEnd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func AddPassenger(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}

func RemovePassenger(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  response := model.HttpResponse{501, "NotImplemented"}
  RespondWith(w, response)
}
