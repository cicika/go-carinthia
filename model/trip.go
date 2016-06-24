package model

type Trip {
  Id int64
  UserId int64
  StartedAt int64 /* Unix timestamp */
  EndedAt int64 /* Unix timestamp */
  TotalPrice float
}

type TripSegment {
  Id int64
  UserId int64
  BeaconId string
  TripId int64
  CheckInType int
  TransportStationId int64
  Latitude float
  Longitude float
  PassedAt int64 /* Unix timestamp */
}

const (
  segmentStart = 0
  segmentCheckIn = 1
  segmentEnd = 2
)
