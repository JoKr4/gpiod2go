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

type lineDirection uint

const (
	LineDirectionAsIs   lineDirection = C.GPIOD_LINE_DIRECTION_AS_IS
	LineDirectionInput  lineDirection = C.GPIOD_LINE_DIRECTION_INPUT
	LineDirectionOutput lineDirection = C.GPIOD_LINE_DIRECTION_OUTPUT
)

type lineValue int

const (
	LineValueError    lineValue = C.GPIOD_LINE_VALUE_ERROR
	LineValueInactive lineValue = C.GPIOD_LINE_VALUE_INACTIVE
	LineValueActive   lineValue = C.GPIOD_LINE_VALUE_ACTIVE
)

func NewLineSettings(offset uint, direction lineDirection) (*lineSettings, error) {
	var nativeRef *C.struct_gpiod_line_settings = C.gpiod_line_settings_new()
	if nativeRef == nil {
		return nil, fmt.Errorf("%s failed: NULL returned", "gpiod_line_settings_new")
	}
	new := lineSettings{
		nativeRef: nativeRef,
		direction: direction,
	}
	err := new.setDirection(direction)
	if err != nil {
		return nil, err
	}
	return &new, nil
}

func (ls *lineSettings) setDirection(direction lineDirection) error {
	var resultC C.int = C.gpiod_line_settings_set_direction(
		ls.nativeRef,
		C.enum_gpiod_line_direction(direction),
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
		C.enum_gpiod_line_value(value),
	)
	if resultC == C.int(-1) {
		return fmt.Errorf("%s failed: -1 returned", "gpiod_line_settings_set_output_value")
	} else if resultC != C.int(0) {
		return fmt.Errorf("%s returned something unexpected", "gpiod_line_settings_set_output_value")
	}

	config, err := NewLineConfig()
	if err != nil {
		return err
	}
	defer config.Free()

	err = config.ApplyLineSettingsForSingleOffset(ls)
	if err != nil {
		return err
	}
	return nil
}

func (ls *lineSettings) Free() {
	C.gpiod_line_settings_free(ls.nativeRef)
}
