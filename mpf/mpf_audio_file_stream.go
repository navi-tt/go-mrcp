package mpf

import (
	"fmt"
	"io"
	"os"
)

/** Audio file stream */
type AudioFileStream struct {
	audioStream *AudioStream

	readHandle  *os.File
	writeHandle *os.File

	eof          bool
	maxWriteSize int64
	curWriteSize int64
}

func AudioFileDestroy(stream *AudioStream) error {
	fileStream, ok := stream.Obj.(*AudioFileStream)
	if !ok {
		return fmt.Errorf("AudioStream.Obj is not *AudioFileStream")
	}
	if fileStream.readHandle != nil {
		err := fileStream.readHandle.Close()
		if err != nil {
			return err
		}
		fileStream.readHandle = nil
	}
	if fileStream.writeHandle != nil {
		err := fileStream.writeHandle.Close()
		if err != nil {
			return err
		}
		fileStream.writeHandle = nil
	}
	return nil
}

func AudioFileReaderOpen(stream *AudioStream, codec *Codec) error {
	return nil
}

func AudioFileReaderClose(stream *AudioStream) error {
	return nil
}

func AudioFileFrameRead(as *AudioStream, frame *Frame) error {
	fileStream, ok := as.Obj.(*AudioFileStream)
	if !ok {
		return fmt.Errorf("AudioStream.Obj is not *AudioFileStream")
	}

	// todo(应该是按frame一帧一帧来读吧, 这里先一次性读完处理, 等确认frame一帧是多长再来改)
	if fileStream.readHandle != nil && !fileStream.eof {
		n, err := io.CopyN(frame.CodecFrame.Buffer, fileStream.readHandle, frame.CodecFrame.Size)
		//n, err := frame.CodecFrame.Buffer.ReadFrom(fileStream.readHandle)
		if err != nil {
			return err
		}
		// 刚好读了1个frame帧
		if n == frame.CodecFrame.Size {
			frame.Type = MEDIA_FRAME_TYPE_AUDIO
		} else {
			fileStream.eof = true
			return AudioFileEventRaise(as, 0, nil)
		}
	}
	return nil
}

func AudioFileWriterOpen(as *AudioStream, codec *Codec) error {
	return nil
}

func AudioFileWriterClose(as *AudioStream) error {
	return nil
}

func AudioFileFrameWrite(as *AudioStream, frame *Frame) error {
	fileStream, ok := as.Obj.(*AudioFileStream)
	if !ok {
		return fmt.Errorf("AudioStream.Obj is not *AudioFileStream")
	}
	if fileStream.writeHandle != nil &&
		(fileStream.maxWriteSize > 0 || fileStream.curWriteSize < fileStream.maxWriteSize) {
		//n, err := frame.CodecFrame.Buffer.WriteTo(fileStream.writeHandle)
		n, err := io.CopyN(fileStream.writeHandle, frame.CodecFrame.Buffer, frame.CodecFrame.Size)
		if err != nil {
			return err
		}
		fileStream.curWriteSize += n
		if fileStream.curWriteSize >= fileStream.maxWriteSize {
			return AudioFileEventRaise(as, 0, nil)
		}
	}
	return nil
}

var vtable AudioStreamVTable = AudioStreamVTable{
	Destroy:    AudioFileDestroy,
	OpenRX:     AudioFileReaderOpen,
	CloseRX:    AudioFileReaderClose,
	ReadFrame:  AudioFileFrameRead,
	OpenTX:     AudioFileWriterOpen,
	CloseTX:    AudioFileWriterClose,
	WriteFrame: AudioFileFrameWrite,
	Trace:      nil,
}

/**
 * Create file stream.
 * @param termination the back pointer to hold
 * @param pool the pool to allocate memory from
 */
func FileStreamCreate(termination *Termination) *AudioStream {
	var (
		fileStream   = &AudioFileStream{}
		capabilities = StreamCapabilitiesCreate(STREAM_DIRECTION_DUPLEX)
		audioStream  = AudioStreamCreate(fileStream, &vtable, capabilities)
	)
	if audioStream == nil {
		return nil
	}
	audioStream.termination = termination
	fileStream.audioStream = audioStream
	return audioStream
}

/**
 * Modify file stream.
 * @param stream file stream to modify
 * @param descriptor the descriptor to modify stream according
 */
func FileStreamModify(as *AudioStream, descriptor *AudioFileDescriptor) error {
	fileStream, ok := as.Obj.(*AudioFileStream)
	if !ok {
		return fmt.Errorf("AudioStream.Obj is not *AudioFileStream")
	}
	if (descriptor.mask & FILE_READER) > 0 {
		if fileStream.readHandle != nil {
			fileStream.readHandle.Close()
		}
		fileStream.readHandle = descriptor.ReadHandle
		fileStream.eof = false
		as.direction |= FILE_READER
	}
	if (descriptor.mask & FILE_WRITER) > 0 {
		if fileStream.writeHandle != nil {
			fileStream.writeHandle.Close()
		}
		fileStream.writeHandle = descriptor.WriteHandle
		fileStream.maxWriteSize = descriptor.MaxWriteSize
		fileStream.curWriteSize = 0
		as.direction |= FILE_WRITER
		as.TXDescriptor = descriptor.CodecDescriptor
	}
	return nil
}

func AudioFileEventRaise(as *AudioStream, eventId int, descriptor interface{}) error {
	if as.termination != nil && as.termination.EventHandler != nil {
		return as.termination.EventHandler(as.termination, eventId, descriptor)
	}
	return nil
}
