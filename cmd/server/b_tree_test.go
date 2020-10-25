package main

import (
	"reflect"
	"testing"
)

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

func TestLeafPageInsert(t *testing.T) {
	tables := []struct {
		leaf   leaf
		input  leafPage
		output leafPage
	}{
		{
			leaf{
				key:   "key",
				value: "value",
			},
			leafPage{make(map[string]leaf, 10)},
			leafPage{make(map[string]leaf, 10)},
		},
	}

	for _, table := range tables {
		table.output.leaves[table.leaf.key] = table.leaf
		table.input.insert(table.leaf.key, table.leaf.value)
		if !reflect.DeepEqual(table.input, table.output) {
			t.Errorf("LeafPage is %+v, wants %+v", table.input, table.output)
		}
	}
}

func TestLeafPageDelete(t *testing.T) {
	tables := []struct {
		leaf   leaf
		input  leafPage
		output leafPage
	}{
		{
			leaf{
				key:   "key",
				value: "value",
			},
			leafPage{make(map[string]leaf, 10)},
			leafPage{make(map[string]leaf, 10)},
		},
	}

	for _, table := range tables {
		table.input.leaves[table.leaf.key] = table.leaf
		table.input.delete(table.leaf.key)
		if !reflect.DeepEqual(table.input, table.output) {
			t.Errorf("LeafPage is %+v, wants %+v", table.input, table.output)
		}
	}
}

func TestLeafPageSearch(t *testing.T) {
	tables := []struct {
		input  leafPage
		output leaf
		err    error
	}{
		{leafPage{leaves: map[string]leaf{},
		}, leaf{
			key:   "",
			value: "",
		}, ErrorKeyDoesNotExist,
		},
		{leafPage{leaves: map[string]leaf{"one": {
			"one", "two",
		}},
		}, leaf{
			key:   "one",
			value: "two",
		}, nil,
		},
	}

	for _, table := range tables {
		leaf, err := table.input.search(table.output.key)
		if err != table.err {
			t.Errorf("Error is %s, wants %s", err, table.err)
		}
		if leaf != table.output {
			t.Errorf("Leaf is %#v, want %#v", leaf, table.output)
		}
	}
}

func TestPointerPageInsert(t *testing.T) {
	tables := []struct {
		name string
		key string
		value string
		input pointerPage
		output pointerPage
	}{
		{

		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			table.input.insert(table.key, table.value)
			if !reflect.DeepEqual(table.input, table.output) {
				t.Errorf("pointerPage is %+v, wants %+v", table.input, table.output)
			}
		})
	}
}