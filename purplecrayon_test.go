package purplecrayon_test

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	_ "github.com/jordan-bonecutter/purplecrayon/svg"
	"log"
	"os"
)

func Example() {
	canv, err := pc.NewCanvas("svg", 1920, 1080, os.Stdout)
	if err != nil {
		log.Fatalf("Failed opening canvas: %s\n", err)
	}
	defer canv.Close()

	rect := canv.Rect()
	rect.Width(100)
	rect.Height(100)
	rect.FillRGB(pc.RGB{255, 0, 0})
	rect.TopLeft(pc.Point{400, 500})
	rect.Close()

	gradient := canv.LinearGradient()
	gradient.SetLine(pc.Point{0, 0}, pc.Point{1, 1})
	stops := gradient.GradientStops()
	stops.Stop().Position(0.1).RGBA(pc.RGBA{255, 120, 0, 127}).Finish()
	stops.Stop().Position(0.9).RGBA(pc.RGBA{255, 0, 120, 255}).Finish()
	stops.Finish()
	ref := gradient.Close()

	circle := canv.Circle()
	circle.Center(pc.Point{1920 / 2, 1080 / 2})
	circle.Radius(400)
	circle.Fill(ref)
	circle.Close()

	// Output:
	// <svg id="pcobj-0" width="1920.000000" height="1080.000000" xmlns="http://www.w3.org/2000/svg">
	// <rect id="pcobj-1" width="100.000000" height="100.000000" fill="rgb(255,0,0)" x="400.000000" y="500.000000"/>
	// <linearGradient id="pcobj-2" x1="0.000000" y1="0.000000" x2="1.000000" y2="1.000000">
	// <stop id="pcobj-3" offset="10.000000%" stop-color="rgba(255,120,0,127)"/>
	// <stop id="pcobj-4" offset="90.000000%" stop-color="rgba(255,0,120,255)"/></linearGradient>
	// <circle id="pcobj-5" cx="960.000000" cy="540.000000" r="400.000000" fill="url(#pcobj-2)"/></svg>
}
