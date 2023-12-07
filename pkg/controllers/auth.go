package controllers

import (
	"github.com/c1tad3l/backend-wedo/pkg/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	auth.CreateUser()
}

func (h *Handler) signIn(c *gin.Context) {
	auth.LoginUser()
}

func (h *Handler) verificationEmail(c *gin.Context) {
	auth.VerificationUser()
}
