package purplecrayon

import (
  "fmt"
  "io"
)

type svg struct {
  writer io.Writer
  width float64
  height float64
  objectCounter uint64
}

func (svg *svg) nextReference() Reference {
  defer func() {
    svg.objectCounter++
  }()
  return Reference(fmt.Sprintf("obj-%d", svg.objectCounter))
}

type svgcanvas struct {
  svg *svg
  onClose func() Reference
}

func topLevelClose() Reference {
  return ""
}

// Creates a new canvas which draws to an svg via the given io.Writer
func NewSVGCanvas(width, height float64, writer io.Writer) (canvas *svgcanvas) {
  root := &svg{
    writer: writer,
    width: width,
    height: height,
  }

  canvas = &svgcanvas{
    svg: root,
    onClose: topLevelClose,
  }

  return
}

func (s *svgcanvas) Width() float64 {
  return s.svg.width
}

func (s *svgcanvas) Height() float64 {
  return s.svg.height
}

type svgStr struct {
  set bool
  val string
}

func (s *svgStr) Set(val string) {
  s.set = true
  s.val = val
}

func (s *svgStr) Str() (str string, ok bool) {
  str = s.val
  ok = s.set
  return
}

type svgFloat struct {
  set bool
  val float64
}

func (s *svgFloat) Set(val float64) {
  s.set = true
  s.val = val
}

func (s *svgFloat) Float64() (f float64, ok bool) {
  f = s.val
  ok = s.set
  return
}

func (color RGB) compiled() string {
  return fmt.Sprintf("rgb(%f, %f, %f)",
    float64(color.R) / 255,
    float64(color.G) / 255,
    float64(color.B) / 255,
  )
}

func (color RGBA) compiled() string {
  return fmt.Sprintf("rgba(%f, %f, %f, %f)",
    float64(color.R) / 255,
    float64(color.G) / 255,
    float64(color.B) / 255,
    float64(color.A) / 255,
  )
}

func (r Reference) compiled() string {
  return fmt.Sprintf("url(#%s)", string(r))
}

type svgPaintable struct {
  fill svgStr
  stroke svgStr
  strokeWidth svgFloat
  fontFamily svgStr
  fontSize svgFloat
  mask svgStr
}

func (s *svgPaintable) FillTransparent() {
  s.fill.Set("none")
}

func (s *svgPaintable) FillRGB(color RGB) {
  s.fill.Set(color.compiled())
}

func (s *svgPaintable) FillRGBA(color RGBA) {
  s.fill.Set(color.compiled())
}

func (s *svgPaintable) Fill(ref Reference) {
  s.fill.Set(ref.compiled())
}

func (s *svgPaintable) StrokeWidth(w float64) {
  s.strokeWidth.Set(w)
}

func (s *svgPaintable) StrokeRGB(color RGB) {
  s.stroke.Set(color.compiled())
}

func (s *svgPaintable) StrokeRGBA(color RGBA) {
  s.stroke.Set(color.compiled())
}

func (s *svgPaintable) StrokeTransparent() {
  s.stroke.Set("none")
}

func (s *svgPaintable) Stroke(ref Reference) {
  s.stroke.Set(ref.compiled())
}

func (s *svgPaintable) FontFamily(f string) {
  s.fontFamily.Set(f)
}

func (s *svgPaintable) FontSize(size float64) {
  s.fontSize.Set(size)
}

func (s *svgPaintable) CompositeMask(ref Reference) {
  s.mask.Set(ref.compiled())
}

type svgTransformable struct {
  translate svgStr
  scale svgFloat
}

func (s *svgTransformable) Translate(delta Point) {
  s.translate.Set(fmt.Sprintf("translate(%f, %f)", delta.X, delta.Y))
}

func (s *svgTransformable) Scale(scale float64) {
  s.scale.Set(scale)
}

type svgRect struct {
  svgPaintable
  svgTransformable
}

