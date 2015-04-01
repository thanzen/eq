///<reference path="../../../libs/flux.d.ts" />
///<reference path='../../../node_modules/immutable/dist/Immutable.d.ts'/>
import events = require("../../../events/events");
import disp = require("../../../dispatcher");
import role = require("../models/role");
import at = require("../eventType");
import Immutable = require('immutable');
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;

export class RoleStore extends events.EventEmitter {
    constructor(private roles: role.Role[] = []) {
        super();
        this.dispatchToken = this.registerEvents();
    }

    getAll(): role.Role[] {
        return this.roles;
    }

    getRole(id: number): role.Role {
        var roles: role.Role[] = this.roles.filter(function(role) {
            return role.id === id;
        });
        return roles.length > 0 ? roles[0] : new role.Role();
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
