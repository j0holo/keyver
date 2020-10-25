// b_tree implements a b tree data structure.
package main

import "errors"

const (
	maxLeavesLength      = 100
	maxPointerPageLength = 256
	maxKeySize           = 512        //bytes
	maxValueSize         = 1024 << 10 // 1 megabyte
)

const (
	equal = iota
	larger
	smaller
)

var ErrorKeyDoesNotExist = errors.New("the key does not exist")
var ErrorKeyTooLarge = errors.New("the key is too large")
var ErrorValueTooLarge = errors.New("the value is too large")
var ErrorPageIsFull = errors.New("the underlying page is full")

// page can be a pointerPage or leafPage
type page interface {
	insert(key, value string) error
	delete(key string)
	search(key string) (leaf, error)
}

// page represents the root page or an internal page.
type pointerPage struct {
	keys []string
	// pagePointers always has one element more then the number of keys.
	pagePointers []page
}

func (pp *pointerPage) insert(key, value string) error {
	for i, pk := range pp.keys {
		switch compareStrings(pk, key) {
		case smaller:
		case equal:
			err := pp.pagePointers[i].insert(key, value)
			if err == ErrorPageIsFull {
				// TODO: split the underlying page and add the new pages to this page.
				pp.splitChild(i)
			} else if err != nil {
				return err
			}
			break
		case larger:
			if len(pp.keys) == i {
				err := pp.pagePointers[len(pp.pagePointers)-1].insert(key, value)
				if err == ErrorPageIsFull {

				} else if err != nil {
					return err
				} else {

				}
			}
			break
		}
	}
	return nil
}

func (pp *pointerPage) splitChild(i int) {
	// This gives an error that pointerPage does not implement the page interface.
	if _, ok := pp.pagePointers[i].(pointerPage); ok {
		pp.pagePointers[i].keys[]
	}

}

func (pp *pointerPage) delete(key string) {
	for i, pk := range pp.keys {
		switch compareStrings(pk, key) {
		case smaller:
		case equal:
			pp.pagePointers[i].delete(key)
			break
		case larger:
			if len(pp.keys) == i {
				pp.pagePointers[len(pp.pagePointers)-1].delete(key)
			}
			break
		}
	}
}

func (pp *pointerPage) search(key string) (leaf, error) {
	var leaf leaf
	var err error
	for i, pk := range pp.keys {
		switch compareStrings(pk, key) {
		case smaller:
		case equal:
			leaf, err = pp.pagePointers[i].search(key)
			break
		case larger:
			if len(pp.keys) == i {
				leaf, err = pp.pagePointers[len(pp.pagePointers)-1].search(key)
			}
			break
		}
	}

	return leaf, err
}

// leafPage represents a page with a sequential set of leaves.
type leafPage struct {
	leaves map[string]leaf
}

func (lp *leafPage) insert(key, value string) error {
	lp.leaves[key] = leaf{key, value}
	return nil
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
