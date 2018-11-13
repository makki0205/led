package led

import (
	"fmt"
	"testing"
)

func TestToByte(t *testing.T) {
	fmt.Println(ToByte(255))
}

func TestNewLED(t *testing.T) {
	led, err := NewLED("")
	if err != nil {
		panic(err)
	}
	led.Send(255, 255, 255)
}
