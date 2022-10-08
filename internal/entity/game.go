package entity

const (
	glassPrice = 300
	icePrice   = 50
	StandPrice = 800
	days       = 7
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
	Days       int64
	Balance    int64
}

func NewSession() *Session {
	return &Session{
		GameParams: *NewGameParams(),
		Days:       days,
		Balance:    balance,
	}
}

type Weather struct {
	Wtype      string
	RainChance int64
}
