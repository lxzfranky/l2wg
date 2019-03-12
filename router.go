package l2wg

import "net/http"

type Router struct {
	trees                      map[string]*node
	RedirectTrailingSlash      bool
	RedirectFixedPath          bool
	HandleMethodNotAllowed     bool
	HandleOPTIONS              bool
	NotFoundUseMidWare         bool
	MethodNotAllowedUseMidware bool

	NotFound         Handler
	MethodNotAllowed Handler
	PanicHandler     func(http.ResponseWriter, *http.Request, interface{})
	Mid              []Handler
}