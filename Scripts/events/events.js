var EventEmitter = (function () {
    function EventEmitter() {
        this.registry = {};
    }
    EventEmitter.prototype.emit = function (name) {
        var args = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            args[_i - 1] = arguments[_i];
        }
        if (!this.registry[name])
            return;
        this.registry[name].forEach(function (x) {
            x.apply(null, args);
        });
    };
    EventEmitter.prototype.removeListener = function (name, fn) {
        if (!this.registry[name])
            return;
        this.registry[name] = this.registry[name].filter(function (f) {
            return f != fn;
        });
    };
    EventEmitter.prototype.addListener = function (name, fn) {
        if (!this.registry[name]) {
            this.registry[name] = [
                fn
            ];
        }
        else {
            this.registry[name].push(fn);
        }
    };
    return EventEmitter;
})();
exports.EventEmitter = EventEmitter;
var PubSub;
(function (PubSub) {
    var registry = {};
    PubSub.Pub = function (name) {
        var args = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            args[_i - 1] = arguments[_i];
        }
        if (!registry[name])
            return;
        registry[name].forEach(function (x) {
            x.apply(null, args);
        });
    };
    PubSub.Sub = function (name, fn) {
        if (!registry[name]) {
            registry[name] = [
                fn
            ];
        }
        else {
            registry[name].push(fn);
        }
    };
})(PubSub = exports.PubSub || (exports.PubSub = {}));
