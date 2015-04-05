var React = require('react/addons');
var classNames = require("classnames");
var Button = require("react-bootstrap/lib/Button");
var Input = require("react-bootstrap/lib/Input");

function getClasses(role, currentSelected) {
    return classNames({
    })
};

var SearchBar = React.createClass({
    componentDidMount: function () {
    },

    componentWillUnmount: function () {
    },

    render: function () {
        return (
            <div>
              <Input type="text"></<Input>
              <Button>Search</Button>
            </div>
        );
    }
});

module.exports = SearchBar;
