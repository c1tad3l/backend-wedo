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
<h2 style="background: white; color: black; padding:1rem">/auth/sing-up</h2>

```typescript
type  userParams={
    	Name:   string
	LastName:   string
	Surname:    string
	Phone:  string
	Email:  string
	EmailVerification:  boolean
	PassportDate:   string
	PassportSeries: string
	PassportNumber: string
	PassportBy: string
	CertificateNumber:  string
	CertificateDate:    string
	CertificateSchoolName:  string
	AveragePoint:   number
	IsGeneralEducation: boolean
	IsCitizenship: boolean
	Role: string
}
```
```typescript

type estmsParams={
    EstmtName :string
    Grade:  string
}
````
```typescript
type parentsParams={
    FirstName :string
    FirstLastName: string
    FirstSurname: string
    SecondName: string
    SecondLastName: string
    SecondSurname : string
}
````

<h2 style="background: white; color: black; padding:1rem">/auth/sign-in</h2>

```typescript
type params={
    Email:  string
    VerificationCode:   string
}
```

<h2 style="background: white; color: black; padding:1rem">/auth/verification</h2>

```typescript
type params = {
    email: string,
    code: string
}
```

<h2 style="background: white; color: black; padding:1rem">/auth/sendCode</h2>

```typescript
type params = {
    email: string
}
```

