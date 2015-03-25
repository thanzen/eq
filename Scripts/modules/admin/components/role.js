var React = require('react');
var action = require("../actions/adminActions")
var store = require("../stores/roleStore")


function getTodoItem(role) {
    return (
      <div>{role.name}</div>
    );
}

function getItems(){
    var rs = store.RoleStoreInstance.getAll();
    return {
        roles:rs
    }
}


var ChatApp = React.createClass({
    getInitialState: function() {
        return {roles:store.RoleStoreInstance.getAll()};
    },

    componentDidMount: function() {
        store.RoleStoreInstance.addListener("change",this._onChange);
    },

    componentWillUnmount: function() {
        store.RoleStoreInstance.removeListener("change",this._onChange);
    },

    render: function() {
        var todoItems = this.state.roles.map(getTodoItem);
        return (
          <div className="chatapp">
            nothing happen yet
             <textarea
        className="message-composer"
        name="message"
        value={this.state.name}
    />
    {todoItems}
    <button onClick={this.onclick}>receive</button>
    <button  onClick={this.btnClick}>reset</button>

  </div>
      );
    },
    _onChange:function(){
        //  this.setState({messages:store.TodoStoreInstance.getAll()});
    },
    btnClick:function(){
        this.setState( getItems());
    },
    onclick:function(){
        this.setState( getItems());
    }
});
module.exports  = ChatApp;