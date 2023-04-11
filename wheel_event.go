package ebievent

import (
	"time"

	"github.com/shiyou0130011/go-event"
)

type WheelEvent interface {
	event.Event
	Wheel() (x, y float64)
}

type nonDispatchedWheelEvent struct {
	x, y     float64
	typename string
}

func (e *nonDispatchedWheelEvent) Wheel() (float64, float64) {
	return e.x, e.y
}

func (e *nonDispatchedWheelEvent) Type() string {
	return e.typename
}

type wheelEvent struct {
	nonDispatchedWheelEvent
	target    event.EventTarget
	timestamp time.Time
}

func (e *wheelEvent) Timestamp() time.Time {
	return e.timestamp
}

func (e *wheelEvent) Target() event.EventTarget {
	return e.target
}
