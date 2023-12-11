package controllers

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/verification", h.verificationEmail)
		auth.POST("/sendCode", h.sendEmailCode)
	}

	return router
}

type Auth interface {
	CreateUser()
	LoginUser()
	verificationEmail()
	sendEmailCode()
}

type Handler struct {
	Auth
}
