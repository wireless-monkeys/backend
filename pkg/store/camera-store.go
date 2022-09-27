package store

import "time"

type CameraStore struct {
	Timestamp          time.Time
	CurrentCameraImage []byte
}

var CameraStoreInstance = CameraStore{}
