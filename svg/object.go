package svg

import (
	"fmt"
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type object struct {
	name string
	ref  core.Reference
	attributes
}

func makeObject(svg *svg, name string) object {
	obj := object{
		name:       name,
		ref:        svg.nextReference(),
		attributes: makeAttributes(svg),
	}

	obj.Begin()
	return obj
}

func (o object) Transform() pc.Transform {
	return o.Attr("transform")
}

func (o object) Begin() {
	o.WriteString(fmt.Sprintf("\n<%s", o.name))
	o.Attr("id").Str(string(o.ref)).Finish()
}

func (o object) AbbreviatedClose() core.Reference {
	o.WriteString("/>")
	return o.ref
}

func (o object) VerboseClose() core.Reference {
	o.WriteString(fmt.Sprintf("</%s>", o.name))
	return o.ref
}

func (o object) Stop() {
	o.WriteString(">")
}
