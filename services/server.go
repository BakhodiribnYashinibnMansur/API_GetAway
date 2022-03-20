package services

import (
	"apiGetaway/config"
	"apiGetaway/genproto/position_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	ProfessionService() position_service.ProfessionServiceClient
}

type grpcClients struct {
	positionService position_service.ProfessionServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connPositionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		positionService: position_service.NewProfessionServiceClient(connPositionService),
	}, nil
}

func (g *grpcClients) ProfessionService() position_service.ProfessionServiceClient {
	return g.positionService
}
