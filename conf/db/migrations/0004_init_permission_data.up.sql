--role permission
insert into permission(name,category) values("role_add_users","role");
insert into permission(name,category) values("role_delete_users","role");
insert into permission(name,category) values("role_add_permissions","role");
insert into permission(name,category) values("role_delete_permissions","role");
insert into permission(name,category) values("role_create","role");
insert into permission(name,category) values("role_update","role");
insert into permission(name,category) values("role_delete","role");
insert into permission(name,category) values("role_delete_users","role");
insert into permission(name,category) values("role_view_all","role");


--user permission
insert into permission(name,category) values("user_fuzzy_search","user");
insert into permission(name,category) values("user_admin_update","user");
insert into permission(name,category) values("user_admin_reset_password","user");












insert into role_permission(role_id, permission_id) select 1, id from permission;