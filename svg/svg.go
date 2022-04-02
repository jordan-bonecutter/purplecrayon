package svg

import (
	"fmt"
	pc "github.com/jordan-bonecutter/purplecrayon"
	core "github.com/jordan-bonecutter/purplecrayon/core"
	"io"
)

const (
	XMLNS_SVG = "http://www.w3.org/2000/svg"
)

// Register to purplecrayon
func init() {
	pc.Register("svg", NewSVGCanvas)
}

type svg struct {
	writer        io.Writer
	objectCounter uint64
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
		writer: writer,
	}
	canv := canvas{
		svg:       root,
		width:     width,
		height:    height,
		object: makeObject(root, "svg"),
	}

	canv.Set("width", fmt.Sprintf("%f", width))
	canv.Set("height", fmt.Sprintf("%f", height))
	canv.Set("xmlns", XMLNS_SVG)
	canv.Open()

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
	return c.ClosingTag()
}

func (c canvas) Rect() pc.Rect {
	r := makeRect(c.svg)
	return &r
}

func (c canvas) Circle() pc.Circle {
	r := makeCircle(c.svg)
	return &r
}

func (c canvas) Cursor() pc.Cursor {
	r := makeCursor(c.svg)
	return r
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
