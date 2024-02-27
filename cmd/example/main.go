package main

import (
	"log"

	"github.com/JoKr4/gpiod2go/pkg/gpiod"
)

func main() {
	log.Println("gpiod api version is:", gpiod.ApiVersion())

	// devicename according to "gpiodetect"
	// use the full path
	d := gpiod.NewDevice("/dev/gpiochip0")
	err := d.Open()
	if err != nil {
		log.Println(err)
		return
	}
	defer d.Close()
	log.Println("successfully opened")

	err = d.AddLine(22, gpiod.LineDirectionOutput)
	if err != nil {
		log.Println(err)
		return
	}

	err = d.SetLineValue(22, gpiod.LineValueActive)
	if err != nil {
		log.Println(err)
		return
	}
}
