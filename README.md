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

## Currently Implemented Types
** The Idea is to eventually implement all of the following: **

Basic Type | Current Implementation Status | Unit Test Status | Benchmark Status
-----------|-------------------------------|------------------|-------------------
uint8 | no | no | no
uint16 | no | no | no
uint32 | yes | yes | no
uint64 | yes | no | no
int8 | no | no | no
int16 | no | no | no
int32 | yes | yes | no
int64 | yes | yes | no
float32 | no | no | no
float64 | no | no | no
complex64* | no | no | no
complex128* | no | no | no
byte | no | no | no
rune | no | no | no


* These may or not ever be implemented

