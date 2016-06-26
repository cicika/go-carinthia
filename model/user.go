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

type PaymentMethod struct {
  Id int64
  UserId int64
  CardNumber int32
  CardHolder string
  Expiration string
  Cvv int32
}

/* add some methods here */