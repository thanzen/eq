var AdminService = (function () {
    function AdminService() {
    }
    AdminService.getAllRoles = function () {
        return Promise.resolve($.get("something"));
    };
    return AdminService;
})();
