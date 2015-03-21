///<reference path="../libs/flux.d.ts" />
import events = require("../events/events");
import disp = require("../dispatcher");
var dispatcher = disp.Dispatcher;
export interface TodoItem {
    name?: string;
    amount?: number;
}

export class TodoStore extends events.EventEmitter {
    constructor(private items:TodoItem[]=[]) {
        super();
        this.items = [{ name: "1" }, { name: "2" }, { name: "3" }]
        this.dispatchToken = this.registerEvents();
    }

    getAll(): TodoItem[]{
        return this.items;
    }

    reSet(): TodoItem[] {
        return [{ name: "4" }, { name: "6" }, {name:"5"}];
    }

    private registerEvents(): string {
        return dispatcher.register((action: any)=>{
            //dispatcher.waitFor();
            switch (action.type) {
                case "get_all":
                    this.emit("change");
                    break;
                case "reset":
                    this.emit("change");
                    break;
                default:
                    break;
            }
        });
    }
}
export var TodoStoreInstance = new TodoStore();
