package main

import (
	"log"

	"github.com/JoKr4/gpiod2go/pkg/gpiod"
)

func main() {
	log.Println("gpiod api version is:", gpiod.ApiVersion())

	// devicename according to "gpiodetect"
	// 'gpiochip0', and '/dev/gpiochip0' all refer to the same chip.
	d := gpiod.NewDevice("gpiochip0")
	err := d.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success")
}
