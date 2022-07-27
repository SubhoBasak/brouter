package routers

import (
	"github.com/SubhoBasak/brouter"

	"test/controllers"
)

func BlogRouter() *brouter.Router {
	router := brouter.Router{}

	router.GET("/", controllers.GetBlog)
	router.POST("/", controllers.PostBlog)
	router.PUT("/", controllers.EditBlog)
	router.DELETE("/", controllers.DelBlog)

	router.GET("/all", controllers.AllBlog)

	return &router
}
