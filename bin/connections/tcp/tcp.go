package tcp

import "net"

func InitTCP(network, address string) (net.Listener, error) {
	return net.Listen(network, address)
}
