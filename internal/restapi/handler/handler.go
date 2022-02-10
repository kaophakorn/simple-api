package handler

import (
	"net/http"
	"simple-go-api/internal/repo/league"
	"simple-go-api/internal/repo/player"
	"simple-go-api/internal/repo/team"
	"simple-go-api/internal/restapi/response"

	"github.com/gin-gonic/gin"
)

func GetPlayers(ginContext *gin.Context) {
	d, err := player.GetAll()
	if err != nil {
		response.JSON(ginContext, http.StatusInternalServerError, 0, "", nil)
	} else {
		response.JSON(ginContext, http.StatusOK, 1, "ok", d)
	}
}

func GetLeagues(ginContext *gin.Context) {
	d, err := league.GetAll()
	if err != nil {
		response.JSON(ginContext, http.StatusInternalServerError, 0, "", nil)
	} else {
		response.JSON(ginContext, http.StatusOK, 1, "ok", d)
	}
}

func GetTeam(ginContext *gin.Context) {
	d, err := team.GetAll()
	if err != nil {
		response.JSON(ginContext, http.StatusInternalServerError, 0, "", nil)
	} else {
		response.JSON(ginContext, http.StatusOK, 1, "ok", d)
	}
}
