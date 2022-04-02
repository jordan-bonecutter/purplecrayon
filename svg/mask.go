package svg

func makeMask(svg *svg, canv canvas) tree {
	return makeTree(svg, canv, "mask")
}
