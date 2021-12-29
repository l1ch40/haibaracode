package config

import (
	"fmt"
	"testing"
)

func TestCookie(t *testing.T) {
	Cookie = (&cookie{}).Load("../../conf/app.ini").Init()
	fmt.Println(Cookie)
	if Cookie == nil {
		t.Fail()
	}
}
