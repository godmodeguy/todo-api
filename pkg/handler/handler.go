package handler

import (
	"learn/todoapi/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) Handler {
	return Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", h.singin)
		auth.POST("/sign-up", h.singup)
	}

	api := r.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.newList)
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			tasks := lists.Group("/:id/item")
			{
				tasks.POST("/", h.newTask)
				tasks.GET("/", h.getTasks)
				tasks.GET("/:task_id", h.getTaskById)
				tasks.PUT("/:task_id", h.updateTask)
				tasks.DELETE("/:task_id", h.deleteTask)
			}
		}
	}

	return r
}
