package auth

import (
	"errors"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/c1tad3l/backend-wedo/initializers"
	sender "github.com/c1tad3l/backend-wedo/pkg/mail"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"github.com/c1tad3l/backend-wedo/pkg/reqBodyData"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser Создания пользователя
func CreateUser(c *gin.Context) {
	id := uuid.New()
	randomString := GenerateRandomString(10)

	///пароль отобразиться в консоли, так как результат будет уже захеширован (удалить когда он уже будет отправляться на почту)
	//fmt.Println(randomString)

	hash, err := bcrypt.GenerateFromPassword([]byte(randomString), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Ошибка хеширования",
		})
	}

	uservals := reqBodyData.UsersVals
	err = c.BindJSON(&uservals)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Не введены данные",
		})
		return
	}

	var userEstimates []users.UserEstimates

	for _, estimate := range uservals.Estimates {
		userEstimates = append(userEstimates, users.UserEstimates{
			Id:    uuid.New(),
			Name:  estimate.Name,
			Grade: estimate.Grade,
		})
	}

	var userParents []users.UserParents

	for _, parents := range uservals.Parents {
		userParents = append(userParents, users.UserParents{
			Id:       uuid.New(),
			Name:     parents.Name,
			LastName: parents.LastName,
			Surname:  parents.Surname,
			Phone:    parents.Phone,
		})
	}

	/// проверка на то сущесвтует ли email в базе данных
	var usr users.User
	checkmail := initializers.DB.First(&usr, "email = ?", uservals.Email).Error
	if !errors.Is(checkmail, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"resutl": "такой email ужe cуществует",
		})
		return
	}

	if uservals.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "пожалуйста присвойте роль пользователю",
		})
		return
	}

	user := users.User{
		Id:                    id,
		Name:                  uservals.Name,
		LastName:              uservals.LastName,
		Surname:               uservals.Surname,
		Password:              string(hash),
		Genre:                 uservals.Genre,
		Birthday:              uservals.Birthday,
		Phone:                 uservals.Phone,
		Email:                 uservals.Email,
		EmailVerification:     uservals.EmailVerification,
		PassportDate:          uservals.PassportDate,
		PassportSeries:        uservals.PassportSeries,
		PassportNumber:        uservals.PassportNumber,
		PassportBy:            uservals.PassportBy,
		PassportAddress:       uservals.PassportAddress,
		CertificateNumber:     uservals.CertificateNumber,
		CertificateDate:       uservals.CertificateDate,
		CertificateSchoolName: uservals.CertificateSchoolName,
		CertificateBy:         uservals.CertificateBy,
		IsGeneralEducation:    uservals.IsGeneralEducation,
		IsCitizenship:         uservals.IsCitizenship,
		Role:                  uservals.Role,
		UserParents:           userParents,
		UserEstimates:         userEstimates,
	}
	userResult := initializers.DB.Create(&user)

	if userResult.Error != nil {
		c.Status(400)
		return
	}

	to := []string{uservals.Email}
	cc := []string{}
	bcc := []string{}

	subject := "Первичный пароль"
	mailtype := "html"
	replyToAddress := ""

	body := `
			<html>
			<body>
			<h1>
				Приветствуем!
			</h1><br>
			<h3>
				Ваш код для входа в аккаунт:
			</h3>
			<h2>` + randomString + ` </h2><br><br><h5>Пароль можно будет сменить после входа в аккаунт<br>На это сообщение не нужно отвечать.</h5> </body> </html>`

	err = sender.SendToMail(subject, body, mailtype, replyToAddress, to, cc, bcc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  true,
			"result": "Произошла какая то непредвиденная ошибка",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"user":  user,
	})

	//c.JSON(http.StatusOK, gin.H{
	//	"error":  false,
	//	"result": "новый пользователь создан"})
}

// LoginUser Авторизация пользователя
func LoginUser(c *gin.Context) {
	loginVals := reqBodyData.LogingVals
	err := c.Bind(&loginVals)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Не введен email или пароль",
		})
		return
	}
	var user users.User

	mailCheck := initializers.DB.First(&user, "email = ?", loginVals.Email).Error
	if errors.Is(mailCheck, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Не правильно введен email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Не правильно введен пароль",
		})
	}

	//jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("we4r5678987654e3w3e456789876yt5rewr5t678765r"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  true,
			"result": "не получилось создать токен",
		})
	}
	c.SetSameSite(http.SameSiteLaxMode)

	//поменять secure parameter в будущем
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"token":  tokenString,
		"userId": user.Id,
	})

}

// VerificationMail Подтверждение почты
func VerificationMail(c *gin.Context) {

	data := users.Verification
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

	initializers.DB.Exec("DELETE FROM emails WHERE code=" + "'" + data.Code + "'" + "AND email=" + "'" + data.Email + "'")

	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": true,
	})
	return
}

// ResetPassword Сброс пароля
func ResetPassword(c *gin.Context) {
	var user users.User
	data := reqBodyData.UserPassword
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Укажите email пользователя",
		})
		return
	}

	answer := checkingEmailReg(data.Email)

	if !answer {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Неверно указана почта",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "error",
		})
	}

	checkEmail := checkingEmailInBD(data.Email)
	if !checkEmail {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Email не совпал с почтой из базы",
		})
		return
	}

	initializers.DB.Model(&user).Updates(users.User{Password: string(hash)})
	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": "Пароль успешно изменён",
	})
}

// SendEmailCode Функция для отправки кода на почту
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

	answerCheck := checkingEmailReg(email.Email)

	if !answerCheck {
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
			<h2>` + code + ` </h2><br><br><h5>Код будет активен в течении 30 минут.<br>На это сообщение не нужно отвечать.</h5> </body> </html>`

	err = sender.SendToMail(subject, body, mailtype, replyToAddress, to, cc, bcc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  true,
			"result": "Произошла какая то непредвиденная ошибка",
		})
		return
	}

	initializers.DB.Create(&users.Email{Email: email.Email, Code: code})

	go deleteCodeAfterTime(code, email.Email)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})
	return

}

// Функции который используется в качестве вспомогательных или обработчиков

func deleteCodeAfterTime(code string, email string) {
	time.Sleep(30 * time.Minute)

	answer := initializers.DB.First(&users.Email{Email: email, Code: code})

	if errors.Is(answer.Error, gorm.ErrRecordNotFound) {
		return
	} else {
		initializers.DB.Exec("DELETE FROM emails WHERE code=" + "'" + code + "'" + "AND email=" + "'" + email + "'")
	}
}

func generationCode() string {

	var randomCode = ""
	for i := 0; i < 4; i++ {
		res := rand.Intn(9)
		randomCode = randomCode + strconv.Itoa(int(res))
	}

	return randomCode
}

func checkingEmailReg(email string) bool {

	matched, _ := regexp.MatchString(`([A-Za-z0-9_\-.])+@([A-Za-z0-9_\-.])+\.([A-Za-z]{2,4})`, email)

	if !matched {

		return false
	}
	return true
}

func checkingEmailInBD(email string) bool {
	var user users.User
	answer := initializers.DB.First(&user, "email = ?", email)

	if errors.Is(answer.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = reqBodyData.Charset[reqBodyData.SeededRand.Intn(len(reqBodyData.Charset))]
	}
	return string(b)
}
