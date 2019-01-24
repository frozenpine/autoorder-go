package autoorder

import (
	"testing"
	"time"
)

func TestTCPHandShake(t *testing.T) {
	if hst, err := TCPHandShake(":1", time.Second*5); err == nil {
		t.Error(hst, err)
	} else {
		t.Log(hst, err)
	}

	if hst, err := TCPHandShake("baidu.com:80", time.Second*5); err != nil {
		t.Error(hst, err)
	} else {
		t.Log(hst, err)
	}
}
