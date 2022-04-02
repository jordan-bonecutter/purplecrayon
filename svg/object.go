package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type Closer interface {
	Close() core.Reference
}

type object struct {
	*svg
	name  string
	attrs map[string]string
	ref   core.Reference
	paintable
	transformable
}

func makeObject(svg *svg, name string) object {
	return object{
		svg:           svg,
		name:          name,
		ref:           svg.nextReference(),
		attrs:         make(map[string]string),
		paintable:     makePaintable(),
		transformable: makeTransformable(),
	}
}

func (o object) Set(k, v string) {
	o.attrs[k] = v
}

func (o object) Open() {
	o.WriteString("\n<")
	o.writeOpeningTagBody()
	o.WriteString(">")
}

func (o object) writeOpeningTagBody() core.Reference {
	o.WriteString(fmt.Sprintf(`%s id="%s"`, o.name, string(o.ref)))
	sortedMapIter(o.attrs, func(k, v string) {
		o.svg.WriteString(fmt.Sprintf(` %s="%s"`, k, v))
	})

	for _, compiled := range o.paintable.compile() {
		o.WriteString(" " + compiled)
	}

	for _, compiled := range o.transformable.compile() {
		o.WriteString(" " + compiled)
	}

	return o.ref
}

func (o object) CloseChildren(children ...Closer) core.Reference {
	o.Open()

	for _, child := range children {
		child.Close()
	}

	return o.ClosingTag()
}

func (o object) Close() core.Reference {
	o.svg.WriteString("\n<")
	o.writeOpeningTagBody()
	o.WriteString(`/>`)

	return o.ref
}

func (o object) ClosingTag() core.Reference {
	o.WriteString(fmt.Sprintf("\n</%s>", o.name))
	return o.ref
}
