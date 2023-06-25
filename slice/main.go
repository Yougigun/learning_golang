package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	// print the address of the slice(variable) on stack
	fmt.Printf("Address of slice: %p\n", &s)

	// Get the pointer, length, and capacity of the slice on stack
	ptr := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
	length := len(s)
	capacity := cap(s)

	fmt.Printf("Pointer to data on heap: %p\n", ptr)
	fmt.Printf("Length: %d\n", length)
	fmt.Printf("Capacity: %d\n", capacity)

	// Get the memory size of the slice header on the stack
	headerSize := unsafe.Sizeof(s)
	fmt.Printf("Slice value size on stack: %d bytes\n", headerSize)
}
