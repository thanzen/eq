package user

type UserType struct {
	Id   int64    `json:"id,string" orm:"column(id);auto;pk"`
	Name string `json:"name" orm:"column(name)"`
	User []*User  `orm:"reverse(many)"`
}

func (ut *UserType) TableName() string {
	return "user_type"
}
