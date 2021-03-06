var __extends = this.__extends || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    __.prototype = b.prototype;
    d.prototype = new __();
};
///<reference path='../../../libs/definitions/flux.d.ts' />
///<reference path='../../../node_modules/immutable/dist/Immutable.d.ts'/>
var events = require("../../../events/events");
var disp = require("../../../dispatcher");
var role = require("../models/role");
var at = require("../eventType");
var Immutable = require('immutable');
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;
exports.ChangeEvent = "CHANGE";
var RoleStore = (function (_super) {
    __extends(RoleStore, _super);
    function RoleStore(roles) {
        if (roles === void 0) { roles = Immutable.List(); }
        _super.call(this);
        this.roles = roles;
        this.dispatchToken = this.registerEvents();
    }
    RoleStore.prototype.getAll = function () {
        return this.roles.toArray();
    };
    RoleStore.prototype.getRole = function (id) {
        var roles = this.roles.filter(function (role) {
            return role.id === id;
        });
        return roles.count() > 0 ? roles.first() : new role.Role();
    };
    RoleStore.prototype.receiveAll = function (roles) {
        this.roles = Immutable.List();
        this.roles = (_a = this.roles).push.apply(_a, roles);
        var _a;
    };
    RoleStore.prototype.registerEvents = function () {
        var _this = this;
        return dispatcher.register(function (action) {
            switch (action.type) {
                case EventType.ROLE_RECEVIVE_ALL:
                    _this.receiveAll(action.roles);
                    _this.emit(exports.ChangeEvent);
                    break;
                case EventType.ROLE_ADD:
                    _this.roles = _this.roles.push(action.role);
                    _this.emit(exports.ChangeEvent);
                    break;
                case EventType.ROlE_UPDATE:
                    _this.roles = _this.roles.withMutations(function (list) {
                        list.forEach(function (r) {
                            if (r.id == action.role.id) {
                                r.description = action.role.description;
                                r.name = action.role.name;
                            }
                        });
                    });
                    _this.emit(exports.ChangeEvent);
                    break;
                default:
                    break;
            }
        });
    };
    return RoleStore;
})(events.EventEmitter);
exports.RoleStore = RoleStore;
exports.RoleStoreInstance = new RoleStore();
