package users

import "github.com/google/uuid"

type UserEstimates struct {
	Id    uuid.UUID `gorm:"PrimaryKey;" json:"id"`
	Name  string    `json:"estmtName"`
	Grade string    `json:"grade"`
}
