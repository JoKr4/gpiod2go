package gpiod

// #include <gpiod.h>
import "C"

func ApiVersion() string {
	return C.GoString(C.gpiod_api_version())
}
