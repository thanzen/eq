var React = require('react/addons');
var classNames = require("classnames");
var Input = require("react-bootstrap/lib/Input");
var Button = require("react-bootstrap/lib/Button");

var SearchBar = React.createClass({
    getInitialState: function () {
        return {value: ''};
    },

    handleChange: function () {
        // This could also be done using ReactLink:
        // http://facebook.github.io/react/docs/two-way-binding-helpers.html
        this.setState({
            value: this.refs.input.getValue()
        });
    },

    handleClick: function () {
        this.props.onClick(this.state.value);
    },

    render: function () {
        return (
            <div>
                <Input
                    type='text'
                    value={this.state.value}
                    placeholder='Enter text'
                    label=''
                    help=''
                    hasFeedback
                    ref='input'
                    groupClassName='group-class'
                    wrapperClassName='wrapper-class'
                    labelClassName='label-class'
                    onChange={this.handleChange}/>
                <Button onClick={this.handleClick}>Search</Button>
            </div>);
    }
});

module.exports = SearchBar;
