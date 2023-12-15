package users

import "github.com/google/uuid"

type UserParents struct {
	Id             uuid.UUID `gorm:"PrimaryKey" json:"id"`
	FirstName      string    `json:"first_name"`
	FirstLastName  string    `json:"first_last_name"`
	FirstSurname   string    `json:"first_surname"`
	SecondName     string    `json:"second_name"`
	SecondLastName string    `json:"second_last_name"`
	SecondSurname  string    `json:"second_surname"`
	UserId         uuid.UUID `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId"`
}
