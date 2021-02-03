package handler

import (
	"net/http"
	"talktalk/authentication/handler/response"

	"github.com/gin-gonic/gin"
)

// NewHTTPAuthentication ...
func NewHTTPAuthentication(cfg Config) Interface {
	h := &handler{}

	h.router = gin.Default()

	h.router.GET("/ping", h.handlePing)

	return h
}

type handler struct {
	router *gin.Engine
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
