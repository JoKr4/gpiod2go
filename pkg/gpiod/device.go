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
	lineSet   map[uint]*lineSettings
}

func NewDevice(path string) *device {
	return &device{
		path:    path,
		lineSet: make(map[uint]*lineSettings),
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

func (d *device) AddLine(offset uint, direction lineDirection, value lineValue) error {

	newLine, err := NewLineSettings(offset, direction, value)
	if err != nil {
		return err
	}

	// TODO already added?

	d.lineSet[offset] = newLine

	return nil
}
