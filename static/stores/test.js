var __extends = this.__extends || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    __.prototype = b.prototype;
    d.prototype = new __();
};
///<reference path="../libs/flux.d.ts" />
var events = require("../events/events");
var disp = require("../dispatcher");
var dispatcher = disp.Dispatcher;
var TodoStore = (function (_super) {
    __extends(TodoStore, _super);
    function TodoStore() {
        _super.call(this);
        this.dispatchToken = this.registerEvents();
    }
    TodoStore.prototype.getAll = function () {
        return ["a", "b", "c"];
    };

    TodoStore.prototype.reSet = function () {
        return ["e", "f", "g"];
    };

    TodoStore.prototype.registerEvents = function () {
        var _this = this;
        return dispatcher.register(function (action) {
            switch (action.type) {
                case "store_change":
                    alert("change");
                    _this.emit("change");
                    break;
                default:
                    alert("default");
                    break;
            }
        });
    };
    return TodoStore;
})(events.EventEmitter);
