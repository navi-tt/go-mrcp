package mpf

import (
	"fmt"
	"github.com/navi-tt/go-mrcp/apr"
	"strconv"
	"strings"
)

/** Opaque codec manager declaration */
type CodecManager struct {
	CodecArr *apr.ArrayHeader // Dynamic (resizable) array of codecs (mpf_codec_t*

	EventDescriptor *CodecDescriptor // Default named event descriptor
}

/** Create codec manager */
func CodecManagerCreate(codecCount int) *CodecManager {
	codecManager := CodecManager{
		CodecArr:        apr.NewArrayHeader(codecCount),
		EventDescriptor: EventDescriptorCreate(8000),
	}
	return &codecManager
}

/** Destroy codec manager */
func CodecManagerDestroy(cm *CodecManager) {}

/** Register codec in codec manager */
func (cm *CodecManager) CodecManagerCodecRegister(codec *Codec) error {
	if codec == nil || codec.Attribs == nil || codec.Attribs.Name == "" {
		return fmt.Errorf(`codec is nil, or codec.Attribs is nil, or codec.Attribs.Name is ""`)
	}
	cm.CodecArr.Stack.Push(codec)
	return nil
}

/** Get (allocate) codec by codec descriptor */
func (cm *CodecManager) CodecManagerCodecGet(descriptor *CodecDescriptor) (*Codec, error) {
	if descriptor == nil {
		return nil, fmt.Errorf("descriptor is nil")
	}
	for i := 0; i < cm.CodecArr.Stack.Size(); i++ {
		codec := cm.CodecArr.ArrayHeaderIndex(i).(*Codec)
		if CodecDescriptorMatchByAttribs(descriptor, codec.StaticDescriptor, codec.Attribs) {
			return codec.CodecClone(), nil
		}
	}
	return nil, nil
}

/** Get (allocate) list of available codecs */
func (cm *CodecManager) CodecManagerCodecListGet(codecList *CodecList) error {
	CodecListInit(codecList, cm.CodecArr.Stack.Size())
	for i := 0; i < cm.CodecArr.Stack.Size(); i++ {
		codec := cm.CodecArr.ArrayHeaderIndex(i).(*Codec)
		if staticDescriptor := codec.StaticDescriptor; staticDescriptor != nil {
			if descriptor := CodecListAdd(codecList); descriptor != nil {
				*descriptor = *staticDescriptor
			}
		}

	}
	return nil
}

func (cm *CodecManager) CodecManagerCodecParse(codecList *CodecList, codecDescStr string) error {
	var (
		codec      *Codec
		descriptor *CodecDescriptor
		separator  = "/"
		str        string
	)

	/* parse codec name */
	codecDescs := strings.Split(codecDescStr, separator)
	if len(codecDescs) > 0 {
		str = codecDescs[0]
	}
	if len(str) > 0 {
		name := str
		/* find codec by name */
		codec = cm.CodecManagerCodecFind(name)

		if codec != nil {
			descriptor = CodecListAdd(codecList)
			descriptor.Name = name

			/* set default attributes */
			if codec.StaticDescriptor != nil {
				descriptor.PayloadType = codec.StaticDescriptor.PayloadType
				descriptor.SamplingRate = codec.StaticDescriptor.SamplingRate
				descriptor.ChannelCount = codec.StaticDescriptor.ChannelCount
			} else {
				descriptor.PayloadType = RTP_PT_DYNAMIC
				descriptor.SamplingRate = 8000
				descriptor.ChannelCount = 1
			}
		} else {
			eventDescriptor := cm.EventDescriptor
			if eventDescriptor != nil && strings.EqualFold(eventDescriptor.Name, name) {
				descriptor = CodecListAdd(codecList)
				*descriptor = *eventDescriptor
			} else {
				return fmt.Errorf("no such codec [%s]", str)
			}
		}

		/* parse optional payload type */
		if len(codecDescs) > 1 {
			str = codecDescs[1]
		}
		if len(str) > 0 {
			payloadType, _ := strconv.Atoi(str)
			descriptor.PayloadType = uint8(payloadType)

			/* parse optional sampling rate */
			if len(codecDescs) > 2 {
				str = codecDescs[2]
			}
			if len(str) > 0 {
				samplingRate, _ := strconv.Atoi(str)
				descriptor.SamplingRate = uint16(samplingRate)

				/* parse optional channel count */
				if len(codecDescs) > 3 {
					str = codecDescs[3]
				}
				if len(str) > 0 {
					channelCount, _ := strconv.Atoi(str)
					descriptor.ChannelCount = uint8(channelCount)
				}
			}
		}
	}
	return nil
}

/** Load (allocate) list of codecs  */
func (cm *CodecManager) CodecManagerCodecListLoad(codecList *CodecList, str string) error {
	codecDescs := strings.Split(str, " ")
	for _, codecDesc := range codecDescs {
		err := cm.CodecManagerCodecParse(codecList, codecDesc)
		if err != nil {
			return err
		}
	}
	return nil
}

/** Find codec by name  */
func (cm *CodecManager) CodecManagerCodecFind(codecName string) *Codec {
	for i := 0; i < cm.CodecArr.Stack.Size(); i++ {
		codec := cm.CodecArr.ArrayHeaderIndex(i).(*Codec)
		if strings.EqualFold(codec.Attribs.Name, codecName) {
			return codec
		}
	}
	return nil
}
