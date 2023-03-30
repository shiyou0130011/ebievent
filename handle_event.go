package ebievent

import (
	"reflect"
	"sync"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/shiyou0130011/go-event"
)

// copied code from BasicEvent in go-event lib
type basicEvent struct {
	event.NonDispatchedEvent
	target    event.EventTarget
	timestamp time.Time
}

func (e *basicEvent) Target() event.EventTarget {
	return e.target
}
func (e *basicEvent) Timestamp() time.Time {
	return e.timestamp
}

// copied code from same function in go-event lib
func (g *Game) AddEventListener(name string, listener event.Listener) {
	if g.listeners == nil {
		g.listeners = make(map[string][]event.Listener)
	}
	lAdded := reflect.ValueOf(listener)
	if list, has := g.listeners[name]; has {
		for _, l := range list {
			l1 := reflect.ValueOf(l)

			if l1.Pointer() == lAdded.Pointer() {
				return
			}
		}
	}
	g.listeners[name] = append(g.listeners[name], listener)
}

// copied code from same function in go-event lib
func (g *Game) RemoveEventListener(name string, listener event.Listener) {
	if g.listeners[name] == nil || len(g.listeners[name]) == 0 {
		return
	}
	lRemoved := reflect.ValueOf(listener)
	for i, l := range g.listeners[name] {
		l1 := reflect.ValueOf(l)

		if l1.Pointer() == lRemoved.Pointer() {
			g.listeners[name] = append(g.listeners[name][:i], g.listeners[name][i+1:]...)

			return
		}
	}
}

func (g *Game) DispatchEvent(origevent event.NonDispatchedEvent) (result bool) {
	result = true
	var cancelable = false
	if ce, isCancelableEvent := origevent.(event.CancelableEvent); isCancelableEvent {
		cancelable = ce.Cancelable()
	}

	list, has := g.listeners[origevent.Type()]
	if !has || len(list) == 0 {
		return
	}
	var e event.Event
	if ke, is := origevent.(*nonDispatchedKeyboardEvent); is {
		_e := &keyboardEvent{}
		_e.nonDispatchedKeyboardEvent = *ke
		_e.timestamp = time.Now()
		e = _e

	} else {
		switch origevent.Type() {
		case EClick:
			_e := &clickEvent{}
			_e.x, _e.y = ebiten.CursorPosition()
			_e.timestamp = time.Now()
			e = _e

		default:
			e = &basicEvent{
				NonDispatchedEvent: origevent,
				timestamp:          time.Now(),
			}
		}

	}

	if cancelable {
		result = g.dispatchCancelableEvent(e)
	} else {
		g.dispatchNonCancelableEvent(e)
	}
	return
}

// copied code from same function in go-event lib
func (t *Game) dispatchCancelableEvent(e event.Event) (result bool) {
	for _, listener := range t.listeners[e.Type()] {
		result = listener(e)
		if !result {
			return
		}
	}
	return
}

// copied code from same function in go-event lib
func (t *Game) dispatchNonCancelableEvent(e event.Event) {
	var wg sync.WaitGroup
	for _, listener := range t.listeners[e.Type()] {
		wg.Add(1)
		go (func(l event.Listener) {
			defer wg.Done()

			l(e)
		})(listener)
	}
	wg.Wait()
}
