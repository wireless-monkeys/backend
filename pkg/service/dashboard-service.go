package service

import (
	"context"
	"fmt"
	"github.com/asaskevich/EventBus"
	api2 "github.com/influxdata/influxdb-client-go/v2/api"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	api "github.com/wireless-monkeys/backend/pkg/api"
	"github.com/wireless-monkeys/backend/pkg/config"
	"github.com/wireless-monkeys/backend/pkg/store"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type dashboardServiceServer struct {
	api.UnimplementedDashboardServiceServer
	config   *config.QdbConfig
	bus      EventBus.Bus
	queryAPI api2.QueryAPI
}

func NewDashboardServiceServer(config *config.Config, bus EventBus.Bus, queryAPI api2.QueryAPI) api.DashboardServiceServer {
	return &dashboardServiceServer{
		config:   config.QdbConfig,
		bus:      bus,
		queryAPI: queryAPI,
	}
}

var queryTemplate = `
from(bucket: "people_count")
 |> range(start: %v, stop: %v)
 |> filter(fn: (r) => r["_measurement"] == "people_count")
 |> aggregateWindow(every: %vm, fn: max, createEmpty: false)
 |> yield(name: "max")
`

func (s *dashboardServiceServer) GetNumberOfPeople(ctx context.Context, in *api.GetNumberOfPeopleRequest) (*api.GetNumberOfPeopleResponse, error) {
	intervalMinutes := in.GetIntervalMinutes()
	if intervalMinutes <= 0 {
		intervalMinutes = 1
	}
	query := fmt.Sprintf(
		queryTemplate,
		in.GetStartTime().GetSeconds(),
		in.GetEndTime().GetSeconds(),
		intervalMinutes,
	)
	log.Println(query)
	result, err := s.queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var rows []*api.NumberOfPeopleRow
	for result.Next() {
		for result.Next() {
			rows = append(rows, &api.NumberOfPeopleRow{
				NumberOfPeople: result.Record().Value().(int64),
				Timestamp:      timestamppb.New(result.Record().Time()),
			})
		}
		if result.Err() != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &api.GetNumberOfPeopleResponse{
		Rows: rows,
	}, nil
}

func (s *dashboardServiceServer) SubscribeCamera(empty *api.Empty, stream api.DashboardService_SubscribeCameraServer) error {
	handler := func(event interface{}) {
		cameraStore := store.CameraStoreInstance
		stream.Send(&api.CameraResponse{
			Timestamp: timestamppb.New(cameraStore.Timestamp),
			Image:     cameraStore.CurrentCameraImage,
		})
	}
	s.bus.Subscribe("edge:setdata", handler)
	defer s.bus.Unsubscribe("edge:setdata", handler)
	<-stream.Context().Done()
	return stream.Context().Err()
}
