package mpf

import "sync"

type FrameBuffer struct {
	RawData    []byte
	Frames     []*Frame
	FrameCount int64
	FrameSize  int64

	WritePos int64
	ReadPos  int64

	guard sync.Mutex
}

/** Create frame buffer */
func FrameBufferCreate(frameSize, frameCount int64) *FrameBuffer {
	return nil
}

/** Destroy frame buffer */
func FrameBufferDestroy(buffer *FrameBuffer) error {
	return nil
}

/** Restart frame buffer */
func (buffer *FrameBuffer) FrameBufferRestart() error {
	return nil
}

/** Write frame to buffer */
func (buffer *FrameBuffer) FrameBufferWrite(frame *Frame) error {
	return nil
}

/** Read frame from buffer */
func (buffer *FrameBuffer) FrameBufferRead(frame *Frame) error {
	return nil
}
