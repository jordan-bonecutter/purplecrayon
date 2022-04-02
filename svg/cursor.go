package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
	"strings"
)

type cursor struct {
	moves []string
	object
}

func makeCursor(svg *svg) *cursor {
	return &cursor{
		object: makeObject(svg, "path"),
	}
}

func (c *cursor) MoveTo(p core.Point) {
	c.moves = append(c.moves, fmt.Sprintf("M %f %f", p.X, p.Y))
}

func (c *cursor) MoveToRel(p core.Point) {
	c.moves = append(c.moves, fmt.Sprintf("m %f %f", p.X, p.Y))
}

func (c *cursor) LineTo(p core.Point) {
	c.moves = append(c.moves, fmt.Sprintf("L %f %f", p.X, p.Y))
}

func (c *cursor) LineToRel(p core.Point) {
	c.moves = append(c.moves, fmt.Sprintf("l %f %f", p.X, p.Y))
}

func (c *cursor) QuadTo(p0, p1 core.Point) {
	c.moves = append(c.moves, fmt.Sprintf("Q %f %f %f %f", p0.X, p0.Y, p1.X, p1.Y))
}

func (c *cursor) QuadToRel(p0, p1 core.Point) {
	c.moves = append(c.moves, fmt.Sprintf("q %f %f %f %f", p0.X, p0.Y, p1.X, p1.Y))
}

func (c *cursor) Zip() {
	c.moves = append(c.moves, "z")
}

func (c *cursor) Close() core.Reference {
	c.Set("d", strings.Join(c.moves, " "))
	return c.object.Close()
}
