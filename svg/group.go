package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	core "github.com/jordan-bonecutter/purplecrayon/core"
)

type group struct {
	canvas
	svgObject
}

func makeGroup(svg *svg, canv canvas) group {
	return group{
		canvas:    canv,
		svgObject: makeSvgObject(svg, "g"),
	}
}

func (g group) Open() pc.Canvas {
	g.svgObject.Open()
	return g
}

func (g group) Close() core.Reference {
	return g.svgObject.ClosingTag()
}
