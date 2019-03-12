package server

import (
	"errors"
	"net"
)

type conn struct {
	net.Conn
	server *Server
}

func (c *conn) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("System Error.")
			}
		}
	}()
	err = c.Conn.Close()
	if err == nil {
		c.server.wg.Done()
	}
	return
}
