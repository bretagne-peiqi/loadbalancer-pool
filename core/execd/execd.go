package execd

import (
	"errors"
	"fmt"
)

const (
	crashBackoff = 3
)

var (
	ErrNotRunning = errors.New("execd: process exited ...")
)

type Daemon struct {
}
