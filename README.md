# backend-wedo


Пакеты для изучения:

github.com/gin-gonic/gin<br>

https://gorm.io/index.html
github.com/spf13/viper
<h1>чтобы запустить сервер<h1> 
<h2 style="background: white; color: black; padding:1rem">
go run cmd/main.go
<h2>

<h1>Чтобы выполнить миграцию</h1>
<h2 style="background: white; color: black; padding:1rem">
go run migration/migration.go</h2>

<h1>Запросы</h1>

<h1>POST</h1>
<h2 style="background: white; color: black; padding:1rem">/auth/sing-up</h2>

<h3>Запрос для регистрации пользователей</h3>

<h3>Typescript request </h3>

```typescript
type  userParams={
    name:   string
	last_name:   string
	surname:    string
	phone:  string
	email:  string
	email_verification:  boolean
	passport_date:   string
	passport_series: string
	passport_number: string
	passport_by: string
	certificate_number:  string
	certificate_date:    string
	certificate_school_name:  string
	average_point:   number
	is_general_education: boolean
	is_citizenship: boolean
	role: string
}
````
<h3>GO model</h3>

```
type User struct {
	Id                    uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Name                  string    `json:"name"`
	Password              string    `json:"password"`
	LastName              string    `json:"last_name"`
	Surname               string    `json:"surname"`
	Phone                 string    `json:"phone"`
	Email                 string    `json:"email"`
	EmailVerification     bool      `json:"email_verification"`
	PassportDate          string    `json:"passport_date"`
	PassportSeries        string    `json:"passport_series"`
	PassportNumber        string    `json:"passport_number"`
	PassportBy            string    `json:"passport_by"`
	CertificateNumber     string    `json:"certificate_number"`
	CertificateDate       string    `json:"certificate_date"`
	CertificateSchoolName string    `json:"certificate_school_name"`
	AveragePoint          float64   `json:"average_point"`
	IsGeneralEducation    bool      `json:"is_general_education"`
	IsCitizenship         bool      `json:"is_citizenship"`
	Role                  string    `json:"role"`
}
````
<h3>Typescipt request</h3>

```typescript
type estmsParams={
    estmt_name :string
    grade:  string
}
````
<h3>GO model</h3>

```
type UserEstimates struct {
	Id     uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Name   string    `json:"name"`
	Grade  string    `json:"grade"`
	UserId uuid.UUID `gorm:"foreignKey:Id" json:"user_id"`
}
````

<h3>Typescipt request</h3>

```typescript
type parentsParams={
    name :string
    last_name: string
    surname: string
}
````
<h3>GO model</h3>

```
type UserParents struct {
	Id       uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Surname  string    `json:"surname"`
	UserId   uuid.UUID `gorm:"foreignKey:Id" json:"user_id"`
}
````

<h2>Варианты ответов:</h2>

<h3>200</h3>

```typescript
response = {
    error: false,
    token: tokenString,
}
```


<h1>POST</h1>
<h2 style="background: white; color: black; padding:1rem">/auth/sign-in</h2>

<h3>Запрос для авторизации пользователя</h3>

<h3>Typescript request</h3>

```typescript
type params={
    email:  string
    password:   string
}
```

<h3>GO model</h3>
```
var UserPassword struct{
	Email                 string    `json:"email"`
	Password              string    `json:"password"`
}
```

<h2>Варианты ответов:</h2>

<h3>200</h3>

```typescript
response = {
    error: false,
    token: tokenString,
}
```

<h3>400</h3>
<h4>Пароль не совпал с паролем в базе данных</h2>

```typescript
response = {
    error: true,
    result: "Не правильно введен Пароль",
}
```

<h3>400</h3>
<h4>Поле email или поле password оказались пустыми</h2>

```typescript
response = {
    error:  true,
    result: "Не введен email или пароль",
}
```

<h3>404</h3>
<h4>Не удалось найти email в базе данных</h2>

```typescript
response = {
    error: true,
    result: "Не правильно введен email",
}
```

<h3>500</h3>
<h4>По каким то причинам сервер не смог сгенирировать токен защиты</h2>

```typescript
response = {
    "error": true,
    "result": "не получилось создать токен",
}
```

<h1>POST</h1>
<h2 style="background: white; color: black; padding:1rem">/auth/reset-password</h2>



<h1>POST</h1>
<h2 style="background: white; color: black; padding:1rem">/auth/verification</h2>

<h3>Запрос для подтверждения почты</h3>

<h3>Typescript request</h3>

```typescript
type params = {
    email: string,
    code: string
}
```
<h3>GO model</h3>

```
var Verification struct {
	Email string
	Code  string
}
```

<h2>Варианты ответов:</h2>

<h3>200</h3>

```typescript
response = {
    error:  false,
    result: true,
}

```

<h3>400</h3>
<h4>Email не был отправлен вообще, или же оказалось пустым</h2>

```typescript
response = {
    error:  true,
    result: "Укажите email пользователя",
}
```

<h3>400</h3>
<h4>Почта не прошла валидацию</h2>

```typescript
response = {
    error:  true,
    result: "Неверно указана почта",
}
```
<h3>404</h3>
<h4>Почта или код не существует в базе данных ( можно доработать и сделать проверку чего конкретно нет )</h2>

```typescript
response = {
    error:  true,
    result: "Не правильно введен email или проверочный код",
}
```

<h1>POST</h1>
<h2 style="background: white; color: black; padding:1rem">/auth/sendCode </h2>

<h3>Запрос для отправки на почту код подтверждения</h3>

<h3>Typescript request</h3>

```typescript
type params = {
    email: string
}
```
<h3>GO model</h3>

```
var EmailType struct {
Email string `json:"email"`
}
```

<h2>Варианты ответов:</h2>

<h3>200</h3>

```typescript
response = {
error: false
}
```

<h3>400</h3>
<h4>Email не был отправлен вообще, или же оказалось пустым</h2>

```typescript
response = {
    error:  true,
    result: "Укажите email пользователя",
}
```

<h3>400</h3>
<h4>Почта не прошла валидацию</h2>

```typescript
response = {
    error:  true,
    result: "Неверно указана почта",
}
```

<h3 id="test">500</h3>
<h4>По какой то причине произошла ошибки при отправке, тяжело сделать обработчик возможно поможет отправить снова, в будущем может что то придумаем если такое случится.D</h2>

```typescript
response = {
    error:  true,
    result: "Произошла какая то непредвиденная ошибка",
}
```
