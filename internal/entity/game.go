package entity

import "go.mongodb.org/mongo-driver/bson"

const (
	glassPrice = 10
	icePrice   = 50
	StandPrice = 10
	days       = 21
	balance    = 2000
	/////////////////
	lives       = 5
	sheets      = 1
	water       = 0
	treeBalance = 2000
)

type User struct {
	ID int64 `json:"ID"`
}

type GameParamsPrices struct {
	Glass int64 `bson:"glass"`
	Ice   int64 `bson:"ice"`
	Stand int64 `bson:"stand"`
}

type GameParamsTree struct {
	Lives  int64 `bson:"lives"`
	Sheets int64 `bson:"sheets"`
	Water  int64 `bson:"water"`
}

func NewGameParams() *GameParamsPrices {
	return &GameParamsPrices{
		Glass: glassPrice,
		Ice:   icePrice,
		Stand: StandPrice,
	}
}

func NewTreeGameParams() *GameParamsTree {
	return &GameParamsTree{
		Lives:  lives,
		Sheets: sheets,
		Water:  water,
	}
}

type Session struct {
	SessionId  string    `bson:"session_id"`
	VKUserId   string    `bson:"vk_user_id"`
	GameParams []byte    `bson:"game_params"`
	Weather    []Weather `bson:"weather"`
	Days       int64     `bson:"days"`
	Balance    int64     `bson:"balance"`
	CurDay     int64     `bson:"cur_day"`
}

func NewSession(userName string, SessionId string) *Session {
	params, _ := bson.Marshal(NewGameParams())
	return &Session{
		SessionId:  SessionId,
		VKUserId:   userName,
		GameParams: params,
		Weather:    make([]Weather, days+1),
		Days:       days,
		Balance:    balance,
		CurDay:     1,
	}
}

func NewTreeSession(userName string, SessionId string) *Session {
	params, _ := bson.Marshal(NewTreeGameParams())
	return &Session{
		SessionId:  SessionId,
		VKUserId:   userName,
		GameParams: params,
		Weather:    make([]Weather, days+1),
		Days:       days,
		Balance:    treeBalance,
		CurDay:     1,
	}
}

type Weather struct {
	Wtype      string `bson:"wtype"`
	RainChance int64  `bson:"rainChance"`
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

type Statistics struct {
	VKUserId string `bson:"vk_user_id"`
	Result   int64  `bson:"result"`
}
