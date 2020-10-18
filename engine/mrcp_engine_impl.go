package engine

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mpf"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/message"
)

/** Create engine */
func MRCPEngineCreate(rid mrcp.MRCPResourceId, obj interface{}, vtable *MRCPEngineMethodVTable, pool *memory.AprPool) *MRCPEngine {
	return &MRCPEngine{}
}

/** Send engine open response */
func (engine *MRCPEngine) MRCPEngineOpenRespond(status bool) error {
	return engine.EventVTable.OnOpen(engine, status)
}

/** Send engine close response */
func (engine *MRCPEngine) MRCPEngineCloseRespond() error {
	return engine.EventVTable.OnClose(engine)
}

/** Get engine config */
func (engine *MRCPEngine) MRCPEngineConfigGet() *MRCPEngineConfig {
	return engine.Config
}

/** Get engine param by name */
func (engine *MRCPEngine) MRCPEngineParamGet(name string) string {
	return ""
}

/** Create engine channel */
func (engine *MRCPEngine) MRCPEngineChannelCreate(methodVTable *MRCPEngineMethodVTable, methodObj interface{}, termination *mpf.MPFTermination, pool *memory.AprPool) *MRCPEngineChannel {
	return nil
}

/** Create audio termination */
func MRCPEngineAudioTerminationCreate(obj interface{}, streamVTable *mpf.MPFAudioStreamVTable, capabilities *mpf.MPFStreamCapabilities, pool *memory.AprPool) *mpf.MPFTermination {
	return nil
}

/** Create engine channel and source media termination
 * @deprecated @see mrcp_engine_channel_create() and mrcp_engine_audio_termination_create()
 */
func (engine *MRCPEngine) MRCPEngineSourceChannelCreate(channelVTable *MRCPEngineChannelMethodVTable, streamVTable *mpf.MPFAudioStreamVTable,
	methodObj interface{}, codecDescriptor *mpf.MPFCodecDescriptor, pool *memory.AprPool) *MRCPEngineChannel {
	return nil
}

/** Create engine channel and sink media termination
 * @deprecated @see mrcp_engine_channel_create() and mrcp_engine_audio_termination_create()
 */
func (engine *MRCPEngine) MRCPEngineSinkChannelCreate(channelVTable *MRCPEngineChannelMethodVTable, streamVTable *mpf.MPFAudioStreamVTable,
	methodObj interface{}, codecDescriptor *mpf.MPFCodecDescriptor, pool *memory.AprPool) *MRCPEngineChannel {
	return nil
}

/** Send channel open response */
func (channel *MRCPEngineChannel) MRCPEngineChannelOpenRespond(status bool) error {
	return channel.EventVTable.OnOpen(channel, status)
}

/** Send channel close response */
func (channel *MRCPEngineChannel) MRCPEngineChannelCloseRespond() error {
	return channel.EventVTable.OnClose(channel)
}

/** Send response/event message */
func (channel *MRCPEngineChannel) MRCPEngineChannelMessageSend(message *message.MRCPMessage) error {
	return channel.EventVTable.OnMessage(channel, message)
}

/** Get channel identifier */
func (channel *MRCPEngineChannel) MRCPEngineChannelIdGet() string {
	return channel.Id
}

/** Get MRCP version channel is created in the scope of */
func (channel *MRCPEngineChannel) MRCPEngineChannelVersionGet() mrcp.Version {
	return channel.Version
}

/** Get codec descriptor of the audio source stream */
func (channel *MRCPEngineChannel) MRCPEngineSourceStreamCodecGet() *mpf.MPFCodecDescriptor {
	return nil
}

/** Get codec descriptor of the audio sink stream */
func (channel *MRCPEngineChannel) MRCPEngineSinkStreamCodecGet() *mpf.MPFCodecDescriptor {
	return nil
}
