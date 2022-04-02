package svg

import (
	"github.com/jordan-bonecutter/purplecrayon/core"
)

// Paintable is an extension on attributes.
// While technically all svg objects will contain these methods, only
// those exposed by purplecrayon will be callable from the public 
// interface!

func (p attributes) FillTransparent() {
  p.Attr("fill").Str("none").Finish()
}

func (p attributes) FillRGB(color core.RGB) {
  p.Attr("fill").RGB(color).Finish()
}

func (p attributes) FillRGBA(color core.RGBA) {
  p.Attr("fill").RGBA(color).Finish()
}

func (p attributes) Fill(ref core.Reference) {
  p.Attr("fill").Ref(ref).Finish()
}

func (p attributes) StrokeWidth(w float64) {
  p.Attr("stroke-width").F64(w).Finish()
}

func (p attributes) StrokeRGB(color core.RGB) {
  p.Attr("stroke").RGB(color).Finish()
}

func (p attributes) StrokeRGBA(color core.RGBA) {
  p.Attr("stroke").RGBA(color).Finish()
}

func (p attributes) StrokeTransparent() {
  p.Attr("stroke").Str("none").Finish()
}

func (p attributes) Stroke(ref core.Reference) {
  p.Attr("stroke").Ref(ref).Finish()
}

func (p attributes) CompositeMask(ref core.Reference) {
  p.Attr("mask").Ref(ref).Finish()
}
