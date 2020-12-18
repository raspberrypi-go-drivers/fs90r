package fs90r

import (
	"math"

	"github.com/stianeikeland/go-rpio/v4"
)

// FS90R instance
type FS90R struct {
	pinID          uint8
	pin            rpio.Pin
	clockFrequence int
	balancePoint   float64
	cycleLen       uint32
}

// NewFS90R creates a new FS90R instance
//
// The pinID specified must be a PWM compatible pin
func NewFS90R(pinID uint8) *FS90R {
	fs90r := FS90R{
		pinID:          pinID,
		clockFrequence: 1000000, // 1MHz
		balancePoint:   1.5,
		cycleLen:       20000, // 20000 clock cycles = 2ms
	}
	fs90r.pin = rpio.Pin(fs90r.pinID)
	fs90r.pin.Mode(rpio.Pwm)
	fs90r.pin.Freq(fs90r.clockFrequence)
	return &fs90r
}

// SetSpeed set the speed using percentage, from -100 to 100
//
// Positive value is clockwise, negative CCW
func (fs90r *FS90R) SetSpeed(speedPct int8) {
	if speedPct < -100 {
		speedPct = -100
	} else if speedPct > 100 {
		speedPct = 100
	}
	multiplier := float64(1)
	speedPctAbs := float64(speedPct)
	if speedPct < 0 {
		multiplier = -1
		speedPctAbs = math.Abs(float64(speedPct))
	}
	pulseDuration := fs90r.balancePoint + ((0.6*speedPctAbs)/100)*multiplier
	ratio := uint32((pulseDuration * float64(fs90r.cycleLen)) / 20.0)
	fs90r.pin.DutyCycle(ratio, uint32(fs90r.cycleLen))
}

// SetBalancePoint set the balance point (pulse duration) at which the servo will stay still. The balance point is in milliseconds and is 1.5 by default
//
// You can use this method if the servo is moving at default position (0% speed) but it's better to adjust this value by using the small screw under the FS90R
//
// You can check servos PWM behavior for more details
// https://www.pololu.com/blog/17/servo-control-interface-in-detail
func (fs90r *FS90R) SetBalancePoint(balancePoint float64) {
	fs90r.balancePoint = balancePoint
}
