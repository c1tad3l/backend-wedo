package reqBodyData

import (
	"math/rand"
	"time"
)

var UsersVals struct {
	Name                  string
	LastName              string
	Surname               string
	Phone                 string
	Email                 string
	EmailVerification     bool
	PassportDate          string `json:"passport_date"`
	PassportSeries        string `json:"passport_series"`
	PassportNumber        string `json:"passport_number"`
	PassportBy            string `json:"passport_by"`
	CertificateNumber     string
	CertificateDate       string
	CertificateSchoolName string
	AveragePoint          float64
	IsGeneralEducation    bool
	IsCitizenship         bool
	Role                  string
	EstmtName             string
	Grade                 string
	ParentFirstName       string `json:"parentsName"`
	ParentFirstLastName   string `json:"parentsLast_name"`
	ParentFirstSurname    string `json:"surname"`
}

var LogingVals struct {
	Email    string
	Password string
}

var UserPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var SeededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
