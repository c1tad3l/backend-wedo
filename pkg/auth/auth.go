package auth

import (
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func CreateUser(c *gin.Context) {

	/// посылаемые запросом данные(body)
	var usersVals struct {
		Id                    int
		Name                  string
		LastName              string
		Surname               string
		Phone                 string
		Email                 string
		EmailVerification     bool
		PassportDate          string
		PassportSeries        string
		PassportNumber        string
		PassportBy            string
		CertificateNumber     string
		CertificateDate       string
		CertificateSchoolName string
		AveragePoint          float64
		IsGeneralEducation    bool
		IsCitizenship         bool
		Role                  string
		EstmtName             string
		Grade                 string
		UserId                int
		FirstName             string
		FirstLastName         string
		FirstSurname          string
		SecondName            string
		SecondLastName        string
		SecondSurname         string
	}

	c.Bind(&usersVals)

	user := users.User{
		Name:                  usersVals.Name,
		LastName:              usersVals.LastName,
		Surname:               usersVals.Surname,
		Phone:                 usersVals.Phone,
		Email:                 usersVals.Email,
		EmailVerification:     usersVals.EmailVerification,
		PassportDate:          usersVals.PassportDate,
		PassportSeries:        usersVals.PassportSeries,
		PassportNumber:        usersVals.PassportNumber,
		PassportBy:            usersVals.PassportBy,
		CertificateNumber:     usersVals.CertificateNumber,
		CertificateDate:       usersVals.CertificateDate,
		CertificateSchoolName: usersVals.CertificateSchoolName,
		AveragePoint:          usersVals.AveragePoint,
		IsGeneralEducation:    usersVals.IsGeneralEducation,
		IsCitizenship:         usersVals.IsCitizenship,
		Role:                  usersVals.Role,
	}
	estms := users.UserEstimates{

		Name:  usersVals.EstmtName,
		Grade: usersVals.Grade,
	}
	parents := users.UserParents{
		FirstName:      usersVals.FirstName,
		FirstLastName:  usersVals.FirstLastName,
		FirstSurname:   usersVals.SecondSurname,
		SecondName:     usersVals.SecondName,
		SecondLastName: usersVals.SecondLastName,
		SecondSurname:  usersVals.SecondSurname,
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
	var logingVals struct {
		Email            string
		VerificationCode string
	}
	c.Bind(&logingVals)

	var user users.User
	initializers.DB.First(&user, "email = ?", logingVals.Email)

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Ошибка": "Не правильно введен email или проверочный код",
		})
	}

	////проверочный код//
	///

	//jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
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
