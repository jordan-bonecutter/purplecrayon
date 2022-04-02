package svg

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	"github.com/jordan-bonecutter/purplecrayon/core"
)

// Similarly to paint, transform is implemnented as an extension of attr.
// It is not an attributes because transform is entirely enclosed within
// one svg tag attribute.

func (t attr) transform(name string, arguments ...float64) attr {
	t.Str(" ")

	t.Str(name).Str("(")
	nArgs := len(arguments)
	switch nArgs {
	case 0:
		break
	case 1:
		t.F64(arguments[0])
	default:
		t.F64(arguments[0])
		for _, arg := range arguments[:nArgs-1] {
			t.Str(",").F64(arg)
		}
		t.F64(arguments[nArgs-1]).Str(")")
	}

	return t
}

func (t attr) Translate(p core.Point) pc.Transform {
	return t.transform("translate", p.X, p.Y)
}

func (t attr) Scale(scale float64) pc.Transform {
	return t.transform("scale", scale)
}

func (t attr) Rotate(degrees float64) pc.Transform {
	return t.transform("rotate", degrees)
}
