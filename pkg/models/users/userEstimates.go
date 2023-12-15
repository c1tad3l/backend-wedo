package users

import "github.com/google/uuid"

type UserEstimates struct {
	Id     uuid.UUID `gorm:"PrimaryKey" json:"id"`
	Name   string    `json:"name"`
	Grade  string    `json:"grade"`
	UserId uuid.UUID `json:"user_id"`
	User   User      `gorm:"foreignKey:UserId"`
}
