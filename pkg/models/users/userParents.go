package users

import "github.com/google/uuid"

type UserParents struct {
	Id       uuid.UUID `gorm:"PrimaryKey;" json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Surname  string    `json:"surname"`
	Phone    string    `json:"phone"`
}
