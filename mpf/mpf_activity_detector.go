package mpf

import (
	"bytes"
	"encoding/binary"
)

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
	detector := ActivityDetector{
		LevelThreshold: 2,    /* 0 .. 255 */
		SpeechTimeout:  300,  /* 0.3 s */
		SilenceTimeout: 300,  /* 0.3 s */
		NoInputTimeout: 5000, /* 5 s */
		State:          DETECTOR_STATE_INACTIVITY,
		Duration:       0,
	}
	return &detector
}

/** Reset activity detector */
func (ad *ActivityDetector) ActivityDetectorReset() error {
	ad.Duration = 0
	ad.State = DETECTOR_STATE_INACTIVITY
	return nil
}

/** Set threshold of voice activity (silence) level */
func (ad *ActivityDetector) ActivityDetectorLevelSet(levelThreshold int64) {
	ad.LevelThreshold = levelThreshold
}

/** Set noinput timeout */
func (ad *ActivityDetector) ActivityDetectorNoInputTimeoutSet(noInputTimeout int64) {
	ad.NoInputTimeout = noInputTimeout
}

/** Set timeout required to trigger speech (transition from inactive to active state) */
func (ad *ActivityDetector) ActivityDetectorSpeechTimeoutSet(speechTimeout int64) {
	ad.SpeechTimeout = speechTimeout
}

/** Set timeout required to trigger silence (transition from active to inactive state) */
func (ad *ActivityDetector) ActivityDetectorSilenceTimeoutSet(silenceTimeout int64) {
	ad.SilenceTimeout = silenceTimeout
}

func (ad *ActivityDetector) ActivityDetectorStateChange(state DetectorState) {
	ad.Duration = 0
	ad.State = state
}

func (ad *ActivityDetector) ActivityDetectorLevelCalculate(frame *Frame) int64 {
	var (
		sum   int64 = 0
		count       = frame.CodecFrame.Buffer.Len() / 2
		cur         = 0
		end         = cur + count
	)
	// todo( 从[2]byte读一个int16 )
	for ; cur < end; cur++ {
		tmp := make([]byte, 2)
		_, _ = frame.CodecFrame.Buffer.Read(tmp)
		br := bytes.NewReader(tmp)
		var utt16 int16
		_ = binary.Read(br, binary.LittleEndian, &utt16)

		if utt16 < 0 {
			sum -= int64(utt16)
		} else {
			sum += int64(utt16)
		}
	}

	return sum / int64(count)
}

/** Process current frame return detected event if any */
func (ad *ActivityDetector) ActivityDetectorProcess(frame *Frame) DetectorEvent {
	var (
		detEvent DetectorEvent = MPF_DETECTOR_EVENT_NONE
		level    int64         = 0
	)

	if (frame.Type & MEDIA_FRAME_TYPE_AUDIO) == MEDIA_FRAME_TYPE_AUDIO {
		level = ad.ActivityDetectorLevelCalculate(frame)
	}

	if ad.State == DETECTOR_STATE_INACTIVITY {
		if level >= ad.LevelThreshold {
			/* start to detect activity */
			ad.ActivityDetectorStateChange(DETECTOR_STATE_ACTIVITY_TRANSITION)
		} else {
			ad.Duration += CODEC_FRAME_TIME_BASE
			if ad.Duration >= ad.NoInputTimeout {
				detEvent = MPF_DETECTOR_EVENT_NOINPUT
			}
		}
	} else if ad.State == DETECTOR_STATE_ACTIVITY_TRANSITION {
		if level >= ad.LevelThreshold {
			ad.Duration += CODEC_FRAME_TIME_BASE
			if ad.Duration >= ad.SpeechTimeout {
				/* finally detected activity */
				detEvent = MPF_DETECTOR_EVENT_ACTIVITY
				ad.ActivityDetectorStateChange(DETECTOR_STATE_ACTIVITY)
			}
		} else {
			/* fallback to inactivity */
			ad.ActivityDetectorStateChange(DETECTOR_STATE_INACTIVITY)
		}
	} else if ad.State == DETECTOR_STATE_ACTIVITY {
		if level >= ad.LevelThreshold {
			ad.Duration += CODEC_FRAME_TIME_BASE
		} else {
			/* start to detect inactivity */
			ad.ActivityDetectorStateChange(DETECTOR_STATE_INACTIVITY_TRANSITION)
		}
	} else if ad.State == DETECTOR_STATE_INACTIVITY_TRANSITION {
		if level >= ad.LevelThreshold {
			/* fallback to activity */
			ad.ActivityDetectorStateChange(DETECTOR_STATE_ACTIVITY)
		} else {
			ad.Duration += CODEC_FRAME_TIME_BASE
			if ad.Duration >= ad.SilenceTimeout {
				/* detected inactivity */
				detEvent = MPF_DETECTOR_EVENT_INACTIVITY
				ad.ActivityDetectorStateChange(DETECTOR_STATE_INACTIVITY)
			}
		}
	}

	return detEvent
}
