var ChatApp = require("./components/todo");
var store = require("./stores/todoStore");
var React = require('react');
window.React = React; // export for http://fb.me/react-devtools

//simulate initilization data stores
store.TodoStoreInstance.getAll();





React.render(
    <ChatApp />,
    document.getElementById('react')
);
