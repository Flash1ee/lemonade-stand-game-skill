package v1

import (
	"context"

	proto "github.com/evrone/go-clean-template/internal/generated/delivery/protobuf"
)

type LemonadeGamerServer interface {
	Create(context.Context, *proto.Nothing) (*proto.CreateResult, error)
}
