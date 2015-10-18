// Use of this source code is governed by a BSD-style license
// which can be found in the LICENSE file
// Package sliceDiff provides functions to determine the difference
// between two slices. There are tests and some benchmark utilities
// THIS IS INTENDED ONLY FOR "SMALL" SLICES. Use on slices greater than
// ~ 30k entries should not be used.

package sliceDiff

// Int64SliceDiff returns the int64 values that are not in
// both source slices
func Int64SliceDiff(sliceOne []int64, sliceTwo []int64) []int64 {
	var diff []int64
	for i := 0; i < 2; i++ {
		for _, s1 := range sliceOne {
			found := false
			for _, s2 := range sliceTwo {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff,s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Int32SliceDiff returns the int32 values that are not in
// both soruce slices
func Int32SliceDiff(sliceOne []int32, sliceTwo []int32) []int32 {
	var diff []int32
	for i := 0; i < 2; i++ {
		for _, s1 := range sliceOne {
			found := false
			for _, s2 := range sliceTwo {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff,s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// StringSliceDiff returns teh string values that are not in
// both source slices
func StringSliceDiff(sliceOne []string, sliceTwo []string) []string {
	var diff []string
	for i := 0; i < 2; i++ {
		for _, s1 := range sliceOne {
			found := false
			for _, s2 := range sliceTwo {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff,s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}