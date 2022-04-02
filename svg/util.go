package svg

import (
  "github.com/jordan-bonecutter/purplecrayon/core"
  "fmt"
)

func svgRGB(color core.RGB) string {
  return fmt.Sprintf("rgb(%d, %d, %d)",
    color.R,
    color.G,
    color.B,
  )
}

func svgRGBA(color core.RGBA) string {
  return fmt.Sprintf("rgba(%d, %d, %d, %d)",
    color.R,
    color.G,
    color.B,
    color.A,
  )
}

func svgRef(ref core.Reference) string {
  return fmt.Sprintf("url(#%s)", string(ref))
}
