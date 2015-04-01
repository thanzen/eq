var jQuery = require("jquery");
var $ = jQuery;
var AdminService = (function () {
    function AdminService() {
    }
    AdminService.getAllRoles = function () {
        return Promise.resolve($.get("/api/role"));
    };
    return AdminService;
})();
exports.AdminService = AdminService;
