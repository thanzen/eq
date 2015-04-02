var React = require('react/addons');
var Modal = require('react-bootstrap/lib/Modal');
var ModalTrigger = require('react-bootstrap/lib/ModalTrigger');
var Button = require('react-bootstrap/lib/Button');
var Role = require('../../models/role').Role;
var store = require("../../stores/roleStore");
var action = require("../../actions/adminActions");

var RoleForm = React.createClass({
    getInitialState: function () {
        return {role: new Role()};
    },

    componentDidMount: function () {
        var role = store.RoleStoreInstance.getRole(this.props.id)
        if (role){
          role =JSON.parse(JSON.stringify(role));
        }
        this.setState({role: role})
    },

    componentWillUnmount: function () {
    },

    handleChange: function (event) {
        this.state.role.name = event.target.value;
        this.setState({role: this.state.role});
    },

    handleOk:function(){
      //todo: add data write back once it successful or maybe action need ed here.
      this.props.onRequestHide();
    },

    handleClose:function(){
      this.props.onRequestHide();
      action.roleResetAll();
    },

    render: function () {
        return (
            <Modal {...this.props} bsStyle='primary' title='Modal heading' animation={false}>
                <div className='modal-body'>
                    <h4>{this.state.role.name}</h4>
                    <input input type="text" value={this.state.role.name} onChange={this.handleChange}/>
                    <p>Duis mollis, est non commodo luctus, nisi erat porttitor ligula.</p>
                </div>
                <div className='modal-footer'>
                    <Button onClick={this.handleOk}>Ok</Button>
                    <Button onClick={this.handleClose}>Close</Button>
                </div>
            </Modal>
        );
    }
});
module.exports = RoleForm;
