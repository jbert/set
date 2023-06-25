package set

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	a := assert.New(t)
	s := New[int]()

	a.True(s.IsEmpty(), "Nil set is empty")
	a.Equal([]int{}, s.ToList(), "Get back nil list for empty set")

	a.True(NewFromList([]int{1, 2, 3}).Equals(NewFromList([]int{3, 2, 1})), "order doesn't matter")

	xs := []int{1, 3, 5, 8}
	x := NewFromList(xs)
	a.Equal(x.Size(), len(xs), "Set same length as list")
	l := x.ToList()
	sort.Ints(l)
	a.Equal(l, xs, "ToList gives the right list")

	x.InsertList(xs)
	a.Equal(x.Size(), len(xs), "Inserting same list twice doesn't change set")

	dups := []int{3, 5, 3, 7, 3, 11}
	undups := []int{3, 5, 7, 11}
	d := NewFromList(dups)
	a.Equal(d.Size(), len(undups), "Sets don't contain duplicates")

	ud := NewFromList(undups)
	a.True(d.Equals(ud), "Set of dups and undups equal")

	a.True(x.Equals(x), "Set equals itself")

	u := s.Union(s)
	a.True(s.Equals(u), "Self union doesn't change the set")

	i := s.Intersect(s)
	a.True(s.Equals(i), "Self intersect doesn't change the set")

	ys := []int{3, 5, 7, -1, 100}
	y := NewFromList(ys)
	a.Equal(x.Intersect(y), NewFromList([]int{3, 5}), "intersect")
	a.Equal(x.Union(y), NewFromList([]int{1, 3, 5, 7, 8, -1, 100}), "union")
}

func TestTake(t *testing.T) {
	a := assert.New(t)
	l := []int{1, 2, 3, 4, 5}
	s := NewFromList(l)

	expectedSize := len(l)
	for i := 0; i < len(l); i++ {
		a.False(s.IsEmpty(), "Set is not yet empty")
		t.Logf("Iteration: %d: S: %s", i, s)
		_, err := s.Take()
		expectedSize--
		a.Nil(err, "No error Taking from non-empty")
		a.Equal(s.Size(), expectedSize, fmt.Sprintf("Iteration %d - check set size", i))
	}
	_, err := s.Take()
	a.Equal(ErrIsEmpty, err, "Can't take from empty set")
	a.Equal(s.Size(), 0, fmt.Sprintf("Set is now empty"))
	a.True(s.IsEmpty(), "Set is now empty")
}
