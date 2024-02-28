package main

import (
	"log"
	"time"

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

	err = d.AddLine(uint(useOffset), gpiod.LineDirectionOutput)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("successfully added line")

	t := time.NewTicker(3.0 * time.Second)
	defer t.Stop()

	stopAfter := 3
	toggleLineValue := gpiod.LineValueActive

	for {
		<-t.C
		err = d.SetLineValue(uint(useOffset), toggleLineValue)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("successfully set line value to:", toggleLineValue.String())
		if toggleLineValue == gpiod.LineValueActive {
			toggleLineValue = gpiod.LineValueInactive
		} else {
			toggleLineValue = gpiod.LineValueActive
		}
		stopAfter--
		if stopAfter == 0 {
			break
		}
	}

	log.Println("all done")
}
