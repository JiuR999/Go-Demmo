package middleware

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !strings.Contains(context.Request.URL.Path, "login") {
			context.Abort()
			context.JSON(http.StatusUnauthorized, models.Result{
				Code: 0,
				Data: http.StatusText(http.StatusUnauthorized),
			})
			return
		}

	}
}
