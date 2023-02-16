package util

import (
	"github.com/imdario/mergo"
)

func MergeSlices[T comparable](withUnique bool, slices ...[]T) []T {
	resultMerged := make([]T, 0)
	resultMergedAndUnique := make([]T, 0)

	for _, s := range slices {
		resultMerged = append(resultMerged, s...)
	}

	if !withUnique {
		return resultMerged
	}

	for _, v := range resultMerged {
		if !CheckSliceContain(resultMergedAndUnique, v) {
			resultMergedAndUnique = append(resultMergedAndUnique, v)
		}
	}

	return resultMergedAndUnique
}

func MergeStructs[T comparable](withOverride bool, structs ...T) T {
	dst := new(T)

	if !withOverride {
		for _, s := range structs {
			mergo.Merge(dst, s)
		}

		return *dst
	}

	for _, s := range structs {
		mergo.Merge(dst, s, mergo.WithOverride)
	}

	return *dst
}
