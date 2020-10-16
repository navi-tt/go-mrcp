package engine

import (
	"github.com/navi-tt/go-mrcp/apr"
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp"
)

type MRCPEngine struct {
	Id         string              // Identifier of the engine
	ResourceId mrcp.MRCPResourceId // Resource identifier
	obj        interface{}         // External object associated with engine

	MethodVTable *MRCPEngineMethodVTable // Table of virtual methods
	EventVTable  *MRCPEngineEventVTable  // Table of virtual event handlers
	eventObj     interface{}             // External object used with event handlers

	CodecManager    *MpfCodecManager  // Codec manager
	DirLayout       *AptDirLayout     // Dir layout structure
	Config          *MRCPEngineConfig // Config of engine
	CurChannelCount int64             // Number of simultaneous channels currently in use
	IsOpen          bool              // Is engine successfully opened
	pool            *memory.AprPool   // Pool to allocate memory from

	/** Create state machine */
	CreateStateMachine func(obj interface{}, version mrcp.Version, pool *memory.AprPool) *MRCPStateMachine
}

/** Table of MRCP engine virtual methods */
type MRCPEngineMethodVTable struct {
	/** Virtual destroy */
	Destroy func(engine *MRCPEngine) bool

	/** Virtual open */
	Open func(engine *MRCPEngine) bool

	/** Virtual close */
	Close func(engine *MRCPEngine) bool

	/** Virtual channel create */
	CreateChannel func(engine *MRCPEngine, pool *memory.AprPool) MRCPEngineChannel
}

/** Table of MRCP engine virtual event handlers */
type MRCPEngineEventVTable struct {
	/** Open event handler */
	OnOpen func(channel *MRCPEngine, status bool) bool

	/** Close event handler */
	OnClose func(channel *MRCPEngine) bool
}

/** Opaque codec manager declaration */
type MpfCodecManager struct {
	pool *memory.AprPool // Memory pool

	CodecArr []apr.AprArrayHeader // Dynamic (resizable) array of codecs (mpf_codec_t*

	EventDescriptor *MpfCodecDescriptor // Default named event descriptor
}

/** Directories layout */
type AptDirLayout struct {
	Paths []string // Array of the directory paths the layout is composed of
	Count int64    // Number of directories in the layout
}

/** MRCP engine config */
type MRCPEngineConfig struct {
	MaxChannelCount int64         // Max number of simultaneous channels
	Params          *apr.AprTable // Table of name/value string params
}

type MRCPEngineChannel struct {
}

type MpfCodec struct {
}

type MpfCodecDescriptor struct {
}
