// Package purplecrayon is a drawing library written in Go.
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

// Point is a point on a canvas
type Point = core.Point

// RGB is a color represented in RGB space
type RGB = core.RGB

// RGBA is a color represented in RGBA space
type RGBA = core.RGBA

// Reference is a reference to an object
type Reference = core.Reference

// Canvas is the main interface for drawing with purplecrayon.
// Only one operation should be open at a time, undefined behavior occurs otherwise.
type Canvas interface {
	Referrable

	// Set parameters for the current backend.
	// Returns nil iff no error occurred.
	Configure(key string, value interface{}) error

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

// Group is a collection of objects which can be transformed and painted together.
// All transform / paint operations must be finished before calling Start.
// The canvas returned by start is a "derived canvas", in that it eventually
// refers to the same canvas from which the group was created but through a
// layer of indirection, namely the group.
type Group interface {
	Transformable
	Paintable
	Open() Canvas
}

// Mask objects contain a group of objects which together may be used as a compositing mask.
// Works similarly to the Group.
type Mask interface {
	Transformable
	Paintable
	Open() Canvas
}

// Transformable objects are any object which can be transformed.
// Note: The transform must be finished before working on other attributes!
type Transformable interface {
	Transform() Transform
}

// Transform objects are instances of a transformations.
type Transform interface {
	Translate(delta Point) Transform
	Scale(float64) Transform
	Rotate(degrees float64) Transform

	// Finish the current transformation.
	// This method is not called Close because it doesn't return a reference!
	// In other words, transforms are not referrable.
	Finish()
}

// Paintable objects are any object which can be painted onto a canvas.
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

// Referrable objects are objects which may return a reference to themselves for others to use.
type Referrable interface {
	// Closes the object and creates a reference for other objects to refer to it.
	Close() Reference
}

// Path objects are generic shapes which can be drawn with cursors.
type Path interface {
	Referrable
	Transformable
	Paintable

	Cursor() Cursor

	// Finish the path.
	Close() Reference
}

// Cursor objects allow arbitrary drawing of shapes on a canvas.
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

// Rect objects are drawn as rectangles on a canvas.
type Rect interface {
	Referrable
	Transformable
	Paintable
	TopLeft(p Point) Rect
	Width(float64) Rect
	Height(float64) Rect
}

// Circle objects are drawn as circles on a canvas.
type Circle interface {
	Referrable
	Transformable
	Paintable
	Center(Point) Circle
	Radius(float64) Circle
}

// Ellipse objects are drawn as ellpises on a canvas.
type Ellipse interface {
	Referrable
	Transformable
	Paintable
	Center(p Point) Ellipse
	Radii(p Point) Ellipse
}

// LinearGradient objects return a reference to a linear gradient which can be used
// as a fill or stroke for Paintable objects.
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

// GradientStop objects represent a fixed color stop in a gradient.
type GradientStop interface {
	RGB(RGB) GradientStop
	RGBA(RGBA) GradientStop
	Position(float64) GradientStop
	Finish()
}

// GradientStops objects allow creation of GradientStop objects for a Gradient.
type GradientStops interface {
	// Add a gradient stop to the parent gradient
	Stop() GradientStop

	// Finish adding stops to the gradient
	Finish()
}

// Driver functions return Canvas objects which implement the Canvas interface.
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

var errDriverNotFound = fmt.Errorf("Driver not found")

// NewCanvas creates a Canvas for the given regsitered Driver
func NewCanvas(driver string, width, height float64, w io.Writer) (c Canvas, err error) {
	if registeredDrivers == nil {
		err = errDriverNotFound
	}

	if lib, ok := registeredDrivers[driver]; ok {
		c = lib(width, height, w)
	} else {
		err = errDriverNotFound
	}
	return
}
