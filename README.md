# FS90R servo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/switchprocontroller)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/fs90r)

This drivers allows to control a FS90R servo connected on a PWM GPIO pin

## Documentation

For full documentation, please visit [![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/switchprocontroller)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/fs90r)

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
