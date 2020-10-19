package mpf

/** Detector states */
type DetectorState = int

const (
	DETECTOR_STATE_INACTIVITY            DetectorState = iota /**< inactivity detected */
	DETECTOR_STATE_ACTIVITY_TRANSITION                        /**< activity detection is in-progress */
	DETECTOR_STATE_ACTIVITY                                   /**< activity detected */
	DETECTOR_STATE_INACTIVITY_TRANSITION                      /**< inactivity detection is in-progress */
)

/** Events of activity detector */
type DetectorEvent = int

const (
	MPF_DETECTOR_EVENT_NONE       DetectorEvent = iota /**< no event occurred */
	MPF_DETECTOR_EVENT_ACTIVITY                        /**< voice activity (transition to activity from inactivity state) */
	MPF_DETECTOR_EVENT_INACTIVITY                      /**< voice inactivity (transition to inactivity from activity state) */
	MPF_DETECTOR_EVENT_NOINPUT                         /**< noinput event occurred */
)

/** Activity detector */
type ActivityDetector struct {
	/* voice activity (silence) level threshold */
	LevelThreshold int64

	/* period of activity required to complete transition to active state */
	SpeechTimeout int64
	/* period of inactivity required to complete transition to inactive state */
	SilenceTimeout int64
	/* noinput timeout */
	NoInputTimeout int64

	/* current state */
	State DetectorState
	/* duration spent in current state  */
	Duration int64
}

/** Create activity detector */
func ActivityDetectorCreate() *ActivityDetector {
	return &ActivityDetector{}
}

/** Reset activity detector */
func (ad *ActivityDetector) ActivityDetectorReset() error {
	return nil
}

/** Set threshold of voice activity (silence) level */
func (ad *ActivityDetector) ActivityDetectorLevelSet(levelThreshold int64) {
}

/** Set noinput timeout */
func (ad *ActivityDetector) ActivityDetectorNoInputTimeoutSet(noInputTimeout int64) {

}

/** Set timeout required to trigger speech (transition from inactive to active state) */
func (ad *ActivityDetector) ActivityDetectorSpeechTimeoutSet(speechTimeout int64) {}

/** Set timeout required to trigger silence (transition from active to inactive state) */
func (ad *ActivityDetector) ActivityDetectorSilenceTimeoutSet(silenceTimeout int64) {}

/** Process current frame return detected event if any */
func (ad *ActivityDetector) ActivityDetectorProcess(frame *Frame) DetectorEvent {
	return 0
}
