package main

import (
	"log"

	"github.com/JoKr4/gpiod2go/pkg/gpiod"
)

func main() {
	log.Println("gpiod api version is:", gpiod.ApiVersion())

	// devicename according to "gpiodetect"
	// use the full path
	useDevice := "/dev/gpiochip0" // pinctrl-bcm2835 on rpi3
	useOffset := 22               // GPIO22 on rpi3

	d := gpiod.NewDevice(useDevice)
	err := d.Open()
	if err != nil {
		log.Println(err)
		return
	}
	defer d.Close()
	log.Println("successfully opened device")

	currentValue, err := d.GetLineValue(uint(useOffset))
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("current value of %s is %s\n", useOffset, currentValue.String())
}
