package botanical_garden

import (
	"crypto/rand"
	"math/big"

	"github.com/evrone/go-clean-template/internal/entity"
	utils "github.com/evrone/go-clean-template/pkg"
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

type BotanicalGardenGameUsecase struct {
	repo GameRepo
}

// New -.
func New(r GameRepo) *BotanicalGardenGameUsecase {
	return &BotanicalGardenGameUsecase{
		repo: r,
	}
}

func (u *BotanicalGardenGameUsecase) CreateUser() string {
	return u.repo.CreateUser()
}

func (u *BotanicalGardenGameUsecase) GetRandomWeather(userID string) (entity.Weather, error) {
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

func (u *BotanicalGardenGameUsecase) GetBalance(userID string) (int64, error) {
	res, err := u.repo.GetBalance(userID)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (u *BotanicalGardenGameUsecase) Calculate(params entity.DayParams, userID string) (entity.DayResult, error) {
	response := entity.DayResult{}
	session, err := u.repo.GetSession(userID)
	if err != nil {
		return response, err
	}
	dayWeather := session.Weather[session.CurDay]
	coef := 0.0
	switch dayWeather.Wtype {
	case hotWeather:
		if params.IceAmount > params.CupsAmount {
			coef += 0.3
		}
	case sunnyWeather:
		if params.IceAmount > params.CupsAmount {
			coef += 0.3
		}
	case cloudyWeather:
		if params.IceAmount > params.CupsAmount {
			coef -= 0.3
		}

	}

	if res, _ := rand.Int(rand.Reader, big.NewInt(100)); res.Int64() < dayWeather.RainChance {
		coef /= 2
	}

	res, _ := rand.Int(rand.Reader, big.NewInt(params.StandAmount*3))
	if params.StandAmount > res.Int64()/2 {
		coef *= 2
	} else {
		coef += 0.2
	}
	res, _ = rand.Int(rand.Reader, big.NewInt(1000))

	coef += float64(res.Int64())
	coef = utils.Min(coef, 1.0)

	profit := params.CupsAmount*params.Price - session.GameParams.Glass*params.CupsAmount - session.GameParams.Ice*params.IceAmount - session.GameParams.Stand*params.StandAmount
	profit = int64(float64(profit) * coef)

	response.Day = session.CurDay
	response.Profit = profit
	response.Balance = session.Balance - profit
	return response, u.repo.NextDay(userID, response.Balance)
}
