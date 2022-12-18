package server

import (
	"github.com/gin-gonic/gin"
	"github.com/takumi2786/takumiweb/controllers"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	ctrl := controllers.HealthzController{}
	r.GET("/healthz", ctrl.Healthz)

	u := r.Group("/users")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.GET("/:id", ctrl.Show)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}
	return r
}
