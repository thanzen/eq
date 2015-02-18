package user

import (
    "github.com/astaxie/beego/orm"
    "github.com/thanzen/eq/models/user"
    . "github.com/thanzen/eq/services"
)

type PermissionService struct {
}

func (this PermissionService) Queryable() orm.QuerySeter {
    return orm.NewOrm().QueryTable("permission").OrderBy("-Id")
}
func (this PermissionService) Insert(p *user.Permission) error {
    if _, err := orm.NewOrm().Insert(p); err != nil {
        return err
    }
    return nil
}

func (this PermissionService) Read(p *user.Permission, fields ...string) error {
    if err := orm.NewOrm().Read(p, fields...); err != nil {
        return err
    }
    return nil
}

func (this PermissionService) Update(p *user.Permission, fields ...string) error {
    if _, err := orm.NewOrm().Update(p, fields...); err != nil {
        return err
    }
    return nil
}

func (this PermissionService) Delete(p *user.Permission) error {
    if _, err := orm.NewOrm().Delete(p); err != nil {
        return err
    }
    return nil
}

func (this PermissionService) Get(roles *[]*user.Permission, options SearchOptions) (int64, error) {
    qs := this.Queryable()
    cond := GenerateCondition(options)
    qs.SetCond(cond)
    return qs.All(roles)
}

func (this PermissionService) GetOne(p *user.Permission, options SearchOptions) error {
    qs := this.Queryable()
    cond := GenerateCondition(options)
    qs.SetCond(cond)
    return qs.One(p)
}






