import perm = require("./permission");
export class Role {
    id: number;
    name: string;
    isSystemRole: boolean;
    description: string;
    permissions: perm.Permission[];
}
