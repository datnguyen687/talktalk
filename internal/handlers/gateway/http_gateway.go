package gateway

import (
	"fmt"
	"net/http"
	"talktalk/internal/handlers"

	"github.com/gin-gonic/gin"
)

type httpGateWay struct {
	engine *gin.Engine
	port   int
}

// NewHTTPGateWayHandler ...
func NewHTTPGateWayHandler() handlers.HandlerInterface {
	return &httpGateWay{}
}

func (gw *httpGateWay) Init(cfg interface{}) error {
	config := cfg.(*Config)
	gw.port = config.Port
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	gw.engine = gin.Default()
	gw.setupRoutes()

	return nil
}

func (gw *httpGateWay) Run() error {
	addr := fmt.Sprintf(`:%d`, gw.port)
	return gw.engine.Run(addr)
}

func (gw *httpGateWay) setupRoutes() {
	gw.engine.GET("/ping", gw.handlePing)
	gw.setupAuthenticationRoutes()
}

func (gw *httpGateWay) handlePing(c *gin.Context) {
	resp := BasicReponse{
		Error: nil,
		Data:  "pong",
	}
	c.JSON(http.StatusOK, &resp)
}
