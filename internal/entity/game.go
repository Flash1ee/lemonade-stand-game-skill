package entity

const (
	glassPrice = 10
	icePrice   = 50
	StandPrice = 10
	days       = 21
	balance    = 2000
)

type User struct {
	ID int64 `json:"ID"`
}

type GameParamsPrices struct {
	Glass int64
	Ice   int64
	Stand int64
}

func NewGameParams() *GameParamsPrices {
	return &GameParamsPrices{
		Glass: glassPrice,
		Ice:   icePrice,
		Stand: StandPrice,
	}
}

type Session struct {
	GameParams GameParamsPrices
	Weather    []Weather
	Days       int64
	Balance    int64
	CurDay     int64
}

func NewSession() *Session {
	return &Session{
		GameParams: *NewGameParams(),
		Weather:    make([]Weather, days+1),
		Days:       days,
		Balance:    balance,
		CurDay:     1,
	}
}

type Weather struct {
	Wtype      string
	RainChance int64
}

type DayParams struct {
	CupsAmount  int64
	IceAmount   int64
	StandAmount int64
	Price       int64
}

type DayResult struct {
	Balance int64
	Profit  int64
	Day     int64
}
