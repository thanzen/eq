///<reference path="../../../libs/bluebird.d.ts" />
///<reference path="../../../libs/jquery.d.ts" />
import role = require("../models/role");
import perm = require("../models/permission");
class AdminService {
    static getAllRoles():  Promise<role.Role[]>{
        return Promise.resolve<role.Role[]>($.get("something"));
    }
}