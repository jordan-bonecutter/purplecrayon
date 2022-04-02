package svg

import (
	"fmt"
	pc "github.com/jordan-bonecutter/purplecrayon"
	core "github.com/jordan-bonecutter/purplecrayon/core"
	"io"
	"strconv"
)

const (
	XMLNS_SVG = "http://www.w3.org/2000/svg"
)

const (
	FLOAT_PRECISION         = "floatPrecision"
	FLOAT_PRECISION_HIGHEST = -1
)

// Register to purplecrayon
func init() {
	pc.Register("svg", NewSVGCanvas)
}

type svg struct {
	writer         io.Writer
	objectCounter  uint64
	floatPrecision int
}

func (svg *svg) FormatF64(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', svg.floatPrecision, 64)
}

var UnknownKeyError = fmt.Errorf("Unknown configuration key")

func (svg *svg) Configure(key string, value interface{}) error {
	switch key {
	case FLOAT_PRECISION:
		return svg.ConfigureFloatPrecision(value)
	default:
		return UnknownKeyError
	}
}

var RequiresIntError = fmt.Errorf("Float Precision must be of type int")

func (svg *svg) ConfigureFloatPrecision(value interface{}) error {
	if i, ok := value.(int); ok {
		svg.floatPrecision = i
		return nil
	}

	return RequiresIntError
}

type canvas struct {
	svg    *svg
	width  float64
	height float64
	object
}

func (svg *svg) nextReference() core.Reference {
	defer func() {
		svg.objectCounter++
	}()
	return core.Reference(fmt.Sprintf("pcobj-%d", svg.objectCounter))
}

func (s *svg) WriteString(str string) {
	io.WriteString(s.writer, str)
}

// Creates a new canvas which draws to an svg via the given io.Writer
func NewSVGCanvas(width, height float64, writer io.Writer) (pcCanvas pc.Canvas) {

	root := &svg{
		writer:         writer,
		floatPrecision: FLOAT_PRECISION_HIGHEST,
	}
	canv := canvas{
		svg:    root,
		width:  width,
		height: height,
		object: makeObject(root, "svg"),
	}

	canv.Attr("width").F64(width).Finish()
	canv.Attr("height").F64(height).Finish()
	canv.Attr("xmlns").Str(XMLNS_SVG).Finish()
	canv.Stop()

	pcCanvas = canv

	return
}

func (c canvas) Width() float64 {
	return c.width
}

func (c canvas) Height() float64 {
	return c.height
}

func (c canvas) Close() core.Reference {
	return c.VerboseClose()
}

func (c canvas) Rect() pc.Rect {
	return makeRect(c.svg)
}

func (c canvas) Circle() pc.Circle {
	return makeCircle(c.svg)
}

func (c canvas) Ellipse() pc.Ellipse {
	return makeEllipse(c.svg)
}

func (c canvas) Path() pc.Path {
	return makePath(c.svg)
}

func (c canvas) LinearGradient() pc.LinearGradient {
	return makeLinearGradient(c.svg)
}

func (c canvas) Group() pc.Group {
	return makeGroup(c.svg, c)
}

func (c canvas) Mask() pc.Mask {
	return makeMask(c.svg, c)
}
