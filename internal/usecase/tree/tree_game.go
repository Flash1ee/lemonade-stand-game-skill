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

type TreeGameUsecase struct {
	repo GameRepo
}

// New -.
func New(r GameRepo) *TreeGameUsecase {
	return &TreeGameUsecase{
		repo: r,
	}
}

func (u *TreeGameUsecase) CreateUser(userName string) (string, error) {
	return u.repo.CreateUser(userName)
}

func (u *TreeGameUsecase) SaveStatistics(userID string, result int64) error {
	return u.repo.SaveResult(userID, result)
}

func (u *TreeGameUsecase) GetStatistics(userID string) ([]entity.Statistics, error) {
	return u.repo.GetResult(userID)
}

func (u *TreeGameUsecase) GetBalance(userID string) (int64, error) {
	res, err := u.repo.GetBalance(userID)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (u *TreeGameUsecase) GetRandomWeather(userID string) (entity.Weather, error) {
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

	return response, u.repo.SetDayWeather(userID, response)
}
