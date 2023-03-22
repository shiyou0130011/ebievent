package ebievent

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shiyou0130011/go-event"
)

func anyMouseButtonPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}

type mouseEvent struct {
	x, y      int
	timestamp time.Time
	target    event.EventTarget
}

func (m *mouseEvent) Timestamp() time.Time {
	return m.timestamp
}
func (m *mouseEvent) Position() (int, int) {
	return m.x, m.y
}
func (m *mouseEvent) Target() event.EventTarget {
	return m.target
}

// The event dispatched when mouse click
type ClickEvent struct{ mouseEvent }

func (c *ClickEvent) Type() string {
	return "click"
}
