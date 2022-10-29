package botanical_garden

import "github.com/evrone/go-clean-template/internal/entity"

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	GameRepo interface {
		CreateUser() string
		GetBalance(userID string) (int64, error)
		GetSession(userID string) (entity.Session, error)
		SetDayWeather(userID string, weather entity.Weather) error
		NextDay(userID string, newBalance int64) error
	}
)
