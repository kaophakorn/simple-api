package middleware

import (
	"simple-go-api/internal/provider"

	"github.com/gin-gonic/gin"
)

const ResponseKey string = "responseBody"

// DefaultMiddleWare : active every incoming request
func DependencyInjection() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		di := provider.GetWebserviceDI()
		ginContext.Set("WebserviceDI", &di)
		ginContext.Next()
	}
}
