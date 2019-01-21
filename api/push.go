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
	//j,err:=json.Marshal(p)
	//fmt.Println(string(j))

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

func LazyPush(cid, title, content, transmission string) error {
	msgStyle := style.GetSystemStyle(content, title)
	var p *push.PushSingleParmar
	if transmission == "" {
		p = GetPushSingleNotification(msgStyle, "", "", "")
	} else {
		p = GetPushSingleMsgTypeTransmission(msgStyle, transmission, "", "")
	}
	pushInfo := new(push.PushInfo)
	pushInfo.Aps.Alert.Title = title
	pushInfo.Aps.Alert.Body = content
	pushInfo.Aps.AutoBadge = "+1"
	pushInfo.Transmission = transmission
	p.PushInfo = pushInfo
	_, err := PushSingle(cid, "", xid.NewXID().Hex(), p)
	if err != nil {
		return err
	}
	return nil
}
