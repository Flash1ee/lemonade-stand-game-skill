package v1

import (
	"net/http"

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
		h.GET("/create", r.createUser)
		h.GET("/weather", r.randomWeather)
	}
}

type createUserResponse struct {
	UserID string `json:"userID"`
}

func (r *lemonadeRoutes) createUser(c *gin.Context) {
	user := r.t.CreateUser()
	c.JSON(http.StatusOK, createUserResponse{user})
}

type randomWeatherResponse struct {
	WeatherName string `json:"weather_name"`
	RainChance  int64  `json:"rain_chance"`
}

func (r *lemonadeRoutes) randomWeather(c *gin.Context) {
	weather, err := r.t.GetRandomWeather()
	if err != nil {
		r.l.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "invalid get weather")
		return
	}
	c.JSON(http.StatusOK, randomWeatherResponse{
		WeatherName: weather.Wtype,
		RainChance:  weather.RainChance,
	})
}

//type doTranslateRequest struct {
//	Source      string `json:"source"       binding:"required"  example:"auto"`
//	Destination string `json:"destination"  binding:"required"  example:"en"`
//	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
//}
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
