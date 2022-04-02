package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	core "github.com/jordan-bonecutter/purplecrayon/core"
)

type tree struct {
	canvas
	svgObject
}

func makeTree(svg *svg, canv canvas, name string) tree {
	return tree{
		canvas:    canv,
		svgObject: makeSvgObject(svg, name),
	}
}

func (t tree) Open() pc.Canvas {
	t.svgObject.Open()
	return t
}

func (t tree) Close() core.Reference {
	return t.svgObject.ClosingTag()
}
