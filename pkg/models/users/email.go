package users

import "github.com/google/uuid"

var EmailType struct {
	Email string `json:"email"`
}

type Email struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Code  string    `json:"code"`
}
