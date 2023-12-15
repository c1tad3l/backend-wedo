package auth

import (
	"errors"
	"github.com/c1tad3l/backend-wedo/initializers"
	sender "github.com/c1tad3l/backend-wedo/pkg/mail"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
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
			"error": "не получилось создать токен",
		})
	}
	c.SetSameSite(http.SameSiteLaxMode)

	//поменять в будущем
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
func VerificationMail(c *gin.Context) {

	data := &users.Verification
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Укажите email пользователя",
		})
		return
	}

	matched, _ := regexp.MatchString(`([A-Za-z0-9_\-.])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})`, data.Email)

	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Неверно указана почта",
		})
		return
	}

	answer := initializers.DB.First(&users.Email{Email: data.Email, Code: data.Code})

	if errors.Is(answer.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Не правильно введен email или проверочный код",
		})
		return
	}

	initializers.DB.Exec("DELETE FROM emails WHERE code=" + "'" + data.Code + "'")

	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": true,
	})
	return
}

func SendEmailCode(c *gin.Context) {

	email := &users.EmailType
	err := c.BindJSON(&email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Укажите email пользователя",
		})
		return
	}

	matched, _ := regexp.MatchString(`([A-Za-z0-9_\-.])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})`, email.Email)

	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Неверно указана почта",
		})
		return
	}

	code := generationCode()

	to := []string{email.Email}
	cc := []string{}
	bcc := []string{}

	subject := "Отправка кода на подтверждения почты"
	mailtype := "html"
	replyToAddress := ""

	body := `
			<html>
			<body>
			<h1>
				Приветствуем!
			</h1><br>
			<h3>
				Код для подтверждения:
			</h3>
			<h2>` + code + ` </h2><br><br><h5>Код будет активен в течении двух часов.<br>На это сообщение не нужно отвечать.</h5> </body> </html>`

	err = sender.SendToMail(subject, body, mailtype, replyToAddress, to, cc, bcc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  true,
			"result": "Произошла какая то непредвиденная ошибка",
		})
		return
	}

	answer := initializers.DB.Create(&users.Email{Email: email.Email, Code: code})

	print(answer.RowsAffected)
	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}

func generationCode() string {

	var randomCode = ""
	for i := 0; i < 10; i++ {
		res := rand.Intn(36)
		randomCode = randomCode + strconv.Itoa(int(res))
	}

	return randomCode
}
