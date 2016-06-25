package model

type Trip struct {
  Id int64
  UserId int64
  StartedAt int64 /* Unix timestamp */
  EndedAt int64 /* Unix timestamp */
  TotalPrice float32
}

type TripSegment struct {
  Id int64
  UserId int64
  BeaconId string
  TripId int64
  CheckInType int
  TransportStationId int64
  Latitude float32
  Longitude float32
  PassedAt int64 /* Unix timestamp */
}

const (
  segmentStart = 0
  segmentCheckIn = 1
  segmentEnd = 2
)
