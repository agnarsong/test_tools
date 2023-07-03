package layer2

import (
	"crypto/sha256"
	"encoding/hex"
)

type MerkleTree struct {
	Leaves [][]byte
	Layers [][][32]byte
}

func NewMerkleTree(leaves [][]byte) *MerkleTree {
	tree := &MerkleTree{Leaves: leaves, Layers: make([][][32]byte, len(leaves)+1)}
	tree.init()
	tree.build()
	return tree
}

func (tree *MerkleTree) init() {
	for i := range tree.Leaves {
		var hash [32]byte
		copy(hash[:], tree.Leaves[i])
		tree.Layers[0] = append(tree.Layers[0], hash)
	}
}

func (tree *MerkleTree) build() {
	for i := 0; i < len(tree.Layers)-1; i++ {
		if len(tree.Layers[i])%2 != 0 {
			tree.Layers[i] = append(tree.Layers[i], sha256.Sum256(make([]byte, 32)))
		}

		var layer [][32]byte
		for j := 0; j < len(tree.Layers[i]); j += 2 {
			var concat [64]byte
			copy(concat[:32], tree.Layers[i][j][:])
			copy(concat[32:], tree.Layers[i][j+1][:])
			hash := sha256.Sum256(concat[:])
			layer = append(layer, hash)
		}
		tree.Layers[i+1] = layer
	}
}

func (tree *MerkleTree) Root() []byte {
	return tree.Layers[len(tree.Layers)-1][0][:]
}

func (tree *MerkleTree) Proof(leaf []byte) [][]byte {
	var index int
	for i := range tree.Layers[0] {
		if string(leaf) == string(tree.Layers[0][i][:]) {
			index = i
			break
		}
	}

	var proof [][]byte
	for i, layer := range tree.Layers {
		if i == len(tree.Layers)-1 {
			break
		}
		if index%2 == 0 && index+1 < len(layer) {
			proof = append(proof, layer[index+1][:])
		} else if index%2 != 0 {
			proof = append(proof, layer[index-1][:])
		}
		index /= 2
	}

	return proof
}

func (tree *MerkleTree) VerifyProof(leaf []byte, proof [][]byte, root []byte) bool {
	hash := sha256.Sum256(leaf)
	for _, sibling := range proof {
		var concat [64]byte
		if sibling[0] < hash[0] {
			copy(concat[:32], sibling)
			copy(concat[32:], hash[:])
		} else {
			copy(concat[:32], hash[:])
			copy(concat[32:], sibling)
		}
		hash = sha256.Sum256(concat[:])
	}
	return hex.EncodeToString(hash[:]) == hex.EncodeToString(root)
}
