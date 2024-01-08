package reqBodyData

import (
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
	"math/rand"
	"time"
)

var UsersVals struct {
	Name                  string `json:"name"`
	LastName              string `json:"last_name"`
	Surname               string `json:"surname"`
	Genre                 string `json:"genre"`
	Birthday              string `json:"birthday"`
	Phone                 string `json:"phone"`
	Email                 string `json:"email"`
	EmailVerification     bool   `json:"email_verification"`
	PassportDate          string `json:"passport_date"`
	PassportSeries        string `json:"passport_series"`
	PassportNumber        string `json:"passport_number"`
	PassportBy            string `json:"passport_by"`
	PassportAddress       string `json:"passport_address"`
	CertificateNumber     string `json:"certificate_number"`
	CertificateDate       string `json:"certificate_date"`
	CertificateBy         string `json:"certificate_by"`
	CertificateSchoolName string `json:"certificate_school_name"`
	AveragePoint          float64
	IsGeneralEducation    bool                  `json:"is_general_education"`
	IsCitizenship         bool                  `json:"is_citizenship"`
	Role                  string                `json:"role"`
	Estimates             []users.UserEstimates `json:"user_estimates"`
	Parents               []users.UserParents   `json:"user_parents"`
}

var LogingVals struct {
	Email    string
	Password string
}
var EstimatesUpdate struct {
	Name  string `json:"estimates_name"`
	Grade string `json:"grade"`
}
var ParentsUpdate struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
}

var UserPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var SeededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
