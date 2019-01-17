package api

import (
	"fmt"
	"github.com/HankWang95/go-getui-api/style"
	"github.com/HankWang95/go-getui-api/token"
	"github.com/smartwalle/xid"
	"testing"
)

var(
	testAppkey = "K18daILRxi84F6otMqDMr7"
	testAppID = "77DMykZUPT8kBKveSUYSj8"
	testMasterSecret = "nsCjdmj0i29s2xeLr0Lks"
	myPhoneCID = "e3e21cb9534dc41fe4b676fc464ae065"
)

var (
	msgStyle = style.GetSystemStyle("title 1333 ", "content 1222 ")
	p = GetPushSingleNotification(msgStyle, "", "", "")
)

func init() {
	err := InitGeTui(testAppkey, testAppID, testMasterSecret)
	fmt.Println(err)
}

// 单推
func TestPushSingle(t *testing.T) {
	result, err := PushSingle(myPhoneCID, "", xid.NewXID().Hex(), p)
	if err != nil {
		t.Error(err, "adsjiasidojaosdja")
	}
	fmt.Println("----------------------TestPushSingle\t---------------------------")
	fmt.Println(result)

	//parmarList := make([]string, 0, 1)
	//parmarList = append(parmarList, "RASS_0114_fd41012da40c1481d1b0f98ec3124f01")
	//pResult, err := query.PushResult(getAppId(), getToken(), &query.PushRESParmar{TaskIdList:parmarList})
	//fmt.Println(pResult, err)
}

// 单推 测试token失效情景
func TestPushSingleCloseToken(t *testing.T) {
	fmt.Println("----------------------TestPushSingleCloseToken\t---------------------------")
	tokenResult, err := token.SetAuthClose(getAppId(), getToken())
	fmt.Println("set token close :",tokenResult, err)
	result, err := PushSingle("", "hh", xid.NewXID().Hex(), p)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

// 判断用户是否在线
func TestCheckUserOnline(t *testing.T) {
	fmt.Println("----------------------TestCheckUserOnline\t---------------------------")
	s, e := CheckUserOnline(myPhoneCID)
	fmt.Println(s, e)
}

func TestLazyPush(t *testing.T) {
	err := LazyPush(myPhoneCID, "lazy Push test", "lazy push content", "sada")
	if err != nil {
		fmt.Println(err)
	}
}