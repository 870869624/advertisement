package mediaorganization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/models/mediaorganization"
)

func Create(context *gin.Context) {
	var media mediaorganization.Mediaorganization
	if err := context.ShouldBindJSON(&media); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数有误",
		})
		fmt.Println(media)
		return
	}
	e := media.Create()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "创建失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func List(context *gin.Context) {

}
func DELETE(context *gin.Context) {

}

func Change(context *gin.Context) {

}
