package mpf

import "os"

/** Audio file stream */
type AudioFileStream struct {
	audioStream *AudioStream

	readHandle  *os.File
	writeHandle *os.File

	eof          bool
	maxWriteSize int64
	curWriteSize int64
}

func AudioFileEventRaise(stream *AudioStream, eventId int, descriptor interface{}) {

}

func AudioFileDestroy(stream *AudioStream) error {
	var (
		fileStream *AudioFileStream = stream.Obj.(*AudioFileStream)
	)
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

func AudioFileFrameRead(stream *AudioStream, frame *Frame) error {
	var (
		fileStream *AudioFileStream = stream.Obj.(*AudioFileStream)
	)
	if fileStream.readHandle != nil && !fileStream.eof {
		_, err := frame.CodecFrame.Buffer.ReadFrom(fileStream.readHandle)
		if err != nil {
			return err
		}

	}

	return nil
}

/**
 * Create file stream.
 * @param termination the back pointer to hold
 * @param pool the pool to allocate memory from
 */
func FileStreamCreate(termination *Termination) *AudioStream {
	return nil
}

/**
 * Modify file stream.
 * @param stream file stream to modify
 * @param descriptor the descriptor to modify stream according
 */
func (as *AudioStream) FileStreamModify(descriptor *AudioFileDescriptor) error {
	return nil
}
