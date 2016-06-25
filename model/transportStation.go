package model

type TransportStation struct {
  Id int64
  Name string
  ProvidersIdentifier int
  Latitude float32
  Longitude float32
  Type string
  ZoneId int64
}

const (
  busStation = "bus_station"
  trainStation = "train_station"
  bicycleStation = "bicycle_station"
  other = "other"
)
