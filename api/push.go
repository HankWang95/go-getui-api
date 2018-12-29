package api

import (
	"github.com/HankWang95/go-getui-api/push"
	"time"
)

// 单推
func PushSingle(cid, requestId string, p *push.PushSingleParmar) (result *push.PushSingleResult, err error) {
	p.RequestId = requestId
	p.Cid = cid

	// Push不成功，重试5次
	for redoTime := 0; redoTime < 5; redoTime++ {
		result, err := push.PushSingle(getAppId(), getToken(), p)
		if err != nil {
			return nil, err
		}
		redo, err := handleResult(result)
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