package svg

import (
  "github.com/jordan-bonecutter/purplecrayon/core"
  "fmt"
)

type svgRect struct {
  svgObject
}

func makeSvgRect(svg *svg) svgRect {
  return svgRect{makeSvgObject(svg, "rect")}
}

func (s svgRect) TopLeft(p core.Point) {
  s.Set("x", fmt.Sprintf("%f", p.X))
  s.Set("y", fmt.Sprintf("%f", p.Y))
}

func (s svgRect) Width(w float64) {
  s.Set("width", fmt.Sprintf("%f", w))
}

func (s svgRect) Height(h float64) {
  s.Set("height", fmt.Sprintf("%f", h))
}

