var Role = React.createClass({
    render: function () {
        return (
            <option value={this.props.id}>
               {this.props.name}
            </option>
        );
    }
});


var RoleComboList = React.createClass({
    getInitialState: function () {
        //  return {data:[{id:1,name:"Regular"},{id:2,name:"Admin"}]};
        return {data: []};
    },
    render: function () {
        this.props.data = [{id: 1, name: "Regular"}, {id: 2, name: "Admin"}];
        var roleNodes = this.props.data.map(function (role, index) {
            return (
                // `key` is a React-specific concept and is not mandatory for the
                // purpose of this tutorial. if you're curious, see more here:
                // http://facebook.github.io/react/docs/multiple-components.html#dynamic-children
                <Role id={role.id} name={role.name} key={index}>
                </Role>
            );
        });
        return (
            <select className="commentList">
    {roleNodes}
            </select>
        );
    }
});


/*    React.render(
 <RoleComboList> </RoleComboList>,
 document.getElementById('roles')
 );*/
