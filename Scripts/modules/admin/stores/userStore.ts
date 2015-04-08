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

export interface UserApiResponse{
  users:user.User[];
  total:number;
}

export class UserStore extends events.EventEmitter {
    constructor(private users: Immutable.Map<number, UserApiResponse> = Immutable.Map<number, UserApiResponse>()) {
        super();
        this.dispatchToken = this.registerEvents();
    }

    getListByRoleId(roleId: number): user.User[] {
        if (!this.users.has(roleId))
            return [];
        return this.users.get(roleId).users;
    }

    getTotalByRoleId(roleId: number): number {
        if (!this.users.has(roleId))
            return 0;
        return this.users.get(roleId).total;
    }

    getUser(userId, roleId: number): user.User {
        var ru = this.users.get(roleId).users;
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
                    if (action.users == null) {
                        action.users = [];
                        action.total = 0;
                    }
                    this.users = this.users.set(action.roleId, {users:action.users,total:action.total});
                    this.emit(ChangeEvent);
                    break;
                default:
                    break;
            }
        });
    }
}
export var UserStoreInstance = new UserStore();
