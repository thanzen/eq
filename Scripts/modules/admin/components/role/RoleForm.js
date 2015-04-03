var React = require('react/addons');
var Modal = require('react-bootstrap/lib/Modal');
var OverlayMixin = require('react-bootstrap/lib/OverlayMixin');
var ModalTrigger = require('react-bootstrap/lib/ModalTrigger');
var Button = require('react-bootstrap/lib/Button');
var Role = require('../../models/role').Role;
var store = require("../../stores/roleStore");
var action = require("../../actions/adminActions");
var dispatcher = require("../../../../dispatcher").Dispatcher;
var EventType = require("../../eventType").EventType;

var RoleForm = React.createClass({
    mixins: [OverlayMixin,React.addons.LinkedStateMixin],

    getInitialState: function () {
        return {role: new Role(), isModalOpen: false};
    },

    componentDidMount: function () {
        this.registerEvents();
    },

    componentWillUnmount: function () {
    },

    update:function(id){
      var role = store.RoleStoreInstance.getRole(id)
      if (role) {
          role = JSON.parse(JSON.stringify(role));
      }
      this.setState({role: role})
    },

    handleToggle: function () {
        this.setState({
            isModalOpen: !this.state.isModalOpen
        });
    },

    handleNameChange: function (event) {
        this.state.role.name = event.target.value;
        this.setState({role: this.state.role});
    },

    handleDescriptionChange: function (event) {
        this.state.role.description = event.target.value;
        this.setState({role: this.state.role});
    },

    handleOk: function () {
        //todo: add data write back once it successful or maybe action need ed here.
        this.handleToggle();
    },

    handleClose: function () {
        this.handleToggle();
    },

    registerEvents: function () {
      var self = this;
      dispatcher.register(function (action) {
          //dispatcher.waitFor();
          switch (action.type) {
              case EventType.UI_OPEN_ROLE_FORM:
                  self.update(action.id);
                  self.handleToggle();
                  break;
              default:
                  break;
          }
      });
    },

    render: function () {
        return null;
    },

    renderOverlay: function () {
        if (!this.state.isModalOpen) {
            return <span/>;
        }
        return (
            <Modal {...this.props} bsStyle='primary' title={this.state.role.id > 0?'Edit role':'New roles'} animation={true} onRequestHide={this.handleToggle}>
                <div className='modal-body'>
                    {'Name:         '}<input input type="text" value={this.state.role.name} onChange={this.handleNameChange}/>
                    <br/>
                    <br/>
                    {'Description:  '}<input input type="text" value={this.state.role.description} onChange={this.handleDescriptionChange}/>
                </div>
                <div className='modal-footer'>
                    <Button onClick={this.handleOk}>Ok</Button>
                    <Button onClick={this.handleClose}>Close</Button>
                </div>
            </Modal>
        )
    }

});
module.exports = RoleForm;
