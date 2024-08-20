package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tshiba06/account_backend/usecase/user"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	Tracer      trace.Tracer
	UserUseCase user.UseCase
}

func (h *Handler) GetUsers(c *gin.Context) {
	_, span := h.Tracer.Start(c, "GetUsers")
	defer span.End()

	h.UserUseCase.Get(c)

	c.JSON(http.StatusOK, &gin.H{"msg": "Hello, world"})
}

func (h *Handler) PostUsers(c *gin.Context) {
	_, span := h.Tracer.Start(c, "PostUsers")
	defer span.End()

	c.JSON(http.StatusOK, &gin.H{"msg": "Test"})
}
