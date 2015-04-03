var React = require('react/addons');
var action = require("../../actions/adminActions");
var store = require("../../stores/roleStore");
var RoleListItem = require('./RoleListItem');
var Button = require('react-bootstrap/lib/Button');
var ListGroupItem = require('react-bootstrap/lib/ListGroupItem');
var ModalTrigger = require('react-bootstrap/lib/ModalTrigger');
var ListGroup = require('react-bootstrap/lib/ListGroup');
var RoleForm = require('./RoleForm');

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
            <RoleListItem key={role.id} role={role} selected={this.state.selected} onClick={this.handleClick}>
            <ModalTrigger modal={<RoleForm id={this.state.selected.id} />}>
                <Button bsStyle='primary' bsSize='small'>Edit</Button>
            </ModalTrigger>
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
                <ModalTrigger modal={<RoleForm id={0} />}>
                    <Button bsStyle='primary' bsSize='large'>Add</Button>
                </ModalTrigger>
            </div>
        );
    },
    _onChange: function () {
         this.setState({roles:store.RoleStoreInstance.getAll()});
    },
    btnClick: function () {
        this.setState(getItems());
    },
    handleClick: function (roleListItem) {
        this.setState({selected: roleListItem.props.role});
    },
    handleDoubleClick: function () {

    }
});
module.exports = RoleComposer;
