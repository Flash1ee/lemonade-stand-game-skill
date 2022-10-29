package usecase

import (
	"crypto/rand"
	"math"
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

func (u *LemonadeGameUsecase) CreateUser(assesToken string, clientId string) (string, error) {
	return u.repo.CreateUser()
}

func (u *LemonadeGameUsecase) SaveStatistics(userID string, result int64) error {
	return u.repo.SaveResult(userID, result)
}

func (u *LemonadeGameUsecase) GetFriendsStatistics(sessionID string) error {
	return nil
}

func (u *LemonadeGameUsecase) GetRandomWeather(userID string) (entity.Weather, error) {
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

func (u *LemonadeGameUsecase) GetBalance(userID string) (int64, error) {
	res, err := u.repo.GetBalance(userID)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (u *LemonadeGameUsecase) Calculate(params entity.DayParams, userID string) (entity.DayResult, error) {
	response := entity.DayResult{}
	session, err := u.repo.GetSession(userID)
	if err != nil {
		return response, err
	}
	dayWeather := session.Weather[session.CurDay]
	users := int64(100)

	switch dayWeather.Wtype {
	case hotWeather:
		if params.IceAmount < 4 {
			users = users / (4 - params.IceAmount)
		} else {
			users = int64(float64(users) * 1.5)
		}
	case sunnyWeather:
		if params.IceAmount < 2 {
			users = users / (2 - params.IceAmount)
		} else if params.IceAmount <= 4 {
			users = int64(float64(users) * 1.5)
		} else {
			users = int64(float64(users) * 0.5)
		}
	case cloudyWeather:
		if res, _ := rand.Int(rand.Reader, big.NewInt(100)); res.Int64() < dayWeather.RainChance {
			if params.IceAmount != 0 {
				users = int64(float64(users) * 0.5)
			}
		} else {
			if params.IceAmount == 0 {
				users = int64(float64(users) / 1.5)
			} else if params.IceAmount <= 2 {
				users = int64(float64(users) * 1.5)
			} else {
				users = int64(float64(users) * 0.5)
			}
		}
	}

	res, _ := rand.Int(rand.Reader, big.NewInt(params.StandAmount*3))
	if params.StandAmount > res.Int64()/2 {
		users *= 2
	} else {
		users = int64(float64(users) * 0.2)
	}

	profit := int64(math.Min(float64(users), float64(params.CupsAmount))*float64(params.Price)) -
		session.GameParams.Glass*params.CupsAmount - session.GameParams.Ice*params.IceAmount - session.GameParams.Stand*params.StandAmount

	response.Day = session.CurDay
	response.Profit = profit
	response.Balance = session.Balance - profit
	return response, u.repo.NextDay(userID, response.Balance)
}
