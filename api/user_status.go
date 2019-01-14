package api

import "github.com/HankWang95/go-getui-api/query"

// 检查用户是否在线
func CheckUserOnline(cid string) (online bool, err error) {
	result, err := query.UserStatus(getAppId(), getToken(), cid)
	if err != nil {
		return false, err
	}
	if result.Status == "online" {
		return true, nil
	}
	return false, nil
}