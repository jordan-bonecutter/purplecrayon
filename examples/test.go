package main

import (
  "os"
  _ "github.com/jordan-bonecutter/purplecrayon/svg"
  pc "github.com/jordan-bonecutter/purplecrayon"
  "log"
)

func main() {
  fi, err := os.Create("test.svg")
  if err != nil {
    log.Fatalf("Failed opening test file: %s\n", err.Error())
  }
  defer fi.Close()

  canv, err := pc.NewCanvas("svg", 100, 100, fi)
  if err != nil {
    log.Fatalf("Failed creating canvas: %s\n", err.Error())
  }

  r := canv.Rect()
  r.Width(20)
  r.Height(50)
  r.FillRGB(pc.RGB{255, 0, 0})
  r.Translate(pc.Point{10, 10})
  r.Close()

  circle := canv.Circle()
  circle.Center(pc.Point{50, 50})
  circle.Radius(20)
  circle.FillRGB(pc.RGB{200, 0, 200})
  circle.Close()

  canv.Close()
}

