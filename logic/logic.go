package logic

import (
	"github.com/ExquisiteCore/LagrangeGo-Template/bot"
	"github.com/ExquisiteCore/LagrangeGo-Template/utils"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/event"
	"github.com/LagrangeDev/LagrangeGo/message"
)

// 定义不同类型事件的处理函数类型

type PrivateMessageHandler func(*client.QQClient, *message.PrivateMessage)
type GroupMessageHandler func(*client.QQClient, *message.GroupMessage)
type NewFriendRequestHandler func(*client.QQClient, *event.NewFriendRequest)

// HandlerManager 管理所有自定义逻辑

type HandlerManager struct {
	privateMessageHandlers []PrivateMessageHandler
	groupMessageHandlers   []GroupMessageHandler
}

// 全局 HandlerManager 实例

var Manager = &HandlerManager{}

// RegisterPrivateMessageHandler 注册私聊消息处理函数
func (hm *HandlerManager) RegisterPrivateMessageHandler(handler PrivateMessageHandler) {
	hm.privateMessageHandlers = append(hm.privateMessageHandlers, handler)
}

// RegisterGroupMessageHandler 注册群消息处理函数
func (hm *HandlerManager) RegisterGroupMessageHandler(handler GroupMessageHandler) {
	hm.groupMessageHandlers = append(hm.groupMessageHandlers, handler)
}

// SetupLogic 设置所有逻辑处理
func SetupLogic() {
	bot.QQClient.PrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
		for _, handler := range Manager.privateMessageHandlers {
			handler(client, event)
		}
	})

	bot.QQClient.GroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
		for _, handler := range Manager.groupMessageHandlers {
			handler(client, event)
		}
	})

	bot.QQClient.SelfGroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
		for _, handler := range Manager.groupMessageHandlers {
			handler(client, event)
		}
	})

	bot.QQClient.GroupNotifyEvent.Subscribe(func(client *client.QQClient, e event.INotifyEvent) {
		utils.Logger.Info("notify.group[gid:%v,content:%v]", e.From(), e.Content())
		switch e.(type) {
		case *event.GroupPokeEvent:
			pokeE := e.(*event.GroupPokeEvent)
			if pokeE.Sender == client.Uin || pokeE.Sender == 2412125282 || e.Receiver == client.Uin {
				return
			}
			_ = client.GroupPoke(pokeE.GroupUin, pokeE.Sender)

		}

	})

}
