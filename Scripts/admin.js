var ChatApp = require("./modules/admin/components/role/RoleComposer");
var action = require("./modules/admin/actions/adminActions");
var React = require('react');


//simulate initilization data stores
action.roleGetAll().then(function(res){
    React.render(
        <ChatApp />,
        document.getElementById('react')
    );
});
