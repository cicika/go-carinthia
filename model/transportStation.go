package model

type TransportStation {
  Id int64
  Name string
  ProvidersIdentifier int
  Latitude float
  Longitude float
  Type string
  ZoneId int64
}

const (
  busStation = "bus_station"
  trainStation = "train_station"
  bicycleStation = "bicycle_station"
  other = "other"
)
