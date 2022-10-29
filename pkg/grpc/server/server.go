package grpcserver

import (
	"net"

	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
	"github.com/evrone/go-clean-template/pkg/logger"
	"google.golang.org/grpc"
)

type GameLemonadeServer struct {
	grpcServer   *grpc.Server
	gameHandlerL proto.LemonadeGameServer
	gameHandlerB proto.LemonadeGameServer

	logger logger.Interface
}

func NewGameLemonadeGRPCServer(logger logger.Interface, grpcServer *grpc.Server,
	gameServerL proto.LemonadeGameServer, gameServerB proto.BotanicalGardenGameServer) *GameLemonadeServer {
	server := &GameLemonadeServer{
		grpcServer:   grpcServer,
		gameHandlerL: gameServerL,
		gameHandlerB: gameServerB,
		logger:       logger,
	}
	return server
}

func (server *GameLemonadeServer) StartGRPCServer(listenUrl string) error {
	lis, err := net.Listen("tcp", listenUrl)
	server.logger.Info("my listen url %s \n", listenUrl)

	if err != nil {
		server.logger.Error("GameLemonadeServer\n")
		server.logger.Error("can not listen url: %s err :%v\n", listenUrl, err)
		return err
	}
	proto.RegisterLemonadeGameServer(server.grpcServer, server.gameHandlerL)
	proto.RegisterBotanicalGardenGameServer(server.grpcServer, server.gameHandlerB)

	server.logger.Info("Start session service\n")
	return server.grpcServer.Serve(lis)
}
