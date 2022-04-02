package svg

import (
  "github.com/jordan-bonecutter/purplecrayon/core"
  "fmt"
)

type svgLinearGradient struct {
  svg *svg
  stops []svgObject
  svgObject
}

func makeSvgLinearGradient(svg *svg) *svgLinearGradient {
  return &svgLinearGradient{
    svg: svg,
    svgObject: makeSvgObject(svg, "linearGradient"),
  }
}

func (g *svgLinearGradient) SetLine(p0, p1 core.Point) {
  g.Set("x1", fmt.Sprintf("%f", p0.X))
  g.Set("y1", fmt.Sprintf("%f", p0.Y))
  g.Set("x2", fmt.Sprintf("%f", p1.X))
  g.Set("y2", fmt.Sprintf("%f", p1.Y))
}

func (g *svgLinearGradient) AddStop(position float64, data string) {
  stop := makeSvgObject(g.svg, "stop")
  stop.Set("offset", fmt.Sprintf("%f%%", 100 * position))
  stop.Set("stop-color", data)
  g.stops = append(g.stops, stop)
}

func (g *svgLinearGradient) AddRGBStop(position float64, color core.RGB) {
  g.AddStop(position, svgRGB(color))
}

func (g *svgLinearGradient) AddRGBAStop(position float64, color core.RGBA) {
  g.AddStop(position, svgRGBA(color))
}

func (g *svgLinearGradient) Close() core.Reference {
  closers := make([]Closer, len(g.stops))
  for idx, stop := range g.stops {
    closers[idx] = stop
  }
  return g.svgObject.CloseChildren(closers...)
}

