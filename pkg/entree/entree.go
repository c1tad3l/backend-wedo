package entree

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllEntries получение всех абитуриентов
func GetAllEntries(c *gin.Context) {
	var user []users.User
	initializers.DB.Preload("UserParents").Preload("UserEstimates").Find(&user)

	for i := 0; i < len(user); i++ {
		user[i].AveragePoint = calcAveragePoints(user[i].UserEstimates)
	}

	c.JSON(200, gin.H{
		"error": false,
		"user":  user,
	})
}

// UpdateEstms Обновление атестата
func UpdateEstms(c *gin.Context) {
	id := c.Param("id")
	estimates := reqBodyData.EstimatesUpdate
	err := c.BindJSON(&estimates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": err,
		})
		return
	}
	var userEstimates users.UserEstimates
	сheckEstmtsId := initializers.DB.First(&userEstimates, "id = ? ", id).Error
	if errors.Is(сheckEstmtsId, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Нет такой записи",
		})
		return
	}
	initializers.DB.Model(&userEstimates).Updates(users.UserEstimates{
		Name:  estimates.Name,
		Grade: estimates.Grade,
	})

	c.JSON(200, gin.H{
		"error":     false,
		"estimates": userEstimates,
	})
}

// UpdateParentsInfo обновление данных о родителях
func UpdateParentsInfo(c *gin.Context) {
	id := c.Param("id")
	parent := reqBodyData.ParentsUpdate
	err := c.Bind(&parent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": err,
		})
		return
	}
	var parents users.UserParents
	сheckParentsId := initializers.DB.First(&parents, "id = ? ", id).Error
	if errors.Is(сheckParentsId, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "такой пользователь не найден",
		})
		return
	}
	initializers.DB.Model(&parents).Updates(users.UserParents{
		Name:     parent.Name,
		LastName: parent.LastName,
		Surname:  parent.Surname,
		Phone:    parent.Phone,
	})
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"result":  "Данные успешно изменены",
		"parents": parents,
	})
}

// UpdatePassport обновление паспортных данных
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
		PassportDate:    pass.PassportDate,
		PassportSeries:  pass.PassportSeries,
		PassportNumber:  pass.PassportNumber,
		PassportBy:      pass.PassportBy,
		PassportAddress: pass.PassportAddress,
	})
	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": "Данные успешно изменены",
		"user":   userpass,
	})
}

// вычисление среднего балла
func calcAveragePoints(userEstimates []users.UserEstimates) float64 {

	if len(userEstimates) != 0 {
		result := float64(0)
		for i := 0; i < len(userEstimates); i++ {
			grade, _ := strconv.ParseFloat(userEstimates[i].Grade, 64)
			result = result + grade
		}

		return result / float64(len(userEstimates))
	} else {
		return 0
	}
}
