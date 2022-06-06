package area

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/models/area"
)

func Add(context *gin.Context) {
	var area area.Area
	if err := context.ShouldBindJSON(&area); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数有误",
		})
		return
	}
	e := area.Add()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}
func List(context *gin.Context) {
	Level := context.DefaultQuery("level", "0")
	Left := context.Query("left")
	Right := context.Query("right")
	Order := context.QueryArray("order")
	ilevel, e := strconv.Atoi(Level)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "地域等级为数字类型",
		})
		return
	}
	ileft, e := strconv.Atoi(Left)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "地域范围为数字类型",
		})
		return
	}
	iright, e := strconv.Atoi(Right)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "地域范围为数字类型",
		})
		return
	}
	Request := area.Request{
		Left:  ileft,
		Right: iright,
		Level: ilevel,
	}
	iOrder := area.Order{
		By: Order,
	}
	Query := area.AreaQuery{
		Request: Request,
		Order:   iOrder,
	}
	response, e := area.List(&Query)
	context.JSON(http.StatusOK, gin.H{
		"message": response.Result,
		"err":     e,
	})
}
func Delete(context *gin.Context) {
	var delete area.Area
	if err := context.ShouldBindJSON(&delete); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	e := delete.Delete()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "删除失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
