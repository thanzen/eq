package user

type Role struct {
	Id           int64         `json:"id" orm:"column(id);pk;auto"`
	Name         string        `json:"name" orm:"column(name)"  valid:"Required;Match(/^\\S*$/)"`
	IsSystemRole bool          `json:"isSystemRole" orm:"column(is_system_role)"`
	Description  string        `json:"description" orm:"column(description)"`
	Users        []*User       `json:"-" orm:"reverse(many)"`
	Permissions  []*Permission `json:"-" orm:"rel(m2m);rel_table(role_permission)"`
}
