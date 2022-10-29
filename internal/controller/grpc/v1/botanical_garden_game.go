package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
	usecase "github.com/evrone/go-clean-template/internal/usecase/botanical_garden"
	"github.com/evrone/go-clean-template/pkg/logger"
)

type botanicalGardenRoutes struct {
	t usecase.BotanicalGardenGameUsecase
	l logger.Interface
}

func NewBotanicalGardenRoutes(t usecase.BotanicalGardenGameUsecase, l logger.Interface) *botanicalGardenRoutes {
	return &botanicalGardenRoutes{
		t: t,
		l: l,
	}
}

func (r botanicalGardenRoutes) Create(_ context.Context, _ *proto.Nothing) (*proto.CreateResult, error) {
	user := r.t.CreateUser()
	return &proto.CreateResult{
		Id: user,
	}, nil
}

func (r botanicalGardenRoutes) RandomWeather(_ context.Context, id *proto.GameID) (*proto.Weather, error) {
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

func (r botanicalGardenRoutes) GetBalance(_ context.Context, id *proto.GameID) (*proto.Balance, error) {
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

func (r botanicalGardenRoutes) Calculate(_ context.Context, data *proto.CalculateRequest) (*proto.CalculateResponse, error) {
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
