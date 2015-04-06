package user

import "time"

type User struct {
	Id               int64     `json:"id,string" orm:"column(id);auto;pk"`
	Username         string    `json:"username" orm:"column(user_name)"`
	Firstname        string    `json:"firstname" orm:"column(first_name)"`
	Lastname         string    `json:"lastname" orm:"column(last_name)"`
	Email            string    `json:"email" orm:"column(email)"`
	Password         string    `json:"-" orm:"column(password)"`
	PasswordSalt     string    `json:"-" orm:"column(password_salt)"`
	AdminComment     string    `json:"adminComment" orm:"column(admin_comment)"`
	Active           bool      `json:"active" orm:"column(active);default(true)"`
	Deleted          bool      `json:"-" orm:"column(deleted)"`
	LastIp           string    `json:"-" orm:"column(last_ip)"`
	Created          time.Time `json:"-" orm:"column(created);auto_now_add"`
	LastLoginDate    time.Time `json:"-" orm:"column(last_login_date);auto_now"`
	LastActivityDate time.Time `json:"-" orm:"column(last_activity_date);auto_now"`
	CellPhone        string    `json:"cellPhone" orm:"column(cell_phone)"`
	Officephone      string    `json:"officePhone" orm:"column(office_phone)"`
	Fax              string    `json:"fax" orm:"column(fax)"`
	Country          string    `json:"country" orm:"column(country)"`
	City             string    `json:"city" orm:"column(city)"`
	Postcode         string    `json:"postCode" orm:"column(post_code)"`
	Lang             string    `json:"lang" orm:"column(lang);default(en-US)"`
	GravatarEmail    string    `json:"-" orm:"column(gravatar_email)"`
	Updated          time.Time `json:"-" orm:"column(updated);auto_now"`
	Company          string    `json:"company" orm:"company"`
	UserType         *UserType `json:"-" orm:"rel(fk);column(user_type_id);default(1)"`
	Roles            []*Role   `json:"-" orm:"rel(m2m);rel_table(user_role);on_delete(cascade)"`
	IsSystemAccount  bool      `json:"-" orm:"-"`
}

func (u *User) TableName() string {
	return "user_info"
}
