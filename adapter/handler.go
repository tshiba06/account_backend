package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, &gin.H{"msg": "Hello, world"})
}

func (h *Handler) PostUsers(c *gin.Context) {
	panic("implement me")
}
