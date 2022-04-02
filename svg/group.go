package svg

func makeGroup(svg *svg, canv canvas) tree {
  return makeTree(svg, canv, "g")
}
