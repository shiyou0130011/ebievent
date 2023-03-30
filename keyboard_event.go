package ebievent

import (
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/shiyou0130011/go-event"
)

type KeyboardEvent interface {
	event.Event

	Keys() []ebiten.Key
	Ctrl() KeyPosition
	Alt() KeyPosition
	Shift() KeyPosition
}

type KeyPosition int

const (
	KeyInLeft KeyPosition = 1 << iota
	KeyInRight
	KeyIsPressed
)
const KeyNotPress KeyPosition = 0

type nonDispatchedKeyboardEvent struct {
	keys     []ebiten.Key
	ctrl     KeyPosition
	alt      KeyPosition
	shift    KeyPosition
	typename string
}

func (e *nonDispatchedKeyboardEvent) Type() string {
	return e.typename
}

func (e *nonDispatchedKeyboardEvent) Keys() []ebiten.Key {
	return e.keys
}

func (e *nonDispatchedKeyboardEvent) Ctrl() KeyPosition {
	return e.ctrl

}
func (e *nonDispatchedKeyboardEvent) Alt() KeyPosition {
	return e.alt
}
func (e *nonDispatchedKeyboardEvent) Shift() KeyPosition {
	return e.shift
}

func newNonDispatchedKeyboardEvent(name string, keys []ebiten.Key) *nonDispatchedKeyboardEvent {
	e := &nonDispatchedKeyboardEvent{
		typename: name,
	}
	for _, k := range keys {
		switch k {
		case ebiten.KeyAlt:
			e.alt |= KeyIsPressed
		case ebiten.KeyAltLeft:
			e.alt |= KeyInLeft
		case ebiten.KeyAltRight:
			e.alt |= KeyInRight
		case ebiten.KeyShift:
			e.shift |= KeyIsPressed
		case ebiten.KeyShiftLeft:
			e.shift |= KeyInLeft | KeyIsPressed
		case ebiten.KeyShiftRight:
			e.shift |= KeyInRight | KeyIsPressed
		case ebiten.KeyControl:
			e.ctrl |= KeyIsPressed
		case ebiten.KeyControlLeft:
			e.ctrl |= KeyInLeft | KeyIsPressed
		case ebiten.KeyControlRight:
			e.ctrl |= KeyInRight | KeyIsPressed
		default:
			e.keys = append(e.keys, k)
		}
	}
	return e
}

type keyboardEvent struct {
	nonDispatchedKeyboardEvent
	target    event.EventTarget
	timestamp time.Time
}

func (e *keyboardEvent) Timestamp() time.Time {
	return e.timestamp
}

func (e *keyboardEvent) Target() event.EventTarget {
	return e.target
}
