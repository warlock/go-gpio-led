package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {

	adaptor := raspi.NewAdaptor()
	led := gpio.NewLedDriver(adaptor, "12")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("superBot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()

}
