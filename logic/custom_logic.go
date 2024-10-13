package logic

import (
	"github.com/ExquisiteCore/LagrangeGo-Template/utils"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/entity"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/sirupsen/logrus"
	"strings"
)

// RegisterCustomLogic 注册所有自定义逻辑
func RegisterCustomLogic() {
	//注册私聊消息处理逻辑

	Manager.RegisterPrivateMessageHandler(func(client *client.QQClient, event *message.PrivateMessage) {
		//client.SendPrivateMessage(event.Sender.Uin, []message.IMessageElement{message.NewText("Hello World!")})
		utils.Logger.Info("message.private[uid:%v,msg:%v]", event.Sender.Uin, event.ToString())
	})

	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
		utils.Logger.Info("message.group[gid:%v,uid:%v,msg:%v]", event.GroupUin, event.Sender.Uin, event.ToString())
	})

	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
		msg := event.ToString()
		if strings.HasPrefix(msg, "sgst") {
			memberInfo := client.GetCachedMemberInfo(client.Uin, event.GroupUin)
			if memberInfo.Permission == entity.Owner {
				title := strings.TrimSpace(strings.TrimPrefix(msg, "sgst"))
				_ = client.GroupSetSpecialTitle(event.GroupUin, event.Sender.Uin, title)
			}
		} else if strings.HasPrefix(msg, "/test") {
			logrus.Info("test")
			node := &message.ForwardNode{Message: []message.IMessageElement{message.NewText("测试```java\nSystem.out.println()\n```"), message.NewText("测试")}}
			_, _ = client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewNodeFoward([]*message.ForwardNode{node})})
		}
	})

}

//
