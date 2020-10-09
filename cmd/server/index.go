package main

import "errors"

type node struct {
	key   string
	size  uint32
	left  *node
	right *node
}

const (
	equal   = 0
	larger  = 1
	smaller = 2
)

var keyAlreadyExists = errors.New("the key already exists")

// Insert inserts a new node into the tree.
func (cn *node) insert(n *node) error {
	var err error
	switch compareStrings(cn.key, n.key) {
	case larger:
		if cn.left == nil {
			cn.left = n
			break
		}
		err = cn.left.insert(n)
		break
	case smaller:
		if cn.right == nil {
			cn.right = n
			break
		}
		err = cn.right.insert(n)
		break
	default:
		err = keyAlreadyExists
	}
	return err
}

// Fetch a with the total size of the complete tree path.
func (cn *node) fetch() (*node, error) {
	return &node{}, nil
}

// Update the node with the given key.
func (cn *node) update() error {
	return nil
}

// compareStrings compares b to a, so the integer reflects the binary order of b to a.
// If a is 0b01 and b is 0b11 then the output should be 'larger' (1). B is a larger
// binary number then a.
func compareStrings(a, b string) int {
	if a < b {
		return larger
	} else if a > b {
		return smaller
	}
	return equal

}
