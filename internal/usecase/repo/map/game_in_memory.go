package _map

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

func (repo *GameRepository) GetSession(userID string) (entity.Session, error) {
	session, ok := repo.data[userID]
	if !ok {
		return entity.Session{}, errors.New("not found user")
	}

	return session, nil
}

func (repo *GameRepository) SetDayWeather(userID string, weather entity.Weather) error {
	session, ok := repo.data[userID]
	if !ok {
		return errors.New("not found user")
	}
	curDay := session.CurDay
	session.Weather[curDay] = weather
	repo.data[userID] = session

	return nil
}

func (repo *GameRepository) NextDay(userID string, newBalance int64) error {
	session, ok := repo.data[userID]
	if !ok {
		return errors.New("not found user")
	}
	session.CurDay += 1
	session.Balance = newBalance
	repo.data[userID] = session
	return nil
}
