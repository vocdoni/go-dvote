package censustreefactory

import (
	"fmt"

	"github.com/vocdoni/arbo"
	"go.vocdoni.io/dvote/censustree"
	"go.vocdoni.io/dvote/censustree/arbotree"
	"go.vocdoni.io/dvote/censustree/gravitontree"
)

const (
	// TreeTypeArboBlake2b defines a tree type that uses arbo with Blake2b
	// hash function. Is thought for being used when computation speed is
	// important.
	TreeTypeArboBlake2b = iota
	// TreeTypeArboPoseidon defines a tree type that uses arbo with Poseidon
	// hash function. Is thought for being used when zkSNARK compatibility
	// is needed (with circomlib).
	TreeTypeArboPoseidon
	// TreeTypeGraviton defines a tree type that uses graviton tree.
	TreeTypeGraviton
)

// NewCensusTree creates a merkle tree using the given storage and hash
// function.  Note that each tree should use an entirely separate namespace for
// its database keys.
func NewCensusTree(treeType int, name, storageDir string) (censustree.Tree, error) {
	var err error
	var tree censustree.Tree
	switch treeType {
	case TreeTypeArboBlake2b:
		if tree, err =
			arbotree.NewTree(name, storageDir, arbo.HashFunctionBlake2b); err != nil {
			return nil, err
		}
	case TreeTypeArboPoseidon:
		if tree, err =
			arbotree.NewTree(name, storageDir, arbo.HashFunctionPoseidon); err != nil {
			return nil, err
		}
	case TreeTypeGraviton:
		if tree, err = gravitontree.NewTree(name, storageDir); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unrecognized tree type (%d)", treeType)
	}

	return tree, nil
}
