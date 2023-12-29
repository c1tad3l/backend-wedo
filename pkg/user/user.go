package user

import (
	"errors"
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetUser Получение любого пользователя
func GetUser(c *gin.Context) {

	_, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  true,
			"result": "Пользователь не авторизован",
		})
		return
	}

	id := c.Param("id")

	var user users.User
	answer := initializers.DB.Preload("UserParents").Preload("UserEstimates").Find(&user, "id=?", id)

	if errors.Is(answer.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Такой пользователь не найден",
		})
		return
	}

	user.AveragePoint = calcAveragePoints(user.UserEstimates)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"user":  user,
	})
	return
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
