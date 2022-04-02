// purplecrayon is a drawing library written in Go.
// While it can support an arbitrary backend, it only supports svg for now.
//
// The main interaction is via the Canvas. A canvas represents the interface between
// the frontend and the backend, allowing the programmer to create primitives like
// rectangles, circles, paths, and gradients. You can create such an SVG canvas via:
//  canv, err := pc.NewCanvas("svg", 1920, 1080, w)
// as long as the svg component has been imported via:
//  import _ "github.com/jordan-bonecutter/purplecrayon/svg"
//
// Always remember to close your objects with Close() so they're drawn to the canvas!
//
// purplecrayon can be extended to support any other backend by:
//
// 1. Implement the defined interfaces
//
// 2. Registering the library via
//  pc.Register("myBackend", NewMyBackendCanvas)
//
// 3. Instantiating your backend with
//  pc.NewCanvas("myBackend", width, height, writer)
package purplecrayon

import (
	"fmt"
	core "github.com/jordan-bonecutter/purplecrayon/core"
	"io"
)

type Point = core.Point
type RGB = core.RGB
type RGBA = core.RGBA
type Reference = core.Reference

// The main interface for drawing with purplecrayon.
// Only one operation should be open at a time, undefined behavior occurs otherwise.
type Canvas interface {
	Referrable

	// Return the width of the canvas
	Width() float64

	// Return the height of the canvas
	Height() float64

	// Draw a rectangle inside the canvas
	Rect() Rect

	// Draw a circle inside the canvas
	Circle() Circle

	// Draw an ellipse inside the canvas
	Ellipse() Ellipse

	// Begin a path.
	Path() Path

	// Create a linear gradient
	LinearGradient() LinearGradient

	// Create a group which draws to a subgroup in this canvas
	Group() Group

	// Create a mask which draws to a submask in this canvas
	Mask() Mask
}

// A group of objects which can be transformed and painted together.
// All transform / paint operations must be finished before calling Start.
// The canvas returned by start is a "derived canvas", in that it eventually
// refers to the same canvas from which the group was created but through a
// layer of indirection, namely the group.
type Group interface {
	Transformable
	Paintable
	Open() Canvas
}

// A group of objects which together may be used as a compositing mask.
// Works similarly to the Group.
type Mask interface {
	Transformable
	Paintable
	Open() Canvas
}

// Any object which can be transformed
// Note: The transform must be finished before working on other attributes!
type Transformable interface {
	Transform() Transform
}

type Transform interface {
	Translate(delta Point) Transform
	Scale(float64) Transform
	Rotate(degrees float64) Transform

	// Finish the current transformation.
	// This method is not called Close because it doesn't return a reference!
	// In other words, transforms are not referrable.
	Finish()
}

// Any object which can be painted
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
	CompositeMask(Reference)
}

// Referrables are objects which can have references.
type Referrable interface {
	// Closes the object and creates a reference for other objects to refer to it.
	Close() Reference
}

// Equivalent to an SVG path
type Path interface {
	Referrable
	Transformable
	Paintable

	Cursor() Cursor

	// Finish the path.
	Close() Reference
}

type Cursor interface {
	// Move the cursor to an absolute point.
	MoveTo(Point) Cursor

	// Move the cursor to a relative point.
	MoveToRel(Point) Cursor

	// Draw a line to an absolute point.
	LineTo(Point) Cursor

	// Draw a line to a relative point.
	LineToRel(Point) Cursor

	// Draw an absolute quadratic bezier curve.
	QuadTo(Point, Point) Cursor

	// Draw a relative quadratic bezier curve.
	QuadToRel(Point, Point) Cursor

	// Zips up the path to where it started.
	Zip() Cursor

	// Finishes the cursor movement.
	// Not close because it returns no reference!
	// The reference is held in the enclosing Path
	Finish()
}

// A rectangle
type Rect interface {
	Referrable
	Transformable
	Paintable
	TopLeft(p Point) Rect
	Width(float64) Rect
	Height(float64) Rect
}

// A circle
type Circle interface {
	Referrable
	Transformable
	Paintable
	Center(Point) Circle
	Radius(float64) Circle
}

type Ellipse interface {
	Referrable
	Transformable
	Paintable
	Center(p Point) Ellipse
	Radii(p Point) Ellipse
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
	SetLine(p0, p1 Point) LinearGradient

	// the object by which stops are added to the gradient.
	GradientStops() GradientStops
}

type GradientStop interface {
	RGB(RGB) GradientStop
	RGBA(RGBA) GradientStop
	Position(float64) GradientStop
	Finish()
}

type GradientStops interface {
	// Add a gradient stop to the parent gradient
	Stop() GradientStop

	// Finish adding stops to the gradient
	Finish()
}

// A function which serves as a driver for a purplecrayon backend
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
