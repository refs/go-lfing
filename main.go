package main

import (
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

func main() {
	hi := "Hello, world"

	// eli5
	addr := fmt.Sprintf("%v", &hi)

	// eli5
	ui64, err := strconv.ParseUint(addr, 0, 64)
	if err != nil {
		log.Panic(err)
	}

	// eli5
	var ptr uintptr = uintptr(ui64)

	printStringFromUnderlyingArray(ptr)
}

// eli5
func printStringFromUnderlyingArray(ptr uintptr) {
	// one could argue using reflect -> StringHeader: https://golang.org/pkg/reflect/#StringHeader
	// but:
	/*
		StringHeader is the runtime representation of a string. It cannot be
		used safely or portably and its representation may change in a later
		release. Moreover, the Data field is not sufficient to guarantee the
		data it references will not be garbage collected, so programs must
		keep a separate, correctly typed pointer to the underlying data.
	*/

	// so here we're essentially initializing a pointer with the address of our underlying []byte
	p := unsafe.Pointer(ptr)

	// and dereferencing it (https://www.golang-book.com/books/intro/8)
	fmt.Println(*(*string)(p))
}
