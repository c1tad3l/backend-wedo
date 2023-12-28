package entree

import (
	"errors"
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// получение всех абитуриентов

func GetAllEntries(c *gin.Context) {
	var user []users.User
	initializers.DB.Preload("UserParents").Preload("UserEstimates").Find(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}

// обновление атестата

func UpdateEstms(c *gin.Context) {
	id := c.Param("id")
	estms := reqBodyData.UsersVals
	err := c.BindJSON(&estms)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": err,
		})
		return
	}
	var estimates users.UserEstimates
	сheckEstmtsId := initializers.DB.First(&estimates, "id = ? ", id).Error
	if errors.Is(сheckEstmtsId, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "такой пользователь не найден",
		})
		return
	}
	initializers.DB.Model(&estimates).Updates(users.UserEstimates{
		Name:  estms.EstmtName,
		Grade: estms.Grade,
	})

	c.JSON(200, gin.H{
		"estms": estimates,
	})
}

// обновление данных о родителях

func UpdateParentsInfo(c *gin.Context) {

}

// обновление паспортных данных

func UpdatePassport(c *gin.Context) {

}
