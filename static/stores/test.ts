///<reference path="../libs/flux.d.ts" />
import events = require("../events/events");
import disp = require("../dispatcher");
var dispatcher = disp.Dispatcher;
class TodoStore extends events.EventEmitter {
    constructor() {
        super();
        this.dispatchToken = this.registerEvents();
    }

    getAll(): string[] {
        return ["a", "b", "c"];
    }

    reSet(): string[] {
        return ["e", "f", "g"];
    }

    private registerEvents(): string {
        return dispatcher.register((action: any)=>{
            //dispatcher.waitFor();
            switch (action.type) {
                case "store_change":
                    alert("change");
                    this.emit("change");
                    break;
                default:
                    alert("default");
                    break;
            }
        });
    }
}