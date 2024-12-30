package combo

import (
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	src := []int{1, 2, 3}
	want := [][]int{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}}

	actual := make([][]int, 0, len(want))
	for perm := range Permutations(src) {
		actual = append(actual, perm)
	}

	slices.SortFunc(want, slices.Compare)
	slices.SortFunc(actual, slices.Compare)

	equal := slices.EqualFunc(actual, want, slices.Equal)

	if !equal {
		t.Errorf("Permutations(%v) = %v, want %v", src, actual, want)
	}
}

func TestUniquePermutations(t *testing.T) {
	src := []int{1, 1, 2}
	want := [][]int{{1,1,2}, {1,2,1}, {2,1,1}}

	actual := make([][]int, 0, len(want))
	for perm := range UniquePermutations(src) {
		actual = append(actual, perm)
	}

	slices.SortFunc(want, slices.Compare)
	slices.SortFunc(actual, slices.Compare)

	equal := slices.EqualFunc(actual, want, slices.Equal)

	if !equal {
		t.Errorf("UniquePermutations(%v) = %v, want %v", src, actual, want)
	}
}
