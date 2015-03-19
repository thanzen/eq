var EventEmitter = (function () {
    function EventEmitter() {
        this.registry = {};
    }
    EventEmitter.prototype.emit = function (name) {
        var args = [];
        for (var _i = 0; _i < (arguments.length - 1); _i++) {
            args[_i] = arguments[_i + 1];
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
            this.registry[name] = [fn];
        } else {
            this.registry[name].push(fn);
        }
    };
    return EventEmitter;
})();
exports.EventEmitter = EventEmitter;

(function (PubSub) {
    var registry = {};

    PubSub.Pub = function (name) {
        var args = [];
        for (var _i = 0; _i < (arguments.length - 1); _i++) {
            args[_i] = arguments[_i + 1];
        }
        if (!registry[name])
            return;
        registry[name].forEach(function (x) {
            x.apply(null, args);
        });
    };

    PubSub.Sub = function (name, fn) {
        if (!registry[name]) {
            registry[name] = [fn];
        } else {
            registry[name].push(fn);
        }
    };
})(exports.PubSub || (exports.PubSub = {}));
var PubSub = exports.PubSub;
