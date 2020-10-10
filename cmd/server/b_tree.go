// b_tree implements a b tree data structure.
package main

import "errors"

const (
	equal = iota
	larger
	smaller
)

var ErrorKeyDoesNotExist = errors.New("the key does not exist")
var ErrorKeyTooLarge = errors.New("the key is too large")
var ErrorValueTooLarge = errors.New("the value is too large")

// page can be a pointerPage or leafPage
type page interface {
	insert(key, value string)
	delete(key string)
	search(key string) (leaf, error)
}

// page represents the root page or an internal page.
type pointerPage struct {
	smaller page
	key     string
	larger  page
}

func (pp *pointerPage) insert(key, value string) {
	switch compareStrings(pp.key, key) {
	case smaller:
	case equal:
		if pp.smaller != nil {
			pp.smaller.insert(key, value)
		} else {
			pp.smaller = &leafPage{leaves: make(map[string]leaf, 1024)}
			pp.smaller.insert(key, value)
		}
		break
	case larger:
		if pp.larger != nil {
			pp.larger.insert(key, value)
		} else {
			pp.larger = &leafPage{leaves: make(map[string]leaf, 1024)}
			pp.larger.insert(key, value)
		}
		break
	}

}

func (pp *pointerPage) delete(key string) {
	switch compareStrings(pp.key, key) {
	case smaller:
	case equal:
		if pp.smaller != nil {
			 pp.smaller.delete(key)
		}
		break
	case larger:
		if pp.larger != nil {
			pp.larger.delete(key)
		}
		break
	}
}

func (pp *pointerPage) search(key string) (leaf, error) {
	var leaf leaf
	var err error
	switch compareStrings(pp.key, key) {
	case smaller:
	case equal:
		if pp.smaller != nil {
			leaf, err = pp.smaller.search(key)
		}
		break
	case larger:
		if pp.larger != nil {
			leaf, err = pp.larger.search(key)
		}
		break
	}

	return leaf, err
}

// leafPage represents a page with a sequential set of leaves.
type leafPage struct {
	leaves map[string]leaf
}

func (lp *leafPage) insert(key, value string) {
	lp.leaves[key] = leaf{key, value}
}

func (lp *leafPage) delete(key string) {
	delete(lp.leaves, key)
}

func (lp *leafPage) search(key string) (leaf, error) {
	var err error
	leaf := lp.leaves[key]
	if leaf.key != "" {
		return leaf, err
	} else {
		err = ErrorKeyDoesNotExist
	}

	return leaf, err
}

// leaf is a simple struct with a key, value, and metadata.
type leaf struct {
	// The length of the key and value should be limited to guarantee a
	// maximum leaf size. Which in turn guarantees a maximum leafPage size.
	key   string
	value string
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
