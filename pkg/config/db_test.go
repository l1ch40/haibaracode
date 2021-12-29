package config

import (
	"fmt"
	"testing"
)

func TestMySQL(t *testing.T) {
	MySQL = (&mysql{}).Load("../../conf/db.ini").Init()
	fmt.Println(MySQL)
	if MySQL == nil {
		t.Fail()
	}
}
