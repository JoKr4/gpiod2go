package gpiod

// #include <gpiod.h>
import "C"
import "fmt"

type lineSettings struct {
	nativeRef *C.struct_gpiod_line_settings
	offset    int
	direction lineDirection
	value     lineValue
}

type lineDirection C.uint32

const (
	lineDirectionAsIs   lineDirection = C.GPIOD_LINE_DIRECTION_AS_IS
	lineDirectionInput  lineDirection = C.GPIOD_LINE_DIRECTION_INPUT
	lineDirectionOutput lineDirection = C.GPIOD_LINE_DIRECTION_OUTPUT
)

type lineValue C.int32

const (
	lineValueError    lineValue = C.GPIOD_LINE_VALUE_ERROR
	lineValueInactive lineValue = C.GPIOD_LINE_VALUE_INACTIVE
	lineValueActive   lineValue = C.GPIOD_LINE_VALUE_ACTIVE
)

func NewLineSettings(offset uint) (*lineSettings, error) {
	var nativeRef *C.struct_gpiod_line_settings = C.gpiod_line_settings_new()
	if nativeRef == nil {
		return nil, fmt.Errorf("%s failed: NULL returned", "gpiod_line_settings_new")
	}
	return &lineSettings{
		nativeRef: nativeRef,
	}, nil
}

func (ls *lineSettings) SetDirection(direction lineDirection) error {
	var resultC C.int = C.gpiod_line_settings_set_direction(
		ls.nativeRef,
		direction,
	)
	if resultC == C.int(-1) {
		return fmt.Errorf("%s failed: -1 returned", "gpiod_line_settings_set_direction")
	}
	if resultC == C.int(0) {
		ls.direction = direction
		return nil
	}
	return fmt.Errorf("%s returned something unexpected", "gpiod_line_settings_set_direction")
}

func (ls *lineSettings) SetOutputValue(value lineValue) error {
	var resultC C.int = C.gpiod_line_settings_set_output_value(
		ls.nativeRef,
		value,
	)
	if resultC == C.int(-1) {
		return fmt.Errorf("%s failed: -1 returned", "gpiod_line_settings_set_output_value")
	}
	if resultC == C.int(0) {
		return nil
	}
	return fmt.Errorf("%s returned something unexpected", "gpiod_line_settings_set_output_value")
}
