package autoorder

import (
	"testing"
	"time"
)

func TestTCPHandShake(t *testing.T) {
	if rrt, err := TCPHandShake(":1", time.Second*5); err == nil {
		t.Error(rrt, err)
	} else {
		t.Log(rrt, err)
	}

	if rrt, err := TCPHandShake("baidu.com:80", time.Second*5); err != nil {
		t.Error(rrt, err)
	} else {
		t.Log(rrt, err)
	}
}
