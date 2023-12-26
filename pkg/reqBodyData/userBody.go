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
	PassportDate          string
	PassportSeries        string
	PassportNumber        string
	PassportBy            string
	CertificateNumber     string
	CertificateDate       string
	CertificateSchoolName string
	AveragePoint          float64
	IsGeneralEducation    bool
	IsCitizenship         bool
	Role                  string
	EstmtName             string
	Grade                 string
	FirstName             string
	FirstLastName         string
	FirstSurname          string
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
