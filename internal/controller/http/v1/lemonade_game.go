package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/gin-gonic/gin"

	"github.com/evrone/go-clean-template/internal/usecase/lemonade"
	"github.com/evrone/go-clean-template/pkg/logger"
)

type lemonadeRoutes struct {
	t usecase.LemonadeGameUsecase
	l logger.Interface
}

func newLemonadeRoutes(handler *gin.RouterGroup, t usecase.LemonadeGameUsecase, l logger.Interface) {
	r := &lemonadeRoutes{t, l}

	h := handler.Group("lemonade")
	{
		h.POST("/id", r.createUser)
		h.GET("/weather", r.randomWeather)
		h.GET("/balance", r.balance)
		h.POST("/calculate", r.calculate)

	}
}

func (r *lemonadeRoutes) createUser(c *gin.Context) {
	user, err := r.t.CreateUser("")
	if err != nil {
		r.l.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "invalid get weather")
		return
	}
	c.JSON(http.StatusOK, createUserResponse{user})
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
