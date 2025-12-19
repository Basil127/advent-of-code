package dsu

import (
	"errors"
)

type DSU[T comparable] struct {
	parent []int
	values map[T]int
	rank   []int
}

func NewDSU[T comparable](values []T) *DSU[T] {
	size := len(values)
	parent := make([]int, size)
	valuesMap := make(map[T]int, size)
	for i, v := range values {
		valuesMap[v] = i
	}
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &DSU[T]{
		parent: parent,
		values: valuesMap,
		rank:   rank,
	}
}

/*
Find returns the root index of the set containing the given value.
An error is returned if the value is not found.
Flattening is applied to optimize future queries.
*/
func (d *DSU[T]) Find(value T) (int, error) {
	index, exists := d.values[value]
	if !exists {
		return -1, errors.New("value not found in DSU")
	}
	if d.parent[index] != index {
		d.parent[index], _ = d.FindByIndex(d.parent[index])
	}
	return d.parent[index], nil
}

func (d *DSU[T]) FindByIndex(index int) (int, error) {
	var err error

	if index < 0 || index >= len(d.parent) {
		return -1, errors.New("index out of bounds")
	}

	if d.parent[index] != index {
		d.parent[index], err = d.FindByIndex(d.parent[index])
		if err != nil {
			return -1, err
		}
	}

	return d.parent[index], nil
}

/*
Union merges the sets containing value1 and value2.
If they are already in the same set, no action is taken.
*/
func (d *DSU[T]) Union(value1, value2 T) error {
	root1, err1 := d.Find(value1)
	root2, err2 := d.Find(value2)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	if root1 != root2 {
		if d.rank[root1] < d.rank[root2] {
			d.parent[root1] = root2
		} else if d.rank[root1] > d.rank[root2] {
			d.parent[root2] = root1
		} else {
			d.parent[root2] = root1
			d.rank[root1]++
		}
	}

	return nil
}

func (d *DSU[T]) Connected(value1, value2 T) (bool, error) {
	root1, err1 := d.Find(value1)
	root2, err2 := d.Find(value2)
	if err1 != nil {
		return false, err1
	}
	if err2 != nil {
		return false, err2
	}
	return root1 == root2, nil
}

/*
Returns a map of root indices to the list of values in that set.
*/
func (d *DSU[T]) GetSets() map[int][]T {
	rootSet := make(map[int][]T)
	for val := range d.values {
		root, err := d.Find(val)
		if err != nil {
			continue
		}
		rootSet[root] = append(rootSet[root], val)
	}
	return rootSet
}

func (d *DSU[T]) CountSets() int {
	sets := d.GetSets()
	return len(sets)
}
