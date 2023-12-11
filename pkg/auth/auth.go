package auth

import (
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context) {

	c.Bind(&reqBodyData.UsersVals)
	user := users.User{
		Name:                  reqBodyData.UsersVals.Name,
		LastName:              reqBodyData.UsersVals.LastName,
		Surname:               reqBodyData.UsersVals.Surname,
		Phone:                 reqBodyData.UsersVals.Phone,
		Email:                 reqBodyData.UsersVals.Email,
		EmailVerification:     reqBodyData.UsersVals.EmailVerification,
		PassportDate:          reqBodyData.UsersVals.PassportDate,
		PassportSeries:        reqBodyData.UsersVals.PassportSeries,
		PassportNumber:        reqBodyData.UsersVals.PassportNumber,
		PassportBy:            reqBodyData.UsersVals.PassportBy,
		CertificateNumber:     reqBodyData.UsersVals.CertificateNumber,
		CertificateDate:       reqBodyData.UsersVals.CertificateDate,
		CertificateSchoolName: reqBodyData.UsersVals.CertificateSchoolName,
		AveragePoint:          reqBodyData.UsersVals.AveragePoint,
		IsGeneralEducation:    reqBodyData.UsersVals.IsGeneralEducation,
		IsCitizenship:         reqBodyData.UsersVals.IsCitizenship,
		Role:                  reqBodyData.UsersVals.Role,
	}
	estms := users.UserEstimates{

		Name:  reqBodyData.UsersVals.EstmtName,
		Grade: reqBodyData.UsersVals.Grade,
	}
	parents := users.UserParents{
		FirstName:      reqBodyData.UsersVals.FirstName,
		FirstLastName:  reqBodyData.UsersVals.FirstLastName,
		FirstSurname:   reqBodyData.UsersVals.SecondSurname,
		SecondName:     reqBodyData.UsersVals.SecondName,
		SecondLastName: reqBodyData.UsersVals.SecondLastName,
		SecondSurname:  reqBodyData.UsersVals.SecondSurname,
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

	c.Bind(&reqBodyData.LogingVals)

	var user users.User
	initializers.DB.First(&user, "email = ?", reqBodyData.LogingVals.Email)

	//if user.Id == 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"Ошибка": "Не правильно введен email или проверочный код",
	//	})
	//}

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
