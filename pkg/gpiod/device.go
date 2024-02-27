package gpiod

// #include <gpiod.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type device struct {
	path      string
	nativeRef *C.struct_gpiod_chip
}

func NewDevice(path string) *device {
	return &device{
		path: path,
	}
}

func (d *device) Open() error {
	charC := C.CString(d.path)
	defer C.free(unsafe.Pointer(charC))
	var nativeRef *C.struct_gpiod_chip = C.gpiod_chip_open(charC)

	// https://stackoverflow.com/questions/56352863/c-null-type-in-cgo
	if nativeRef == nil {
		return fmt.Errorf("%s failed: NULL returned", "gpiod_chip_open")
	}
	d.nativeRef = nativeRef
	return nil
}

func (d *device) Close() {
	C.gpiod_chip_close(d.nativeRef)
}
