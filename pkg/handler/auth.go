package handler

import (
	"learn/todoapi/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authCokieName = "X-Auth"
	authCokieAge = 12 * 60 * 60 * time.Hour
)


func (h *Handler) singup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		raiseResposeError(c, http.StatusBadRequest, "sing-un invalid input")
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		raiseResposeError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": id})
}

type signInInput struct {
	Password	string	`json:"password" binding:"required"`
	Username	string	`json:"username" binding:"required"`
}

func (h *Handler) singin(c *gin.Context) {
	var input signInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		raiseResposeError(c, http.StatusBadRequest, "sing-in invalid input")
		return
	}

	logrus.Debugln("trying sign-in as ", input.Password, input.Username)

	jwt, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		raiseResposeError(c, http.StatusUnauthorized, "sing-in error " + err.Error())
		return
	}

	c.SetCookie(authCokieName, jwt, int(authCokieAge.Seconds()), "/", ".", true, true)
	c.JSON(http.StatusOK, gin.H{"token": jwt})
}
