var React = require('react/addons');
var action = require("../../actions/adminActions")
var store = require("../../stores/roleStore")

var RoleListItem = require('./RoleListItem');

var Button = require('react-bootstrap/lib/Button');

var ListGroupItem = require('react-bootstrap/lib/ListGroupItem');

var ListGroup = require('react-bootstrap/lib/ListGroup');


//init data
function initData() {
    var rs = store.RoleStoreInstance.getAll();
    var first;
    if (rs != null && rs.length > 0) {
        first = rs[0];
    } else {
        first = {
            id: -1,
            name: '',
            'description': ''
        }
    }
    return {
        roles: rs,
        selected: first
    };
}

var RoleComposer = React.createClass({
    getRoleItem: function (role) {
        return (
            <RoleListItem role={role} selected={this.state.selected} onClick={this.handleClick}/>
        )
    },

    getInitialState: function () {
        return initData();
    },

    componentDidMount: function () {
        store.RoleStoreInstance.addListener("change", this._onChange);
    },

    componentWillUnmount: function () {
        store.RoleStoreInstance.removeListener("change", this._onChange);
    },

    render: function () {
        var roles = this.state.roles.map(this.getRoleItem);
        return (
            <ListGroup>
            {roles}
            </ListGroup>
        );
    },
    _onChange: function () {
        //  this.setState({messages:store.TodoStoreInstance.getAll()});
    },
    btnClick: function () {
        this.setState(getItems());
    },
    handleClick: function (roleListItem) {
        this.setState({selected: roleListItem.props.role});
    }
});
module.exports = RoleComposer;