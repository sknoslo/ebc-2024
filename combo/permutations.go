package combo

import (
	"iter"
	"slices"
)

func Permutations[T any](src []T) iter.Seq[[]T] {
	src = slices.Clone(src)
	return func(yield func([]T) bool) {
		var permute func(idx int)
		permute = func(idx int) {
			if idx == len(src) {
				yield(slices.Clone(src))
			}

			for i := idx; i < len(src); i++ {
				src[idx], src[i] = src[i], src[idx]
				permute(idx + 1)
				src[idx], src[i] = src[i], src[idx]
			}
		}
		permute(0)
	}
}

// based on: http://paddy3118.blogspot.com/2012/07/a-permutation-journey.html
func UniquePermutations[T comparable](src []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(src) == 1 {
			yield(slices.Clone(src))
		} else {
			dir := 1
			curr := src[0]
			for perm := range UniquePermutations(src[1:]) {
				maxIndex := slices.Index(perm, curr)
				if maxIndex == -1 {
					maxIndex = len(perm)
				}
				if dir == 1 {
					for i := maxIndex; i >= 0; i-- {
						items := make([]T, 0, len(src))
						items = append(items, perm[:i]...)
						items = append(items, curr)
						items = append(items, perm[i:]...)
						yield(items)
					}
				} else {
					for i := 0; i <= maxIndex; i++ {
						items := make([]T, 0, len(src))
						items = append(items, perm[:i]...)
						items = append(items, curr)
						items = append(items, perm[i:]...)
						yield(items)
					}
				}
				dir = -dir
			}
		}
	}
}
