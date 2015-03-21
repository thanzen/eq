var disp = require("../dispatcher");
var dispatcher = disp.Dispatcher;
exports.getAll = function () {
    dispatcher.dispatch({ type: "get_all", b: "cc" });
    //var message = ChatMessageUtils.getCreatedMessageData(text, currentThreadID);
    //ChatWebAPIUtils.createMessage(message);
};
exports.reSet = function () {
    dispatcher.dispatch({ type: "reset", b: "cc" });
};
