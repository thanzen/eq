


var Role = React.createClass({
    render: function(){
        return (
            <div>
                <div hidden>this.props.id</div>
                <div>{this.props.name}</div>
            </div>
            );
    }
});


var RoleComboList = React.createClass({
    getInitialState:function(){
        //  return {data:[{id:1,name:"Regular"},{id:2,name:"Admin"}]};
        return {data:[]};
    },
    render: function() {
        this.props.data = [{id:1,name:"Regular"},{id:2,name:"Admin"}];
        var roleNodes = this.props.data.map(function(role, index) {
            return (
              // `key` is a React-specific concept and is not mandatory for the
              // purpose of this tutorial. if you're curious, see more here:
              // http://facebook.github.io/react/docs/multiple-components.html#dynamic-children
               <Role id={role.id} name={role.name} key={index}>
              </Role>
      );
});
return (
  <div className="commentList">
    {roleNodes}
  </div>
    );
}
});




    React.render(
        <RoleComboList> </RoleComboList>,
        document.getElementById('roles')
        );
