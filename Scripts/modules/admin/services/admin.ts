///<reference path="../../../libs/bluebird.d.ts" />
///<reference path="../../../libs/jquery.d.ts" />
import rs = require("../stores/roleStore");
import role = require("../models/role");
import perm = require("../models/permission");
export class AdminService {
    static getAllRoles(): Promise<role.Role[]>{
        return Promise.resolve<role.Role[]>($.get("/api/role"));
    }
}