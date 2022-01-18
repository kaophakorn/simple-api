package webservice

import (
	"myapp/src/httpresp"

	"github.com/gin-gonic/gin"
)

type APICovid struct{}

func APICovidHandler(ginContext *gin.Context) {
	var respF httpresp.RespF
	respData := struct {
		TotalDeath uint32 `json:"total_death"`
	}{
		0,
	}
	respF = httpresp.OK
	respF.Data = respData
	ginContext.PureJSON(respF.ToHttpResp())
}
