package gateway

import "github.com/gin-gonic/gin"

func (gw *httpGateWay) setupAuthenticationRoutes() {
	// authentication group
	authGroup := gw.engine.Group("/auth/")

	// user group
	adminGroup := authGroup.Group("/user")
	adminGroup.POST("/sign-up", gw.handleSignUp)
	adminGroup.GET("/activate", gw.handleActivate)
	adminGroup.GET("/resend-code", gw.handleResendCode)
	adminGroup.POST("/log-in", gw.handleLogIn)
}

func (gw *httpGateWay) handleSignUp(c *gin.Context) {
	// TODO
}

func (gw *httpGateWay) handleActivate(c *gin.Context) {
	// TODO
}

func (gw *httpGateWay) handleResendCode(c *gin.Context) {
	// TODO
}

func (gw *httpGateWay) handleLogIn(c *gin.Context) {
	// TODO
}
