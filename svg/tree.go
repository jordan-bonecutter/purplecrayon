package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	core "github.com/jordan-bonecutter/purplecrayon/core"
)

type tree struct {
	canvas
	object
}

func makeTree(svg *svg, canv canvas, name string) tree {
	return tree{
		canvas: canv,
		object: makeObject(svg, name),
	}
}

func (t tree) Open() pc.Canvas {
	t.object.Open()
	return t
}

func (t tree) Close() core.Reference {
	return t.object.ClosingTag()
}
