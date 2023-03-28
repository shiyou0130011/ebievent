package ebievent

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shiyou0130011/go-event"
)

type Game struct {
	listeners        map[string][]event.Listener
	isMouseParseing  bool
	numOfKeysPressed int
}

type noneDispatchEvent string

func (e noneDispatchEvent) Type() string {
	return string(e)
}

// Update implements the ebiten.Game interface
func (g *Game) Update() error {
	if anyMouseButtonPressed() && !g.isMouseParseing {
		g.DispatchEvent(noneDispatchEvent("click"))
		g.isMouseParseing = true
	} else if !anyMouseButtonPressed() {
		g.isMouseParseing = false
	}
	if keys := inpututil.AppendPressedKeys(nil); len(keys) > 0 && len(keys) > g.numOfKeysPressed {
		g.numOfKeysPressed = len(keys)
		g.DispatchEvent(newNonDispatchedKeyboardEvent("keydown", keys))
	} else if len(keys) < g.numOfKeysPressed {
		g.numOfKeysPressed = len(keys)
		g.DispatchEvent(newNonDispatchedKeyboardEvent("keyup", keys))
	}
	return nil
}
