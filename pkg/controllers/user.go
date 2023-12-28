package controllers

import (
	"github.com/c1tad3l/backend-wedo/pkg/user"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	user.GetUser(c)

}
