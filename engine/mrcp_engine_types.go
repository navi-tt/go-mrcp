package engine

import (
	"github.com/navi-tt/go-mrcp/apr"
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mpf"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/message"
	"github.com/navi-tt/go-mrcp/toolkit"
)

/** Table of channel virtual methods */
type MRCPEngineChannelMethodVTable struct {
	/** Virtual destroy */
	Destroy func(channel *MRCPEngineChannel) error
	/** Virtual open */
	Open func(channel *MRCPEngineChannel) error
	/** Virtual close */
	Close func(channel *MRCPEngineChannel) error
	/** Virtual process_request */
	ProcessRequest func(channel *MRCPEngineChannel, request *message.MRCPMessage) error
}

/** Table of channel virtual event handlers */
type MRCPEngineChannelEventVTable struct {
	/** Open event handler */
	OnOpen func(channel *MRCPEngineChannel, status bool) error
	/** Close event handler */
	OnClose func(channel *MRCPEngineChannel) error
	/** Message event handler */
	OnMessage func(channel *MRCPEngineChannel, message *message.MRCPMessage) error
}

/** MRCP engine channel declaration */
type MRCPEngineChannel struct {
	MethodVTable *MRCPEngineChannelMethodVTable // Table of virtual methods
	MethodObj    interface{}                    // External object used with virtual methods
	EventVTable  *MRCPEngineChannelEventVTable  // Table of virtual event handlers
	EventObj     interface{}                    // External object used with event handlers
	Termination  *mpf.MPFTermination            // Media termination todo(未完成)
	engine       *MRCPEngine                    // Back pointer to engine
	Id           string                         // Unique identifier to be used in traces
	Version      mrcp.Version                   // MRCP version
	IsOpen       bool                           // Is channel successfully opened
	pool         *memory.AprPool                // Pool to allocate memory from
}

/** Table of MRCP engine virtual methods */
type MRCPEngineMethodVTable struct {
	/** Virtual destroy */
	Destroy func(engine *MRCPEngine) error

	/** Virtual open */
	Open func(engine *MRCPEngine) error

	/** Virtual close */
	Close func(engine *MRCPEngine) error

	/** Virtual channel create */
	CreateChannel func(engine *MRCPEngine, pool *memory.AprPool) MRCPEngineChannel
}

/** Table of MRCP engine virtual event handlers */
type MRCPEngineEventVTable struct {
	/** Open event handler */
	OnOpen func(engine *MRCPEngine, status bool) error
	/** Close event handler */
	OnClose func(engine *MRCPEngine) error
}

/** MRCP engine */
type MRCPEngine struct {
	Id         string              // Identifier of the engine
	ResourceId mrcp.MRCPResourceId // Resource identifier
	obj        interface{}         // External object associated with engine

	MethodVTable *MRCPEngineMethodVTable // Table of virtual methods
	EventVTable  *MRCPEngineEventVTable  // Table of virtual event handlers
	eventObj     interface{}             // External object used with event handlers

	CodecManager    *mpf.MPFCodecManager  // Codec manager
	DirLayout       *toolkit.AptDirLayout // Dir layout structure
	Config          *MRCPEngineConfig     // Config of engine
	CurChannelCount int64                 // Number of simultaneous channels currently in use
	IsOpen          bool                  // Is engine successfully opened
	pool            *memory.AprPool       // Pool to allocate memory from

	/** Create state machine */
	CreateStateMachine func(obj interface{}, version mrcp.Version, pool *memory.AprPool) *MRCPStateMachine
}

/** MRCP engine config */
type MRCPEngineConfig struct {
	MaxChannelCount int64         // Max number of simultaneous channels
	Params          *apr.AprTable // Table of name/value string params todo(map[string]string ???)
	//Params        map[string]string // Table of name/value string params
}
