package mpf

import "github.com/navi-tt/go-mrcp/apr/memory"

/**
 * Create bridge of audio streams.
 * @param source the source audio stream
 * @param sink the sink audio stream
 * @param codec_manager the codec manager
 * @param name the informative name used for debugging
 * @param pool the pool to allocate memory from
 */
func BridgeCreate(source *AudioStream, sink *AudioStream, manager *CodecManager, name string, pool *memory.AprPool) *Object {
	return nil
}
