package usecase

import (
	"crypto/rand"
	"math/big"

	"github.com/evrone/go-clean-template/internal/entity"
)

const (
	sunnyWeather  = "sunny"
	hotWeather    = "hot"
	cloudyWeather = "cloudy"
)

var weather = map[int64]string{
	0: sunnyWeather,
	1: hotWeather,
	2: cloudyWeather,
}

type LemonadeGameUsecase struct {
	repo GameRepo
}

// New -.
func New(r GameRepo) *LemonadeGameUsecase {
	return &LemonadeGameUsecase{
		repo: r,
	}
}

func (u *LemonadeGameUsecase) CreateUser() string {
	return u.repo.CreateUser()
}

func (u *LemonadeGameUsecase) GetRandomWeather() (entity.Weather, error) {
	weatherType, err := rand.Int(rand.Reader, big.NewInt(3))
	if err != nil {
		return entity.Weather{}, err
	}
	response := entity.Weather{
		Wtype: weather[weatherType.Int64()],
	}
	if response.Wtype == cloudyWeather {
		chance, err := rand.Int(rand.Reader, big.NewInt(101))
		if err != nil {
			return entity.Weather{}, err
		}
		response.RainChance = chance.Int64()
	}
	return response, nil
}
