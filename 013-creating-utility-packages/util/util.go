// Package util is the mistake: the name says nothing about what it provides,
// so each function must carry the context itself (NewStringSet, SortStringSet).
package util

import (
	"maps"
	"slices"
)

func NewStringSet(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func SortStringSet(set map[string]struct{}) []string {
	return slices.Sorted(maps.Keys(set))
}
