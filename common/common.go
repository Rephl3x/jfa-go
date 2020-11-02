package common

import (
	"fmt"
	"log"
)

// TimeoutHandler recovers from an http timeout.
type TimeoutHandler func()

// NewTimeoutHandler returns a new Timeout handler.
func NewTimeoutHandler(name, addr string, noFail bool) TimeoutHandler {
	return func() {
		out := fmt.Sprintf("Failed to authenticate with %s @ %s: Timed out", name, addr)
		if noFail {
			log.Print(out)
		} else {
			log.Fatalf(out)
		}
	}
}
