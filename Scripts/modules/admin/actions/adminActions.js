var disp = require("../../../dispatcher");
var et = require("../eventType");
var admin = require("../services/admin");
var service = admin.AdminService;
var eventType = et.EventType;
var dispatcher = disp.Dispatcher;
exports.roleGetAll = function () {
    return service.getAllRoles().then(function (response) {
        dispatcher.dispatch({ type: eventType.ROLE_RECEVIVE_ALL, roles: response });
    });
};
exports.roleSave = function (role) {
    return service.saveRole(role).then(function (response) {
        if (role.id > 0) {
            dispatcher.dispatch({ type: eventType.ROlE_UPDATE, role: role });
        }
        else {
            role = response;
            dispatcher.dispatch({ type: eventType.ROLE_ADD, role: role });
        }
    }).catch(function (e) {
        console.log(e.statusText);
    });
};
