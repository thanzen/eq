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
      var rows = this.props.users.map(function(user){
        return getRow(user);
      });
        return (
            <Table>
                <thead>
                <tr>
                    <th>{"First Name"}</th>
                    <th>{"Last Name"}</th>
                    <th>{"Username"}</th>
                    <th>{"Email"}</th>
                    <th>{"Company"}</th>
                </tr>
                </thead>
                {rows}
            </Table>
        );
    }
});

module.exports = UserTable;
