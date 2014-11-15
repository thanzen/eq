package user

type User struct {
	Id        int    `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"first_name"`
	Lastname  string `json:"lastname" db:"last_name"`
	Username  string `json:"username" db:"user_name"`
	//Password    string `json:"password" db:"password"`
	Email       string `json:"email" db:"email"`
	Age         int    `json:"age" db:"age"`
	Country     string `json:"country" db:"country"`
	City        string `json:"city" db:"city"`
	Postcode    string `json:"postcode" db:"post_code"`
	Cellphone   string `json:"cellphone" db:"cell_phone"`
	Homephone   string `json:"homephone" db:"home_phone"`
	Officephone string `json:"officephone" db:"office_phone"`
	Fax         string `json:"fax" db:"fax"`
}
