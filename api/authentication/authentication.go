package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/models/authentication"
)

func Signin(context *gin.Context) {
	var auth authentication.Authentication
	if err := context.ShouldBindJSON(&auth); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	token, e := auth.Signin()
	if token == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
