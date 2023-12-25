package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/verification", h.verificationEmail)
		auth.POST("/sendCode", h.sendEmailCode)
		auth.POST("/reset-password", h.resetPassword)
	}

	entree := router.Group("/entree")
	{
		entree.GET("/", h.getAllEntree)
		entree.PUT("/update-esmts/:id", h.updateEstmts)
		entree.PUT("/update-parents/:id", h.updateParents)
		entree.PUT("/update-passport/:id", h.updatePassword)
	}
	return router
}

type Auth interface {
	CreateUser()
	LoginUser()
	verificationEmail()
	sendEmailCode()
	ResetPassword()
}

type Entree interface {
	GetAllEntries()
	UpdateEstms()
	UpdateParentsInfo()
	UpdatePassport()
}
type Handler struct {
	Auth
	Entree
}
