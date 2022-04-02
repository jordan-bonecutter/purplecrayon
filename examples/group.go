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
	gradient.AddRGBStop(0.1, pc.RGB{
		255, 120, 0,
	})
	gradient.AddRGBStop(0.8, pc.RGB{
		255, 0, 200,
	})
	ref := gradient.Close()

  g := canv.Group()
  g.Rotate(45)
  gCanv := g.Open()

	r := gCanv.Rect()
	r.Width(20)
	r.Height(50)
	r.FillRGB(pc.RGB{255, 0, 0})
	r.Translate(pc.Point{10, 10})
	r.Close()

	circle := gCanv.Circle()
	circle.Center(pc.Point{50, 50})
	circle.Radius(20)
	circle.Fill(ref)
	circle.Close()

  gCanv.Close()

	canv.Close()
}
