package gpiod

// #include <gpiod.h>
import "C"
import "fmt"

type lineConfig struct {
	nativeRef *C.struct_gpiod_line_config
}

func NewLineConfig() (*lineConfig, error) {
	var nativeRef *C.struct_gpiod_line_config = C.gpiod_line_config_new()
	if nativeRef == nil {
		return nil, fmt.Errorf("%s failed: NULL returned", "gpiod_line_config_new")
	}
	return &lineConfig{
		nativeRef: nativeRef,
	}, nil
}

func (lc *lineConfig) ApplyLineSettingsForSingleOffset(ls *lineSettings) error {

	var offsetsC C.uint = C.uint(ls.offset)
	var numOffsetsC C.size_t = 1

	var resultC C.int = C.gpiod_line_config_add_line_settings(
		lc.nativeRef,
		&offsetsC,
		numOffsetsC,
		ls.nativeRef,
	)
	if resultC == C.int(-1) {
		return fmt.Errorf("%s failed: -1 returned", "gpiod_line_config_add_line_settings")
	}
	if resultC == C.int(0) {
		return nil
	}
	return fmt.Errorf("%s returned something unexpected", "gpiod_line_config_add_line_settings")
}

func (lc *lineConfig) Free() {
	C.gpiod_line_config_free(lc.nativeRef)
}
