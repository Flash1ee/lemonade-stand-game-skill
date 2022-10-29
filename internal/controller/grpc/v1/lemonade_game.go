package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
)

type lemonadeRoutes struct {
	t usecase.LemonadeGameUsecase
	l logger.Interface
}

func NewLemonadeRoutes(t usecase.LemonadeGameUsecase, l logger.Interface) *lemonadeRoutes {
	return &lemonadeRoutes{
		t: t,
		l: l,
	}
}

//func newLemonadeRoutes(baseServer grpc.ServiceRegistrar, gameServer LemonadeGamerServer) {
//	proto.RegisterLemonadeGameServer(baseServer, gameServer)
//}

func (r lemonadeRoutes) Create(_ context.Context, _ *proto.Nothing) (*proto.CreateResult, error) {
	user, err := r.t.CreateUser("", "")
	if err != nil {
		r.l.Error(fmt.Errorf("grpc - v1 - Create, err: %w", err))
		return nil, err
	}
	return &proto.CreateResult{
		Id: user,
	}, nil
}

func (r lemonadeRoutes) RandomWeather(_ context.Context, id *proto.GameID) (*proto.Weather, error) {
	if id == nil {
		r.l.Error("grpc - v1 - RandomWeather invalid request")
		return nil, errors.New("invalid game id")
	}
	userID := id.Id
	weather, err := r.t.GetRandomWeather(userID)
	if err != nil {
		r.l.Error(fmt.Errorf("grpc - v1 - RandomWeather, err: %w", err))
		return nil, errors.New("invalid get weather")
	}
	return &proto.Weather{
		WeatherName: weather.Wtype,
		RainChance:  weather.RainChance,
	}, nil
}

func (r lemonadeRoutes) GetBalance(_ context.Context, id *proto.GameID) (*proto.Balance, error) {
	if id == nil {
		r.l.Error("grpc - v1 - GetBalance invalid request")
		return nil, errors.New("invalid game id")
	}
	userID := id.Id
	res, err := r.t.GetBalance(userID)
	if err != nil {
		r.l.Error(fmt.Errorf("grpc - v1 - GetBalance, err: %w", err))
		return nil, errors.New("invalid get balance")
	}
	return &proto.Balance{
		Balance: res,
	}, nil
}

func (r lemonadeRoutes) Calculate(_ context.Context, data *proto.CalculateRequest) (*proto.CalculateResponse, error) {
	if data == nil {
		r.l.Error("grpc - v1 - Calculate invalid request")
		return nil, errors.New("invalid game data")
	}
	if data.Game == nil {
		r.l.Error("grpc - v1 - Calculate invalid request")
		return nil, errors.New("invalid game id")
	}
	userID := data.Game.Id

	res, err := r.t.Calculate(entity.DayParams{
		CupsAmount:  data.CupsAmount,
		IceAmount:   data.IceAmount,
		StandAmount: data.StandAmount,
		Price:       data.Price,
	}, userID)
	if err != nil {
		r.l.Error(err, "grpc - v1 - Calculate")
		return nil, errors.New("error calculate")
	}
	return &proto.CalculateResponse{
		Balance: res.Balance,
		Day:     res.Day,
		Profit:  res.Profit,
	}, nil
}
