package main

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	_ "github.com/jordan-bonecutter/purplecrayon/svg"
	"log"
	"os"
)

func main() {
	fi, err := os.Create("mask.svg")
	if err != nil {
		log.Fatalf("Failed opening test file: %s\n", err.Error())
	}
	defer fi.Close()

	canv, err := pc.NewCanvas("svg", 1920, 1080, fi)
	if err != nil {
		log.Fatalf("Failed creating canvas: %s\n", err.Error())
	}

	grad := canv.LinearGradient()
	grad.SetLine(pc.Point{0, 0}, pc.Point{1, 1})
	stops := grad.GradientStops()
	stops.Stop().Position(0).RGB(pc.RGB{255, 0, 0}).Finish()
	stops.Stop().Position(1).RGB(pc.RGB{0, 0, 255}).Finish()
	stops.Finish()
	ref := grad.Close()

	m := canv.Mask().Open()
	circle := m.Circle()
	circle.Center(pc.Point{1920 / 2, 1080 / 2})
	circle.Radius(400)
	circle.FillRGB(pc.RGB{255, 255, 255})
	circle.Close()
	maskRef := m.Close()

	rect := canv.Rect()
	rect.TopLeft(pc.Point{100, 100})
	rect.Width(1500)
	rect.Height(900)
	rect.Fill(ref)
	rect.CompositeMask(maskRef)
	rect.Close()

	canv.Close()
}
