package purplecrayon

import (
  core "github.com/jordan-bonecutter/purplecrayon/core"
  "io"
  "fmt"
)

type Point = core.Point
type RGB = core.RGB
type RGBA = core.RGBA
type Reference = core.Reference

type Canvas interface {
  Referrable

  // Get the width of the canvas
  Width() float64

  // Get the height of the canvas
  Height() float64

  // Draw a rectangle inside the canvas
  Rect() Rect

  // Draw a circle inside the canvas
  Circle() Circle

  // Begin a path.
  Cursor() Cursor
}

type Transformable interface {
  Translate(delta Point)
  Scale(float64)
  Rotate(degrees float64)
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
}

// Referrables are objects which can have references.
type Referrable interface {
  // Closes the object and creates a reference for other objects to refer to it.
  Close() Reference
}

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

type Rect interface {
  Referrable
  Transformable
  Paintable
  TopLeft(p Point)
  Width(float64)
  Height(float64)
}

type Circle interface {
  Referrable
  Transformable
  Paintable
  Center(Point)
  Radius(float64)
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

// A function which returns a driver for some backend
type Driver func(width, height float64, w io.Writer) Canvas

// All drivers which have been registered
var registeredDrivers map[string]Driver

// Register a new driver
func Register(name string, d Driver) {
  if registeredDrivers == nil {
    registeredDrivers = make(map[string]Driver)
  }

  registeredDrivers[name] = d
}

var driverNotFoundError = fmt.Errorf("Driver not found!")

// Create a new canvas for the given driver
func NewCanvas(driver string, width, height float64, w io.Writer) (c Canvas, err error) {
  if registeredDrivers == nil {
    err = driverNotFoundError
  }

  if lib, ok := registeredDrivers[driver]; ok {
    c = lib(width, height, w) 
  } else {
    err = driverNotFoundError
  }
  return
}

