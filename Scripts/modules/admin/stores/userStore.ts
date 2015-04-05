///<reference path='../../../libs/definitions/flux.d.ts' />
///<reference path='../../../node_modules/immutable/dist/Immutable.d.ts'/>
import events = require("../../../events/events");
import disp = require("../../../dispatcher");
import user = require("../models/user");
import at = require("../eventType");
import Immutable = require('immutable');
var EventType = at.EventType;
var dispatcher = disp.Dispatcher;
export const ChangeEvent = "CHANGE";

export class UserStore extends events.EventEmitter {
    constructor(private users: Immutable.Map<number, user.User[]> = Immutable.Map<number, user.User[]>()) {
        super();
        this.dispatchToken = this.registerEvents();
    }

    getListByRoleId(roleId: number): user.User[] {
        if (!this.users.has(roleId))
            return [];
        return this.users.get(roleId);
    }

    getUser(userId, roleId: number): user.User {
        var ru = this.users.get(roleId);
        var users = ru.filter(function(user) {
            return user.id === userId;
        });
        return users.length > 0 ? users[0] : new user.User();
    }

    private registerEvents(): string {
        return dispatcher.register((action: any) => {
            //dispatcher.waitFor();
            switch (action.type) {
                case EventType.USER_GET_LIST:
                    this.users = this.users.set(action.roleId, action.users);
                    this.emit(ChangeEvent);
                    break;
                default:
                    break;
            }
        });
    }
}
export var UserStoreInstance = new UserStore();
