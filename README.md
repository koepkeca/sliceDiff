# SliceDiff

SliceDiff is a [go] (http://www.golang.org)  library to find the difference between slices.

## Documentation

* [API Reference] (https://godoc.org/github.com/koepkeca/sliceDiff)

## Library Status

This library is still under developemnt. The functions Int32SliceDiff, Int64SliceDiff, and StringSliceDiff are fully functional and have testing / benchmarking libraries

## Installation

go get gtihub.com/koepkeca/sliceDiff

## Notes

This library is not intended for massive slices. You can modify the benchmarking files located in sliceDiff_test.go to see how long operations take with this implementation. I personally wouldn't use it for any slices containing more than 30,000 elements.

