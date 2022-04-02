package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type path struct {
	basicObject
}

func makePath(svg *svg) path {
	return path{makeBasicObject(svg, "path")}
}

func (p path) Cursor() pc.Cursor {
	return p.Attr("d")
}

// The cursor is an extension of attr.
func (c attr) cursorMove(action string, arguments ...float64) {
	c.Str(" ").Str(action)
	for _, arg := range arguments {
		c.F64(arg).Str(" ")
	}
}

func (c attr) MoveTo(p core.Point) pc.Cursor {
	c.cursorMove("M", p.X, p.Y)
  return c
}

func (c attr) MoveToRel(p core.Point) pc.Cursor {
	c.cursorMove("m", p.X, p.Y)
  return c
}

func (c attr) LineTo(p core.Point) pc.Cursor {
	c.cursorMove("L", p.X, p.Y)
  return c
}

func (c attr) LineToRel(p core.Point) pc.Cursor {
	c.cursorMove("l", p.X, p.Y)
  return c
}

func (c attr) QuadTo(p0, p1 core.Point) pc.Cursor {
	c.cursorMove("Q", p0.X, p0.Y, p1.X, p1.Y)
  return c
}

func (c attr) QuadToRel(p0, p1 core.Point) pc.Cursor {
	c.cursorMove("q", p0.X, p0.Y, p1.X, p1.Y)
  return c
}

func (c attr) Zip() pc.Cursor {
	c.cursorMove("z")
  return c
}
