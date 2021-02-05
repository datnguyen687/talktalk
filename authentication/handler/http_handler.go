package handler

import (
	"net/http"
	"talktalk/authentication/handler/response"
	"talktalk/authentication/usecase"

	"github.com/gin-gonic/gin"
)

// NewHTTPAuthentication ...
func NewHTTPAuthentication(cfg Config, uc usecase.Interface) Interface {
	h := &handler{
		usecase: uc,
	}

	h.router = gin.Default()

	h.router.GET("/ping", h.handlePing)

	return h
}

type handler struct {
	router  *gin.Engine
	usecase usecase.Interface
}

func (h *handler) Run() error {
	return nil
}

func (h *handler) handlePing(c *gin.Context) {
	resp := response.BasicResponse{
		Code:  http.StatusOK,
		Error: "",
		Data:  "pong",
	}

	c.JSON(http.StatusOK, &resp)
}
