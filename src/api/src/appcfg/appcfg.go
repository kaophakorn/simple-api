package appcfg

import (
	"myapp/src/utils"
)

var ServicePort string = "5000" // go port

func Setup() {
	ServicePort = utils.GetEnv("SERVICE_PORT", ServicePort)
}
