package model

type Beacon struct {
  Id int64
  Identifier string
  UserId int64
  TransportType string
  ServiceProviderId int64
  IsStationary bool
}

const (
  Bus = "bus"
  Train = "train"
  Bicycle = "bicycle"
  Unknown = "unknown"
)
