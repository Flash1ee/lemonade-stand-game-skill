package tree

import (
	"context"

	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
)

type TreeGame struct {
	client proto.TreeGameClient
}

func (l *TreeGame) Create(ctx context.Context, username string) (string, error) {
	res, err := l.client.Create(ctx, &proto.User{
		Username: username,
	})
	if err != nil {
		return "", err
	}
	return res.Id, err
}

func (l *TreeGame) GetBalance(ctx context.Context, userID string) (int64, error) {
	protoGameID := &proto.GameID{Id: userID}
	res, err := l.client.GetBalance(ctx, protoGameID)
	if err != nil {
		return 0, err
	}
	return res.Balance, nil
}

func (l *TreeGame) SaveResult(ctx context.Context, userID string, result int64) error {
	protoResultData := &proto.SaveResultMessage{
		ID: &proto.GameID{
			Id: userID,
		},
		Result: result,
	}
	_, err := l.client.SaveResult(ctx, protoResultData)
	return err
}

func (l *TreeGame) GetResult(ctx context.Context, userID string) ([]StatResult, error) {
	protoGameID := &proto.GameID{Id: userID}

	res, err := l.client.GetResult(ctx, protoGameID)
	if err != nil {
		return []StatResult{}, err
	}
	return exported(res.Results), nil
}

func exported(data []*proto.Result) []StatResult {
	res := make([]StatResult, 0, len(data))
	for _, val := range data {
		if val != nil {
			res = append(res, StatResult{
				Result: val.Result,
			})
		}
	}
	return res
}
