package server

import "net"

type listener struct {
	net.Listener
	stop    chan error
	stopped bool
	server  *Server
}

func newListener()  {

}

func Accept()  {

}

func Close()  {

}