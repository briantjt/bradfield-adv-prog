package main

import (
	"fmt"
	"math"
	"unsafe"
)

type Point struct {
	x int
	y int
}

func main() {
	// String
	s := "Hello World"
	sPtr := unsafe.Pointer(&s)
	sLenLoc := unsafe.Add(sPtr, 8)
	length := *(*uint64)(sLenLoc)
	fmt.Printf("Length of string a: %d\n", length)

	// Point
	point := Point{x: 5, y: 10}
	pPtr := unsafe.Pointer(&point)
	yLocation := unsafe.Add(pPtr, unsafe.Sizeof(point.x))
	yVal := *(*int)(yLocation)
	fmt.Printf("Value of y: %d\n", yVal)

	// Int array
	intList := []int{1, 2, 3, 4, 5, 6, 7}
	listPtr := unsafe.Pointer(&intList)
	arrayLocation := unsafe.Pointer(*(*uintptr)(listPtr))
	lenLocation := unsafe.Add(listPtr, 8)
	list_len := *(*uint64)(lenLocation)
	sum := 0
	for i := uint64(0); i < list_len; i++ {
		loc := unsafe.Add(arrayLocation, uintptr(i)*unsafe.Sizeof(int(0)))
		sum += *(*int)(loc)
	}
	fmt.Printf("Sum of array: %d\n", sum)

	// Hashmap
	hashmap := make(map[int]int)
	for i := 0; i < 10; i++ {
		hashmap[i] = i
	}
	hashPtr := unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&hashmap)))
	bucketsOffset := unsafe.Sizeof(int(0)) + 2*unsafe.Sizeof(uint8(0)) + unsafe.Sizeof(uint16(0)) + unsafe.Sizeof(uint32(0))
	numBuckets := 1 << (*(*uint8)(unsafe.Add(hashPtr, unsafe.Sizeof(int(0))+unsafe.Sizeof(uint8(0)))))
	fmt.Printf("Num buckets: %d\n", numBuckets)
	bucketPtr := unsafe.Add(hashPtr, bucketsOffset)
	buckets := unsafe.Pointer(*(*uintptr)(bucketPtr))
	bucketSize := unsafe.Sizeof([8]uint8{}) + 2*unsafe.Sizeof([8]int{}) + unsafe.Sizeof(uintptr(0))
	highest := math.MinInt
	for i := 0; i < numBuckets; i++ {
		bucket := unsafe.Add(buckets, uintptr(i)*bucketSize)
		// Iterate through tophash
		for j := 0; j < 8; j++ {
			tophashPtr := unsafe.Add(bucket, uintptr(j)*unsafe.Sizeof(uint8(0)))
			tophash := *(*uint8)(tophashPtr)
			if tophash == 0 {
				break
			}
			valuePtr := unsafe.Add(bucket, unsafe.Sizeof([8]uint8{})+unsafe.Sizeof([8]int{})+uintptr(j)*unsafe.Sizeof(int(0)))
			value := *(*int)(valuePtr)
			if value > highest {
				highest = value
			}
			fmt.Printf("Tophash value: %d, value: %d\n", tophash, value)
		}
	}
	fmt.Printf("Highest value found: %d\n", highest)
}
