var $ = require("jquery");
var AdminService = (function () {
    function AdminService() {
    }
    AdminService.getAllRoles = function () {
        return Promise.resolve($.get("/api/role"));
    };
    AdminService.saveRole = function (role) {
        return Promise.resolve($.ajax({
            type: role.id > 0 ? "PUT" : "POST",
            url: "/api/role",
            data: JSON.stringify(role),
            dataType: "json"
        }));
    };
    return AdminService;
})();
exports.AdminService = AdminService;
