var React = require('react/addons');
var action = require("../actions/adminActions")
var store = require("../stores/roleStore")
var ButtonGroup = require('react-bootstrap/lib/ButtonGroup');
var Button = require('react-bootstrap/lib/Button');
var DropdownButton = require('react-bootstrap/lib/DropdownButton');
var MenuItem = require('react-bootstrap/lib/MenuItem');
var Table = require('react-bootstrap/lib/Table');
var ListGroupItem = require('react-bootstrap/lib/ListGroupItem');

var ListGroup = require('react-bootstrap/lib/ListGroup');


function getTodoItem(role) {
    return (
      <div>{role.name}</div>
    );
}

var cx = React.addons.classSet;
var classes = cx({
    'success': true,
    'active': true
});


const tableInstance = (
  <Table striped bordered condensed hover>
    <thead>
      <tr success>
        <th></th>
        <th>First Name</th>
        <th>Last Name</th>
        <th>Username</th>
      </tr>
    </thead>
    <tbody>
      <tr className={classes}>
        <td>1</td>
        <td>Mark</td>
        <td>Otto</td>
        <td>@mdo</td>
      </tr>
      <tr className={classes}>
        <td>2</td>
        <td>Jacob</td>
        <td>Thornton</td>
        <td>@fat</td>
      </tr>
      <tr>
        <td>3</td>
        <td colSpan="2">Larry the Bird</td>
        <td>@twitter</td>
      </tr>
    </tbody>
  </Table>
);

var isFirst =true;
var isSecond = false;

const listgroupInstance = (
  <ListGroup>
    <ListGroupItem  className={classes}>Link 1</ListGroupItem>
    <ListGroupItem  disabled>Link 2</ListGroupItem>
    <ListGroupItem  disabled>Link 3</ListGroupItem>
  </ListGroup>
);




function getItems(){
    var rs = store.RoleStoreInstance.getAll();
    return {
        roles:rs
    }
}

const buttonGroupInstance = (
  <ButtonGroup>
    <Button>1</Button>
    <Button>2</Button>
    <DropdownButton title='Dropdown'>
      <MenuItem eventKey='1'>Dropdown link</MenuItem>
      <MenuItem eventKey='2'>Dropdown link</MenuItem>
    </DropdownButton>
  </ButtonGroup>
);

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
    {buttonGroupInstance}
    {tableInstance}
    {listgroupInstance}
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