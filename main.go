package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinghaijun.com/advertisement-management/api/area"
	"github.com/jinghaijun.com/advertisement-management/api/authentication"
	"github.com/jinghaijun.com/advertisement-management/api/media"
	"github.com/jinghaijun.com/advertisement-management/api/mediaorganization"
	"github.com/jinghaijun.com/advertisement-management/api/user"
	"github.com/jinghaijun.com/advertisement-management/middleware/authorzition"
)

func main() {

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user", user.Signup)
	r.POST("/authentication", authentication.Signin)

	user_group := r.Group("/user")
	user_group.Use(authorzition.Auth)
	{
		user_group.PATCH("/update", user.Update)
	}
	media_group := r.Group("/media")
	media_group.Use(authorzition.Auth)
	{
		media_group.POST("/add", media.Add)
		media_group.PATCH("/change", media.Change)
		media_group.GET("/list", media.List)
	}
	// area_group := r.Group("/area")
	// area_group.Use(authorzition.Auth)
	// {

	// }

	r.POST("/area", area.Create)
	r.GET("/area/list", area.List)
	r.DELETE("/area/delete", area.Delete)
	r.PATCH("/area/change", area.Change)

	r.POST("/mediaor", mediaorganization.Create)
	r.GET("/mediaor/list", mediaorganization.List)
	r.DELETE("/mediaor/delete", mediaorganization.DELETE)
	r.PATCH("/mediaor/change", mediaorganization.Change)

	r.Run(":8000")

}
