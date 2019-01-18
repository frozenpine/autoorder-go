package api

import (
	"fmt"
	"time"

	"gitlab.quantdo.cn/yuanyang/autoorder"

	. "gitlab.quantdo.cn/yuanyang/qfmatch4go"
)

// QfMatchTraer 交易接口对象
type QfMatchTraer struct {
	frontAddr string
	instance  *QFMatchTraderAPI
}

// Connect 前置连接
func (td *QfMatchTraer) Connect(addr string) error {
	td.frontAddr = addr

	if _, err := autoorder.TCPHandShake(":1234", time.Second*5); err != nil {
		return fmt.Errorf("failed connect to %s: %s", addr, err.Error())
	}

	return nil
}
