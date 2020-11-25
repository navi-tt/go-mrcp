package mpf

import (
	"bytes"
	"fmt"
)

/** MPF bridge derived from MPF object */
type Bridge struct {
	/** MPF bridge base */
	base *Object
	/** Audio stream source */
	source *AudioStream
	/** Audio stream sink */
	sink *AudioStream
	/** Codec used in case of null bridge */
	codec *Codec
	/** Media frame used to read data from source and write it to sink */
	frame Frame
}

func BridgeProcess(object *Object) error {
	bridge := &Bridge{
		base: object,
	}
	bridge.frame.Type = MEDIA_FRAME_TYPE_NONE
	bridge.frame.Marker = MPF_MARKER_NONE
	// todo(source当前是空指针, panic)
	err := bridge.source.VTable.ReadFrame(bridge.source, &bridge.frame)
	if err != nil {
		return err
	}

	if (bridge.frame.Type & MEDIA_FRAME_TYPE_AUDIO) == 0 {
		bridge.frame.CodecFrame.Buffer.Reset()
	}

	return bridge.sink.VTable.WriteFrame(bridge.sink, &bridge.frame)
}
func NullBridgeProcess(object *Object) error {
	bridge := &Bridge{
		base: object,
	}
	bridge.frame.Type = MEDIA_FRAME_TYPE_NONE
	bridge.frame.Marker = MPF_MARKER_NONE
	// todo(source当前是空指针, panic)
	err := bridge.source.VTable.ReadFrame(bridge.source, &bridge.frame)
	if err != nil {
		return err
	}

	if (bridge.frame.Type & MEDIA_FRAME_TYPE_AUDIO) == 0 {
		/* generate silence frame */
		err = bridge.codec.CodecInitialize(&bridge.frame.CodecFrame)
		if err != nil {
			return err
		}
	}

	return bridge.sink.VTable.WriteFrame(bridge.sink, &bridge.frame)
}

func BridgeDestroy(object *Object) error {
	bridge := &Bridge{
		base: object,
	}
	// todo(source, sink 当前是空指针, panic)
	err := bridge.source.AudioStreamRXClose()
	if err != nil {
		return err
	}
	return bridge.sink.AudioStreamTXClose()
}

func BridgeBaseCreate(source *AudioStream, sink *AudioStream, name string) (*Bridge, error) {
	if source == nil || sink == nil {
		return nil, fmt.Errorf("source or sink is nil")
	}

	bridge := &Bridge{
		base:   ObjectInit(name),
		source: source,
		sink:   sink,
		codec:  nil,
		frame:  Frame{},
	}

	bridge.base.Destroy = BridgeDestroy
	bridge.base.Process = BridgeProcess

	return bridge, nil
}

func LinearBridgeCreate(source, sink *AudioStream, codecManager *CodecManager, name string) (*Object, error) {
	var (
		descriptor *CodecDescriptor
		frameSize  int64
		bridge     *Bridge
		err        error
	)

	bridge, err = BridgeBaseCreate(source, sink, name)
	if err != nil {
		return nil, err
	}

	descriptor = source.RXDescriptor
	frameSize = CodecLinearFrameSizeCalculate(descriptor.SamplingRate, descriptor.ChannelCount)
	bridge.frame.CodecFrame.Buffer = bytes.NewBuffer(make([]byte, frameSize))

	if err = source.AudioStreamRXOpen(nil); err != nil {
		return nil, err
	}
	if err = sink.AudioStreamTXOpen(nil); err != nil {
		source.AudioStreamRXClose()
		return nil, err
	}

	return bridge.base, nil
}

func NullBridgeCreate(source, sink *AudioStream, codecManager *CodecManager, name string) (*Object, error) {
	var (
		codec     *Codec
		frameSize int64
		bridge    *Bridge
		err       error
	)

	bridge, err = BridgeBaseCreate(source, sink, name)
	if err != nil {
		return nil, err
	}
	bridge.base.Process = NullBridgeProcess

	codec, err = codecManager.CodecManagerCodecGet(source.RXDescriptor)
	if err != nil {
		return nil, err
	}

	frameSize = source.RXDescriptor.CodecFrameSizeCalculate(codec.Attribs)
	bridge.codec = codec
	bridge.frame.CodecFrame.Buffer = bytes.NewBuffer(make([]byte, frameSize))

	if err = source.AudioStreamRXOpen(nil); err != nil {
		return nil, err
	}
	if err = sink.AudioStreamTXOpen(nil); err != nil {
		source.AudioStreamRXClose()
		return nil, err
	}

	return bridge.base, nil
}

/**
 * Create bridge of audio streams.
 * @param source the source audio stream
 * @param sink the sink audio stream
 * @param codec_manager the codec manager
 * @param name the informative name used for debugging
 * @param pool the pool to allocate memory from
 */
func BridgeCreate(source, sink *AudioStream, manager *CodecManager, name string) (*Object, error) {
	if source == nil || sink == nil {
		return nil, fmt.Errorf("source or sink is nil")
	}

	err := source.AudioStreamRXValidate(sink.TXDescriptor, sink.TXEventDescriptor)
	if err != nil {
		return nil, err
	}
	err = sink.AudioStreamTXValidate(source.RXDescriptor, source.RXEventDescriptor)
	if err != nil {
		return nil, err
	}

	if CodecDescriptorsMatch(source.RXDescriptor, sink.TXDescriptor) {
		return NullBridgeCreate(source, sink, manager, name)
	}

	if !CodecLPcmDescriptorMatch(source.RXDescriptor) {
		codec, err := manager.CodecManagerCodecGet(sink.TXDescriptor)
		if err != nil {
			return nil, err
		}
		encoder := EncoderCreate(sink, codec)
		sink = encoder
	}

	if source.RXDescriptor.SamplingRate != sink.TXDescriptor.SamplingRate {
		/* set resampler before bridge */
		resampler, err := ReSamplerCreate(source, sink)
		if err != nil {
			return nil, err
		}
		source = resampler
	}

	return LinearBridgeCreate(source, sink, manager, name)
}
