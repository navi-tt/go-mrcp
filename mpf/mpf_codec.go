package mpf

/** Codec */
type Codec struct {

	/** Codec manipulators (encode, decode, dissect) */
	//vtable *CodecVTable
	/** Codec attributes (capabilities) */
	//const mpf_codec_attribs_t    *attribs;
	/** Optional static codec descriptor (pt < 96) */
	StaticDescriptor *CodecDescriptor
}
