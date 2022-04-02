package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type rect struct {
	basicObject
}

func makeRect(svg *svg) rect {
	return rect{makeBasicObject(svg, "rect")}
}

func (r rect) TopLeft(p core.Point) pc.Rect {
	r.Attr("x").F64(p.X).Finish()
	r.Attr("y").F64(p.Y).Finish()
  return r
}

func (r rect) Width(w float64) pc.Rect {
	r.Attr("width").F64(w).Finish()
  return r
}

func (r rect) Height(h float64) pc.Rect {
	r.Attr("height").F64(h).Finish()
  return r
}
