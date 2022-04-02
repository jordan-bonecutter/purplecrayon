package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

type transformable struct {
	attrs map[string]string
}

func makeTransformable() transformable {
	return transformable{
		attrs: make(map[string]string),
	}
}

func (t transformable) Translate(p core.Point) {
	t.attrs["translate"] = fmt.Sprintf("translate(%f, %f)", p.X, p.Y)
}

func (t transformable) Scale(scale float64) {
	t.attrs["scale"] = fmt.Sprintf("scale(%f)", scale)
}

func (t transformable) Rotate(degrees float64) {
	t.attrs["rotate"] = fmt.Sprintf("rotate(%f)", degrees)
}

func (t transformable) compile() []string {
	transform := ""
	sortedMapIter(t.attrs, func(k, v string) {
		transform += " " + v
	})
	return []string{fmt.Sprintf(`transform="%s"`, transform)}
}
