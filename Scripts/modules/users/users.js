var UserRow = React.createClass({
    render: function () {
        return (
            <div>
                <span>{this.props.user.firstName}</span> &nbsp;
                <span>{this.props.user.lastName}</span>
            </div>
        );
    }
});


var UserList = React.createClass({
    /*getInitialState: function () {
        //  return {data:[{id:1,name:"Regular"},{id:2,name:"Admin"}]};
        return {data: [{id: 1, firstName: "Billy", lastName: "Bob"}, {id: 2, firstName: "Joe", lastName: "James"}]};
    },*/
    componentWillMount: function(){
        console.log("componentWillMount")
    },
    render: function () {
        //this.props.data = [{id: 1, firstName: "Billy", lastName: "Bob"}, {id: 2, firstName: "Joe", lastName: "James"}];
        var userNodes = this.props.data.map(function (user, index) {
            return (
                // `key` is a React-specific concept and is not mandatory for the
                // purpose of this tutorial. if you're curious, see more here:
                // http://facebook.github.io/react/docs/multiple-components.html#dynamic-children
                <UserRow id={user.id} user={user} key={index}>
                </UserRow>
            );
        });
        return (
            <div className="commentList">
                {userNodes}
            </div>
        );
    }
});


/*    React.render(
 <RoleComboList> </RoleComboList>,
 document.getElementById('roles')
 );*/
