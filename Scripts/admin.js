var ChatApp = require("./modules/admin/components/role/RoleComposer");
var setting = require("./setting");
var action = require("./modules/admin/actions/adminActions");
var React = require('react');
var Promise = require("bluebird");

//simulate initilization data stores

Promise.props({
  roles:action.roleGetAll(),
  users:action.userGetList({query:"",roleId:"0",offset:0,limit:setting.TableLimit,includeTotal:true})
}).then(function(result){
  React.render(
      <ChatApp />,
      document.getElementById('react')
  );
});

//
// action.roleGetAll().then(function(res){
//
//     React.render(
//         <ChatApp />,
//         document.getElementById('react')
//     );
//     action.userGetList({query:"a",roleId:2,offset:0,limit:50});
// });
