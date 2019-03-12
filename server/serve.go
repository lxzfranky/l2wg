package server

import (
	"net"
	"net/http"
	"os"
	"sync"
)

type Server struct {
	*http.Server
	Listener      net.Listener
	SignalHooks      map[int]map[os.Signal][]func()
	tlsInnerListener *listener
	wg               sync.WaitGroup
	sigChan          chan os.Signal
	isChild          bool
	state            uint8
	Network          string
}
