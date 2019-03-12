package l2wg

//import (
//	"net/http"
//	"github.com/lxzfranky/l2wg/server"
//)
//
//type server interface {
//	ServeHTTP(w http.ResponseWriter, req *http.Request)
//}
//
//// 启动一个server
//// mod start/reload...
//func Run(mod Server, addr ...string) error {
//	address := resolveAddress(addr)
//	err := http.ListenAndServe(address, mod)
//	return err
//}