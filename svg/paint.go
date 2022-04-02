package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type paintable struct {
	attrs map[string]string
}

func makePaintable() paintable {
	return paintable{
		attrs: make(map[string]string),
	}
}

func (p paintable) compile() []string {
	ret := make([]string, len(p.attrs))
	idx := 0
	sortedMapIter(p.attrs, func(k, v string) {
		ret[idx] = fmt.Sprintf(`%s="%s"`, k, v)
		idx++
	})
	return ret
}

func (p paintable) FillTransparent() {
	p.attrs["fill"] = "none"
}

func (p paintable) FillRGB(color core.RGB) {
	p.attrs["fill"] = svgRGB(color)
}

func (p paintable) FillRGBA(color core.RGBA) {
	p.attrs["fill"] = svgRGBA(color)
}

func (p paintable) Fill(ref core.Reference) {
	p.attrs["fill"] = svgRef(ref)
}

func (p paintable) StrokeWidth(w float64) {
	p.attrs["stroke-width"] = fmt.Sprintf("%f", w)
}

func (p paintable) StrokeRGB(color core.RGB) {
	p.attrs["stroke"] = svgRGB(color)
}

func (p paintable) StrokeRGBA(color core.RGBA) {
	p.attrs["stroke"] = svgRGBA(color)
}

func (p paintable) StrokeTransparent() {
	p.attrs["stroke"] = "none"
}

func (p paintable) Stroke(ref core.Reference) {
	p.attrs["stroke"] = svgRef(ref)
}

func (p paintable) CompositeMask(ref core.Reference) {
	p.attrs["mask"] = svgRef(ref)
}
