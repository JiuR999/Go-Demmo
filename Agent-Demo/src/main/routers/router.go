package routers

import (
	"github.com/gin-gonic/gin"
)

var Router *router

type router struct{}

func InitRouter() {
	Router = &router{}
}

func (router *router) Init() *gin.Engine {
	r := gin.Default()
	/*gin.SetMode("debug")
	if gin.IsDebugging() {
		pprof.RouteRegister(r, "/debug/pprof")
	}*/

	networkApi(r.Group("/network"))
	cpuApi(r.Group("/cpu"))
	diskApi(r.Group("/disk"))
	return r
}
