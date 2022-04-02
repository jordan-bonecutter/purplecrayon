package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type rect struct {
	object
}

func makeRect(svg *svg) rect {
	return rect{makeObject(svg, "rect")}
}

func (r rect) TopLeft(p core.Point) {
	r.Set("x", fmt.Sprintf("%f", p.X))
	r.Set("y", fmt.Sprintf("%f", p.Y))
}

func (r rect) Width(w float64) {
	r.Set("width", fmt.Sprintf("%f", w))
}

func (r rect) Height(h float64) {
	r.Set("height", fmt.Sprintf("%f", h))
}
