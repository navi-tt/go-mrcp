package mpf

import "github.com/navi-tt/go-mrcp/toolkit"

/** Prototype of termination event handler */
type TerminationEventHandler func(termination *Termination, eventId int, descriptor interface{}) error

/** Table of termination virtual methods */
type TerminationVTable struct {

	/** Virtual termination destroy method */
	Destroy func(termination *Termination) error

	/** Virtual termination add method */
	Add func(termination *Termination, descriptor interface{}) error
	/** Virtual termination modify method */
	Modify func(termination *Termination, descriptor interface{}) error
	/** Virtual termination subtract method */
	Subtract func(termination *Termination) error
}

/** MPF Termination */
type Termination struct {

	/** Informative name used for debugging */
	Name string
	/** External object */
	Obj interface{}
	/** Media engine to send events to */
	MediaEngine interface{}
	/** Event handler */
	EventHandler TerminationEventHandler
	/** Codec manager */
	codecManager *CodecManager
	/** Timer queue */
	timerQueue toolkit.TimerQueue
	/** Termination factory entire termination created by */
	terminationFactory *TerminationFactory
	/** Table of virtual methods */
	vtable *TerminationVTable
	/** Slot in context */
	slot int64

	/** Audio stream */
	audioStream *AudioStream
	/** Video stream */
	videoStream *VideoStream
}

/**
 * Create MPF termination base.
 * @param termination_factory the termination factory
 * @param obj the external object associated with termination
 * @param vtable the table of virtual functions of termination
 * @param audio_stream the audio stream
 * @param video_stream the video stream
 * @param pool the pool to allocate memory from
 */
func TerminationBaseCreate(terminationFactory *TerminationFactory, obj interface{},
vtable *TerminationVTable, audioStream *AudioStream, videoStream *VideoStream) *Termination {
	return nil
}

/**
 * Add MPF termination.
 * @param termination the termination to add
 * @param descriptor the termination specific descriptor
 */
func (t *Termination) TerminationAdd(descriptor interface{}) error {
	return nil
}

/**
 * Modify MPF termination.
 * @param termination the termination to modify
 * @param descriptor the termination specific descriptor
 */
func (t *Termination) TerminationModify(descriptor interface{}) error {
	return nil
}

/**
 * Subtract MPF termination.
 * @param termination the termination to subtract
 */
func (t *Termination) TerminationSubtract() error {
	return nil
}
