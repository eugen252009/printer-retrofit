package main

import (
	"fmt"
	"io"
	"log"

	"github.com/eugen252009/printer-retrofit/bin/connections"
	"github.com/eugen252009/printer-retrofit/bin/connections/serial"
	"github.com/eugen252009/printer-retrofit/bin/connections/tcp"
)

func main() {
	tcpListener, err := tcp.InitTCP("tcp", ":5000")
	if err != nil {
		log.Fatal("TCP Init Error:", err)
	}

	serialconn, err := serial.InitSerial()
	if err != nil {
		log.Fatal("Serial Init Error:", err)
	}
	defer serialconn.Close()

	fmt.Println("Warte auf TCP-Verbindung...")
	netconn, err := tcpListener.Accept()
	if err != nil {
		log.Fatal("Accept Error:", err)
	}
	defer netconn.Close()

	conn := connections.TCPSerialStrategy{
		TCP:    netconn,
		Serial: serialconn,
	}

	fmt.Println("Bridge aktiv: Serial <-> TCP")

	errChan := make(chan error, 2)

	go func() {
		_, err := io.Copy(conn.TCP, conn.Serial)
		errChan <- err
	}()

	go func() {
		_, err := io.Copy(conn.Serial, conn.TCP)
		errChan <- err
	}()

	fatalErr := <-errChan
	if fatalErr != nil {
		log.Printf("Bridge geschlossen: %v", fatalErr)
	}
}
