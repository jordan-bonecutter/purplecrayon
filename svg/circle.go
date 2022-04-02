package svg

import (
  "github.com/jordan-bonecutter/purplecrayon/core"
  "fmt"
)

type svgCircle struct {
  svgObject
}

func makeSvgCircle(svg *svg) svgCircle {
  return svgCircle{makeSvgObject(svg, "circle")}
}

func (s svgCircle) Center(p core.Point) {
  s.Set("cx", fmt.Sprintf("%f", p.X))
  s.Set("cy", fmt.Sprintf("%f", p.Y))
}

func (s svgCircle) Radius(r float64) {
  s.Set("r", fmt.Sprintf("%f", r))
}
