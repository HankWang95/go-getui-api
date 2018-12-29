package api

import (
	"fmt"
	"testing"
)

func TestInitGeTui(t *testing.T) {
	InitGeTui("XH93kDE2AZ6x3pCGwEQNn", "XH93kDE2AZ6x3pCGwEQNn", "tT1khrhlup8vskHi5iVpk4")
	p := GetNotification()
	result, err := PushSingle("sb", "requestid", p)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
