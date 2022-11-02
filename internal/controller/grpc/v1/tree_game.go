package v1

import (
	"context"
	"errors"
	"fmt"

	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
	usecase "github.com/evrone/go-clean-template/internal/usecase/tree"
	"github.com/evrone/go-clean-template/pkg/logger"
)

type treeRoutes struct {
	t usecase.TreeGameUsecase
	l logger.Interface
}

func NewTreeRoutes(t usecase.TreeGameUsecase, l logger.Interface) *treeRoutes {
	return &treeRoutes{
		t: t,
		l: l,
	}
}

func (r treeRoutes) Create(_ context.Context, user *proto.User) (*proto.CreateResult, error) {
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

func (r treeRoutes) GetBalance(_ context.Context, id *proto.GameID) (*proto.Balance, error) {
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

func (r treeRoutes) SaveResult(_ context.Context, result *proto.SaveResultMessage) (*proto.Nothing, error) {
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

func (r treeRoutes) GetResult(_ context.Context, userID *proto.GameID) (*proto.ResultResponses, error) {
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
