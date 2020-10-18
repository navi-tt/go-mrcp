package mpf

import (
	"github.com/navi-tt/go-mrcp/apr"
	"github.com/navi-tt/go-mrcp/apr/memory"
)

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 4:58 下午
 */

/** Opaque codec manager declaration */
type CodecManager struct {
	pool *memory.AprPool // Memory pool

	CodecArr []apr.AprArrayHeader // Dynamic (resizable) array of codecs (mpf_codec_t*

	EventDescriptor *CodecDescriptor // Default named event descriptor
}
