package v1

type createUserResponse struct {
	UserID string `json:"userID"`
}

type randomWeatherResponse struct {
	WeatherName string `json:"weather_name"`
	RainChance  int64  `json:"rain_chance"`
}

type balanceResponse struct {
	Balance int64 `json:"balance"`
}

type calculateRequest struct {
	CupsAmount  int64 `json:"cups_amount"      binding:"required"`
	IceAmount   int64 `json:"ice_amount"       binding:"required"`
	StandAmount int64 `json:"stand_amount"     binding:"required"`
	Price       int64 `json:"price"            binding:"required"`
}

type calculateResponse struct {
	Balance int64 `json:"balance"`
	Profit  int64 `json:"profit"`
	Day     int64 `json:"day"`
}
