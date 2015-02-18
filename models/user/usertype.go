package user

type UserType struct {
	Id   int    `json:"id" orm:"column(id);auto;pk"`
	Name string `json:"name" orm:"column(name)"`
	User []*User  `orm:"reverse(many)"`
}

func (ut *UserType) TableName() string {
	return "user_type"
}
