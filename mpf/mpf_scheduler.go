package mpf

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

type SchedulerProc func(scheduler *Scheduler, obj interface{})
