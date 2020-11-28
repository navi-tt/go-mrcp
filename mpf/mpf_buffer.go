package mpf

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"sync"
)

type Chunk struct {
	frame Frame
}

type Buffer struct {
	head               *list.Element // todo(ring 的 head)
	link               *list.List    // todo(ring 的 link)
	CurChunk           *Chunk
	RemainingChunkSize int64
	guard              sync.Mutex
	size               int64 // total size
}

/** Create buffer */
func BufferCreate() *Buffer {
	return &Buffer{
		head:               nil,
		link:               list.New(),
		CurChunk:           nil,
		RemainingChunkSize: 0,
		size:               0,
	}
}

/** Destroy buffer */
func BufferDestroy(buffer *Buffer) error {
	return nil
}

/** Restart buffer */
func (buffer *Buffer) BufferRestart() error {
	buffer.guard.Lock()
	defer buffer.guard.Unlock()
	buffer.link = list.New()
	buffer.head = buffer.link.Front()
	return nil
}

func (buffer *Buffer) BufferChunkWrite(chunk *Chunk) error {
	if chunk == nil {
		return fmt.Errorf("chunk is nil")
	}
	buffer.link.PushBack(chunk)
	return nil
}

func (buffer *Buffer) BufferChunkRead() *Chunk {
	chunkEle := buffer.link.Front()
	if chunkEle != nil {
		if chunkEle.Value != nil {
			if chunk, ok := chunkEle.Value.(*Chunk); ok {
				buffer.link.Remove(chunkEle)
				return chunk
			}
		}
	}
	return nil
}

/** Write audio chunk to buffer */
func (buffer *Buffer) BufferAudioWrite(data []byte) error {
	buffer.guard.Lock()
	defer buffer.guard.Unlock()

	chunk := &Chunk{frame: Frame{
		Type:   MEDIA_FRAME_TYPE_AUDIO,
		Marker: 0,
		CodecFrame: CodecFrame{
			Buffer: bytes.NewBuffer(data),
			Size:   int64(len(data)),
		},
		EventFrame: NamedEventFrame{},
	}}

	err := buffer.BufferChunkWrite(chunk)
	if err != nil {
		return err
	}
	buffer.size += int64(len(data))
	return nil
}

/** Write event to buffer */
func (buffer *Buffer) BufferEventWrite(eventType FrameType) error {
	buffer.guard.Lock()
	defer buffer.guard.Unlock()
	chunk := &Chunk{frame: Frame{
		Type:   eventType,
		Marker: 0,
		CodecFrame: CodecFrame{
			Buffer: nil,
		},
		EventFrame: NamedEventFrame{},
	}}
	err := buffer.BufferChunkWrite(chunk)
	if err != nil {
		return err
	}
	return nil
}

/** Read media frame from buffer */
func (buffer *Buffer) BufferFrameRead(mediaFrame *Frame) error {
	var (
		dest, src          *CodecFrame
		remainingFrameSize = int64(mediaFrame.CodecFrame.Size)
	)
	buffer.guard.Lock()
	for {
		if buffer.CurChunk == nil {
			buffer.CurChunk = buffer.BufferChunkRead()
			if buffer.CurChunk == nil {
				break
			}
			buffer.RemainingChunkSize = int64(buffer.CurChunk.frame.CodecFrame.Size)
		}

		dest = &mediaFrame.CodecFrame
		src = &buffer.CurChunk.frame.CodecFrame

		mediaFrame.Type |= buffer.CurChunk.frame.Type

		if remainingFrameSize < buffer.RemainingChunkSize {
			/* copy remaining_frame_size */
			//todo(做测试下)
			_, err := io.CopyN(dest.Buffer, src.Buffer, remainingFrameSize)
			if err != nil {
				return err
			}
			buffer.RemainingChunkSize -= remainingFrameSize
			remainingFrameSize = 0
		} else {
			/* copy remaining_chunk_size and proceed to the next chunk */
			//todo(做测试下)
			_, err := io.CopyN(dest.Buffer, src.Buffer, buffer.RemainingChunkSize)
			if err != nil {
				return err
			}
			remainingFrameSize -= buffer.RemainingChunkSize
			buffer.size -= buffer.RemainingChunkSize
			buffer.RemainingChunkSize = 0
			buffer.CurChunk = nil
		}

		if remainingFrameSize <= 0 {
			break
		}
	}

	if remainingFrameSize > 0 {
		//offset := int64(mediaFrame.CodecFrame.Buffer.Len()) - remainingFrameSize
		//mediaFrame.CodecFrame.Buffer.Write() // todo(这里bytebuffer不需要把多余的置空)
	}

	return nil
}

/** Get size of buffer **/
func (buffer *Buffer) BufferGetSize() int64 {
	return buffer.size
}
