package users

type UserParents struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	FirstLastName  string `json:"first_last_name"`
	FirstSurname   string `json:"first_surname"`
	SecondName     string `json:"second_name"`
	SecondLastName string `json:"second_last_name"`
	SecondSurname  string `json:"second_surname"`
	UserId         int    `json:"user_id"`
}
