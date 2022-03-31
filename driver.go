package purplecrayon

type Canvas interface {
  Referrable

  // Get the width of the canvas
  Width() float64

  // Get the height of the canvas
  Height() float64

  // Draw a rectangle inside the canvas
  Rect() Rect

  // Begin a path.
  Cursor() Cursor

  // Start writing text.
  Text() Text

  // Create a derived canvas which draws to a group.
  Group() Group

  // Create a derived canvas which draws to a mask.
  Mask() Mask
}

type Group interface {
  Canvas
  Transformable
  Paintable
}

type Mask interface {
  Canvas
  Transformable
  Paintable
}

type Transformable interface {
  Translate(delta Point)
  Scale(float64)
}

type Paintable interface {
  FillTransparent()
  FillRGB(RGB)
  FillRGBA(RGBA)
  Fill(Reference)
  StrokeWidth(float64)
  StrokeRGB(RGB)
  StrokeRGBA(RGBA)
  StrokeTransparent()
  Stroke(Reference)
  FontFamily(string)
  FontSize(float64)
  CompositeMask(Reference)
}

// Referrables are objects which can have references.
type Referrable interface {
  // Sets the reference for the object.
  Reference(string)

  // Closes the object and creates a reference for other objects to refer to it.
  Close() Reference
}

type Reference string

type Cursor interface {
  Referrable
  Transformable
  Paintable

  // Move the cursor to an absolute point.
  MoveTo(Point)

  // Move the cursor to a relative point.
  MoveToRel(Point)

  // Draw a line to an absolute point.
  LineTo(Point)

  // Draw a line to a relative point.
  LineToRel(Point)

  // Draw an absolute quadratic bezier curve.
  QuadTo(Point, Point)

  // Draw a relative quadratic bezier curve.
  QuadToRel(Point, Point)

  // Zips up the path to where it started.
  Zip()

  // Finish the path.
  Close() Reference
}

type Text interface {
  Referrable
  Transformable
  Paintable
  Text(string)
  Path(Reference)
}

type Rect interface {
  Referrable
  Transformable
  Paintable
  TopLeft(p Point)
  Width(float64)
  Height(float64)
}

// A linear gradient may have multiple color stops along a line.
// Modifications to a gradient may only occur before it has been used by Set
type LinearGradient interface {
  Referrable

  // Sets the line along which the gradient varies.
  // This line is always in a "hypothetical space" where the top left corner
  // is at coordinate (0, 0) and the bottom right is at (1, 1).
  // Once the gradient is used (0, 0) and (1, 1) are mapped to the bounds
  // in an affine manner.
  SetLine(p0, p1 Point)

  // Add an rgb color stop at the given position along the line.
  // Position varies from 0 to 1, where 0 will lie at p0 and 1 at p1.
  AddRGBStop(position float64, rgb RGB)

  // Add an rgba color stop at the given position along the line.
  // Position varies from 0 to 1, where 0 will lie at p0 and 1 at p1.
  AddRGBAStop(position float64, rgba RGBA)

  Close() Reference
}

// A color in RGB space.
type RGB struct {
  R uint8
  G uint8
  B uint8
}

// A color in RGBA space.
type RGBA struct {
  R uint8
  G uint8
  B uint8
  A uint8
}

// A point.
type Point struct {
  X float64
  Y float64
}

// Add two points.
func (p Point) Add(o Point) Point {
  return Point{
    X: p.X + o.X, Y: p.Y + o.Y,
  }
}

func (p Point) Sub(o Point) Point {
  return Point{
    X: p.X - o.X, Y: p.Y - o.Y,
  }
}

