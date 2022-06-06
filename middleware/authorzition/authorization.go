package authorzition

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/models/token"
)

func Auth(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization")
	if tokenString == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请登录",
		})
		return
	}
	//解析令牌
	claim, e := token.Parse(tokenString)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请登录",
		})
		return
	}
	context.Request.Header.Set("x-consumer-kind", fmt.Sprintf("%d", claim.Kind))
}
