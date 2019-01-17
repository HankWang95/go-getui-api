package api

import (
	"github.com/HankWang95/go-getui-api/push"
	"github.com/HankWang95/go-getui-api/style"
	"github.com/smartwalle/xid"
	"time"
)

// 单推，cid alias 二者只会生效一个，cid权重大
func PushSingle(cid, alias, requestId string, p *push.PushSingleParmar) (result *push.PushSingleResult, err error) {
	if getAppId() == "" {
		return nil, nil
	}

	p.RequestId = requestId
	p.Cid = cid
	p.Alias = alias

	// Push不成功，重试5次
	for redoTime := 0; redoTime < 5; redoTime++ {
		result, err := push.PushSingle(getAppId(), getToken(), p)
		if err != nil {
			return nil, err
		}
		redo, err := handlePushSingleResult(result)
		if err != nil {
			return nil, err
		}
		if redo {
			time.Sleep(300 * time.Microsecond)
			continue
		} else {
			return result, nil
		}
	}
	return nil, PushErr
}

//func PushList(cids, alias []string,  taskId string, p *push.PushListParmar) (result *push.PushListResult, err error) {
//	p.Cid = cids
//	p.Alias = alias
//	p.TaskId = taskId
//
//	// Push不成功，重试5次
//	for redoTime := 0; redoTime < 5; redoTime++ {
//		result, err := push.PushList(getAppId(), getToken(), p)
//		if err != nil {
//			return nil, err
//		}
//		redo, err := handlePushListResult(result)
//		if err != nil {
//			return nil, err
//		}
//		if redo {
//			time.Sleep(300 * time.Microsecond)
//			continue
//		} else {
//			return result, nil
//		}
//	}
//	return nil, PushErr
//}

func LazyPush(cid, title, content, transmission string) error {
	msgStyle := style.GetSystemStyle(content, title)
	var p *push.PushSingleParmar
	if transmission == "" {
		p = GetPushSingleNotification(msgStyle, "", "", "")
	} else {
		p = GetPushSingleMsgTypeTransmission(msgStyle, transmission, "", "")
	}
	_, err := PushSingle(cid, "", xid.NewXID().Hex(), p)
	if err != nil {
		return err
	}
	return nil
}
//
//func LazyPushList(cids []string, title, content string) error {
//	msgStyle := style.GetSystemStyle(content, title)
//	p := GetPushSingleNotification(msgStyle, "", "", "")
//	_, err := PushList(cids, nil, xid.NewXID().Hex(), p)
//	if err != nil {
//		return err
//	}
//	return nil
//}