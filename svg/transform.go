package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type svgTransformable struct {
	attrs map[string]string
}

func makeSvgTransformable() svgTransformable {
	return svgTransformable{
		attrs: make(map[string]string),
	}
}

func (s svgTransformable) Translate(p core.Point) {
	s.attrs["translate"] = fmt.Sprintf("translate(%f, %f)", p.X, p.Y)
}

func (s svgTransformable) Scale(scale float64) {
	s.attrs["scale"] = fmt.Sprintf("scale(%f)", scale)
}

func (s svgTransformable) Rotate(degrees float64) {
	s.attrs["rotate"] = fmt.Sprintf("rotate(%f)", degrees)
}

func (s svgTransformable) compile() []string {
	transform := ""
	for _, v := range s.attrs {
		transform += " " + v
	}
	return []string{fmt.Sprintf(`transform="%s"`, transform)}
}
