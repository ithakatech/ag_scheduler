package agscheduler

type Scheduler struct{}

func NewScheduler() Scheduler {
	return Scheduler{}
}

func (s *Scheduler) After(delay int32) *Delay {
	return &Delay{Value: delay}
}
