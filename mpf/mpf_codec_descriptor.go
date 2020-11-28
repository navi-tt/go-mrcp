package mpf

import (
	"bytes"
	"github.com/navi-tt/go-mrcp/apr"
	"strings"
)

/** Codec frame time base in msec */
const CODEC_FRAME_TIME_BASE = 10

const LPCM_CODEC_NAME = "LPCM"

const LPCM_CODEC_NAME_LENGTH = len(LPCM_CODEC_NAME) - 1

/** Bytes per sample for linear pcm */
const BYTES_PER_SAMPLE = 2

/** Bits per sample for linear pcm */
const BITS_PER_SAMPLE = 16

var lpcmAttribs = CodecAttribs{
	Name:          LPCM_CODEC_NAME,
	BitsPerSample: 16,
	SampleRates:   MPF_SAMPLE_RATE_8000 | MPF_SAMPLE_RATE_16000 | MPF_SAMPLE_RATE_32000 | MPF_SAMPLE_RATE_48000,
}

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
	Size int64
}

/** Get sampling rate mask (mpf_sample_rate_e) by integer value  */
func SampleRateMaskGet(samplingRate uint16) int {
	switch samplingRate {
	case 8000:
		return MPF_SAMPLE_RATE_8000
	case 16000:
		return MPF_SAMPLE_RATE_16000
	case 32000:
		return MPF_SAMPLE_RATE_32000
	case 48000:
		return MPF_SAMPLE_RATE_48000
	}
	return MPF_SAMPLE_RATE_NONE
}

func SamplingRateCheck(samplingRate uint16, mask int) bool {
	return SampleRateMaskGet(samplingRate)&mask > 0
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
	var descriptor = CodecDescriptorCreate()
	descriptor.PayloadType = uint8(RTP_PT_UNKNOWN)
	descriptor.Name = lpcmAttribs.Name
	descriptor.SamplingRate = samplingRate
	descriptor.ChannelCount = channelCount
	return nil
}

/** Create codec descriptor by capabilities */
func CodecDescriptorCreateByCapabilities(capabilities *CodecCapabilities, peer *CodecDescriptor) *CodecDescriptor {
	var (
		descriptor *CodecDescriptor
		attribs    *CodecAttribs
	)

	if capabilities != nil && peer != nil {
		attribs = CodecCapabilitiesAttribsFind(capabilities, peer)
	}

	if attribs == nil || peer == nil || capabilities == nil {
		return CodecLPcmDescriptorCreate(8000, 1)
	}

	descriptor = &CodecDescriptor{
		PayloadType:  peer.PayloadType,
		Name:         peer.Name,
		SamplingRate: peer.SamplingRate,
		ChannelCount: peer.ChannelCount,
		Format:       peer.Format,
		Enabled:      peer.Enabled,
	}

	if !strings.EqualFold(peer.Name, attribs.Name) {
		descriptor.PayloadType = RTP_PT_UNKNOWN
		descriptor.Name = attribs.Name
	}
	return descriptor
}

/** Match two codec descriptors */
func CodecDescriptorsMatch(descriptor1, descriptor2 *CodecDescriptor) bool {
	var match bool
	if descriptor1.PayloadType < RTP_PT_DYNAMIC && descriptor2.PayloadType < RTP_PT_DYNAMIC {
		if descriptor1.PayloadType == descriptor2.PayloadType {
			match = true
		}
	} else {
		if strings.EqualFold(descriptor1.Name, descriptor2.Name) {
			if descriptor1.SamplingRate == descriptor2.SamplingRate && descriptor1.ChannelCount == descriptor2.ChannelCount {
				match = true
			}
		}
	}
	return match
}

/** Match specified codec descriptor and the default lpcm one */
func CodecLPcmDescriptorMatch(descriptor *CodecDescriptor) bool {
	return strings.EqualFold(descriptor.Name, lpcmAttribs.Name)
}

/** Add default (linear PCM) capabilities */
func CodecDefaultCapabilitiesAdd(capabilities *CodecCapabilities) {
	capabilities.CodecCapabilitiesAdd(MPF_SAMPLE_RATE_8000, lpcmAttribs.Name)
}

/** Match codec descriptor by attribs specified */
func CodecDescriptorMatchByAttribs(descriptor, staticDescriptor *CodecDescriptor, attribs *CodecAttribs) bool {
	var match bool
	if descriptor.PayloadType < RTP_PT_DYNAMIC {
		if staticDescriptor != nil && staticDescriptor.PayloadType == descriptor.PayloadType {
			descriptor.Name = staticDescriptor.Name
			descriptor.SamplingRate = staticDescriptor.SamplingRate
			descriptor.ChannelCount = staticDescriptor.ChannelCount
			match = true
		}
	} else {
		if strings.EqualFold(attribs.Name, descriptor.Name) {
			if SamplingRateCheck(descriptor.SamplingRate, attribs.SampleRates) {
				match = true
			}
		}
	}
	return match
}

/** Find matched descriptor in codec list */
func CodecListDescriptorFind(codecList *CodecList, descriptor *CodecDescriptor) *CodecDescriptor {
	var (
		matchedDescriptor *CodecDescriptor
	)

	for i := 0; i < codecList.DescriptorArr.Stack.Size(); i++ {
		matchedDescriptor = codecList.DescriptorArr.ArrayHeaderIndex(i).(*CodecDescriptor)
		if CodecDescriptorsMatch(descriptor, matchedDescriptor) {
			return matchedDescriptor
		}
	}
	return nil
}

/** Find matched attribs in codec capabilities by descriptor specified */
func CodecCapabilitiesAttribsFind(capabilities *CodecCapabilities, descriptor *CodecDescriptor) *CodecAttribs {
	var attribs *CodecAttribs
	for i := 0; i < capabilities.AttribArr.Stack.Size(); i++ {
		attribs = capabilities.AttribArr.ArrayHeaderIndex(i).(*CodecAttribs)
		if SamplingRateCheck(descriptor.SamplingRate, attribs.SampleRates) {
			return attribs
		}
	}
	return nil
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
	var (
		descriptor *CodecDescriptor
		status     bool
	)
	if capabilities == nil {
		return false
	}
	for i := 0; i < c.DescriptorArr.Stack.Size(); i++ {
		descriptor = c.DescriptorArr.ArrayHeaderIndex(i).(*CodecDescriptor)
		if !descriptor.Enabled {
			continue
		}
		if CodecCapabilitiesAttribsFind(capabilities, descriptor) != nil {
			status = true
		} else {
			descriptor.Enabled = false
		}
	}

	return status
}

/** Intersect two codec lists */
func (cl *CodecList) CodecListsIntersect(c *CodecList) bool {
	var descriptor1, descriptor2 *CodecDescriptor
	cl.PrimaryDescriptor = nil
	cl.EventDescriptor = nil
	c.PrimaryDescriptor = nil
	c.EventDescriptor = nil

	/* find only one match for primary and named event descriptors,
	set the matched descriptors as preferred, disable the others */
	for i := 0; i < cl.DescriptorArr.Stack.Size(); i++ {
		descriptor1 = cl.DescriptorArr.ArrayHeaderIndex(i).(*CodecDescriptor)
		if descriptor1.Enabled == false {
			/* this descriptor has been already disabled, process only enabled ones */
			continue
		}

		/* check whether this is a named event descriptor */

		if EventDescriptorCheck(descriptor1) {
			/* named event descriptor */
			if cl.EventDescriptor != nil {
				/* named event descriptor has been already set, disable this one */
				descriptor1.Enabled = false
			} else {
				/* find if there is a match */
				descriptor2 = CodecListDescriptorFind(c, descriptor1)
				if descriptor2 != nil && descriptor2.Enabled {
					descriptor1.Enabled = true
					cl.EventDescriptor = descriptor1
					c.EventDescriptor = descriptor2
				} else {
					/* no match found, disable this descriptor */
					descriptor1.Enabled = false
				}
			}
		} else {
			/* primary descriptor */
			if cl.PrimaryDescriptor != nil {
				/* primary descriptor has been already set, disable this one */
				descriptor1.Enabled = false
			} else {
				/* find if there is a match */
				descriptor2 = CodecListDescriptorFind(c, descriptor1)
				if descriptor2 != nil && descriptor2.Enabled {
					descriptor1.Enabled = true
					cl.PrimaryDescriptor = descriptor1
					c.PrimaryDescriptor = descriptor2
				} else {
					/* no match found, disable this descriptor */
					descriptor1.Enabled = false
				}
			}
		}
	}

	for i := 0; i < c.DescriptorArr.Stack.Size(); i++ {
		descriptor2 = c.DescriptorArr.ArrayHeaderIndex(i).(*CodecDescriptor)
		if descriptor2 == c.PrimaryDescriptor || descriptor2 == c.EventDescriptor {
			descriptor2.Enabled = true
		} else {
			descriptor2.Enabled = false
		}
	}

	/* if primary descriptor is disabled or not set, return false */
	if cl.PrimaryDescriptor == nil || !cl.PrimaryDescriptor.Enabled {
		return false
	}

	return true
}

/** Compare two codec lists */
func (cl1 *CodecList) CodecListsCompare(cl2 *CodecList) bool {
	var (
		exit                     = false
		descriptor1, descriptor2 *CodecDescriptor
		i, j                     int
	)

	for !exit {
		for i < cl1.DescriptorArr.Stack.Size() {
			descriptor1 = cl1.DescriptorArr.ArrayHeaderIndex(i).(*CodecDescriptor)
			if descriptor1.Enabled {
				break
			}
			i++
		}

		for j < cl2.DescriptorArr.Stack.Size() {
			descriptor2 = cl2.DescriptorArr.ArrayHeaderIndex(j).(*CodecDescriptor)
			if descriptor2.Enabled {
				break
			}
			j++
		}

		if i < cl1.DescriptorArr.Stack.Size() && j < cl2.DescriptorArr.Stack.Size() {
			if !CodecDescriptorsMatch(descriptor1, descriptor2) {
				return false
			}
			i++
			j++
		} else {
			exit = true
		}
	}

	if i != cl1.DescriptorArr.Stack.Size() || j != cl2.DescriptorArr.Stack.Size() {
		return false
	}

	return true
}
