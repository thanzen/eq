var React = require('react/addons');
var classNames = require("classnames");
var Table = require("react-bootstrap/lib/Table");
var Button = require("react-bootstrap/lib/Button");
var UserRow = require("./UserRow");


function getRow(user) {
    return <UserRow key={user.id} user={user}/>;
}
function getButtons(prev, next) {

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

    handlePrevClick: function () {
        this.props.onPrevClick();
    },

    handleNextClick: function () {
        this.props.onNextClick();
    },

    render: function () {
        var rows = null;
        if (this.props.users && this.props.users.length > 0) {
            rows = this.props.users.map(function (user) {
                return getRow(user);
            });
        }

        return (
            <div>
                <Table striped bordered condensed hover responsive>
                    <thead>
                    <tr>
                        <th>{"First Name"}</th>
                        <th>{"Last Name"}</th>
                        <th>{"Username"}</th>
                        <th>{"Email"}</th>
                        <th>{"Company"}</th>
                    </tr>
                    </thead>
                    <tbody>
                    {rows}
                    </tbody>
                </Table>
                {this.props.currentPage > 1 ? (
                    <Button bsStyle='link' onClick={this.handlePrevClick}>prev</Button>) : null}
                {this.props.currentPage < this.props.maxPage ? (
                    <Button bsStyle='link' onClick={this.handleNextClick}>next</Button>) : null}
            </div>

        );
    }
});

module.exports = UserTable;
