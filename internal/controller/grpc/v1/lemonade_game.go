package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
	usecase "github.com/evrone/go-clean-template/internal/usecase/lemonade"
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

func (r lemonadeRoutes) Create(_ context.Context, user *proto.User) (*proto.CreateResult, error) {
	if user == nil {
		r.l.Error("grpc - v1 - Create invalid request")
		return nil, errors.New("invalid user")
	}
	u, err := r.t.CreateUser(user.Username)
	if err != nil {
		r.l.Error(fmt.Errorf("grpc - v1 - Create, err: %w", err))
		return nil, err
	}
	return &proto.CreateResult{
		Id: u,
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
		r.l.Error(fmt.Errorf("grpc - v1 - Calculate, err: %w", err))
		return nil, errors.New("error calculate")
	}
	return &proto.CalculateResponse{
		Balance: res.Balance,
		Day:     res.Day,
		Profit:  res.Profit,
	}, nil
}

func (r lemonadeRoutes) SaveResult(_ context.Context, result *proto.SaveResultMessage) (*proto.Nothing, error) {
	if result == nil {
		r.l.Error("grpc - v1 - SaveResult invalid request")
		return &proto.Nothing{}, errors.New("invalid result data")
	}
	if result.ID == nil {
		r.l.Error("grpc - v1 - SaveResult invalid request")
		return &proto.Nothing{}, errors.New("invalid game id")
	}
	userID := result.ID.Id
	resultGame := result.Result
	if err := r.t.SaveStatistics(userID, resultGame); err != nil {
		r.l.Error(fmt.Errorf("grpc - v1 - SaveResult, err: %w", err))
		return &proto.Nothing{}, errors.New("error SaveResult")
	}
	return &proto.Nothing{}, nil
}

func (r lemonadeRoutes) GetResult(_ context.Context, userID *proto.GameID) (*proto.ResultResponses, error) {
	if userID == nil {
		r.l.Error("grpc - v1 - GetResult invalid request")
		return &proto.ResultResponses{}, errors.New("invalid game id")
	}
	data, err := r.t.GetStatistics(userID.Id)
	if err != nil {
		r.l.Error(fmt.Errorf("grpc - v1 - GetResult, err: %w", err))
		return &proto.ResultResponses{}, errors.New("internal error")
	}
	res := &proto.ResultResponses{
		Results: convertResults(data),
	}
	return res, nil
}

func convertResults(stats []entity.Statistics) []*proto.Result {
	res := make([]*proto.Result, 0, len(stats))
	for _, val := range stats {
		res = append(res, &proto.Result{
			Username: val.VKUserId,
			Result:   val.Result,
		})
	}
	return res
}
