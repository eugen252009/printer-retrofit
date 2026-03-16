package connections

import (
	"io"
	"net"

	"github.com/tarm/serial"
)

type TCPSerialStrategy struct {
	TCP    net.Conn
	Serial *serial.Port
}

func (t *TCPSerialStrategy) SendToMachine(w io.Writer, gcode string) error {
	_, err := w.Write([]byte(gcode + "\n"))
	return err
}

func (t *TCPSerialStrategy) ReadFromMachine(r io.Reader, buffer []byte) (int, error) {
	return r.Read(buffer)
}
