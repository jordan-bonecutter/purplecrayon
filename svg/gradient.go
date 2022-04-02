package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type gradientStops struct {
	*svg
}

func (s gradientStops) Stop() pc.GradientStop {
	return makeGradientStop(s.svg)
}

func (s gradientStops) Finish() {}

type gradientStop struct {
	basicObject
}

func makeGradientStop(svg *svg) gradientStop {
	return gradientStop{makeBasicObject(svg, "stop")}
}

func (s gradientStop) RGB(rgb core.RGB) pc.GradientStop {
	s.Attr("stop-color").RGB(rgb).Finish()
	return s
}

func (s gradientStop) RGBA(rgba core.RGBA) pc.GradientStop {
	s.Attr("stop-color").RGBA(rgba).Finish()
	return s
}

func (s gradientStop) Position(p float64) pc.GradientStop {
	s.Attr("offset").F64(p * 100).Str("%").Finish()
	return s
}

func (s gradientStop) Finish() {
	s.Close()
}
