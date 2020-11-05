package mpf

import "strings"

/** MPF media state */
type MediaState = int

const (
	MPF_MEDIA_DISABLED MediaState = iota /**< disabled media */
	MPF_MEDIA_ENABLED                    /**< enabled media */
)

/** RTP media (local/remote) descriptor */
type RtpMediaDescriptor struct {

	/** Media state (disabled/enabled)*/
	state MediaState
	/** IP address */
	ip string
	/** External (NAT) IP address */
	extIp string
	/** Port */
	port uint16
	/** Stream mode (send/receive) */
	direction StreamDirection
	/** Packetization time */
	ptime uint16
	/** Codec list */
	codecList CodecList
	/** Media identifier */
	mid int64
	/** Position, order in SDP message (0,1,...) */
	id int64
}

/** RTP stream descriptor */
type RtpStreamDescriptor struct {

	/** Stream capabilities */
	capabilities *StreamCapabilities
	/** Local media descriptor */
	local *RtpMediaDescriptor
	/** Remote media descriptor */
	remote *RtpMediaDescriptor
	/** Settings loaded from config */
	settings *RtpSettings
}

/** RTP termination descriptor */
type RtpTerminationDescriptor struct {

	/** Audio stream descriptor */
	audio RtpStreamDescriptor
	/** Video stream descriptor */
	video RtpStreamDescriptor
}

/** Jitter buffer configuration */
type JbConfig struct {

	/** Min playout delay in msec */
	minPlayOutDelay uint32
	/** Initial playout delay in msec */
	initialPlayOutDelay uint32
	/** Max playout delay in msec */
	maxPlayOutDelay uint32
	/** Mode of operation of the jitter buffer: static - 0, adaptive - 1 */
	adaptive byte
	/** Enable/disable time skew detection */
	timeSkewDetection byte
}

/** RTCP BYE transmission policy */
type ByePolicy = int

const (
	RTCP_BYE_DISABLE       ByePolicy = iota /**< disable RTCP BYE transmission */
	RTCP_BYE_PER_SESSION                    /**< transmit RTCP BYE at the end of session */
	RTCP_BYE_PER_TALKSPURT                  /**< transmit RTCP BYE at the end of each talkspurt (input) */
)

/** RTP factory config */
type RtpConfig struct {
	/** Local IP address to bind to */
	ip string
	/** External (NAT) IP address */
	extIp string
	/** Min RTP port */
	rtpPortMin uint16
	/** Max RTP port */
	rtpPortMax uint16
	/** Current RTP port */
	rtpPortCur uint16
}

/** RTP settings */
type RtpSettings struct {

	/** Packetization time */
	ptime uint16
	/** Codec list */
	codecList CodecList
	/** Preference in offer/anwser: 1 - own(local) preference, 0 - remote preference */
	ownPreference bool
	/** Enable/disable RTCP support */
	rtcp bool
	/** RTCP BYE policy */
	rtcpByePolicy ByePolicy
	/** RTCP report transmission interval */
	rtcpTXInterval uint16
	/** RTCP rx resolution (timeout to check for a new RTCP message) */
	rtcpRXResolution uint16
	/** Jitter buffer config */
	jbConfig JbConfig
}

/** Initialize RTP media descriptor */
func RtpMediaDescriptorInit(media *RtpMediaDescriptor) {
	media.state = MPF_MEDIA_DISABLED
	media.ip = ""
	media.extIp = ""
	media.port = 0
	media.direction = STREAM_DIRECTION_NONE
	media.ptime = 0
	CodecListReset(&media.codecList)
	media.mid = 0
	media.id = 0
}

/** Initialize RTP stream descriptor */
func RtpStreamDescriptorInit(descriptor *RtpStreamDescriptor) {
	descriptor.capabilities = nil
	descriptor.local = nil
	descriptor.remote = nil
	descriptor.settings = nil
}

/** Initialize RTP termination descriptor */
func RtpTerminationDescriptorInit(rtpDescriptor *RtpTerminationDescriptor) {
	RtpStreamDescriptorInit(&rtpDescriptor.audio)
	RtpStreamDescriptorInit(&rtpDescriptor.video)
}

/** Initialize JB config */
func JbConfigInit(jbConfig *JbConfig) {
	jbConfig.adaptive = 0
	jbConfig.initialPlayOutDelay = 0
	jbConfig.minPlayOutDelay = 0
	jbConfig.maxPlayOutDelay = 0
	jbConfig.timeSkewDetection = 1
}

/** Allocate RTP config */
func RtpConfigAlloc() *RtpConfig {
	rtpConfig := RtpConfig{
		ip:         "",
		extIp:      "",
		rtpPortMin: 0,
		rtpPortMax: 0,
		rtpPortCur: 0,
	}
	return &rtpConfig
}

/** Allocate RTP settings */
func RtpSettingsAlloc() *RtpSettings {
	rtpSettings := RtpSettings{}

	rtpSettings.ptime = 0
	CodecListInit(&rtpSettings.codecList, 0)
	rtpSettings.rtcpByePolicy = RTCP_BYE_DISABLE
	rtpSettings.rtcpTXInterval = 0
	rtpSettings.rtcpRXResolution = 0
	JbConfigInit(&rtpSettings.jbConfig)
	return &rtpSettings
}

/** Allocate RTP termination descriptor */
func RtpTerminationDescriptorAlloc() *RtpTerminationDescriptor {
	rtpDescriptor := &RtpTerminationDescriptor{}
	RtpTerminationDescriptorInit(rtpDescriptor)
	return rtpDescriptor
}

/** Allocate RTP media descriptor */
func RtpMediaDescriptorAlloc() *RtpMediaDescriptor {
	media := &RtpMediaDescriptor{}
	RtpMediaDescriptorInit(media)
	return media
}

/** Copy RTP media descriptor */
func RtpMediaDescriptorCopy(srcMedia *RtpMediaDescriptor) *RtpMediaDescriptor {
	media := &RtpMediaDescriptor{}
	if srcMedia != nil {
		media.state = srcMedia.state
		media.ip = srcMedia.ip
		media.extIp = srcMedia.extIp
		media.port = srcMedia.port
		media.direction = srcMedia.direction
		media.ptime = srcMedia.ptime
		media.codecList = srcMedia.codecList
		media.mid = srcMedia.mid
		media.id = srcMedia.id
	}
	return media
}

/** Compare RTP media descriptors (return TRUE, if identical) */
func RtpMediaDescriptorsCompare(media1, media2 *RtpMediaDescriptor) bool {
	if media1 == nil || media2 == nil {
		return false
	}

	if media1.state != media2.state {
		return false
	}

	if !strings.EqualFold(media1.ip, media2.ip) {
		return false
	}

	if !strings.EqualFold(media1.extIp, media2.extIp) {
		return false
	}

	if media1.port != media2.port {
		return false
	}

	if media1.direction != media2.direction {
		return false
	}

	if media1.ptime != media2.ptime {
		return false
	}

	if !CodecListsCompare(&media1.codecList, &media2.codecList) {
		return false
	}
	if media1.mid != media2.mid {
		return false
	}

	if media1.id != media2.id {
		return false
	}

	return true
}
