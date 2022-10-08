package repo

import (
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/google/uuid"
)

const initSize = 1000

type GameRepository struct {
	data map[string]entity.Session
}

func NewGameRepository() *GameRepository {
	return &GameRepository{
		data: make(map[string]entity.Session, initSize),
	}
}

func (repo *GameRepository) CreateUser() string {
	userID := uuid.New().String()

	repo.data[userID] = *entity.NewSession()

	return userID
}
