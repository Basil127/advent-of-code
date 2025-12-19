package dsu_test

import (
	"testing"

	"github.com/basil127/advent-of-code/dsu"
)

func TestDSUFindBasic(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	d := dsu.NewDSU(values)
	found := make(map[string]bool)
	for _, v := range values {
		index, err := d.Find(v)
		if err != nil {
			t.Errorf("Unexpected error for Find(%s): %v", v, err)
		}
		found[v] = (index >= 0)
	}
	if len(found) != len(values) {
		t.Errorf("Expected to find all values, found: %v", found)
	}
}

func TestDSUUnionAndFind(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	d := dsu.NewDSU(values)

	// Call union to create 2 connected components: (1,2,3, 4) and (5)
	err := d.Union(1, 2)
	if err != nil {
		t.Errorf("Unexpected error during Union(1, 2): %v", err)
	}
	err = d.Union(3, 4)
	if err != nil {
		t.Errorf("Unexpected error during Union(3, 4): %v", err)
	}
	err = d.Union(2, 3)
	if err != nil {
		t.Errorf("Unexpected error during Union(2, 3): %v", err)
	}

	// Ensure root of 1,2,3,4 are the same
	root1, err := d.Find(1)
	if err != nil {
		t.Errorf("Unexpected error during Find(1): %v", err)
	}
	root2, err := d.Find(2)
	if err != nil {
		t.Errorf("Unexpected error during Find(2): %v", err)
	}
	root3, err := d.Find(3)
	if err != nil {
		t.Errorf("Unexpected error during Find(3): %v", err)
	}
	root4, err := d.Find(4)
	if err != nil {
		t.Errorf("Unexpected error during Find(4): %v", err)
	}
	if root1 != root2 || root2 != root3 || root3 != root4 {
		t.Errorf("Expected 1, 2, 3, and 4 to be in the same set, got roots %d, %d, %d, %d", root1, root2, root3, root4)
	}
}

func TestDSUFindNonExistent(t *testing.T) {
	values := []string{"x", "y", "z"}
	d := dsu.NewDSU(values)
	_, err := d.Find("a")
	if err == nil {
		t.Errorf("Expected error for Find on non-existent value, got nil")
	}
}

func TestDSUUnionNonExistent(t *testing.T) {
	values := []int{10, 20, 30}
	d := dsu.NewDSU(values)
	err := d.Union(10, 40)
	if err == nil {
		t.Errorf("Expected error for Union with non-existent value, got nil")
	}
}

func TestDSUFindByIndexOutOfBounds(t *testing.T) {
	values := []int{100, 200, 300}
	d := dsu.NewDSU(values)
	_, err := d.FindByIndex(5)
	if err == nil {
		t.Errorf("Expected error for FindByIndex with out-of-bounds index, got nil")
	}
	_, err = d.FindByIndex(-1)
	if err == nil {
		t.Errorf("Expected error for FindByIndex with negative index, got nil")
	}
}

func TestDSUUnionSameSet(t *testing.T) {
	values := []string{"p", "q", "r"}
	d := dsu.NewDSU(values)
	err := d.Union("p", "q")
	if err != nil {
		t.Errorf("Unexpected error during Union(p, q): %v", err)
	}
	err = d.Union("p", "q") // Union again
	if err != nil {
		t.Errorf("Unexpected error during Union(p, q) second time: %v", err)
	}
}

func TestDSUConnected(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	d := dsu.NewDSU(values)
	err := d.Union(1, 2)
	if err != nil {
		t.Errorf("Unexpected error during Union(1, 2): %v", err)
	}
	err = d.Union(2, 3)
	if err != nil {
		t.Errorf("Unexpected error during Union(2, 3): %v", err)
	}
	err = d.Union(4, 5)
	if err != nil {
		t.Errorf("Unexpected error during Union(4, 5): %v", err)
	}

	res, err := d.Connected(1, 3)
	if err != nil {
		t.Errorf("Unexpected error during Connected(1, 3): %v", err)
	}
	if !res {
		t.Errorf("Expected 1 and 3 to be connected")
	}

	res, err = d.Connected(1, 4)
	if err != nil {
		t.Errorf("Unexpected error during Connected(1, 4): %v", err)
	}
	if res {
		t.Errorf("Expected 1 and 4 to be not connected")
	}
	res, err = d.Connected(5, 6)
	if err != nil {
		t.Errorf("Unexpected error during Connected(5, 6): %v", err)
	}
	if res {
		t.Errorf("Expected 5 and 6 to be not connected")
	}
}
