package main

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
import "C"

func main() {
	s := "Hello World!"
	cs := C.CString(s)
	C.free(unsafe.Pointer(cs))
	fmt.Println(s)
}
