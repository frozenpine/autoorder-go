package autoorder

import (
	"fmt"
	"math"
	"net"
	"time"
)

// ValidateVolume 校验Volume合法性, <=0 非法
func ValidateVolume(vol int64) bool {
	return vol > 0
}

// ValidatePrice 校验价格合法性, 0 | MaxFloat64 非法
func ValidatePrice(price float64) bool {
	return price != 0 && price != math.MaxFloat64
}

type roundMode uint8

const (
	// RoundDefault 默认四舍五入
	RoundDefault roundMode = iota
	// RoundUp 向上取整
	RoundUp
	// RoundDown 向下取整
	RoundDown
)

// NormalizePrice 将价格以指定的方式取整至tickPrice的整数倍
func NormalizePrice(price, tickPrice float64, round roundMode) float64 {
	multiple := price / tickPrice

	switch round {
	case RoundDefault:
		multiple = math.Round(multiple)
	case RoundUp:
		multiple = math.Ceil(multiple)
	case RoundDown:
		multiple = math.Floor(multiple)
	default:
		multiple = math.Round(multiple)
	}

	return tickPrice * multiple
}

// MaxFloat64 查找一组float64中的最大值
func MaxFloat64(f ...float64) float64 {
	var max float64
	max = 0
	for _, v := range f {
		max = math.Max(max, v)
	}
	return max
}

// MinFloat64 查找一组float64中的最小值
func MinFloat64(f ...float64) float64 {
	var min float64
	min = math.MaxFloat64
	for _, v := range f {
		min = math.Min(min, v)
	}
	return min
}

// TCPHandShake TCP三次握手测试
// 如连接成功, 将返回 三次握手时间 / 3, 作为近似的单程物理延时
func TCPHandShake(addr string, timeout time.Duration) (time.Duration, error) {
	var sockErr error
	ch := make(chan time.Duration)

	go func() {
		start := time.Now()
		conn, err := net.Dial("tcp", addr)
		dur := time.Now().Sub(start)

		if err != nil {
			sockErr = err
			ch <- time.Duration(0)
		} else {
			conn.Close()
			ch <- dur / 3
		}
	}()

	select {
	case dur := <-ch:
		return dur, sockErr
	case <-time.After(timeout):
		return time.Duration(0), fmt.Errorf("connect timeout: %v", timeout)
	}
}
