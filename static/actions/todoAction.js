﻿var disp = require("../dispatcher");
var dispatcher = disp.Dispatcher;
function AddTodo(item) {
    dispatcher.dispatch({ type: "store_change", b: "cc" });
    //var message = ChatMessageUtils.getCreatedMessageData(text, currentThreadID);
    //ChatWebAPIUtils.createMessage(message);
}
