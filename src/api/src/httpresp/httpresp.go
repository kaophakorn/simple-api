package httpresp

import (
	"net/http"
)

var Messages = map[string]string{
	"ok":   "ok",
	"fail": "fail",
}

var (
	OK   = RespF{Code: "ok", HTTPStatus: http.StatusOK}
	Fail = RespF{Code: "fail", HTTPStatus: http.StatusOK}
)

type RespF struct {
	HTTPStatus int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (f RespF) ToHttpResp() (int, interface{}) {
	if f.Message == "" {
		f.Message = Messages[f.Code]
	}
	return f.HTTPStatus, f
}
