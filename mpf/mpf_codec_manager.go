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

/** Create codec manager */
func CodecManagerCreate(codecCount int64) *CodecManager {
	return nil
}

/** Destroy codec manager */
func CodecManagerDestroy(cm *CodecManager) {}

/** Register codec in codec manager */
func (cm *CodecManager) CodecManagerCodecRegister(codec *Codec) error {
	return nil
}

/** Get (allocate) codec by codec descriptor */
func (cm *CodecManager) CodecManagerCodecGet(descriptor *CodecDescriptor) *Codec {
	return nil
}

/** Get (allocate) list of available codecs */
func (cm *CodecManager) CodecManagerCodecListGet(codecList *CodecList) error {
	return nil
}

/** Load (allocate) list of codecs  */
func (cm *CodecManager) CodecManagerCodecListLoad(codecList *CodecList, str string) error {
	return nil
}

/** Find codec by name  */
func (cm *CodecManager) CodecManagerCodecFind(codecName string) *Codec {
	return nil
}
