package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./build -Wl,-rpath,./build -lhello
#include "hello.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	str := C.hello()
	fmt.Println(C.GoString(str))
}
