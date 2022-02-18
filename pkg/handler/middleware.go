package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	userIdCtx = "userId"
)

func (h *Handler) userIdentify(c *gin.Context) {
	token := c.GetHeader(authCokieName)
	logrus.Debug("token in header: ", token)

	if token == "" {
		raiseResposeError(c, http.StatusUnauthorized, "provide access token in header: " + authCokieName)
		return
	}

	id, err := h.service.Authorization.ParseToken(token)
	if err != nil {
		raiseResposeError(c, http.StatusUnauthorized, "invalid token")
		return
	}

	c.Set(userIdCtx, id)
}