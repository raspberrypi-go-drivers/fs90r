package fs90r_test

import (
	"os"

	"github.com/raspberrypi-go-drivers/fs90r"
	"github.com/stianeikeland/go-rpio/v4"
)

func Example_ccw() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	servo := fs90r.NewFS90R(18)
	servo.SetSpeed(50) // Set servo speed to 50% (CCW)
}

func Example_cw() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	servo := fs90r.NewFS90R(18)
	servo.SetSpeed(-50) // Set servo speed to 50% (CW)
}

func ExampleNewFS90R() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	servo := fs90r.NewFS90R(18)
	servo.SetSpeed(100) // Set servo speed to 100% (CCW)
}

func ExampleFS90R_SetSpeed() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	servo := fs90r.NewFS90R(18)
	servo.SetSpeed(-25) // Set servo speed to 25% (CW)
}

func ExampleFS90R_SetBalancePoint() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	servo := fs90r.NewFS90R(18)
	servo.SetSpeed(0) // if the servo is moving with this setup you can use SetBalancePoint
	servo.SetBalancePoint(1.6)
	// if the value used in SetBalancePoint is correct, the servo should now stay still
}
