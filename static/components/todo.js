
var React = require('react');
var action = require("../actions/todoAction")
var store = require("../stores/todoStore")


function getTodoItem(message) {
    return (
      <div>{message.name}</div>
    );
}

function getItems(){
    var mess = store.TodoStoreInstance.reSet();
    return {
        messages:mess
    }
}


var ChatApp = React.createClass({
    getInitialState: function() {
        return {messages:store.TodoStoreInstance.getAll()};
    },

    componentDidMount: function() {
        store.TodoStoreInstance.addListener("change",this._onChange);
    },

    componentWillUnmount: function() {
        store.TodoStoreInstance.removeListener("change",this._onChange);
    },

    render: function() {
        var todoItems = this.state.messages.map(getTodoItem);
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