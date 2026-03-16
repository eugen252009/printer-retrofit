package serial

import (
	"time"

	"github.com/tarm/serial"
)

func InitSerial() (*serial.Port, error) {
	config := &serial.Config{
		Name:        "/dev/ttyUSB0",  // default USB Port
		Baud:        150200,          // Default USB Baud
		ReadTimeout: 5 * time.Second, // 5 Seconds Timeout
	}
	return serial.OpenPort(config)
}
