package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/gin-gonic/gin"

	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
)

type lemonadeRoutes struct {
	t usecase.LemonadeGameUsecase
	l logger.Interface
}

func newLemonadeRoutes(handler *gin.RouterGroup, t usecase.LemonadeGameUsecase, l logger.Interface) {
	r := &lemonadeRoutes{t, l}

	h := handler.Group("")
	{
		h.POST("/id", r.createUser)
		h.GET("/weather", r.randomWeather)
		h.GET("/balance", r.balance)
		h.POST("/calculate", r.calculate)

	}
}

type createUserResponse struct {
	UserID string `json:"userID"`
}

func (r *lemonadeRoutes) createUser(c *gin.Context) {
	user, err := r.t.CreateUser("", "")
	if err != nil {
		r.l.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "invalid get weather")
		return
	}
	c.JSON(http.StatusOK, createUserResponse{user})
}

type randomWeatherResponse struct {
	WeatherName string `json:"weather_name"`
	RainChance  int64  `json:"rain_chance"`
}

func (r *lemonadeRoutes) randomWeather(c *gin.Context) {
	userID := c.Query("id")
	if userID == "" {
		r.l.Error("http - v1 - randomWeather")
		errorResponse(c, http.StatusForbidden, "userID not found")
		return
	}
	weather, err := r.t.GetRandomWeather(userID)
	if err != nil {
		r.l.Error(err, "http - v1 - randomWeather")
		errorResponse(c, http.StatusBadRequest, "invalid get weather")
		return
	}
	c.JSON(http.StatusOK, randomWeatherResponse{
		WeatherName: weather.Wtype,
		RainChance:  weather.RainChance,
	})
}

type balanceResponse struct {
	Balance int64 `json:"balance"`
}

func (r *lemonadeRoutes) balance(c *gin.Context) {
	userID := c.Query("id")
	if userID == "" {
		r.l.Error("http - v1 - balance")
		errorResponse(c, http.StatusForbidden, "userID not found")
		return
	}
	res, err := r.t.GetBalance(userID)
	if err != nil {
		r.l.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "error get balance")
		return
	}
	c.JSON(http.StatusOK, balanceResponse{
		Balance: res,
	})
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

func (r *lemonadeRoutes) calculate(c *gin.Context) {
	userID := c.Query("id")
	if userID == "" {
		r.l.Error("http - v1 - balance")
		errorResponse(c, http.StatusForbidden, "userID not found")
		return
	}
	var req calculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - calculate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	res, err := r.t.Calculate(entity.DayParams{
		CupsAmount:  req.CupsAmount,
		IceAmount:   req.IceAmount,
		StandAmount: req.StandAmount,
		Price:       req.Price,
	}, userID)
	if err != nil {
		r.l.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "error get balance")
		return
	}
	c.JSON(http.StatusOK, calculateResponse{
		Balance: res.Balance,
		Day:     res.Day,
		Profit:  res.Profit,
	})
}

//
//// @Summary     Translate
//// @Description Translate a text
//// @ID          do-translate
//// @Tags  	    translation
//// @Accept      json
//// @Produce     json
//// @Param       request body doTranslateRequest true "Set up translation"
//// @Success     200 {object} entity.Translation
//// @Failure     400 {object} response
//// @Failure     500 {object} response
//// @Router      /translation/do-translate [post]
//func (r *translationRoutes) doTranslate(c *gin.Context) {
//	var request doTranslateRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		r.l.Error(err, "http - v1 - doTranslate")
//		errorResponse(c, http.StatusBadRequest, "invalid request body")
//
//		return
//	}
//
//	translation, err := r.t.Translate(
//		c.Request.Context(),
//		entity.Translation{
//			Source:      request.Source,
//			Destination: request.Destination,
//			Original:    request.Original,
//		},
//	)
//	if err != nil {
//		r.l.Error(err, "http - v1 - doTranslate")
//		errorResponse(c, http.StatusInternalServerError, "translation service problems")
//
//		return
//	}
//
//	c.JSON(http.StatusOK, translation)
//}
