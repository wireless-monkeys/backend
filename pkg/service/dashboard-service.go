package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/asaskevich/EventBus"
	_ "github.com/lib/pq"
	api "github.com/wireless-monkeys/backend/pkg/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type dashboardServiceServer struct {
	api.UnimplementedDashboardServiceServer
	config *QdbConfig
	bus    EventBus.Bus
}

type QdbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	SslMode  string
}

func NewDashboardServiceServer(qdbConfig *QdbConfig, bus EventBus.Bus) api.DashboardServiceServer {
	return &dashboardServiceServer{
		config: qdbConfig,
		bus:    bus,
	}
}

func (s *dashboardServiceServer) GetNumberOfPeople(ctx context.Context, in *api.GetNumberOfPeopleRequest) (*api.GetNumberOfPeopleResponse, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		s.config.Host, s.config.Port, s.config.User, s.config.Password, s.config.Dbname, s.config.SslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlTimeFormat := "2006-01-02 15:04:05"
	query := fmt.Sprintf("SELECT timestamp, number_of_people FROM people WHERE timestamp >= '%s' AND timestamp <= '%s';",
		in.StartTime.AsTime().Format(sqlTimeFormat), in.EndTime.AsTime().Format(sqlTimeFormat))
	fmt.Print(query)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]*api.NumberOfPeopleRow, 0)

	for rows.Next() {
		var timestamp time.Time
		var number_of_people int64
		err = rows.Scan(&timestamp, &number_of_people)
		if err != nil {
			return nil, err
		}
		row := &api.NumberOfPeopleRow{
			Timestamp:      timestamppb.New(timestamp),
			NumberOfPeople: number_of_people,
		}
		people = append(people, row)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &api.GetNumberOfPeopleResponse{
		Rows: people,
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
