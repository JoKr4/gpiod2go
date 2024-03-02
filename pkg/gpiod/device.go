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
	for _, l := range d.lineSet {
		l.Free()
	}
	C.gpiod_chip_close(d.nativeRef)
}

func (d *device) AddLine(offset uint, direction lineDirection) error {

	newLine, err := NewLineSettings(offset, direction)
	if err != nil {
		return err
	}

	// TODO already added?
	d.lineSet[offset] = newLine

	return nil
}

func (d *device) SetLineValue(offset uint, value lineValue) error {

	// TODO catch if offset not found
	err := d.lineSet[offset].SetOutputValue(value)
	if err != nil {
		return err
	}

	config, err := NewLineConfig()
	if err != nil {
		return err
	}
	defer config.Free()

	err = config.ApplyLineSettingsForSingleOffset(d.lineSet[offset])
	if err != nil {
		return err
	}

	err = lineRequestSetValueForSingleOffset(d, config)
	if err != nil {
		return err
	}

	return nil
}

func (d *device) GetLineValue(offset uint) (lineValue, error) {

	config, err := NewLineConfig()
	if err != nil {
		return LineValueError, err
	}
	defer config.Free()

	// TODO catch if offset not found
	err = config.ApplyLineSettingsForSingleOffset(d.lineSet[offset])
	if err != nil {
		return LineValueError, err
	}

	value, err := lineRequestGetValueForSingleOffset(d, config)
	if err != nil {
		return LineValueError, err
	}

	return value, nil
}

func (d *device) GetLineDirection(offset uint) (lineDirection, error) {

	lineRef := C.gpiod_chip_get_line_info(d.nativeRef, C.uint(offset))
	if lineRef == nil {
		return LineDirectionUnknown, fmt.Errorf("%s failed: NULL returned", "gpiod_chip_get_line_info")
	}

	resultC := C.gpiod_line_info_get_direction(lineRef)

	return NewLineDirection(resultC), nil
}
