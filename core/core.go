package core

type Reference string

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


