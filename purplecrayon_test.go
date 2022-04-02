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
	gradient.AddRGBAStop(0.1, pc.RGBA{255, 120, 0, 127})
	gradient.AddRGBAStop(0.9, pc.RGBA{255, 0, 120, 255})
	ref := gradient.Close()

	circle := canv.Circle()
	circle.Center(pc.Point{1920 / 2, 1080 / 2})
	circle.Radius(400)
	circle.Fill(ref)
	circle.Close()
}
