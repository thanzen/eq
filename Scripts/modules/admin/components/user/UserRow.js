var React = require('react/addons');
var classNames = require("classnames");

function getClasses(role, currentSelected) {
    return classNames({
    })
};

var UserRole = React.createClass({
    componentDidMount: function () {

    },

    componentWillUnmount: function () {

    },

    render: function () {
        return (
            <tr>
                <td>{this.props.user.firstname}</td>
                <td>{this.props.user.lastname}</td>
                <td>{this.props.user.username}</td>
                <td>{this.props.user.email}</td>
                <td>{this.props.user.company}</td>
            </tr>
        );
    }
});

module.exports = UserRole;
