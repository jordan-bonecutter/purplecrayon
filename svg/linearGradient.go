package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type linearGradient struct {
	svg   *svg
	stops []object
	object
}

func makeLinearGradient(svg *svg) *linearGradient {
	return &linearGradient{
		svg:    svg,
		object: makeObject(svg, "linearGradient"),
	}
}

func (g *linearGradient) SetLine(p0, p1 core.Point) {
	g.Set("x1", fmt.Sprintf("%f", p0.X))
	g.Set("y1", fmt.Sprintf("%f", p0.Y))
	g.Set("x2", fmt.Sprintf("%f", p1.X))
	g.Set("y2", fmt.Sprintf("%f", p1.Y))
}

func (g *linearGradient) AddStop(position float64, data string) {
	stop := makeObject(g.svg, "stop")
	stop.Set("offset", fmt.Sprintf("%f%%", 100*position))
	stop.Set("stop-color", data)
	g.stops = append(g.stops, stop)
}

func (g *linearGradient) AddRGBStop(position float64, color core.RGB) {
	g.AddStop(position, svgRGB(color))
}

func (g *linearGradient) AddRGBAStop(position float64, color core.RGBA) {
	g.AddStop(position, svgRGBA(color))
}

func (g *linearGradient) Close() core.Reference {
	closers := make([]Closer, len(g.stops))
	for idx, stop := range g.stops {
		closers[idx] = stop
	}
	return g.CloseChildren(closers...)
}
