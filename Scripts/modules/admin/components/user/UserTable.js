var React = require('react/addons');
var classNames = require("classnames");
var Table = require("react-bootstrap/lib/Table");
var UserRow = require("./UserRow");

function getRow(user){
  return <UserRow user={user}/>;
}

function getClasses(role, currentSelected) {
    return classNames({
        // active: role.id === currentSelected.id
    })
};

var UserTable = React.createClass({
    getInitialState: function () {
        return {};
    },

    componentDidMount: function () {

    },

    componentWillUnmount: function () {
    },

    render: function () {
        return (
            <Table>
                <thead>
                <tr>
                    <th>{"Username"}</th>
                    <th>{"First Name"}</th>
                    <th>{"Last Name"}</th>
                    <th>{"Email"}</th>
                    <th>{"Company"}</th>
                </tr>
                </thead>
            </Table>
        );
    }
});

module.exports = UserTable;
