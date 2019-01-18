package eventbus

import (
	"time"
)

// Event 基础事件接口
type Event interface {
	GetTimestamp() time.Time
	GetEventID() string
	GetEventMsg() string
	GetData() interface{}
}

// EventBus 事件总线接口
type EventBus interface {
	PushEvent(event Event) error
	PopEvent() Event
	Identity() string
	GetMaxEventID() int64
}
