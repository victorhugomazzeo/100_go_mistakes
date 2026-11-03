// Package stringset is the fix: same code as util, but the name says what it
// provides, so the functions drop the StringSet suffix (stringset.New).
package stringset

import (
	"maps"
	"slices"
)

func New(values ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func Sort(set map[string]struct{}) []string {
	return slices.Sorted(maps.Keys(set))
}
