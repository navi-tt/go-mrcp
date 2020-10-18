package mpf

/** Codec */
type MPFCodec struct {

	/** Codec manipulators (encode, decode, dissect) */
	//vtable *MPFCodecVTable
	/** Codec attributes (capabilities) */
	//const mpf_codec_attribs_t    *attribs;
	/** Optional static codec descriptor (pt < 96) */
	StaticDescriptor *MPFCodecDescriptor
}
