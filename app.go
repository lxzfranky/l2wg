package l2wg

import "net/http"

type app struct {
	Handlers *ControllerRegister
	Server   *http.Server
}