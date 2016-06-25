package model

type User struct {
  Id int64
  Email string
  Name string
  BirthYear int32
  AuthToken string
  Type string
  BeaconIdentifier string
}

const (
  Passenger = "passenger"
  ServiceProvider = "service_provider"
  PassiveBeacon = "passive_beacon"
)

/* add some methods here */