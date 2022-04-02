package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type ellipse struct {
	basicObject
}

func makeEllipse(svg *svg) ellipse {
	return ellipse{makeBasicObject(svg, "ellipse")}
}

func (e ellipse) Center(p core.Point) pc.Ellipse {
	e.Attr("cx").F64(p.X).Finish()
	e.Attr("cy").F64(p.Y).Finish()
	return e
}

func (e ellipse) Radii(p core.Point) pc.Ellipse {
	e.Attr("rx").F64(p.X).Finish()
	e.Attr("ry").F64(p.Y).Finish()
	return e
}
