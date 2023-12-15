package users

var EmailType struct {
	Email string `json:"email"`
}

type Email struct {
	Id    uint64 `gorm:"primaryKey"`
	Email string `json:"email"`
	Code  string `json:"code"`
}
