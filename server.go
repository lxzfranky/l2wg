package l2wg

import (
	"net"
	"net/http"
	"os"
	"sync"
	"github.com/lxzfranky/l2wg/server"
)


type Server interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// 启动一个server
// mod start/reload...
func Run(mod Server, addr ...string) error {

}