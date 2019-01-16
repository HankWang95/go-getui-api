package api

import "github.com/HankWang95/go-getui-api/token"

var (
	appId        string
	appKey       string
	masterSecret string
	authToken    string
)

func InitGeTui(appkey, appid, mastersecret string) (err error) {
	appKey = appkey
	appId = appid
	masterSecret = mastersecret
	return initToken()

}

func initToken() error {
	result, err := token.GetAuthSign(appId, appKey, masterSecret)
	if err != nil {
		return err
	}
	if result.Result != "ok" {
		return GetTokenErr
	}
	authToken = result.AuthToken
	return nil
}

func getToken() string {
	return authToken
}

func getAppId() string {
	return appId
}
