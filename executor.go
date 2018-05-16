package agscheduler

import (
	"time"
)

type Executor struct {
	TimeDelay time.Duration
}

func (e *Executor) Run(f func()) bool {
	time.AfterFunc(e.TimeDelay, f)
	return true
}
