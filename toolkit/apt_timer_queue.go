package toolkit

import "container/list"

/** Prototype of timer callback */
type TimerProc func(timer *Timer, obj interface{})

/** Timer queue */
type TimerQueue struct {
	/** Ring head */
	Head *list.Element

	Link *list.List

	/** Elapsed time */
	elapsedTime uint32
	/** Whether elapsed_time is reset or not */
	Reset bool
}

/** Timer */
type Timer struct {

	/** Ring entry */
	link *list.List

	/** Back pointer to queue */
	queue *TimerQueue
	/** Time next report is scheduled at */
	scheduledTime uint32

	/** Timer proc */
	proc TimerProc
	/** Timer object */
	obj interface{}
}
