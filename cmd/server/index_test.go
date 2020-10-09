package main

import (
	"testing"
)

func TestNode_insert(t *testing.T) {
	tables := []struct {
		input  *node
		output *node
		error  error
	}{
		{&node{
			key:   "a",
			size:  4,
			left:  nil,
			right: nil,
		}, &node{
			key:   "a",
			size:  4,
			left:  nil,
			right: nil,
		}, keyAlreadyExists},

		{&node{
			key:   "b",
			size:  4,
			left:  nil,
			right: nil,
		}, &node{
			key:  "a",
			size: 4,
			left: nil,
			right: &node{
				key:   "b",
				size:  4,
				left:  nil,
				right: nil,
			},
		}, nil},

		{&node{
			key:   "b",
			size:  4,
			left:  nil,
			right: nil,
		}, &node{
			key:  "c",
			size: 4,
			left: nil,
			right: &node{
				key:   "b",
				size:  4,
				left:  nil,
				right: nil,
			},
		}, nil},
	}

	for _, table := range tables {
		n := &node{
			key:   "a",
			size:  0,
			left:  nil,
			right: nil,
		}

		err := n.insert(table.input)
		if err != table.error {
			t.Errorf("Error: '%v', does not match wanted error: '%v'", err, table.error)
			// if the key already exists there is no point in checking if an insert has been done
		} else if table.error != keyAlreadyExists {
			if n.left == table.output.left {
				t.Fatalf("The left node is %v, not %v", table.input, table.output)
			}

			if n.right == table.output.right {
				t.Fatalf("The left node is %v, not %v", table.input, table.output)
			}

			if n.left.key != table.input.key {
				t.Errorf("The input key '%s' is not %s", table.input.key, n.left.key)
			}
		}
	}
}

func TestCompareStrings(t *testing.T) {
	tables := []struct {
		first  string
		second string
		want   int
	}{
		{"a", "a", equal},
		{"a", "b", larger},
		{"c", "b", smaller},
		{"hello", "world", larger},
	}

	for _, table := range tables {
		returnValue := compareStrings(table.first, table.second)
		if returnValue != table.want {
			t.Errorf("%s and %s returns %d, wants %d", table.first, table.second, returnValue, table.want)
		}
	}
}

func BenchmarkCompareStringsShort(b *testing.B) {
	a := "short"
	c := "horts"
	for i := 0; i < b.N; i++ {
		compareStrings(a, c)
	}
}

func BenchmarkCompareStringsMedium(b *testing.B) {
	a := "abitlongererthenfirst"
	c := "cbitlongererthenfirst"
	for i := 0; i < b.N; i++ {
		compareStrings(a, c)
	}
}

func BenchmarkCompareStringsLong(b *testing.B) {
	a := "TC7xRB0rh7IdDhQiX33GePYRC9xufBklQya2lYHMGV3rNqp0SPJWLJAYAHGc7pOSTTUsNOvpRsOruogkAFdPLpDi2DV2hD6vbpRY"
	c := "wmefxkYGHa11to6PlcP0KCnQWomSLhFbKauxD8KdLrowYKjB3nVJnSEi9xnqqZa8sGyaV17aBCrEEC3sQqQhUskLQNvLd7zNif5L"
	for i := 0; i < b.N; i++ {
		compareStrings(a, c)
	}
}
