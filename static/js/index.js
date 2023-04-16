/*增加监听*/
var ws = new WebSocket("ws://127.0.0.1:8080/api/chatSocket")
// token是否存在,不存在则回退到登录页
var token = $.cookie("token")
if(token === undefined){
    location.href = '/login'
}
// 登录的用户信息
var arr = token.split(".")
var UserInfo = JSON.parse(atob(arr[1]))
// 好友列表
var FriendList
// 消息列表
var ChatList = []
// 当前选中的好友
var SelectUserId = 0
// 聊天好友
var ChatUserId = 0
// 翻页
var page = 0
// 好友列表dom
var FriendsDom = $(".list_area .bottom .item ul").eq(1)
// 聊天列表dom
var ChatListDom = $(".list_area .bottom .item ul").eq(0)
// 聊天框dom
var TalkDom = $(".talk_area ul")
// 好友信息dom
var FriendInfoDom = $(".user_info .bottom .info")
// 左下角设置dom
var configDom = $(".config")

// 更改为用户头像
$(".faceimg img").attr("src",UserInfo.image)
$('.sub_menu').on('click', 'a', function () {
    let index = $(this).index()
    $('.list_area .bottom .item').eq(index).show().siblings().hide()
    $(".user-box").children().eq(index).show().siblings().hide();
    $('.sub_menu a').eq(index).addClass('active').siblings().removeClass('active')
})

// 显示隐藏左下角设置
$(".icon-gengduo1").on('click',function (){
    if(configDom.css("display") === "none") {
        configDom.css("display", "block")
    }else{
        configDom.css("display", "none")
    }
})

// 获取好友列表信息
$.ajax({
    url:'/api/MyFriends',
    type:'post',
    async:false,
    success:function (e){
        if (e.code === 200) {
            FriendList = JSON.parse(e.data)
            for (let i = 0; i < FriendList.length; i++) {
                // 遍历好友列表,插入到dom中
                FriendsDom.append(`<li onclick="ChangeFriendsInfo(`+i+`)"><span class="img"><img src="` + FriendList[i].image + `"/></span><span class="title"><h3>` + FriendList[i].name + `</h3></span></li>`)
            }
        }
    }
})

// 点击发消息按钮
$(".btn").on('click',function (){
    // 切换到聊天界面
    $('.sub_menu a').eq(0).addClass('active')
    $('.sub_menu a').eq(1).removeClass('active')
    $(".user_info").hide()
    $(".user_area").show()
    $('.list_area .bottom .item').eq(0).show()
    $('.list_area .bottom .item').eq(1).hide()
    // 加入消息列表
    AddChatList(SelectUserId)
    // 调用更改聊天对象方法
    ChangeChatUser(SelectUserId)
})

// 更改选中好友的信息
function ChangeFriendsInfo(i) {
    SelectUserId = FriendList[i].id
    $(".user_info .bottom .default-info").hide()
    $(".user_info .bottom img").attr("src", FriendList[i].image)
    FriendInfoDom.find(".name").text(FriendList[i].name)
    FriendInfoDom.find(".wxid").text(FriendList[i].username)
    FriendInfoDom.find(".area").text(FriendList[i].area)
}

// 更改聊天对象
function ChangeChatUser(index){
    $(".user_area .default-info").hide()
    // 更改聊天框名字
    $(".user-box h2").text(ChatList["id"+index].name)
    // 重置翻页
    page = 1
    // 更改选中的聊天对象
    ChatUserId = index
    // 清空聊天框
    TalkDom.html("")
    // 获取聊天记录
    getChatHistory()
    // 聊天框滚动至底部
    ScrollToBottom()
    // 注册鼠标滚轮事件
    TalkDom.on('wheel',function (){
        if(TalkDom.scrollTop() === 0){
            page++
            getChatHistory()
        }
    })
}

// 获取聊天记录
function getChatHistory(){
    $.ajax({
        url : '/api/getChatHistory',
        type : 'post',
        async:false,
        data:{
            chatId : ChatUserId,
            page:page,
        },
        success : function (e) {
            if (e.code === 200) {
                var res = e.data
                // 没有更多数据了,关闭这个事件优化性能
                if(res.length < 20){
                    TalkDom.off('wheel')
                }
                // 从最后一个开始往前遍历,因为是加在数据最前面(prepend)
                for (let i = res.length-1; i >= 0; i--) {
                    let id = parseInt(res[i].split(':')[0])
                    let content = res[i].split(':')[1]
                    let data = {
                        userId: id,
                        content: content,
                    }
                    // 判断是否是用户发的消息
                    if (UserInfo.id === id) {
                        OldChat(data, true)
                    } else {
                        OldChat(data, false)
                    }
                }
            }
        }
    })
}


// 添加消息列表
function AddChatList(index,text = ''){
    // 第一次直接添加,后面不再重复添加
    if(ChatList["id"+index] === undefined) {
        let FriendInfo = GetFriendInfoById(index)
        // 使用index.toString()后数组key依旧是数字,导致有很多空数据,所以加id改成字符串
        ChatList["id"+index] = GetFriendInfoById(index)
        ChatListDom.append(`<li data-id="`+index+`" onclick="ChangeChatUser(`+index+`)"><span class="img"><img src="` + FriendInfo.image + `"/></span><span class="title"><div class="name"><h3>` + FriendInfo.name + `</h3><span class="date"></span></div><span class="small_txt"></span></span></li>`)
    }
    // 有文字则只更改文字消息,不新增
    if(text !== ''){
        ChatListDom.children("li").each(function (i,item){
            let userId = parseInt(item.getAttribute("data-id"))
            if(index === userId){
                item.querySelector(".date").innerText = getNowTime()
                item.querySelector(".small_txt").innerText = text
            }
        })
    }
}

// 将消息加入聊天框,加在最前面
function OldChat(data,me){
    if(me) {
        $(".talk_area ul").prepend(`<li class="right"><div class="content"><p>` + data.content + `</p></div><span class="img"><img src="` + UserInfo.image + `"/></span></li>`)
    }else{
        $(".talk_area ul").prepend(`<li class="left"><span class="img"><img src="` + ChatList["id"+data.userId].image + `"/></span><div class="content"><p>` + data.content + `</p></div></li>`)
    }
}

// 将消息加入聊天框,加在后面
function NewChat(data,me){
    if(me) {
        $(".talk_area ul").append(`<li class="right"><div class="content"><p>` + data.content + `</p></div><span class="img"><img src="` + UserInfo.image + `"/></span></li>`)
    }else{
        $(".talk_area ul").append(`<li class="left"><span class="img"><img src="` + ChatList["id"+data.userId].image + `"/></span><div class="content"><p>` + data.content + `</p></div></li>`)
    }
}

// 聊天框滚动至底部
function ScrollToBottom(){
    TalkDom.scrollTop(TalkDom.prop("scrollHeight"));
}

// 通过ID查找好友信息
function GetFriendInfoById(id){
    for (let i = 0;i < FriendList.length;i++){
        if(FriendList[i].id === id){
            return FriendList[i]
        }
    }
}

// 获取当前时分
function getNowTime(){
    let myDate = new Date();
    let h = myDate.getHours(); // 获取当前小时数(0-23)
    let m = myDate.getMinutes(); // 获取当前分钟数(0-59)
    h = h < 10?"0"+h:h
    m = m < 10?"0"+m:m
    return h+":"+m
}