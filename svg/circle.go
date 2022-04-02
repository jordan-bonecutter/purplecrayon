package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type circle struct {
	object
}

func makeCircle(svg *svg) circle {
	return circle{makeObject(svg, "circle")}
}

func (s circle) Center(p core.Point) {
	s.Set("cx", fmt.Sprintf("%f", p.X))
	s.Set("cy", fmt.Sprintf("%f", p.Y))
}

func (s circle) Radius(r float64) {
	s.Set("r", fmt.Sprintf("%f", r))
}
