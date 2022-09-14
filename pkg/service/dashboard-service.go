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
	"google.golang.org/protobuf/types/known/timestamppb"
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
 |> aggregateWindow(every: 1m, fn: last, createEmpty: false)
 |> yield(name: "last")
`

func (s *dashboardServiceServer) GetNumberOfPeople(ctx context.Context, in *api.GetNumberOfPeopleRequest) (*api.GetNumberOfPeopleResponse, error) {
	query := fmt.Sprintf(queryTemplate, in.GetStartTime().AsTime().Unix(), in.GetEndTime().AsTime().Unix())
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
		in := event.(*api.SetDataRequest)
		stream.Send(&api.CameraResponse{
			Timestamp: in.Timestamp,
			Image:     in.CameraImage,
		})
	}
	s.bus.Subscribe("edge:setdata", handler)
	defer s.bus.Unsubscribe("edge:setdata", handler)
	<-stream.Context().Done()
	return stream.Context().Err()
}
