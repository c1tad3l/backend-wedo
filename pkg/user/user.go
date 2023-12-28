package user

import (
	"errors"
	"fmt"
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

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
	answer := initializers.DB.First(&user, "id = ?", id)

	if errors.Is(answer.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Такой пользователь не найден",
		})
		return
	}

	fmt.Print(user.Surname)

	_, _ = calcAveragePoints(user.Id.String())

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"user":  user,
	})
	return
}

// вычисление среднего балла
func calcAveragePoints(userId string) (result int, err bool) {

	var estimates users.UserEstimates
	answer := initializers.DB.Limit(99).Take(&estimates, "id = ?", userId)

	fmt.Println(estimates)
	if errors.Is(answer.Error, gorm.ErrRecordNotFound) {
		return 0, true
	}

	//fmt.Println(estimates)
	return 0, false
}
