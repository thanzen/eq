import disp = require("../../../dispatcher");
import et = require("../eventType");
import admin = require("../services/admin");
import role = require("../models/role");
var service = admin.AdminService;
var eventType = et.EventType;
var dispatcher = disp.Dispatcher;

export var roleGetAll = function() {
    return service.getAllRoles().then(
        (response: any) => {
            dispatcher.dispatch({ type: eventType.ROLE_RECEVIVE_ALL, roles: response });
        });
}

export var roleSave = function(role: role.Role) {
    return service.saveRole(role).then(
        (response: role.Role) => {
            if (role.id > 0) {
                dispatcher.dispatch({ type: eventType.ROlE_UPDATE, role: role });
            } else {
                role = response;
                dispatcher.dispatch({ type: eventType.ROLE_ADD, role: role });
            }
        }).catch(function(e) {
        console.log(e.statusText);
    })
}


export var userGetList = function(param: admin.UserListSearchParam) {
    return service.getUserList(param).then(
        (response: any) => {
            dispatcher.dispatch({ type: eventType.USER_GET_LIST, users: response.users, roleId: param.roleId });
        });
}
