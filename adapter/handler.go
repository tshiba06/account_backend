package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	Tracer trace.Tracer
}

func (h *Handler) GetUsers(c *gin.Context) {
	_, span := h.Tracer.Start(c.Request.Context(), "GetUsers")
	defer span.End()

	c.JSON(http.StatusOK, &gin.H{"msg": "Hello, world"})
}

func (h *Handler) PostUsers(c *gin.Context) {
	_, span := h.Tracer.Start(c.Request.Context(), "PostUsers")
	defer span.End()

	c.JSON(http.StatusOK, &gin.H{"msg": "Test"})
}
