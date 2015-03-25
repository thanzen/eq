import disp = require("../dispatcher");
var dispatcher = disp.Dispatcher;
export var getAll = function () {   
    dispatcher.dispatch({ type:"get_all",b:"cc"});
    //var message = ChatMessageUtils.getCreatedMessageData(text, currentThreadID);
    //ChatWebAPIUtils.createMessage(message);
}
export var reSet = function () {
    dispatcher.dispatch({ type: "reset", b: "cc" });
}
