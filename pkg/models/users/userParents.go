package users

import "github.com/google/uuid"

type UserParents struct {
	Id       uuid.UUID `gorm:"PrimaryKey;" json:"id"`
	Name     string    `json:"parents_name"`
	LastName string    `json:"parentsLast_name"`
	Surname  string    `json:"parents_surname"`
}
