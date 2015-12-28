# SliceDiff

SliceDiff is a [go] (http://www.golang.org)  library to find the difference between slices.

## Documentation

* [API Reference] (https://godoc.org/github.com/koepkeca/sliceDiff)

## Library Status

This library is still under development. Function names will **not** change so you can feel free to use this for testing / production. [You can view the currently implemented primitive types here.] (#types)  I will add expanded benchmark tests as time permits.

## Notes on DRY, Interfaces and Generics in Go

I do understand that this library violates the DRY principle, however, after researching reflect, go generate and interfaces, it still seems to be the best option to just hammer these out. 
I wanted to have a generic function that accepted an interface but the issue of the return type was still a problem.
## Installation

go get gtihub.com/koepkeca/sliceDiff

## Notes

This library is not intended for massive slices. You can modify the benchmarking files located in sliceDiff_test.go to see how long operations take with this implementation. I personally wouldn't use it for any slices containing more than 30,000 elements.

## Currently Implemented Type [types] ##

** The Idea is to eventually implement all of the following: **

Basic Type | Current Implementation Status | Unit Test Status | Benchmark Status
-----------|-------------------------------|------------------|-------------------
uint8 | yes | yes | no
uint16 | yes | yes | no
uint32 | yes | yes | no
uint64 | yes | yes | no
int8 | yes | yes | no
int16 | yes | yes | no
int32 | yes | yes | partial
int64 | yes | yes | yes
float32 | yes | yes | no
float64 | yes | yes | no
complex64 | yes | yes | no
complex128 | yes | yes | no
byte | no | no | no
rune | no | no | no
string | yes | yes | partial
