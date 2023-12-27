package entree

import (
	"errors"
	"net/http"

	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	id := c.Param("id")
	parent := reqBodyData.UsersVals
	err := c.Bind(&parent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": err,
		})
		return
	}
	var perents users.UserParents
	сheckParentsId := initializers.DB.First(&perents, "id = ? ", id).Error
	if errors.Is(сheckParentsId, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "такой пользователь не найден",
		})
		return
	}
	initializers.DB.Model(&perents).Updates(users.UserParents{
		Name:     parent.FirstName,
		LastName: parent.FirstLastName,
		Surname:  parent.FirstSurname,
	})
	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": "Данные успешно изменены",
	})
	c.JSON(200, gin.H{
		"perent": perents,
	})
}

// обновление паспортных данных

func UpdatePassport(c *gin.Context) {
	id := c.Param("id")
	pass := reqBodyData.UsersVals
	err := c.BindJSON(&pass)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": err,
		})
		return
	}
	var userpass users.User
	сheckParentsId := initializers.DB.First(&userpass, "id = ? ", id).Error
	if errors.Is(сheckParentsId, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "такой пользователь не найден",
		})
		return
	}
	initializers.DB.Model(&userpass).Updates(users.User{
		PassportDate:  pass.PassportDate,
		PassportSeries: pass.PassportSeries,
		PassportNumber:   pass.PassportNumber,
		PassportBy:    pass.PassportBy,
	})
	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": "Данные успешно изменены",
	})
	c.JSON(200, gin.H{
		"user": userpass,
	})
}
