package user

import "time"

type User struct {
	Id               int       `json:"id" db:"id"`
	Username         string    `json:"username" db:"user_name"`
	Firstname        string    `json:"firstname" db:"first_name"`
	Lastname         string    `json:"lastname" db:"last_name"`
	Email            string    `json:"email" db:"email"`
	Password         string    `json:"password" db:"password"`
	PasswordSalt     string    `json:"passwordSalt" db:"password_salt"`
	AdminComment     string    `json:"adminComment" db:"admin_comment"`
	Active           bool      `json:"active" db:"active"`
	Verified         bool      `json:"verified" db:"verified"`
	Deleted          bool      `json:"deleted" db:"deleted"`
	IsSystemAccount  bool      `json:"isSystemAccount" db:"is_system_account"`
	LastIp           string    `json:"lastIp" db:"last_ip"`
	Created          time.Time `json:"created" db:"created"`
	LastLoginDate    time.Time `json:"lastLoginDate" db:"last_login_date"`
	LastActivityDate time.Time `json:"lastActivityDate" db:"last_activity_date"`
	CellPhone        string    `json:"cellPhone" db:"cell_phone"`
	Officephone      string    `json:"officePhone" db:"office_phone"`
	Fax              string    `json:"fax" db:"fax"`
	Country          string    `json:"country" db:"country"`
	City             string    `json:"city" db:"city"`
	Postcode         string    `json:"postCode" db:"post_code"`
}
type UpdateAccount struct {
	Id              int       `json:"id" db:"id"`
	Username        string    `json:"username" db:"user_name"`
	Firstname       string    `json:"firstname" db:"first_name"`
	Lastname        string    `json:"lastname" db:"last_name"`
	Email           string    `json:"email" db:"email"`
	AdminComment    string    `json:"adminComment" db:"admin_comment"`
	Active          bool      `json:"active" db:"active"`
	Verified        bool      `json:"verified" db:"verified"`
	Deleted         bool      `json:"deleted" db:"deleted"`
	IsSystemAccount bool      `json:"isSystemAccount" db:"is_system_account"`
	LastIp          string    `json:"lastIp" db:"last_ip"`
	Created         time.Time `json:"created" db:"created"`
	CellPhone       string    `json:"cellPhone" db:"cell_phone"`
	Officephone     string    `json:"officePhone" db:"office_phone"`
	Fax             string    `json:"fax" db:"fax"`
	Country         string    `json:"country" db:"country"`
	City            string    `json:"city" db:"city"`
	Postcode        string    `json:"postCode" db:"post_code"`
}
type LoginAccount struct {
	Id           int    `json:"id" db:"id"`
	Username     string `json:"userName" db:"user_name" binding:"required"`
	Password     string `json:"password" db:"password"`
	PasswordSalt string `json:"passwordSalt" db:"password_salt"`
}
