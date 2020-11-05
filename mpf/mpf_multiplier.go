package mpf

/**
 * Create audio stream multiplier.
 * @param source the audio source
 * @param sink_arr the array of audio sinks
 * @param sink_count the number of audio sinks
 * @param codec_manager the codec manager
 * @param name the informative name used for debugging
 * @param pool the pool to allocate memory from
 */
func MultiplierCreate(source *AudioStream, sinkArr []*AudioStream, sinkCount int64, codecManager *CodecManager, name string) *Object {
	return nil
}
