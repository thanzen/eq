package user

type Role struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	IsSystemRole bool   `json:"isSystemRole" db:"is_system_role"`
	SystemName   string `json:"systemName" db:"system_name"`
}
