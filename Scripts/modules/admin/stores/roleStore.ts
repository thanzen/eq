///<reference path="../../../libs/flux.d.ts" />
import events = require("../../../events/events");
import disp = require("../../../dispatcher");
import role = require("../models/role");
import at = require("../eventType");
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;

export class RoleStore extends events.EventEmitter {
    constructor(private roles: role.Role[]= []) {
        super();
        this.dispatchToken = this.registerEvents();
    }

    getAll(): role.Role[] {
        return this.roles;
    }

    private receiveAll(roles: role.Role[]) {
        this.roles = [];
        this.roles = this.roles.concat(roles);
    }

    private registerEvents(): string {
        return dispatcher.register((action: any) => {
            //dispatcher.waitFor();
            switch (action.type) {
                case EventType.ROLES_RECEVIVE_ALL:
                    this.receiveAll(action.roles);
                    this.emit("change");
                    break;
                case EventType.ROLES_RECEVIVE_CREATE:
                    this.roles.push(action.role);
                    this.emit("change");
                    break;
                default:
                    break;
            }
        });
    }
}
export var RoleStoreInstance = new RoleStore();
