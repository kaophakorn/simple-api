package handler

import (
	"simple-go-api/internal/repo/player"

	"github.com/gin-gonic/gin"
)

func GetPlayers(ginContext *gin.Context) {
	p, err := player.GetAll()
	if err != nil {
		ginContext.PureJSON(500, nil)
	} else {
		ginContext.PureJSON(200, p)
	}

}
