package svg

import (
	"fmt"
	"github.com/jordan-bonecutter/purplecrayon/core"
	"sort"
)

func sortedMapIter(m map[string]string, f func(k, v string)) {
	sortedKeys := make([]string, len(m))
	idx := 0
	for k := range m {
		sortedKeys[idx] = k
		idx++
	}

	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		f(k, m[k])
	}
}

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
