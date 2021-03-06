// Package sliceDiff provides functions to determine the difference
// between two slices. There are tests and limited benchmarks for this package.
//
// THIS IS INTENDED ONLY FOR "SMALL" SLICES.
//
// Using this package on slices > 30,000 entries can get slow.. I would
// suggest benchmarking the number / type of slices you intend to use
// in sliceDiff_test.go *before* using on slices with a large number
// of entries.
//
// Use of this source code is governed by a BSD-style license
// which can be found in the LICENSE file
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Uint64SliceDiff returns the uint64 values that are not in
// both source slices
func Uint64SliceDiff(sliceOne []uint64, sliceTwo []uint64) []uint64 {
	var diff []uint64
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
				diff = append(diff, s1)
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Uint32SliceDiff returns the uint32 values that are not in
// both soruce slices
func Uint32SliceDiff(sliceOne []uint32, sliceTwo []uint32) []uint32 {
	var diff []uint32
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Int16SliceDiff returns the int16 values that are not in
// both soruce slices
func Int16SliceDiff(sliceOne []int16, sliceTwo []int16) []int16 {
	var diff []int16
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Uint16SliceDiff returns the uint16 values that are not in
// both soruce slices
func Uint16SliceDiff(sliceOne []uint16, sliceTwo []uint16) []uint16 {
	var diff []uint16
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Int8SliceDiff returns the int8 values that are not in
// both soruce slices
func Int8SliceDiff(sliceOne []int8, sliceTwo []int8) []int8 {
	var diff []int8
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Uint8SliceDiff returns the uint8 values that are not in
// both soruce slices
func Uint8SliceDiff(sliceOne []uint8, sliceTwo []uint8) []uint8 {
	var diff []uint8
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Float32SliceDiff returns the float32 values that are not in
// both soruce slices
func Float32SliceDiff(sliceOne []float32, sliceTwo []float32) []float32 {
	var diff []float32
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Float64SliceDiff returns the float64 values that are not in
// both soruce slices
func Float64SliceDiff(sliceOne []float64, sliceTwo []float64) []float64 {
	var diff []float64
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Complex64SliceDiff returns the complex64 values that are not in
// both soruce slices
func Complex64SliceDiff(sliceOne []complex64, sliceTwo []complex64) []complex64 {
	var diff []complex64
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// Complex128SliceDiff returns the complex128 values that are not in
// both soruce slices
func Complex128SliceDiff(sliceOne []complex128, sliceTwo []complex128) []complex128 {
	var diff []complex128
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}

// StringSliceDiff returns the string values that are not in
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
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			sliceOne, sliceTwo = sliceTwo, sliceOne
		}
	}
	return diff
}
