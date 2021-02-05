package handler

import (
	"fmt"
	"net/http"
	"talktalk/authentication/handler/dto"
	"talktalk/authentication/handler/response"
	"talktalk/authentication/usecase"

	"github.com/gin-gonic/gin"
)

// NewHTTPAuthentication ...
func NewHTTPAuthentication(cfg Config, uc usecase.Interface) Interface {
	h := &handler{
		usecase: uc,
		port:    cfg.Port,
	}

	h.router = gin.Default()

	h.router.GET("/ping", h.handlePing)

	h.router.POST("/register", h.handleRegister)

	return h
}

type handler struct {
	router  *gin.Engine
	usecase usecase.Interface
	port    int
}

func (h *handler) Run() error {
	return h.router.Run(fmt.Sprintf(`:%d`, h.port))
}

func (h *handler) handlePing(c *gin.Context) {
	resp := response.BasicResponse{
		Code:  http.StatusOK,
		Error: "",
		Data:  "pong",
	}

	c.JSON(http.StatusOK, &resp)
}

func (h *handler) handleRegister(c *gin.Context) {
	resp := response.BasicResponse{}

	var d dto.UserDTO
	err := c.BindJSON(&d)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	err = d.SelfValidate()
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, &resp)
		return
	}

	data, err := h.usecase.RegisterNewUser(&d)

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, &resp)
		return
	}

	resp.Data = data
	c.JSON(http.StatusOK, &resp)
}
