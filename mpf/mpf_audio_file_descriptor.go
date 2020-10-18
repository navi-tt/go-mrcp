package mpf

import "os"

const (
	/** FILE_READER defined as a stream source */
	FILE_READER = STREAM_DIRECTION_RECEIVE
	/** FILE_WRITER defined as a stream sink */
	FILE_WRITER = STREAM_DIRECTION_SEND
)

/** Audio file descriptor */
type AudioFileDescriptor struct {
	/** Indicate descriptor type (reader and/or writer) */
	mask StreamDirection
	/** Codec descriptor to use for audio file read/write */
	CodecDescriptor *CodecDescriptor
	/** File handle to read audio stream */
	ReadHandle *os.File
	/** File handle to write audio stream */
	WriteHandle *os.File
	/** Max size of file  */
	MaxWriteSize int64
}
