package restapi

import (
	"fmt"
	"simple-go-api/internal/provider"
	"simple-go-api/internal/restapi/handler"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
)

func Serve() {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(nice.Recovery(func(c *gin.Context, err interface{}) {
		c.JSON(500, gin.H{
			"error": "something went wrong...",
		})
	}))
	router := g.Group("/")
	router.GET("players", handler.GetPlayers)
	cfg := provider.GetConfig()

	g.Run(fmt.Sprintf(":%s", cfg.Port))
}
