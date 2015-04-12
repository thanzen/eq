package user

type Permission struct {
    Id       int64   `json:"id,string" orm:"column(id);auto;pk"`
    Name     string  `json:"name" orm:"column(name)"`
    Category string  `json:"category" orm:"column(category)"`
    Roles    []*Role `json:"-" orm:"reverse(many)"`
}

func (u *Permission) TableName() string {
    return "permission"
}
