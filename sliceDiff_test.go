package sliceDiff

import "testing"

// GenerateInt64Slice generates random slices of int64's from 0 to s 
// incrementing in incSize steps
func GenerateInt64Slice(s int64, incSize int64) []int64 {
	var n int64
    var testSlice []int64
	for n = 0; n < s; n += incSize {
		testSlice = append(testSlice,n)
	}
	return(testSlice)
}

func BenchmarkInt64Diff_XXL(b *testing.B) {
    testThree := GenerateInt64Slice(32768,4)
    testFour := GenerateInt64Slice(65535,2)
    for i := 0; i < b.N; i++ {
		Int64SliceDiff(testThree,testFour)
	}
}

func BenchmarkInt64Diff_S(b *testing.B) {
	testOne := []int64{1,2,3,4,5,6,7,8,9}
	testTwo := []int64{1,2,3,5,6,7,8,10}
    for i := 0; i < b.N; i++ {
		Int64SliceDiff(testOne,testTwo)
	}
}

func BenchmarkInt64Diff_M(b *testing.B) {
	testOne := GenerateInt64Slice(256,1)
	testTwo := GenerateInt64Slice(768,6)
	for i := 0; i < b.N; i++ {
		Int64SliceDiff(testOne,testTwo)
	}
}

func BenchmarkInt64Diff_L(b *testing.B) {
	testOne := GenerateInt64Slice(4096,2)
	testTwo := GenerateInt64Slice(8192,4)
	for i := 0; i < b.N; i++ {
		Int64SliceDiff(testOne,testTwo)
	}
}
func BenchmarkInt64Diff_XL(b *testing.B) {
	testOne := GenerateInt64Slice(16384,2)
	testTwo := GenerateInt64Slice(7000,4)
	for i := 0; i < b.N; i++ {
		Int64SliceDiff(testOne,testTwo)
	}
}

func BenchmarkStringDiff_S(b *testing.B) {
	testOne := []string {"example","one","two","three","four","five"}
	testTwo := []string {"one","four","five"}
	for i := 0; i < b.N; i++ {
		StringSliceDiff(testOne,testTwo)
	}
}

func TestInt64Diff_Basic(t *testing.T) {
	testOne := []int64 {1,2,3,5,6,7}
	testTwo := []int64 {1,2,3,4,5,6,7,8}
    diff := Int64SliceDiff(testOne,testTwo)
	diffLen := len(diff)
    if (diffLen != 2) {
		t.Fatal("Basic Int64 Test failed.")
	}
}

func TestStringDiff_Basic(t *testing.T) {
	testOne := []string {"example","test","one","two","three"}
	testTwo := []string {"example","test","one","three"}

    diff := StringSliceDiff(testOne,testTwo)
	diffLen := len(diff)
	if (diffLen != 1) {
		t.Fatal("Basic String Test Failed")
	}
}

func TestInt32Diff_Basic(t *testing.T) {
	testOne := []int32 {120,35,228,67,9,0,32760}
	testTwo := []int32 {120,35,228,67,32760}
    diff := Int32SliceDiff(testOne,testTwo)
    diffLen := len(diff)
	if (diffLen != 2) {
		t.Fatal("Basic Int32 Test failed.")
	}
}

