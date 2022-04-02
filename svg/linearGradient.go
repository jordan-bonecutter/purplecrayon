package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type linearGradient struct {
	object
}

func makeLinearGradient(svg *svg) *linearGradient {
	return &linearGradient{
		object: makeObject(svg, "linearGradient"),
	}
}

func (g *linearGradient) SetLine(p0, p1 core.Point) pc.LinearGradient {
	g.Attr("x1").F64(p0.X).Finish()
	g.Attr("y1").F64(p0.Y).Finish()
	g.Attr("x2").F64(p1.X).Finish()
	g.Attr("y2").F64(p1.Y).Finish()
  return g
}

func (g *linearGradient) GradientStops() pc.GradientStops {
	g.Stop()
	return gradientStops{g.object.svg}
}

func (g *linearGradient) Close() core.Reference {
	return g.VerboseClose()
}
