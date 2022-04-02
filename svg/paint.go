package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type svgPaintable struct {
	attrs map[string]string
}

func makeSvgPaintable() svgPaintable {
	return svgPaintable{
		attrs: make(map[string]string),
	}
}

func (s svgPaintable) compile() []string {
	ret := make([]string, len(s.attrs))
	idx := 0
	sortedMapIter(s.attrs, func(k, v string) {
		ret[idx] = fmt.Sprintf(`%s="%s"`, k, v)
		idx++
	})
	return ret
}

func (s svgPaintable) FillTransparent() {
	s.attrs["fill"] = "none"
}

func (s svgPaintable) FillRGB(color core.RGB) {
	s.attrs["fill"] = svgRGB(color)
}

func (s svgPaintable) FillRGBA(color core.RGBA) {
	s.attrs["fill"] = svgRGBA(color)
}

func (s svgPaintable) Fill(ref core.Reference) {
	s.attrs["fill"] = svgRef(ref)
}

func (s svgPaintable) StrokeWidth(w float64) {
	s.attrs["stroke-width"] = fmt.Sprintf("%f", w)
}

func (s svgPaintable) StrokeRGB(color core.RGB) {
	s.attrs["stroke"] = svgRGB(color)
}

func (s svgPaintable) StrokeRGBA(color core.RGBA) {
	s.attrs["stroke"] = svgRGBA(color)
}

func (s svgPaintable) StrokeTransparent() {
	s.attrs["stroke"] = "none"
}

func (s svgPaintable) Stroke(ref core.Reference) {
	s.attrs["stroke"] = svgRef(ref)
}
