var React = require('react/addons');
//var action = require("../../actions/adminActions");
var roleStore = require("../../stores/roleStore");
var userStore = require("../../stores/userStore");
var RoleListItem = require('./RoleListItem');
var Button = require('react-bootstrap/lib/Button');
var ListGroup = require('react-bootstrap/lib/ListGroup');
var Grid = require('react-bootstrap/lib/Grid');
var Row = require('react-bootstrap/lib/Row');
var Col = require('react-bootstrap/lib/Col');
var TabbedArea = require('react-bootstrap/lib/TabbedArea');
var TabPane = require('react-bootstrap/lib/TabPane');
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
    var maxPage = util.calculatePage(setting.TableLimit, userStore.UserStoreInstance.getTotalByRoleId(0))
    var userTableState = {currentPage: 1, maxPage: maxPage};

    var users = userStore.UserStoreInstance.getListByRoleId(first.id);
    return {roles: rs, selected: first, users: users, tableSetting: userTableState};
}

function prepareAllRoles() {
    var role = new Role();
    role.name = "[ALL Roles]";
    role.id = "0";
    role.description = "";
    var roles = roleStore.RoleStoreInstance.getAll();
    roles.splice(0, 0, role);
    return roles;
}


var RoleComposer = React.createClass({

    prevSearchText: "",

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
        if (roleListItem.props.role.id !== this.state.selected.id) {
            this.setState({selected: roleListItem.props.role});
            this.search(this.prevSearchText, 1, roleListItem.props.role.id);
        }
    },

    handleSearchClick: function (text) {
        if (this.prevSearchText == text) return;
        var self = this;
        this.prevSearchText = text;
        this.search(text, 1, this.state.selected.id);
    },
    search: function (query, currentPage, roleId) {
        var self = this;
        action.userGetList({
            query: query,
            roleId: roleId,
            offset: (currentPage - 1) * setting.TableLimit,
            limit: setting.TableLimit,
            includeTotal: true
        }).then(function () {
            var maxPage = util.calculatePage(setting.TableLimit, userStore.UserStoreInstance.getTotalByRoleId(roleId));
            self.setState({tableSetting: {currentPage: currentPage, maxPage: maxPage}});
        });
    },

    handleDoubleClick: function (roleListItem) {
        this.handleClick(roleListItem);
        dispatcher.dispatch({type: EventType.UI_OPEN_ROLE_FORM, id: this.state.selected.id});
    },

    handlePrevClick: function () {
        this.search(this.prevSearchText, this.state.tableSetting.currentPage - 1, this.state.selected.id);
    },

    handleNextClick: function () {
        this.search(this.prevSearchText, this.state.tableSetting.currentPage + 1, this.state.selected.id);
    },

    render: function () {
        var roles = this.state.roles.map(this.getRoleItem);
        return (
            <Grid>
                <Row className='show-grid'>
                    <Col md={4}>
                        <ListGroup>
                            {roles}
                        </ListGroup>
                        <Button onClick={this.btnClick}>{'Add'}</Button>
                    </Col>
                    <Col md={8}>
                        <TabbedArea defaultActiveKey={1}>
                            <TabPane eventKey={1} tab='Users'>
                                <Row className='show-grid'>
                                    <SearchBar onClick={this.handleSearchClick}/>
                                </Row>
                                <Row className='show-grid'>
                                    <UserTable users={this.state.users}
                                               currentPage={this.state.tableSetting.currentPage}
                                               maxPage={this.state.tableSetting.maxPage}
                                               onPrevClick={this.handlePrevClick} onNextClick={this.handleNextClick}/>
                                    <RoleForm/>
                                </Row>
                            </TabPane>

                            <TabPane eventKey={2} tab='Permissions'>TabPane 2 content</TabPane>
                        </TabbedArea>

                    </Col>
                </Row>

            </Grid>
        );
    }

});

module.exports = RoleComposer;
