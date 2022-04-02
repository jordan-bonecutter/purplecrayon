package svg

import (
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type rect struct {
  basicObject
}

func makeRect(svg *svg) rect {
	return rect{makeBasicObject(svg, "rect")}
}

func (r rect) TopLeft(p core.Point) {
  r.Attr("x").F64(p.X).Finish()
  r.Attr("y").F64(p.Y).Finish()
}

func (r rect) Width(w float64) {
  r.Attr("width").F64(w).Finish()
}

func (r rect) Height(h float64) {
  r.Attr("height").F64(h).Finish()
}
