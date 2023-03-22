package ebievent

import "github.com/shiyou0130011/go-event"

type Game struct {
	listeners       map[string][]event.Listener
	isMouseParseing bool
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
	return nil
}
