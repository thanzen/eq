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

//init data
function initData() {
    var rs = roleStore.RoleStoreInstance.getAll();
    var first;
    if (rs != null && rs.length > 0) {
        first = rs[0];
    } else {
        first = {id: -1, name: '', description: ''}
    }
    var users = userStore.UserStoreInstance.getListByRoleId(2);
    return {roles: rs, selected: first, selectedRoleUsers:users};
}


var RoleComposer = React.createClass({
    getRoleItem: function (role) {
        return (
            <RoleListItem key={role.id} role={role} selected={this.state.selected} onClick={this.handleClick} onDoubleClick={this.handleDoubleClick}>
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
         this.setState({roleStore:store.RoleStoreInstance.getAll()});
    },

    onUserStoreChange: function () {
         this.setState({selectedRoleUsers:userStore.UserStoreInstance.getListByRoleId(this.state.selected.role.id)});
    },

    btnClick: function () {
        dispatcher.dispatch({ type: EventType.UI_OPEN_ROLE_FORM, id:0});
    },

    handleClick: function (roleListItem) {
        this.setState({selected: roleListItem.props.role});
    },

    handleDoubleClick: function (roleListItem) {
       this.handleClick(roleListItem);
       dispatcher.dispatch({ type: EventType.UI_OPEN_ROLE_FORM, id:this.state.selected.id});
    },

    render: function () {
        var roles = this.state.roles.map(this.getRoleItem);
        return (
            <div>
                <ListGroup>
                  {roles}
                </ListGroup>
                <Button onClick={this.btnClick}>{'Add'}</Button>
                <UserTable users={this.state.selectedRoleUsers}/>
                <RoleForm/>
            </div>
        );
    }

});

module.exports = RoleComposer;
