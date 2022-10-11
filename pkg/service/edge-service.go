package service

import (
	"context"
	"github.com/asaskevich/EventBus"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	api2 "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/pkg/errors"
	"github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/store"
)

type edgeServiceServer struct {
	api.UnimplementedEdgeServiceServer

	bus      EventBus.Bus
	writeAPI api2.WriteAPIBlocking
}

func NewEdgeServiceServer(bus EventBus.Bus, writeAPI api2.WriteAPIBlocking) api.EdgeServiceServer {
	return &edgeServiceServer{
		bus:      bus,
		writeAPI: writeAPI,
	}
}

func (s *edgeServiceServer) Heartbeat(ctx context.Context, in *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

func (s *edgeServiceServer) SetData(ctx context.Context, in *api.SetDataRequest) (*api.Empty, error) {
	store.CameraStoreInstance.Timestamp = in.Timestamp.AsTime()
	store.CameraStoreInstance.CurrentCameraImage = in.CameraImage
	store.CameraStoreInstance.NumberOfPeople = in.NumberOfPeople
	s.bus.Publish("edge:setdata", 0)
	p := influxdb2.NewPointWithMeasurement("people_count").
		AddField("count", in.NumberOfPeople).
		SetTime(in.GetTimestamp().AsTime())
	err := s.writeAPI.WritePoint(ctx, p)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &api.Empty{}, nil
}
