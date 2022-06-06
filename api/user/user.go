package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/models/user"
)

func Signup(context *gin.Context) {
	fmt.Println("xxxx")
	var user user.Users
	if err := context.ShouldBindJSON(&user); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("sssss", user)
	result := user.Create()
	if result != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, result.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

func Update(context *gin.Context) {
	var newInfomation user.Users
	if err := context.ShouldBindJSON(&newInfomation); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	e := newInfomation.Update()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}
func Cancelletion(context *gin.Context) {

}
