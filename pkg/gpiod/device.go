package gpiod

//#include <stdlib.h>
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
	charC := C.CBytes(d.path)
	defer C.free(charC)
	structC := C.gpiod_chip_open((*C.char)(charC))
	if structC == C.NULL {
		return fmt.Errorf("%s failed: NULL returned", "gpiod_chip_open")
	}
	d.nativeRef = structC
	return nil
}
