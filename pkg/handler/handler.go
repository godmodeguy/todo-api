package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {

}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")
	}


	api := r.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			tasks := lists.Group("/")
			{
				tasks.POST("/")
				tasks.GET("/")
				tasks.GET("/:task_id")
				tasks.PUT("/:task_id")
				tasks.DELETE("/:task_id")
			}
		}
	}

	return r
}
