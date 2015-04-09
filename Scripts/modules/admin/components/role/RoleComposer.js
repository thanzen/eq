var React = require('react/addons');
//var action = require("../../actions/adminActions");
var roleStore = require("../../stores/roleStore");
var userStore = require("../../stores/userStore");
var RoleListItem = require('./RoleListItem');
var Button = require('react-bootstrap/lib/Button');
var ListGroup = require('react-bootstrap/lib/ListGroup');
var RoleForm = require('./RoleForm');
var UserTable = require('../user/UserTable');
var dispatcher = require("../../../../dispatcher").Dispatcher;
var EventType = require("../../eventType").EventType;
var Role = require('../../models/role').Role;
var action = require("../../actions/adminActions");
var setting = require("../../../../setting");
var util = require("../../../../libs/util");
var SearchBar = require("./SearchBar");

//init data
function initData() {
    var rs = prepareAllRoles();
    var first;
    if (rs != null && rs.length > 0) {
        first = rs[0]; //fake role([All Roles])
    } else {
        first = {id: -1, name: '', description: ''}
    }
    var maxPage = util.calculatePage(setting.TableLimit,userStore.UserStoreInstance.getTotalByRoleId(0))
    var userTableState = {currentPage:1,maxPage:maxPage};

    var users = userStore.UserStoreInstance.getListByRoleId(0);
    return {roles: rs, selected: first, users: users,tableSetting:userTableState};
}

function prepareAllRoles() {
    var role = new Role();
    role.name = "[ALL Roles]";
    role.id = 0;
    role.description = "";
    var roles = roleStore.RoleStoreInstance.getAll();
    roles.splice(0, 0, role);
    return roles;
}


var RoleComposer = React.createClass({

    prevSearchText:"",

    getRoleItem: function (role) {
        return (
            <RoleListItem key={role.id} role={role} selected={this.state.selected} onClick={this.handleClick}
                          onDoubleClick={this.handleDoubleClick}>
            </RoleListItem>
        )
    },

    getInitialState: function () {
        return initData();
    },

    componentDidMount: function () {
        roleStore.RoleStoreInstance.addListener(roleStore.ChangeEvent, this.onRoleStoreChange);
        userStore.UserStoreInstance.addListener(roleStore.ChangeEvent, this.onUserStoreChange);

    },

    componentWillUnmount: function () {
        userStore.RoleStoreInstance.removeListener(roleStore.ChangeEvent, this.onRoleStoreChange);
        userStore.UserStoreInstance.removeListener(userStore.ChangeEvent, this.onUserStoreChange);
    },

    onRoleStoreChange: function () {

        this.setState({roleStore: prepareAllRoles()});
    },

    onUserStoreChange: function () {
        var users = userStore.UserStoreInstance.getListByRoleId(this.state.selected.id);
        this.setState({users: users});
    },

    btnClick: function () {
        dispatcher.dispatch({type: EventType.UI_OPEN_ROLE_FORM, id: 0});
    },

    handleClick: function (roleListItem) {
        if (roleListItem.id !== this.state.selected.id) {
            action.userGetList({
                query: "",
                roleId: roleListItem.props.role.id,
                offset: 0,
                limit: setting.TableLimit,
                includeTotal: true
            });
        }
        this.setState({selected: roleListItem.props.role});
    },

    handleSearchClick:function(text){
      if(this.prevSearchText == text) return;
      var self = this;
      this.prevSearchText = text;
      action.userGetList({
          query: text,
          roleId: this.state.selected.id,
          offset: 0,
          limit: setting.TableLimit,
          includeTotal: true
      }).then(function(){
        var maxPage = util.calculatePage(setting.TableLimit,userStore.UserStoreInstance.getTotalByRoleId(self.state.selected.id));
        self.setState({tableSetting:{currentPage:1,maxPage:maxPage}});
      });
    },

    handleDoubleClick: function (roleListItem) {
        this.handleClick(roleListItem);
        dispatcher.dispatch({type: EventType.UI_OPEN_ROLE_FORM, id: this.state.selected.id});
    },

    handlePrevClick:function(){
      var currentPage=this.state.tableSetting.currentPage;
      var maxPage =this.state.tableSetting.maxPage;
      if(currentPage > 1){
        this.setState({tableSetting:{currentPage:currentPage-1,maxPage:maxPage}});
        action.userGetList({
            query: "",
            roleId: this.state.selected.id,
            offset: (currentPage-2)*setting.TableLimit,
            limit: setting.TableLimit,
            includeTotal: true
        });
      }
    },

    handleNextClick:function(){
      var currentPage=this.state.tableSetting.currentPage;
      var maxPage =this.state.tableSetting.maxPage;
      if(currentPage < maxPage){
        this.setState({tableSetting:{currentPage:currentPage+1,maxPage:maxPage}});
        action.userGetList({
            query: "",
            roleId: this.state.selected.id,
            offset: currentPage*setting.TableLimit,
            limit: setting.TableLimit,
            includeTotal: true
        });
      }
    },

    render: function () {
        var roles = this.state.roles.map(this.getRoleItem);
        return (
            <div>
                <ListGroup>
                    {roles}
                </ListGroup>
                <Button onClick={this.btnClick}>{'Add'}</Button>
                <UserTable users={this.state.users} currentPage={this.state.tableSetting.currentPage} maxPage={this.state.tableSetting.maxPage}
                 onPrevClick={this.handlePrevClick} onNextClick={this.handleNextClick}/>
                <RoleForm/>
                <SearchBar onClick={this.handleSearchClick}/>
            </div>
        );
    }

});

module.exports = RoleComposer;
