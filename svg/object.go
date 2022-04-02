package svg

import (
  "github.com/jordan-bonecutter/purplecrayon/core"
  "fmt"
)

type Closer interface {
  Close() core.Reference
}

type svgObject struct {
  *svg
  name string
  attrs map[string]string
  svgPaintable
  svgTransformable
}

func makeSvgObject(svg *svg, name string) svgObject {
  return svgObject{
    svg: svg,
    name: name,
    attrs: make(map[string]string),
    svgPaintable: makeSvgPaintable(),
    svgTransformable: makeSvgTransformable(),
  }
}

func (o svgObject) Set(k, v string) {
  o.attrs[k] = v
}

func (o svgObject) CloseChildren(children ...Closer) core.Reference {
  ref := o.svg.nextReference()
  o.svg.WriteString(fmt.Sprintf("\n<%s", o.name))
  for k, v := range o.attrs {
    o.svg.WriteString(fmt.Sprintf(` %s="%s"`, k, v))
  }
  o.svg.WriteString(fmt.Sprintf(` id="%s"`, string(ref)))

  for _, compiled := range o.svgPaintable.compile() {
    o.WriteString(" " + compiled)
  }

  for _, compiled := range o.svgTransformable.compile() {
    o.WriteString(" " + compiled)
  }
  o.WriteString(`>`)

  for _, child := range children {
    child.Close()
  }

  o.WriteString(fmt.Sprintf("\n</%s>", o.name))

  return ref
}

func (o svgObject) Close() core.Reference {
  o.svg.WriteString(fmt.Sprintf("\n<%s", o.name))
  for k, v := range o.attrs {
    o.svg.WriteString(fmt.Sprintf(` %s="%s"`, k, v))
  }

  for _, compiled := range o.svgPaintable.compile() {
    o.WriteString(" " + compiled)
  }

  for _, compiled := range o.svgTransformable.compile() {
    o.WriteString(" " + compiled)
  }
  o.WriteString(`/>`)

  return o.svg.nextReference()
}
