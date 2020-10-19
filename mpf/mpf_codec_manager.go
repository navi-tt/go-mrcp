package mpf

import (
	"github.com/navi-tt/go-mrcp/apr"
)

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 4:58 下午
 */

/** Opaque codec manager declaration */
type CodecManager struct {
	// Memory pool

	CodecArr []apr.ArrayHeader // Dynamic (resizable) array of codecs (mpf_codec_t*

	EventDescriptor *CodecDescriptor // Default named event descriptor
}
