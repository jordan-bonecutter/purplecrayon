package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type Closer interface {
	Close() core.Reference
}

type svgObject struct {
	*svg
	name  string
	attrs map[string]string
  ref core.Reference
	svgPaintable
	svgTransformable
}

func makeSvgObject(svg *svg, name string) svgObject {
	return svgObject{
		svg:              svg,
		name:             name,
    ref:              svg.nextReference(),
		attrs:            make(map[string]string),
		svgPaintable:     makeSvgPaintable(),
		svgTransformable: makeSvgTransformable(),
	}
}

func (o svgObject) Set(k, v string) {
	o.attrs[k] = v
}

func (o svgObject) Open() {
  o.WriteString("\n<")
  o.writeOpeningTagBody()
  o.WriteString(">")
}

func (o svgObject) writeOpeningTagBody() core.Reference {
  o.WriteString(fmt.Sprintf(`%s id="%s"`, o.name, string(o.ref)))
	for k, v := range o.attrs {
		o.svg.WriteString(fmt.Sprintf(` %s="%s"`, k, v))
	}

	for _, compiled := range o.svgPaintable.compile() {
		o.WriteString(" " + compiled)
	}

	for _, compiled := range o.svgTransformable.compile() {
		o.WriteString(" " + compiled)
	}

  return o.ref
}

func (o svgObject) CloseChildren(children ...Closer) core.Reference {
  o.Open()

	for _, child := range children {
		child.Close()
	}

	return o.ClosingTag()
}

func (o svgObject) Close() core.Reference {
	o.svg.WriteString("\n<")
  o.writeOpeningTagBody()
	o.WriteString(`/>`)

	return o.ref
}

func (o svgObject) ClosingTag() core.Reference {
	o.WriteString(fmt.Sprintf("\n</%s>", o.name))
  return o.ref
}
