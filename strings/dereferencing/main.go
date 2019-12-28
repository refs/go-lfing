package dereferencing

import (
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

// Show does the thing
func Show() {
	hi := "Hello, world"

	// store the address of hi into `addr` for future dereferencing
	addr := fmt.Sprintf("%v", &hi)
	fmt.Printf("[1]\ncontents: %v\naddress: %v\n---\n", hi, &hi)

	// need to convert from string to uint64
	ui64, err := strconv.ParseUint(addr, 0, 64)
	if err != nil {
		log.Panic(err)
	}

	// all we need for unsafe to get the value of a pointer is a uint64:
	// uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
	var ptr uintptr = uintptr(ui64)

	// print
	printStringFromUnderlyingArray(ptr)
}

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

	// so here we're essentially initializing a pointer with the address of our underlying byte array ([]byte)
	p := unsafe.Pointer(ptr)

	// and dereferencing it (https://www.golang-book.com/books/intro/8) knowing the type
	// this step bypases golang type safety.
	// void warranty. don't do.
	fmt.Println(*(*string)(p))
}
