package reqBodyData

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
	SecondName            string
	SecondLastName        string
	SecondSurname         string
}

var LogingVals struct {
	Email            string
	VerificationCode string
}
