package l2wg

import (
	"github.com/lxzfranky/l2wg/utils"
	"net/http"
)


type Server interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// 启动一个server
// mod start/reload...
func Run(mod Server, addr ...string) error {
	address := utils.ResolveAddress(addr)
	err := http.ListenAndServe(address, mod)
	return err
}