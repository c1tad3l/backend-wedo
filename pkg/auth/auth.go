package auth

import (
	"errors"
	"github.com/c1tad3l/backend-wedo/pkg/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context) {
	uservals := reqBodyData.UsersVals
	err := c.BindJSON(&uservals)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Не введены данные",
		})
		return
	}

	user := users.User{
		Name:                  uservals.Name,
		LastName:              uservals.LastName,
		Surname:               uservals.Surname,
		Phone:                 uservals.Phone,
		Email:                 uservals.Email,
		EmailVerification:     uservals.EmailVerification,
		PassportDate:          uservals.PassportDate,
		PassportSeries:        uservals.PassportSeries,
		PassportNumber:        uservals.PassportNumber,
		PassportBy:            uservals.PassportBy,
		CertificateNumber:     uservals.CertificateNumber,
		CertificateDate:       uservals.CertificateDate,
		CertificateSchoolName: uservals.CertificateSchoolName,
		AveragePoint:          uservals.AveragePoint,
		IsGeneralEducation:    uservals.IsGeneralEducation,
		IsCitizenship:         uservals.IsCitizenship,
		Role:                  uservals.Role,
	}
	estms := users.UserEstimates{

		Name:  uservals.EstmtName,
		Grade: uservals.Grade,
	}
	parents := users.UserParents{
		FirstName:      uservals.FirstName,
		FirstLastName:  uservals.FirstLastName,
		FirstSurname:   uservals.SecondSurname,
		SecondName:     uservals.SecondName,
		SecondLastName: uservals.SecondLastName,
		SecondSurname:  uservals.SecondSurname,
	}

	userResult := initializers.DB.Create(&user)
	etsmtsResult := initializers.DB.Create(&estms)
	parentsResult := initializers.DB.Create(&parents)

	if userResult.Error != nil {
		c.Status(400)
		return
	}

	if etsmtsResult.Error != nil {
		c.Status(400)
		return
	}

	if parentsResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user":    user,
		"estms":   estms,
		"parents": parents,
	})

}

func LoginUser(c *gin.Context) {
	loginVals := reqBodyData.LogingVals
	err := c.Bind(&loginVals)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Не ввден email или пароль",
		})
		return
	}
	var user users.User
	mailCheck := initializers.DB.First(&user, "email = ?", loginVals.Email).Error

	if errors.Is(mailCheck, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Не правильно введен email или проверочный код",
		})
		return
	}
	////проверочный код//
	///
	//jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("we4r5678987654e3w3e456789876yt5rewr5t678765r"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "не получилось создать токен",
		})
	}
	c.SetSameSite(http.SameSiteLaxMode)

	//поменять в будущем
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func VerificationUser() {

}
