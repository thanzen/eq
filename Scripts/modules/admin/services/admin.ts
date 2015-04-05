///<reference path="../../../libs/definitions/bluebird.d.ts" />
///<reference path="../../../libs/definitions/jquery.d.ts" />
import rs = require("../stores/roleStore");
import role = require("../models/role");
import user = require("../models/user");
import perm = require("../models/permission");
import $ = require("jquery");

export interface UserListSearchParam {
    query: string;
    roleId: number;
    offset: number;
    limit: number;
}

export class AdminService {
    static getAllRoles(): Promise<role.Role[]> {
        return Promise.resolve<role.Role[]>($.get("/api/role"));
    }

    static saveRole(role: role.Role): Promise<role.Role> {
        return Promise.resolve<role.Role>($.ajax({
            type: role.id > 0 ? "PUT" : "POST",
            url: "/api/role",
            data: JSON.stringify(role),
            dataType: "json"
        }));
    }

    static getUserList(searchParm: UserListSearchParam): Promise<user.User> {
        var url = "/api/admin/user/getusers";
        return Promise.resolve<user.User>($.ajax({
            type: "POST",
            url: url,
            data: JSON.stringify(searchParm),
            dataType: "json"
        }));
    }
}
