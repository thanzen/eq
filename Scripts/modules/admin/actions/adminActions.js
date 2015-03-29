var disp = require("../../../dispatcher");
var et = require("../eventType");
var admin = require("../services/admin");
var service = admin.AdminService;
var eventType = et.EventType;
var dispatcher = disp.Dispatcher;
exports.roleGetAll = function () {
    return service.getAllRoles().then(function (response) {
        dispatcher.dispatch({ type: eventType.ROLES_RECEVIVE_ALL, roles: response });
    });
};
