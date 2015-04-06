var React = require('react/addons');
//var action = require("../../actions/adminActions");
var roleStore = require("../../stores/roleStore");
var userStore = require("../../stores/userStore");
var RoleListItem = require('./RoleListItem');
var Button = require('react-bootstrap/lib/Button');
var ListGroup = require('react-bootstrap/lib/ListGroup');
var RoleForm = require('./RoleForm');
var UserTable = require('../user/UserTable');
var dispatcher = require("../../../../dispatcher").Dispatcher;
var EventType = require("../../eventType").EventType;
var Role = require('../../models/role').Role;
var action = require("../../actions/adminActions");
var setting = require("../../../../setting");


//init data
function initData() {
    var rs = prepareAllRoles();
    var first;
    if (rs != null && rs.length > 0) {
        first = rs[0]; //fake role([All Roles])
    } else {
        first = {id: -1, name: '', description: ''}
    }
    var users = userStore.UserStoreInstance.getListByRoleId(0);
    return {roles: rs, selected: first, users: users};
}

function prepareAllRoles() {
    var role = new Role();
    role.name = "[ALL Roles]";
    role.id = 0;
    role.description = "";
    var roles = roleStore.RoleStoreInstance.getAll();
    roles.splice(0, 0, role);
    return roles;
}


var RoleComposer = React.createClass({
    getRoleItem: function (role) {
        return (
            <RoleListItem key={role.id} role={role} selected={this.state.selected} onClick={this.handleClick}
                          onDoubleClick={this.handleDoubleClick}>
            </RoleListItem>
        )
    },

    getInitialState: function () {
        return initData();
    },

    componentDidMount: function () {
        roleStore.RoleStoreInstance.addListener(roleStore.ChangeEvent, this.onRoleStoreChange);
        userStore.UserStoreInstance.addListener(roleStore.ChangeEvent, this.onUserStoreChange);

    },

    componentWillUnmount: function () {
        userStore.RoleStoreInstance.removeListener(roleStore.ChangeEvent, this.onRoleStoreChange);
        userStore.UserStoreInstance.removeListener(userStore.ChangeEvent, this.onUserStoreChange);
    },

    onRoleStoreChange: function () {

        this.setState({roleStore: prepareAllRoles()});
    },

    onUserStoreChange: function () {
        var users = userStore.UserStoreInstance.getListByRoleId(this.state.selected.id);
        this.setState({users: users});
    },

    btnClick: function () {
        dispatcher.dispatch({type: EventType.UI_OPEN_ROLE_FORM, id: 0});
    },

    handleClick: function (roleListItem) {
        if (roleListItem.id !== this.state.selected.id) {
            action.userGetList({
                query: "",
                roleId: roleListItem.props.role.id,
                offset: 0,
                limit: setting.TableLimit,
                includeTotal: true
            });
        }
        this.setState({selected: roleListItem.props.role});
    },

    handleDoubleClick: function (roleListItem) {
        this.handleClick(roleListItem);
        dispatcher.dispatch({type: EventType.UI_OPEN_ROLE_FORM, id: this.state.selected.id});
    },

    render: function () {
        var roles = this.state.roles.map(this.getRoleItem);
        return (
            <div>
                <ListGroup>
                    {roles}
                </ListGroup>
                <Button onClick={this.btnClick}>{'Add'}</Button>
                <UserTable users={this.state.users}/>
                <RoleForm/>
            </div>
        );
    }

});

module.exports = RoleComposer;
