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
	d := gpiod.NewDevice("/dev/gpiochip0")
	err := d.Open()
	if err != nil {
		log.Println(err)
		return
	}
	defer d.Close()
	log.Println("successfully opened device")

	err = d.AddLine(22, gpiod.LineDirectionOutput)
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
		err = d.SetLineValue(22, toggleLineValue)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("successfully set line value")
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
