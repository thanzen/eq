var React = require('react/addons');
var ListGroupItem = require('react-bootstrap/lib/ListGroupItem');
var classNames = require("classnames");

function getClasses(role, currentSelected) {
    return classNames({
        active: role.id === currentSelected.id
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
            <ListGroupItem className={getClasses(this.props.role,this.props.selected)} onClick={this.handleClick}
                           onDoubleClick={this.handleDoubleClick}>
                {this.props.children}
                {this.props.role.name + '   ' + this.props.role.description}
            </ListGroupItem>
        );
    },

    handleClick: function () {
        this.props.onClick(this);
    },

    handleDoubleClick: function () {
        this.props.onDoubleClick(this);
    }
});

module.exports = RoleListItem;
