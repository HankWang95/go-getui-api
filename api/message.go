package api

import (
	"github.com/HankWang95/go-getui-api/push"
	"github.com/HankWang95/go-getui-api/tool"
)

type messageType string

var (
	MsgTypeNotification messageType = "notification" // 点开通知打开应用
	MsgTypeLink         messageType = "link"         // 点开通知打开网页
	MsgTypeNotypopload  messageType = "notypopload"  // 点击通知弹窗下载
	MsgTypeTransmission messageType = "transmission" // 透穿消息
)

func getDefaultMessage(msgType messageType) *tool.Message {
	msg := tool.GetMessage()
	msg.SetAppKey(appKey)
	if msgType != "" {
		msg.SetMsgType(string(msgType))
	}
	return msg
}

func GetUndefinedTypeMessage() *tool.Message {
	return getDefaultMessage("")
}

// style 通知栏消息布局样式, 必填
// transmissionContent : 透传内容，没有传 ""
// durationBegin, druationEnd 设定展示开始时间\结束时间
func GetPushSingleNotification(style interface{}, transmissionContent, durationBegin, durationEnd string) *push.PushSingleParmar {
	m := getDefaultMessage(MsgTypeNotification)
	n := tool.GetNotification()

	if transmissionContent != "" {
		n.SetTransmissionContent(transmissionContent)
	}
	if durationBegin != "" {
		n.SetDurationBegin(durationBegin)
	}
	if durationEnd != "" {
		n.SetDurationEnd(durationEnd)
	}
	n.SetNotifyStyle(style)
	p := &push.PushSingleParmar{
		Message:      m,
		Notification: n,
	}
	return p
}

