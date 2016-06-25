package handler

import (
  "github.com/julienschmidt/httprouter"
  "github.com/cicika/go-carinthia/model"
  "github.com/cicika/go-carinthia/business"
  "net/http"
)

func TripStart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  var response model.HttpResponse
  var extracted map[string]string

  required := []string{"StartedAt"}

  if ValidateRequired(required, r) {
    forExtraction := []string{"StartedAt"}
    extracted = ExtractParams(forExtraction, r)
    extracted["UserId"] = p.ByName("uid")

    business.StartTrip(extracted)
    response = model.HttpResponse{200, ""}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
  RespondWith(w, response)
}

func TripSegment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  var response model.HttpResponse
  var extracted map[string]string

  required := []string{"TransportStation", "BeaconIdentifier",
                       "Latitude", "Longitude", "CreatedAt", "CheckInType"}

  if ValidateRequired(required, r) {
    forExtraction := []string{"TransportStation", "TransportType", "BeaconIdentifier",
                       "Latitude", "Longitude", "CreatedAt", "CheckInType"}
    extracted = ExtractParams(forExtraction, r)
    extracted["UserId"] = p.ByName("uid")

    business.AddTripSegment(extracted)
    response = model.HttpResponse{200, ""}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }

  RespondWith(w, response)
}

func TripEnd(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  required := []string{"EndedAt"}

  var response model.HttpResponse
  var extracted map[string]string

  if ValidateRequired(required, r) {
    forExtraction := []string{"EndedAt"}
    extracted = ExtractParams(forExtraction, r)
    extracted["UserId"] = p.ByName("uid")
    
    business.EndTrip(extracted)
    response = model.HttpResponse{200, ""}  
  } else {
    response = model.HttpResponse{400, "BadRequest"}
  }
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
