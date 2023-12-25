package controllers

import (
	"github.com/c1tad3l/backend-wedo/pkg/roles/entree"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllEntree(c *gin.Context) {
	entree.GetAllEntries(c)

}
func (h *Handler) updateEstmts(c *gin.Context) {
	entree.UpdateEstms(c)

}
func (h *Handler) updateParents(c *gin.Context) {
	entree.UpdateParentsInfo(c)

}
func (h *Handler) updatePassword(c *gin.Context) {
	entree.UpdatePassport(c)

}
