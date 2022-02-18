package handler

import (
	"learn/todoapi/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) newList(c *gin.Context) {
	var input models.TodoList
	if err := c.ShouldBindJSON(&input); err != nil {
		raiseResposeError(c, http.StatusBadRequest, "invalid list input" + err.Error())
	}

	id, err := h.service.TodoList.CreateList(input)
	if err != nil {
		raiseResposeError(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"list_id": id})
}

func (h *Handler) deleteList(c *gin.Context) {
var input models.TodoList
	if err := c.ShouldBindJSON(&input); err != nil {
		raiseResposeError(c, http.StatusBadRequest, "invalid list input" + err.Error())
	}

	id, err := h.service.TodoList.CreateList(input)
	if err != nil {
		raiseResposeError(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"list_id": id})
}

type listUpdateInput struct {
	Id          int    `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) updateList(c *gin.Context) {
	var input listUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		raiseResposeError(c, http.StatusBadRequest, "invalid list input" + err.Error())
		return
	}

	id, err := h.service.TodoList.UpdateList(models.TodoList{
		Id: input.Id, 
		Title: input.Title, 
		Description: input.Description,
	})
	if err != nil {
		raiseResposeError(c, http.StatusBadRequest, err.Error())
		return	
	}

	c.JSON(http.StatusOK, gin.H{"list_id": id})
}

func (h *Handler) getLists(c *gin.Context) {
	lists, err := h.service.TodoList.GetAll()
	if err != nil {
		raiseResposeError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"lists": lists})
}

func (h *Handler) getListById(c *gin.Context) {
	id_str, _ := c.Params.Get("id")
	id, err := strconv.ParseUint(id_str, 10, 64)
	if err != nil {
		raiseResposeError(c, http.StatusBadRequest, "invalid id")
		return
	}

	list, err := h.service.TodoList.GetById(int(id))
	if err != nil {
		raiseResposeError(c, http.StatusBadRequest, "invalid id" + err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": list})
}
