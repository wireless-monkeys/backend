package store

import "time"

type CameraStore struct {
	Timestamp          time.Time
	CurrentCameraImage []byte
	NumberOfPeople     int64
}

var CameraStoreInstance = CameraStore{}
