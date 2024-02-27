package gpiod

// #include <gpiod.h>
// #include <stdlib.h>
import "C"
import "fmt"

type device struct {
	path      string
	nativeRef C.gpiod_chip
}

func NewDevice(path string) *device {
	return &device{
		path: path,
	}
}

func (d *device) Open() error {
	charC := C.CString(d.path)
	defer C.free(charC)
	structC := C.gpiod_chip_open(charC)
	if structC == C.NULL {
		return fmt.Errorf("%s failed: NULL returned", "gpiod_chip_open")
	}
	d.nativeRef = structC
	return nil
}
