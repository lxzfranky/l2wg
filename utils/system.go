package utils

import (
	"os"
)

func ResolveAddress(addr []string) string {
	cnt := len(addr)
	switch cnt {
	case 0:
		if port := os.Getenv("PORT"); len(port) > 0 {
			return ":" + port
		}
		return ":8080"

	case 1:
		return addr[0]
	default:
		panic("too many params")
	}
}
