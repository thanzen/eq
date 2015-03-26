var ChatApp = require("./modules/admin/components/role");
var action = require("./modules/admin/actions/adminActions");
var React = require('react');
global.$=global.jQuery = require("jquery");
window.React = React; // export for http://fb.me/react-devtools

//simulate initilization data stores
action.roleGetAll();




React.render(
    <ChatApp />,
    document.getElementById('react')
);
