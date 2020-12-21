package tree

import (
	"github.com/deadsy/sdfx/sdf"
)

// Tree represents a Christmas tree.
type Tree struct {
	Layers              int
	Height, Base, Curve float64
}

// SDF creates the SDF representation of Tree.
func (t Tree) SDF() sdf.SDF3 {
	// Slice of tree layers to unify before returning.
	cones := make([]sdf.SDF3, 0, t.Layers)

	// Decrement factors for each tree layer, found through experimentation.
	decrementHeight := ((t.Height / float64(t.Layers)) / 2)
	decrementBase := ((t.Base / float64(t.Layers)) / 2)

	// Define temporary variables for loop.
	var cone sdf.SDF3
	currHeight := t.Height
	currBase := t.Base
	currLift := 0.0

	// Create each layer and lift upwards by amount we decreased height.
	for i := 0; i < t.Layers; i++ {
		cone, _ = sdf.Cone3D(currHeight, currBase, 0, t.Curve)
		cone = sdf.Transform3D(cone, sdf.Translate3d(sdf.V3{0, 0, currLift}))
		cones = append(cones, cone)

		currHeight -= decrementHeight
		currBase -= decrementBase
		currLift += decrementHeight
	}

	return sdf.Union3D(cones...)
}
