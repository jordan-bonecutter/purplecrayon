package svg

import (
	"github.com/jordan-bonecutter/purplecrayon/core"
	pc "github.com/jordan-bonecutter/purplecrayon"
)

type circle struct {
	basicObject
}

func makeCircle(svg *svg) circle {
	return circle{makeBasicObject(svg, "circle")}
}

func (c circle) Center(p core.Point) pc.Circle {
	c.Attr("cx").F64(p.X).Finish()
	c.Attr("cy").F64(p.Y).Finish()
  return c
}

func (c circle) Radius(r float64) pc.Circle {
	c.Attr("r").F64(r).Finish()
  return c
}
