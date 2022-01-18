package middleware

import (
	"myapp/src/httpresp"
	"myapp/src/provider"

	"github.com/gin-gonic/gin"
)

func DIMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var err error
		di, err := (provider.DI{}).New()
		if err != nil {
			ginContext.PureJSON(httpresp.Fail.ToHttpResp())
			ginContext.Abort()
			return
		}
		ginContext.Set("DI", &di)
		ginContext.Next()
	}
}
