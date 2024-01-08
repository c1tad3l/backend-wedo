package users

import (
	"github.com/google/uuid"
)

type User struct {
	Id                    uuid.UUID       `gorm:"PrimaryKey" json:"id"`
	Name                  string          `json:"name"`
	Password              string          `json:"password"`
	LastName              string          `json:"last_name"`
	Surname               string          `json:"surname"`
	Birthday              string          `json:"birthday"`
	Genre                 string          `json:"genre"`
	Phone                 string          `json:"phone"`
	Email                 string          `json:"email"`
	EmailVerification     bool            `json:"email_verification"`
	PassportDate          string          `json:"passport_date"`
	PassportSeries        string          `json:"passport_series"`
	PassportNumber        string          `json:"passport_number"`
	PassportBy            string          `json:"passport_by"`
	PassportAddress       string          `json:"passport_address"`
	CertificateNumber     string          `json:"certificate_number"`
	CertificateDate       string          `json:"certificate_date"`
	CertificateSchoolName string          `json:"certificate_school_name"`
	CertificateBy         string          `json:"certificate_by"`
	IsGeneralEducation    bool            `json:"is_general_education"`
	IsCitizenship         bool            `json:"is_citizenship"`
	Role                  string          `json:"role"`
	AveragePoint          float64         `json:"average_point,omitempty"`
	UserParents           []UserParents   `gorm:"many2many:user_user_parents;" json:"user_parents,omitempty"`
	UserEstimates         []UserEstimates `gorm:"many2many:user_user_estimates;" json:"user_estimates,omitempty"`
}
