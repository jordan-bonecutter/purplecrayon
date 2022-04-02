package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
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
