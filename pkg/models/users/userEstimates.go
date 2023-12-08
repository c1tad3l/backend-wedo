package users

type UserEstimates struct {
	Id     int    ` gorm:"PrimaryKey" json:"id"`
	Name   string `json:"name"`
	Grade  string `json:"grade"`
	UserId int    `gorm:"foreignKey:Id" json:"user_id"`
}
