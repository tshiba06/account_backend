package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s Server) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, &gin.H{"msg": "Hello, world"})
}

func (s Server) PostUsers(c *gin.Context) {

}
