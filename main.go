package main

import (
	"flag"
	"github.com/deadsy/sdfx/render"
	"github.com/matthewjwhite/jolly-gen/tree"
)

func main() {
	// Define CLI parameters.
	layerPtr := flag.Int("layers", 5, "Number of layers in tree")
	heightPtr := flag.Float64("height", 50, "Height of tree")
	basePtr := flag.Float64("base", 30, "Radius of tree base")
	curvePtr := flag.Float64("curve", 5, "Layer curve factor")
	filenamePtr := flag.String("filename", "out.stl", "Filename")
	flag.Parse()

	// Prepare tree SDF for rendering.
	treeSDF := tree.Tree{*layerPtr, *heightPtr, *basePtr, *curvePtr}.SDF()

	// 300 is a decent amount of mesh cells - this defines, for a lack of a
	// better word, how smooth the rendering is.
	render.RenderSTL(treeSDF, 300, *filenamePtr)
}
