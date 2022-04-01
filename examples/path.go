package main

import (
  "os"
  pc "github.com/jordan-bonecutter/purplecrayon"
  "log"
)

func main() {
  fi, err := os.Create("path.svg")
  if err != nil {
    log.Fatalf("Failed opening test file: %s", err.Error())
  }
  defer fi.Close()

  canv := pc.NewSVGCanvas(1920, 1080, fi)

  cursor := canv.Cursor()
  cursor.MoveTo(pc.Point{500, 500})
  cursor.LineToRel(pc.Point{100, 100})
  cursor.QuadToRel(pc.Point{-50, -50}, pc.Point{-70, -100})
  cursor.Zip()
  cursor.StrokeWidth(5)
  cursor.StrokeRGB(pc.RGB{200, 0, 200})
  cursor.FillTransparent()
  cursor.Close()

  canv.Close()
}

