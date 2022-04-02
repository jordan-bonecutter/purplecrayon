package main

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	_ "github.com/jordan-bonecutter/purplecrayon/svg"
	"log"
	"os"
)

func main() {
	fi, err := os.Create("group.svg")
	if err != nil {
		log.Fatalf("Failed opening test file: %s\n", err.Error())
	}
	defer fi.Close()

	canv, err := pc.NewCanvas("svg", 1920, 1080, fi)
	if err != nil {
		log.Fatalf("Failed creating canvas: %s\n", err.Error())
	}

	gradient := canv.LinearGradient()
	gradient.SetLine(pc.Point{
		X: 0, Y: 0,
	}, pc.Point{
		X: 1, Y: 1,
	})
  gradientStops := gradient.GradientStops()
	gradientStops.Stop().Position(0.1).RGB(pc.RGB{
		255, 120, 0,
	}).Finish()
	gradientStops.Stop().Position(0.8).RGB(pc.RGB{
		255, 0, 200,
	}).Finish()
  gradientStops.Finish()
	ref := gradient.Close()

	g := canv.Group()
  g.Transform().Rotate(45).Finish()
	gCanv := g.Open()

	r := gCanv.Rect()
	r.Width(20)
	r.Height(50)
	r.FillRGB(pc.RGB{255, 0, 0})
  r.Transform().Translate(pc.Point{10, 10}).Finish()
	r.Close()

	circle := gCanv.Circle()
	circle.Center(pc.Point{50, 50})
	circle.Radius(20)
	circle.Fill(ref)
	circle.Close()

	gCanv.Close()

	canv.Close()
}
