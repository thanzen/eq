package user

import (
	"github.com/astaxie/beego/orm"
	"github.com/thanzen/eq/cachemanager"
	"github.com/thanzen/eq/models/user"
	. "github.com/thanzen/eq/services"
)

const ROLE_ALL_KEY = "eq_role_all"

type RoleService struct {
}

func (this RoleService) Queryable() orm.QuerySeter {
	return orm.NewOrm().QueryTable("role").OrderBy("-Id")
}
func (this RoleService) Insert(r *user.Role) error {
	if _, err := orm.NewOrm().Insert(r); err != nil {
		return err
	}
	return nil
}

func (this RoleService) Read(r *user.Role, fields ...string) error {
	if err := orm.NewOrm().Read(r, fields...); err != nil {
        return err
    }
	return nil
}

func (this RoleService) Update(r *user.Role, fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
    cachemanager.Delete(ROLE_ALL_KEY)
	return nil
}

func (this RoleService) Delete(r *user.Role) error {
	if _, err := orm.NewOrm().Delete(r); err != nil {
		return err
	}
    cachemanager.Delete(ROLE_ALL_KEY)
	return nil
}

func (this RoleService) Get(roles []*user.Role, options SearchOptions) (int64, error) {
	qs := this.Queryable()
	cond := GenerateCondition(options)
	qs.SetCond(cond)
	return qs.All(roles)
}

//todo add go routine to fetch permissions for  multiple roles
func (this RoleService) GetAllRoles() (roles map[int64]*user.Role) {
    //one week cache time
	temp, hit := cachemanager.GetWithExpireTime(ROLE_ALL_KEY, 604800, func(params ...interface{}) interface{} {
		var allRoles []*user.Role
		o := orm.NewOrm()
		_, err := o.QueryTable("Role").All(&allRoles)
		if err == nil {
			roles = make(map[int64]*user.Role)
			for _, r := range allRoles {
				o.LoadRelated(r, "Permissions")
				roles[r.Id] = r
			}
		}
		if err == nil {
			return roles
		}
		return nil
	})
	if temp != nil && hit {
		roles = temp.(map[int64]*user.Role)
	}
	return roles
}

func (this RoleService) HasPermission(role *user.Role, permission string) bool {
	for _, p := range role.Permissions {
		if p.Name == permission {
			return true
		}
	}
	return false
}

func (this RoleService) InsertUsers(r *user.Role, users ...*user.User) error {
	return this.InsertUsersWithScope(orm.NewOrm(), r, users...)
}

func (this RoleService) InsertUsersWithScope(tr orm.Ormer, r *user.Role, users ...*user.User) error {
	_, err := tr.QueryM2M(r, "Users").Add(users)
	return err
}

func (this RoleService) DeleteUsers(r *user.Role, users ...*user.User) error {
	_, err := orm.NewOrm().QueryM2M(r, "Users").Remove(users)
	return err
}

func (this RoleService) InsertPermissions(r *user.Role, ps ...*user.Permission) error {
	_, err := orm.NewOrm().QueryM2M(r, "Permissions").Add(ps)
    cachemanager.Delete(ROLE_ALL_KEY)
	return err
}

func (this RoleService) DeletePermissions(r *user.Role, ps ...*user.Permission) error {
	_, err := orm.NewOrm().QueryM2M(r, "Permissions").Remove(ps)
    cachemanager.Delete(ROLE_ALL_KEY)
	return err
}
