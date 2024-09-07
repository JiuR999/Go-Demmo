package routers

import (
	_ "AndroidToolServer-Go/docs"
	"AndroidToolServer-Go/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"sync"
)

var (
	Router *router
	once   sync.Once
)

func init() {
	Router = &router{}
	once = sync.Once{}
}

type router struct{}

func (router *router) Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.JWTMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Use(myHandler())
	articleApi(r.Group("/article"))
	hardwareApi(r.Group("/hardware"))
	zdDreamApi(r.Group("/dream"))
	generalFaluireApi(r.Group("/failure"))
	downloadApi(r.Group("/down"))
	//r := gin.New()
	//gin.SetMode()
	return r
}

// 拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			log.Fatal("没有授权 进行拦截!")
			context.Abort()
		}
		context.Next()
	}
}
