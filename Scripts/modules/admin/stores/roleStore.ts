///<reference path='../../../libs/definitions/flux.d.ts' />
///<reference path='../../../node_modules/immutable/dist/Immutable.d.ts'/>
import events = require("../../../events/events");
import disp = require("../../../dispatcher");
import role = require("../models/role");
import at = require("../eventType");
import Immutable = require('immutable');
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;
export const ChangeEvent = "CHANGE";
export class RoleStore extends events.EventEmitter {
    constructor(private roles: Immutable.List<role.Role> = Immutable.List<role.Role>()) {
        super();
        this.dispatchToken = this.registerEvents();
    }

    getAll(): role.Role[] {
        return this.roles.toArray();
    }

    getRole(id: number): role.Role {
        var roles: Immutable.Iterable<number,role.Role>  = this.roles.filter(function(role) {
            return role.id === id;
        });
        return roles.count() > 0 ? roles.first() : new role.Role();
    }

    private receiveAll(roles: role.Role[]) {
        this.roles = Immutable.List<role.Role>();
        this.roles = this.roles.push(...roles);
    }

    private registerEvents(): string {
        return dispatcher.register((action: any) => {
            //dispatcher.waitFor();
            switch (action.type) {
                case EventType.ROLES_RECEVIVE_ALL:
                    this.receiveAll(action.roles);
                    this.emit(ChangeEvent);
                    break;
                case EventType.ROLES_RECEVIVE_CREATE:
                    this.roles.push(action.role);
                    this.emit(ChangeEvent);
                    break;
                default:
                    break;
            }
        });
    }
}
export var RoleStoreInstance = new RoleStore();
