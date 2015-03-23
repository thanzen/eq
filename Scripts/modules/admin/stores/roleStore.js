﻿var __extends = this.__extends || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    __.prototype = b.prototype;
    d.prototype = new __();
};
///<reference path="../../../libs/flux.d.ts" />
var events = require("../../../events/events");
var disp = require("../../../dispatcher");

var at = require("../eventType");
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;

var RoleStore = (function (_super) {
    __extends(RoleStore, _super);
    function RoleStore(roles) {
        if (typeof roles === "undefined") { roles = []; }
        _super.call(this);
        this.roles = roles;
        this.dispatchToken = this.registerEvents();
    }
    RoleStore.prototype.getAll = function () {
        return this.roles;
    };

    RoleStore.prototype.receiveAll = function (roles) {
        this.roles.concat(roles);
    };

    RoleStore.prototype.registerEvents = function () {
        var _this = this;
        return dispatcher.register(function (action) {
            switch (action.type) {
                case EventType.ROLES_RECEVIVE_ALL:
                    _this.receiveAll(action.roles);
                    _this.emit("change");
                    break;
                case EventType.ROLES_RECEVIVE_CREATE:
                    _this.roles.push(action.role);
                    _this.emit("change");
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