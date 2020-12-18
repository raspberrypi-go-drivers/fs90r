# FS90R servo


[![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/fs90r.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/fs90r)
![golangci-lint](https://github.com/raspberrypi-go-drivers/fs90r/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/fs90R)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/fs90R)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This drivers allows to control a FS90R servo connected on a PWM GPIO pin

## Documentation

For full documentation, please visit [![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/fs90r.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/fs90r)

## Quick start

```go
import (
	"os"

	"github.com/raspberrypi-go-drivers/fs90r"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	servo := fs90r.NewFS90R(18)
	servo.SetSpeed(50) // Set servo speed to 50% (CCW)
}
```

## Raspberry Pi compatibility

This driver has has only been tested on an Raspberry Pi Zero WH using integrated bluetooth but may work well on other Raspberry Pi having integrated Bluetooth

## License

MIT License

---

Special thanks to @stianeikeland

This driver is based on his work in [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio/)
