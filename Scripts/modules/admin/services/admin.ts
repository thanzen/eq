///<reference path="../../../libs/definitions/bluebird.d.ts" />
///<reference path="../../../libs/definitions/jquery.d.ts" />
import rs = require("../stores/roleStore");
import role = require("../models/role");
import perm = require("../models/permission");
import $ = require("jquery");

export class AdminService {
    static getAllRoles(): Promise<role.Role[]>{
        return Promise.resolve<role.Role[]>($.get("/api/role"));
    }
}
