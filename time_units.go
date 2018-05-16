package agscheduler

import (
	"time"
)

type Delay struct {
	Value int32
}

func generateExecutor(multiplier int32, unit time.Duration) *Executor {
	return &Executor{TimeDelay: time.Duration(multiplier) * unit}
}

func (d *Delay) Milliseconds() *Executor {
	return generateExecutor(d.Value, time.Millisecond)
}

func (d *Delay) Seconds() *Executor {
	return generateExecutor(d.Value, time.Second)
}

func (d *Delay) Minutes() *Executor {
	return generateExecutor(d.Value, time.Minute)
}
