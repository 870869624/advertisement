package media

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/models/media"
)

func Add(context *gin.Context) {
	var ad media.Media
	if err := context.ShouldBindJSON(&ad); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数有误",
		})
		fmt.Println(ad)
		return
	}
	e := ad.Create()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}
func Change(context *gin.Context) {
	var new media.Media
	if err := context.ShouldBindJSON(&new); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数有误",
		})
		return
	}
	e := new.Update()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}
func List(context *gin.Context) {
	title := context.Query("title")
	size := context.DefaultQuery("size", "10")
	current := context.DefaultQuery("current", "1")
	order_by := context.QueryArray("order")
	mediatype := context.DefaultQuery("mediatype", "0")
	imediatype, e := strconv.Atoi(mediatype)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "媒体类型为数字类型",
		})
		return
	}
	isize, e := strconv.Atoi(size)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "size为数字类型",
		})
		return
	}
	icurrent, e := strconv.Atoi(current)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "current为数字类型",
		})
		return
	}
	pagenition := media.Pagination{
		Size:    isize,
		Current: icurrent,
	}
	order := media.Order{
		Order_by: order_by,
	}
	request := media.Query{
		Pagination: pagenition,
		Order:      order,
		Mediatype:  imediatype,
	}
	if title != "" {
		request.Title = title
	}
	result, e := media.List(&request)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
