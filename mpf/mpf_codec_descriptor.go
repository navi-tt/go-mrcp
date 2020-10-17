package mpf

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 4:58 下午
 */

/** Codec descriptor */
type MPFCodecDescriptor struct {
	PayloadType  uint8  // Payload type used in RTP packet
	Name         string // Codec name
	SamplingRate uint16 // Sampling rate
	ChannelCount uint8  // Channel count
	Format       string // Codec dependent additional format
	Enabled      bool   // Enabled/disabled state
}
