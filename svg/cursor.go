package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
	"strings"
)

type svgCursor struct {
	moves []string
	svgObject
}

func makeSvgCursor(svg *svg) *svgCursor {
	return &svgCursor{
		svgObject: makeSvgObject(svg, "path"),
	}
}

func (s *svgCursor) MoveTo(p core.Point) {
	s.moves = append(s.moves, fmt.Sprintf("M %f %f", p.X, p.Y))
}

func (s *svgCursor) MoveToRel(p core.Point) {
	s.moves = append(s.moves, fmt.Sprintf("m %f %f", p.X, p.Y))
}

func (s *svgCursor) LineTo(p core.Point) {
	s.moves = append(s.moves, fmt.Sprintf("L %f %f", p.X, p.Y))
}

func (s *svgCursor) LineToRel(p core.Point) {
	s.moves = append(s.moves, fmt.Sprintf("l %f %f", p.X, p.Y))
}

func (s *svgCursor) QuadTo(p0, p1 core.Point) {
	s.moves = append(s.moves, fmt.Sprintf("Q %f %f %f %f", p0.X, p0.Y, p1.X, p1.Y))
}

func (s *svgCursor) QuadToRel(p0, p1 core.Point) {
	s.moves = append(s.moves, fmt.Sprintf("q %f %f %f %f", p0.X, p0.Y, p1.X, p1.Y))
}

func (s *svgCursor) Zip() {
	s.moves = append(s.moves, "z")
}

func (s *svgCursor) Close() core.Reference {
	s.Set("d", strings.Join(s.moves, " "))
	return s.svgObject.Close()
}
