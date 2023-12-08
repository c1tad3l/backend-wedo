package users

type User struct {
	Id                    int     `gorm:"PrimaryKey" json:"id"`
	Name                  string  `json:"name"`
	LastName              string  `json:"last_name"`
	Surname               string  `json:"surname"`
	Phone                 string  `json:"phone"`
	Email                 string  `json:"email"`
	EmailVerification     bool    `json:"email_verification"`
	PassportDate          string  `json:"passport_date"`
	PassportSeries        string  `json:"passport_series"`
	PassportNumber        string  `json:"passport_number"`
	PassportBy            string  `json:"passport_by"`
	CertificateNumber     string  `json:"certificate_number"`
	CertificateDate       string  `json:"certificate_date"`
	CertificateSchoolName string  `json:"certificate_school_name"`
	AveragePoint          float64 `json:"average_point"`
	IsGeneralEducation    bool    `json:"is_general_education"`
	IsCitizenship         bool    `json:"is_citizenship"`
	Role                  string  `json:"role"`
}
