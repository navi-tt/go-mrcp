package mpf

/** Prototype of scheduler callback */
type SchedulerProc func(scheduler *Scheduler, obj interface{})

type Scheduler struct {
	resolution uint64 /* scheduler resolution */

	mediaResolution uint64
	mediaProc       SchedulerProc
	mediaObj        interface{}

	timerResolution  uint64
	timerElapsedTime uint64
	timerProc        SchedulerProc
	timerObj         interface{}
	timerId          uint
}

/** Create scheduler */
func SchedulerCreate() *Scheduler {
	return nil
}

/** Destroy scheduler */
func SchedulerDestroy(scheduler *Scheduler) error {
	return nil
}

/** Set media processing clock */
func (s *Scheduler) SchedulerMediaClockSet(resolution uint64, proc SchedulerProc, obj interface{}) error {
	return nil
}

/** Set timer clock */
func (s *Scheduler) SchedulerTimerClockSet(resolution uint64, proc SchedulerProc, obj interface{}) error {
	return nil
}

/** Set scheduler rate (n times faster than real-time) */
func (s *Scheduler) SchedulerRateSet(rate uint64) error {
	return nil
}

/** Start scheduler */
func (s *Scheduler) SchedulerStart() error {
	return nil
}

/** Stop scheduler */
func (s *Scheduler) SchedulerStop() error {
	return nil
}
