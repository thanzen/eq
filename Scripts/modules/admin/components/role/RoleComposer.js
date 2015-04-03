var React = require('react/addons');
//var action = require("../../actions/adminActions");
var store = require("../../stores/roleStore");
var RoleListItem = require('./RoleListItem');
var Button = require('react-bootstrap/lib/Button');
var ListGroup = require('react-bootstrap/lib/ListGroup');
var RoleForm = require('./RoleForm');
var dispatcher = require("../../../../dispatcher").Dispatcher;
var EventType = require("../../eventType").EventType;

//init data
function initData() {
    var rs = store.RoleStoreInstance.getAll();
    var first;
    if (rs != null && rs.length > 0) {
        first = rs[0];
    } else {
        first = {id: -1, name: '', description: ''}
    }
    return {roles: rs, selected: first};
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
        store.RoleStoreInstance.addListener(store.ChangeEvent, this._onChange);
    },

    componentWillUnmount: function () {
        store.RoleStoreInstance.removeListener(store.ChangeEvent, this._onChange);
    },

    render: function () {
        var roles = this.state.roles.map(this.getRoleItem);
        return (
            <div>
                <ListGroup>
                  {roles}
                </ListGroup>
                <Button onClick={this.btnClick}>{'Test'}</Button>
                <RoleForm/>
            </div>
        );
    },

    _onChange: function () {
         this.setState({roles:store.RoleStoreInstance.getAll()});
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
    }
});

module.exports = RoleComposer;
