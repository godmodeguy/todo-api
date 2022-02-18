package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


type Error struct {
	msg	string
}

func raiseResposeError(c *gin.Context, stautsCode int, msg string) {
	logrus.Debugf("response error: %v", msg)
	c.AbortWithStatus(stautsCode)
}