package repo

import (
	"errors"

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

func (repo *GameRepository) GetBalance(userID string) (int64, error) {
	session, ok := repo.data[userID]
	if !ok {
		return 0, errors.New("not found user")
	}

	return session.Balance, nil
}
