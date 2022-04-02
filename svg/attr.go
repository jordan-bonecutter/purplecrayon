package svg

import (
  "fmt"
  "github.com/jordan-bonecutter/purplecrayon/core"
)

type attributes struct {
  *svg
}

func makeAttributes(svg *svg) attributes {
  return attributes{svg}
}

type attr struct {
  key string
  *svg
}

func makeAttr(svg *svg, key string) attr {
  attr := attr{
    key: key,
    svg: svg,
  }
  attr.begin()
  return attr
}

func (a attr) begin() attr {
  a.WriteString(fmt.Sprintf(` %s="`, a.key)) 
  return a
}

func (a attr) U8(i uint8) attr {
  a.Str(fmt.Sprintf("%d", i))
  return a
}

func (a attr) Ref(ref core.Reference) attr {
  a.Str("url(#").Str(string(ref)).Str(")")
  return a
}

func (a attr) RGB(rgb core.RGB) attr {
  a.Str("rgb(").U8(rgb.R).Str(",").U8(rgb.G).Str(",").U8(rgb.B).Str(")")
  return a
}

func (a attr) RGBA(rgba core.RGBA) attr {
  a.Str("rgba(").U8(rgba.R).Str(",").U8(rgba.G).Str(",").U8(rgba.B).Str(",").U8(rgba.A).Str(")")
  return a
}

func (a attr) F64(f64 float64) attr {
  a.WriteString(a.FormatF64(f64))
  return a
}

func (a attr) Str(str string) attr {
  a.WriteString(str)
  return a
}

func(a attr) Finish() {
  a.WriteString(`"`)
}

func (a attributes) Attr(key string) attr {
  return makeAttr(a.svg, key)
}
