package routers

import (
	"github.com/SubhoBasak/brouter"

	"test/controllers"
)

func PostRouter() *brouter.Router {
	router := brouter.Router{}

	router.GET("/", controllers.GetPost)
	router.POST("/", controllers.PostPost)
	router.PUT("/", controllers.EditPost)
	router.DELETE("/", controllers.DelPost)

	router.GET("/all", controllers.AllPost)

	return &router
}
