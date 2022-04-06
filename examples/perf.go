// This file provides a basic performance test against another popular svg drawing library.
// I tried to make the test as unbiased as possible, drawing the same SVG as I would with both libraries.
// Depending on the parameters, different libraries will win.
// The main area where purplecrayon wins is with long paths.
// This is because the path needs no intermediary (namely, a string which is appended to on each cursor move),
// it goes straight to the writer!
// For simpler objects, like rectangles, circles, and ellipses svgo wins.
// I believe this is due to the more complex function call structure of purplecrayon as compred to svgo.
//
// To alleviate this, I will add some methods to configure multiple parameters at once.
//
// Some notes on the test:
// I try to be as fair as possible by doing the following 3 things:
//   1. Attempt to draw identically rendering SVGs with both libraries
//   2. Call GC immediately before the drawing begins so that GC has a fair chance at negatively affecting both.
//   3. Wrap the drawing functions and pass them off to the same timer, they both must perform all the same tasks within said function.
// I'm no expert in performance analysis so I may be doing some things wrong here!
package main

import (
	"bytes"
	"fmt"
	svgo "github.com/ajstarks/svgo"
	pc "github.com/jordan-bonecutter/purplecrayon"
	_ "github.com/jordan-bonecutter/purplecrayon/svg"
	"io"
	"os"
	"runtime"
	"time"
)

const (
	WARMUP_RUNS   = 10
	NUM_RECTS     = 1000
	NUM_GRADIENTS = 100
	NUM_ELLIPSES  = 100
	NUM_PATHS     = 1000
	PATH_LEN      = 100
	CANV_WIDTH    = 1920
	CANV_HEIGHT   = 1080
)

func main() {
	// Do a few warmup runs
	for i := 0; i < WARMUP_RUNS; i++ {
		perf(drawSVGO)
		perf(drawPC)
	}

	// The actual run!
	svgoTime, svgoBytes := perf(drawSVGO)
	pcTime, pcBytes := perf(drawPC)

	fmt.Printf("Took %v to draw in purplecrayon with %d bytes\n", pcTime, pcBytes.Len())
	fmt.Printf("Took %v to draw in svgo with %d bytes\n", svgoTime, svgoBytes.Len())

	fi, _ := os.Create("pc.svg")
	io.Copy(fi, pcBytes)
	fi.Close()

	fi, _ = os.Create("svgo.svg")
	io.Copy(fi, svgoBytes)
	fi.Close()
}

func perf(draw func(*bytes.Buffer)) (duration time.Duration, drawing *bytes.Buffer) {
	buf := new(bytes.Buffer)
	runtime.GC()

	start := time.Now()
	draw(buf)
	end := time.Now()

	return end.Sub(start), buf
}

func svgoRGB(r, g, b uint8) string {
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

func svgoFillStyle(color string) string {
	return fmt.Sprintf(`fill="%s"`, color)
}

func svgoStrokeStyle(color string) string {
	return fmt.Sprintf(`stroke="%s"`, color)
}

func drawSVGO(buf *bytes.Buffer) {
	canv := svgo.New(buf)
	canv.Start(1920, 1080)
	var idCount uint64
	nextId := func() string {
		idCount++
		return fmt.Sprintf("obj-%d", idCount)
	}

	for nr := 0; nr < NUM_RECTS; nr++ {
		canv.Rect(20+(nr/2), 30+(nr/2), 100, 200, svgoFillStyle(svgoRGB(200, 100, 50+uint8(nr/10))))
	}

	for ng := 0; ng < NUM_GRADIENTS; ng++ {
		canv.LinearGradient(nextId(), 0, 0, 100, 100, []svgo.Offcolor{
			{
				Offset:  10,
				Color:   svgoRGB(200, 100, 50),
				Opacity: 1.0,
			}, {
				Offset:  90,
				Color:   svgoRGB(50, 100, 200),
				Opacity: 1.0,
			},
		})
	}

	for ne := 0; ne < NUM_ELLIPSES; ne++ {
		canv.Ellipse(300, 400, 100, 50, svgoFillStyle("none"))
	}

	for np := 0; np < NUM_PATHS; np++ {
		pathD := "M 100 200"
		for pi := 0; pi < PATH_LEN; pi++ {
			pathD += fmt.Sprintf(" q %.0f %.0f %.0f %.0f", 25.0, 25.0, 33.0, 33.0)
		}
		pathD += " z"
		canv.Path(pathD, svgoStrokeStyle(svgoRGB(200, 50, 32)))
	}
	canv.End()
}

func drawPC(buf *bytes.Buffer) {
	canv, _ := pc.NewCanvas("svg", CANV_WIDTH, CANV_HEIGHT, buf)

	for nr := 0; nr < NUM_RECTS; nr++ {
		rect := canv.Rect().TopLeft(pc.Point{20 + float64(nr/2), 30 + float64(nr/2)}).Width(100).Height(200)
		rect.FillRGB(pc.RGB{200, 100, 50 + uint8(nr/10)})
		rect.Close()
	}

	for ng := 0; ng < NUM_GRADIENTS; ng++ {
		grad := canv.LinearGradient()
		grad.SetLine(pc.Point{0, 0}, pc.Point{1, 1})
		stops := grad.GradientStops()
		stops.Stop().RGB(pc.RGB{200, 100, 50}).Position(0.1).Finish()
		stops.Stop().RGB(pc.RGB{50, 100, 200}).Position(0.9).Finish()
		stops.Finish()
		grad.Close()
	}

	for ne := 0; ne < NUM_ELLIPSES; ne++ {
		ellipse := canv.Ellipse().Center(pc.Point{300, 400}).Radii(pc.Point{100, 50})
		ellipse.FillTransparent()
		ellipse.Close()
	}

	for np := 0; np < NUM_PATHS; np++ {
		path := canv.Path()
		cursor := path.Cursor().MoveTo(pc.Point{100, 200})
		for pi := 0; pi < PATH_LEN; pi++ {
			cursor.QuadToRel(pc.Point{25, 25}, pc.Point{33, 33})
		}
		cursor.Zip()
		cursor.Finish()
		path.StrokeRGB(pc.RGB{200, 50, 32})
		path.Close()
	}
	canv.Close()
}
