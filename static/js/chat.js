// 连接时触发
ws.onopen = function (e){
    // 验证token
    var data = {
        token : token,
    }
    ws.send(json(data))
}

//接收到消息时触发
ws.onmessage = function(e) {
    var scroll = true
    let data = JSON.parse(e.data)
    if(data.code === 500){
        $.removeCookie("token")
        alert(data.content)
        location.href = '/login'
    }
    // 收到消息时是否在聊天框底部
    if(TalkDom.scrollTop()+TalkDom.prop("clientHeight") !== TalkDom.prop("scrollHeight")){
        scroll = false
    }
    // 新增聊天框,聊天消息
    AddChatList(data.userId,data.content)
    if(data.userId === ChatUserId) {
        NewChat(data, false)
    }
    // 收到新消息后滚动到底部
    if(scroll) {
        ScrollToBottom()
    }
};
/*
连接关闭时触发
ws.onclose = function(e) {
    var data = {
        type:'close',
        token : token,
    }
    ws.send(json(data))
};
*/

// 发送消息
textarea = $("textarea")
textarea.on('keydown',function(e){
    if(e.key === "Enter"){
        if(textarea.val() === ""){
            alert("发送消息不能为空")
            return
        }
        var data = {
            type:'text',
            token : token,
            chat_id : ChatUserId,
            send_id : ChatList[ChatUserId].username,
            content : textarea.val(),
        }
        ws.send(json(data))
        // 添加自己发送的消息
        NewChat(data,true)
        // 更改左侧聊天列表的最后一条聊天记录
        AddChatList(ChatUserId,textarea.val())
        // 清空输入框
        textarea.val('')
        // 聊天框滚动至底部
        ScrollToBottom()
        return false
    }
});

// 转换成json字符串
function json(data){
    return JSON.stringify(data)
}