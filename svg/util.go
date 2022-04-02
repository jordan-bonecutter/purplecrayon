package svg

import (
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type basicObject struct {
	object
}

func makeBasicObject(svg *svg, name string) basicObject {
	return basicObject{
		object: makeObject(svg, name),
	}
}

func (b basicObject) Close() core.Reference {
	return b.AbbreviatedClose()
}
