package mpf

/** MPF termination factory */
type TerminationFactory struct {
	/** Virtual create */
	CreateTermination func(factory *TerminationFactory, obj interface{}) *Termination
	/** Virtual assign engine */
	AssignEngine func(factory *TerminationFactory, mediaEngine *Engine) error
}

/**
 * Assign media engine to termination factory.
 * @param termination_factory the termination factory to assign media engine to
 * @param media_engine the media engine to assign
 */
func TerminationFactoryEngineAssign(terminationFactory *TerminationFactory, mediaEngine *Engine) error {
	return nil
}

/**
 * Create MPF termination from termination factory.
 * @param termination_factory the termination factory to create termination from
 * @param obj the external object associated with termination
 * @param pool the pool to allocate memory from
 */
func (tf *TerminationFactory) TerminationCreate(obj interface{}) *Termination {
	return nil
}

/**
 * Create raw MPF termination.
 * @param obj the external object associated with termination
 * @param audio_stream the audio stream of the termination
 * @param video_stream the video stream of the termination
 * @param pool the pool to allocate memory from
 */
func RawTerminationCreate(obj interface{}, audioStream *AudioStream, videoStream *VideoStream) *Termination {
	return nil
}

/**
 * Destroy MPF termination.
 * @param termination the termination to destroy
 */
func TerminationDestroy(termination *Termination) error {
	return nil
}

/**
 * Get termination name.
 * @param termination the termination to get name of
 */
func TerminationNameGet(termination *Termination) string {
	return ""
}

/**
 * Get associated object.
 * @param termination the termination to get object from
 */
func (t *Termination) TerminationObjectGet() interface{} {
	return nil
}

/**
 * Get audio stream.
 * @param termination the termination to get audio stream from
 */
func (t *Termination) TerminationAudioStreamGet() *AudioStream {
	return t.audioStream
}

/**
 * Get video stream.
 * @param termination the termination to get video stream from
 */
func (t *Termination) TerminationVideoStreamGet() *VideoStream {
	return t.videoStream
}
