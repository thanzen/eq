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
    function TodoStore(items) {
        if (typeof items === "undefined") { items = []; }
        _super.call(this);
        this.items = items;
        this.items = [{ name: "1" }, { name: "2" }, { name: "3" }];
        this.dispatchToken = this.registerEvents();
    }
    TodoStore.prototype.getAll = function () {
        return this.items;
    };

    TodoStore.prototype.reSet = function () {
        return [{ name: "4" }, { name: "6" }, { name: "5" }];
    };

    TodoStore.prototype.registerEvents = function () {
        var _this = this;
        return dispatcher.register(function (action) {
            switch (action.type) {
                case "get_all":
                    _this.emit("change");
                    break;
                case "reset":
                    _this.emit("change");
                    break;
                default:
                    break;
            }
        });
    };
    return TodoStore;
})(events.EventEmitter);
exports.TodoStore = TodoStore;
exports.TodoStoreInstance = new TodoStore();
