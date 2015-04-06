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
var user = require("../models/user");
var at = require("../eventType");
var Immutable = require('immutable');
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;
exports.ChangeEvent = "CHANGE";
var UserStore = (function (_super) {
    __extends(UserStore, _super);
    function UserStore(users) {
        if (users === void 0) { users = Immutable.Map(); }
        _super.call(this);
        this.users = users;
        this.dispatchToken = this.registerEvents();
    }
    UserStore.prototype.getListByRoleId = function (roleId) {
        if (!this.users.has(roleId))
            return [];
        return this.users.get(roleId);
    };
    UserStore.prototype.getUser = function (userId, roleId) {
        var ru = this.users.get(roleId);
        var users = ru.filter(function (user) {
            return user.id === userId;
        });
        return users.length > 0 ? users[0] : new user.User();
    };
    UserStore.prototype.registerEvents = function () {
        var _this = this;
        return dispatcher.register(function (action) {
            switch (action.type) {
                case EventType.USER_GET_LIST:
                    if (action.users == null) {
                        action.users = [];
                    }
                    _this.users = _this.users.set(action.roleId, action.users);
                    _this.emit(exports.ChangeEvent);
                    break;
                default:
                    break;
            }
        });
    };
    return UserStore;
})(events.EventEmitter);
exports.UserStore = UserStore;
exports.UserStoreInstance = new UserStore();
