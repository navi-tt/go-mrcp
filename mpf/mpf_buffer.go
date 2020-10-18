package mpf

import (
	"container/list"
	"sync"

	"github.com/navi-tt/go-mrcp/apr/memory"
)

type Chunk struct {
	link  *list.List
	frame Frame
}

type Buffer struct {
	head               *list.Element
	CurChunk           *Chunk
	RemainingChunkSize int64
	guard              sync.Mutex
	pool               *memory.AprPool
	size               int64 /* total size */
}

/** Create buffer */
func BufferCreate(pool *memory.AprPool) *Buffer {
	return &Buffer{}
}

/** Destroy buffer */
func BufferDestroy(buffer *Buffer) error {
	return nil
}

/** Restart buffer */
func (buffer *Buffer) BufferRestart() error {
	return nil
}

/** Write audio chunk to buffer */
func (buffer *Buffer) BufferAudioWrite(data []byte) error {
	return nil
}

/** Write event to buffer */
func (buffer *Buffer) BufferEventWrite(eventType FrameType) error {
	return nil
}

/** Read media frame from buffer */
func (buffer *Buffer) BufferFrameRead(mediaFrame *Frame) error {
	return nil
}

/** Get size of buffer **/
func (buffer *Buffer) BufferGetSize() int64 {
	return 0
}
