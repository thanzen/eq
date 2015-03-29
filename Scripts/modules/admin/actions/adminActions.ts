import disp = require("../../../dispatcher");
import et = require("../eventType");
import admin = require("../services/admin");
var service = admin.AdminService;
var eventType = et.EventType;
var dispatcher = disp.Dispatcher;
export var roleGetAll = function () {
  return  service.getAllRoles().then(
        (response: any) => {
            dispatcher.dispatch({ type: eventType.ROLES_RECEVIVE_ALL, roles:response});
        });
}