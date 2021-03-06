package main

import (
	pc "github.com/jordan-bonecutter/purplecrayon"
	_ "github.com/jordan-bonecutter/purplecrayon/svg"
	"log"
	"os"
)

func main() {
	fi, err := os.Create("path.svg")
	if err != nil {
		log.Fatalf("Failed opening test file: %s\n", err.Error())
	}
	defer fi.Close()

	canv, err := pc.NewCanvas("svg", 1920, 1080, fi)
	if err != nil {
		log.Fatalf("Failed creating canvas: %s\n", err.Error())
	}

	path := canv.Path()
	path.Cursor().MoveTo(pc.Point{500, 500}).
		LineToRel(pc.Point{100, 100}).
		QuadToRel(pc.Point{-50, -50}, pc.Point{-70, -100}).
		Zip().
		Finish()
	path.StrokeWidth(5)
	path.StrokeRGB(pc.RGB{200, 0, 200})
	path.FillTransparent()
	path.Close()

	canv.Close()
}
