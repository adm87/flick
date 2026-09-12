//go:build pprof

package diagnostics

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	log.Println("pprof enabled")
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()
}
