var React = require('react/addons');
var ListGroupItem = require('react-bootstrap/lib/ListGroupItem');
var ListGroup = require('react-bootstrap/lib/ListGroup');

var cx = React.addons.classSet;
function getClasses(role, currentSelected) {
    return cx({
        'active': role.id === currentSelected.id
    })
};

var RoleListItem = React.createClass({
    getInitialState: function () {
        return {};
    },

    componentDidMount: function () {

    },

    componentWillUnmount: function () {

    },

    render: function () {
        return (
            <ListGroupItem className={getClasses(this.props.role,this.props.selected)} onClick={this.handleClick}>{this.props.role.name + '   ' + this.props.role.description}</ListGroupItem>
        );
    },

    handleClick: function () {
        this.props.onClick(this);
    }
});

module.exports = RoleListItem;
