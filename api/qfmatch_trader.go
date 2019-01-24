package api

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/frozenpine/qfmatch4go"
	"gitlab.quantdo.cn/yuanyang/autoorder"
)

// QfMatchTrader 交易接口对象
type QfMatchTrader struct {
	connString string
	instance   *qfmatch4go.QFMatchTraderAPI
}

// Connect 前置连接
func (td *QfMatchTrader) Connect(conn string) (err error) {
	connPattern := regexp.MustCompile(`^(tcp|udp)://([^:]+):(\d+)$`)

	match := connPattern.FindStringSubmatch(conn)

	if len(match) != 4 {
		panic("invalid connection string.")
	}

	protocol, addr, port := match[1], match[2], match[3]

	td.connString = conn

	switch strings.ToLower(protocol) {
	case "tcp":
		addrString := fmt.Sprintf("%s:%s", addr, port)

		if _, hsErr := autoorder.TCPHandShake(addrString, time.Second*5); hsErr != nil {
			err = fmt.Errorf("failed connect to [%s]: %s", addrString, hsErr.Error())
			break
		}

		// todo: TCP client instance initiation
	case "udp":
		// todo: UDP client instance initiation
	}

	return
}

// Disconnect 断开前置连接
func (td *QfMatchTrader) Disconnect() (err error) {
	td.instance = nil
	return
}
