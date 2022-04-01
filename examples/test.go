package main

import (
  "os"
  pc "github.com/jordan-bonecutter/purplecrayon"
  "log"
)

func main() {
  fi, err := os.Create("test.svg")
  if err != nil {
    log.Fatalf("Failed opening test file: %s", err.Error())
  }
  defer fi.Close()

  canv := pc.NewSVGCanvas(100, 100, fi)

  r := canv.Rect()
  r.Width(20)
  r.Height(50)
  r.FillRGB(pc.RGB{255, 0, 0})
  r.Translate(pc.Point{10, 10})
  r.Close()

  canv.Close()
}

