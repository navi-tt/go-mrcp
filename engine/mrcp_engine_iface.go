package engine

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/message"
)

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 5:47 下午
 */

/** Destroy engine */
func MRCPEngineVirtualDestroy(e *MRCPEngine) error {
	return nil
}

/** Open engine */
func MRCPEngineVirtualOpen(e *MRCPEngine) error {
	return nil
}

/** Response to open engine request */
func MRCPEngineOnOpen(e *MRCPEngine, status bool) {

}

/** Close engine */
func MRCPEngineVirtualClose(e *MRCPEngine) error {
	return nil
}

/** Create engine channel */
func MRCPEngineChannelVirtualCreate(engine *MRCPEngine, version mrcp.Version, pool *memory.AprPool) *MRCPEngineChannel {
	return nil
}

/** Destroy engine channel */
func MRCPEngineChannelVirtualDestroy(channel *MRCPEngineChannel) error {
	return nil
}

/** Open engine channel */
func MRCPEngineChannelVirtualOpen(channel *MRCPEngineChannel) error {
	if !channel.IsOpen {
		err := channel.MethodVTable.Open(channel)
		if err != nil {
			channel.IsOpen = false
			return err
		}
	}
	channel.IsOpen = true
	return nil
}

/** Close engine channel */
func MRCPEngineChannelVirtualClose(channel *MRCPEngineChannel) error {
	if channel.IsOpen {
		err := channel.MethodVTable.Close(channel)
		if err != nil {
			return err
		}
		channel.IsOpen = false
	}
	return nil
}

/** Process request */
func MRCPEngineChannelRequestProcess(channel *MRCPEngineChannel, message *message.MRCPMessage) error {
	return channel.MethodVTable.ProcessRequest(channel, message)
}

/** Allocate engine config */
func MRCPEngineConfigAlloc(pool *memory.AprPool) *MRCPEngineConfig {
	return &MRCPEngineConfig{}
}
