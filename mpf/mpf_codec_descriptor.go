package mpf

import (
	"bytes"
	"github.com/navi-tt/go-mrcp/apr"
)

/** Codec frame time base in msec */
const CODEC_FRAME_TIME_BASE = 10

/** Bytes per sample for linear pcm */
const BYTES_PER_SAMPLE = 2

/** Bits per sample for linear pcm */
const BITS_PER_SAMPLE = 16

/** Supported sampling rates */
type SampleRates = int

const (
	MPF_SAMPLE_RATE_NONE  SampleRates = 0x00
	MPF_SAMPLE_RATE_8000  SampleRates = 0x01
	MPF_SAMPLE_RATE_16000 SampleRates = 0x02
	MPF_SAMPLE_RATE_32000 SampleRates = 0x04
	MPF_SAMPLE_RATE_48000 SampleRates = 0x08

	MPF_SAMPLE_RATE_SUPPORTED = MPF_SAMPLE_RATE_8000 | MPF_SAMPLE_RATE_16000 |
		MPF_SAMPLE_RATE_32000 | MPF_SAMPLE_RATE_48000
)

/** Codec descriptor */
type CodecDescriptor struct {
	PayloadType  uint8  // Payload type used in RTP packet
	Name         string // Codec name
	SamplingRate uint16 // Sampling rate
	ChannelCount uint8  // Channel count
	Format       string // Codec dependent additional format
	Enabled      bool   // Enabled/disabled state
}

/** List of codec descriptors */
type CodecList struct {
	/** Dynamic array of codec descriptors (mpf_codec_descriptor_t) */
	DescriptorArr *apr.ArrayHeader
	/** Preffered primary (audio/video codec) descriptor from descriptor_arr */
	PrimaryDescriptor *CodecDescriptor
	/** Preffered named event (telephone-event) descriptor from descriptor_arr */
	EventDescriptor *CodecDescriptor
}

/** Codec attributes */
type CodecAttribs struct {
	/** Codec name */
	Name string
	/** Bits per sample */
	BitsPerSample uint8
	/** Supported sampling rates (mpf_sample_rates_e) */
	SampleRates int
}

/** List of codec attributes (capabilities) */
type CodecCapabilities struct {
	/** Dynamic array of codec attributes (mpf_codec_attrribs_t) */
	AttribArr *apr.ArrayHeader
	/** Allow/support named events */
	AllowNamedEvents bool
}

/** Codec frame */
type CodecFrame struct {
	/** Raw buffer, which may contain encoded or decoded data */
	Buffer *bytes.Buffer
	/** Buffer size */
	//size int64
}

/** Initialize codec descriptor */
func CodecDescriptorInit(descriptor *CodecDescriptor) {
	descriptor.PayloadType = 0
	descriptor.Name = ""
	descriptor.SamplingRate = 0
	descriptor.ChannelCount = 0
	descriptor.Format = ""
	descriptor.Enabled = true

}

/** Initialize codec descriptor */
func CodecDescriptorCreate() *CodecDescriptor {
	d := &CodecDescriptor{}
	CodecDescriptorInit(d)
	return d
}

/** Calculate encoded frame size in bytes */
func (d *CodecDescriptor) CodecFrameSizeCalculate(attribs *CodecAttribs) int64 {
	return int64(d.ChannelCount) * int64(attribs.BitsPerSample) * CODEC_FRAME_TIME_BASE * int64(d.SamplingRate) / 1000 / 8
	/* 1000 - msec per sec, 8 - bits per byte */
}

/** Calculate samples of the frame (ts) */
func (d *CodecDescriptor) CodecFrameSamplesCalculate() int64 {
	return int64(d.ChannelCount) * CODEC_FRAME_TIME_BASE * int64(d.SamplingRate) / 1000
}

/** Calculate linear frame size in bytes */
func CodecLinearFrameSizeCalculate(samplingRate uint16, channelCount uint8) int64 {
	return int64(channelCount) * BYTES_PER_SAMPLE * CODEC_FRAME_TIME_BASE * int64(samplingRate) / 1000
}

/** Reset list of codec descriptors */
func CodecListReset(c *CodecList) {
	c.DescriptorArr = nil
	c.PrimaryDescriptor = nil
	c.EventDescriptor = nil
}

/** Initialize list of codec descriptors */
func CodecListInit(c *CodecList, initialCount int) {
	c.DescriptorArr = &apr.ArrayHeader{Stack: apr.NewStack(initialCount)}
	c.PrimaryDescriptor = nil
	c.EventDescriptor = nil
}

/** Copy list of codec descriptors */
func CodecListCopy(c *CodecList, srcCodecList *CodecList) {
	c.DescriptorArr = apr.ArrayCopy(srcCodecList.DescriptorArr)
}

func CodecListsCompare(codecList1, codecList2 *CodecList) bool {
	var (
		exit                     = false
		descriptor1, descriptor2 *CodecDescriptor
		i, j                     int
	)

	for !exit {
		for i < codecList1.DescriptorArr.Stack.Size() {
			descriptor1 = codecList1.DescriptorArr.Stack.Index(i).(*CodecDescriptor)
			if descriptor1.Enabled {
				break
			}
			i++
		}
		for j < codecList2.DescriptorArr.Stack.Size() {
			descriptor2 = codecList1.DescriptorArr.Stack.Index(j).(*CodecDescriptor)
			if descriptor2.Enabled {
				break
			}
			j++
		}

		if i < codecList1.DescriptorArr.Stack.Size() && j < codecList2.DescriptorArr.Stack.Size() {
			if !CodecDescriptorsMatch(descriptor1, descriptor2) {
				return false
			}

			i++
			j++
		} else {
			exit = true
		}
	}

	if i != codecList1.DescriptorArr.Stack.Size() || j != codecList2.DescriptorArr.Stack.Size() {
		return false
	}

	return true
}

/** Increment number of codec descriptors in the list and return the descriptor to fill */
func CodecListAdd(c *CodecList) *CodecDescriptor {
	descriptor := &CodecDescriptor{}
	c.DescriptorArr.Stack.Push(descriptor)
	CodecDescriptorInit(descriptor)
	return descriptor
}

/** Determine if codec list is empty */
func (c *CodecList) CodecListIsEmpty() bool {
	return c.DescriptorArr.Stack.IsEmpty()
}

/** Get codec descriptor by index */
func (c *CodecList) CodecListDescriptorGet(id int) *CodecDescriptor {
	if id >= c.DescriptorArr.Stack.Size() {
		return nil
	}
	e := c.DescriptorArr.ArrayHeaderIndex(id)
	if e == nil {
		return nil
	}
	return e.(*CodecDescriptor)
}

/** Create linear PCM descriptor */
func CodecLPcmDescriptorCreate(samplingRate uint16, channelCount uint8) *CodecDescriptor {
	return nil
}

/** Create codec descriptor by capabilities */
func CodecDescriptorCreateByCapabilities(capabilities *CodecCapabilities, peer *CodecDescriptor) *CodecDescriptor {
	return nil
}

/** Match two codec descriptors */
func CodecDescriptorsMatch(descriptor1, descriptor2 *CodecDescriptor) bool {
	return false
}

/** Match specified codec descriptor and the default lpcm one */
func CodecLPcmDescriptorMatch(descriptor *CodecDescriptor) bool {
	return false
}

/** Match codec descriptor by attribs specified */
func CodecDescriptorMatchByAttribs(descriptor, staticDescriptor *CodecDescriptor, attribs *CodecAttribs) bool {
	return false
}

/** Initialize codec capabilities */
func (c *CodecCapabilities) CodecCapabilitiesInit(initialCount int) {
	c.AttribArr = apr.NewArrayHeader(initialCount)
	c.AllowNamedEvents = true
}

/** Clone codec capabilities */
func (c *CodecCapabilities) CodecCapabilitiesClone(src *CodecCapabilities) {
	c.AttribArr = apr.ArrayCopy(src.AttribArr)
	c.AllowNamedEvents = src.AllowNamedEvents
}

/** Merge codec capabilities */
func (c *CodecCapabilities) CodecCapabilitiesMerge(src *CodecCapabilities) {
	if !c.AllowNamedEvents && src.AllowNamedEvents {
		c.AllowNamedEvents = src.AllowNamedEvents
	}
	c.AttribArr.Append(src.AttribArr)
}

/** Add codec capabilities */
func (c *CodecCapabilities) CodecCapabilitiesAdd(sampleRates int, codecName string) {
	attribs := &CodecAttribs{
		Name:          codecName,
		BitsPerSample: 0,
		SampleRates:   sampleRates,
	}
	c.AttribArr.Stack.Push(attribs)
}

/** Add default (linear PCM) capabilities */
func (c *CodecCapabilities) CodecDefaultCapabilitiesAdd() {
}

/** Validate codec capabilities */
func (c *CodecCapabilities) CodecCapabilitiesValidate() {
	if c.AttribArr.Stack.IsEmpty() {
		c.CodecDefaultCapabilitiesAdd()
	}
}

/** Find matched descriptor in codec list */
func (c *CodecList) CodecListDescriptorFind(descriptor *CodecDescriptor) *CodecDescriptor {
	return nil
}

/** Match codec list with specified capabilities */
func (c *CodecList) CodecListMatch(capabilities *CodecCapabilities) bool {
	return false
}

/** Intersect two codec lists */
func (cl *CodecList) CodecListsIntersect(c *CodecList) bool {
	return false
}

/** Compare two codec lists */
func (cl *CodecList) CodecListsCompare(c *CodecList) bool {
	return false
}

/** Get sampling rate mask (mpf_sample_rate_e) by integer value  */
func SampleRateMaskGet(samplingRate uint16) int {
	return 0
}
