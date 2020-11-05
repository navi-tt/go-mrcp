package mpf

/**
 * Create audio stream mixer.
 * @param source_arr the array of audio sources
 * @param source_count the number of audio sources
 * @param sink the audio sink
 * @param codec_manager the codec manager
 * @param name the informative name used for debugging
 * @param pool the pool to allocate memory from
 */
func MixerCreate(sourceArr []*AudioStream, sourceCount int64, sink *AudioStream, codecManager *CodecManager, name string) *Object {
	return nil
}
