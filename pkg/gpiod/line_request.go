package gpiod

// #include <gpiod.h>
import "C"
import "fmt"

type lineRequest struct {
	nativeRef *C.struct_gpiod_line_request
}

func newLineRequest(d *device, lc *lineConfig) (*lineRequest, error) {

	var nativeRef *C.struct_gpiod_line_request = C.gpiod_chip_request_lines(
		d.nativeRef,
		nil, /* NULL / default request config*/
		lc.nativeRef,
	)
	if nativeRef == nil {
		return nil, fmt.Errorf("%s failed: NULL returned", "gpiod_chip_request_lines")
	}
	return &lineRequest{
		nativeRef: nativeRef,
	}, nil
}

func (lr *lineRequest) free() {
	C.gpiod_line_request_release(lr.nativeRef)
}

func lineRequestSetValueForSingleOffset(d *device, lc *lineConfig) error {

	req, err := newLineRequest(d, lc)
	if err != nil {
		return err
	}
	defer req.free()

	var resultC C.int = C.gpiod_line_request_set_value(
		req.nativeRef,
		C.uint(lc.lineSet.offset),
		C.enum_gpiod_line_value(lc.lineSet.value),
	)
	if resultC == C.int(-1) {
		return fmt.Errorf("%s failed: -1 returned", "gpiod_line_request_set_value")
	}
	if resultC == C.int(0) {
		return nil
	}
	return fmt.Errorf("%s returned something unexpected", "gpiod_line_request_set_value")
}

func lineRequestGetValueForSingleOffset(d *device, lc *lineConfig) (lineValue, error) {

	req, err := newLineRequest(d, lc)
	if err != nil {
		return LineValueError, err
	}
	defer req.free()

	var resultC C.enum_gpiod_line_value = C.gpiod_line_request_get_value(
		req.nativeRef,
		C.uint(lc.lineSet.offset),
	)
	if resultC == C.GPIOD_LINE_VALUE_ERROR {
		return LineValueError, fmt.Errorf("%s failed: -1 returned", "gpiod_line_request_get_value")
	}
	if resultC == C.GPIOD_LINE_VALUE_ACTIVE {
		return LineValueActive, nil
	}
	if resultC == C.GPIOD_LINE_VALUE_INACTIVE {
		return LineValueInactive, nil
	}
	return LineValueError, fmt.Errorf("%s returned something unexpected", "gpiod_line_request_get_value")
}
