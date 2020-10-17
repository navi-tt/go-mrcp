package control

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message"
	"github.com/navi-tt/go-mrcp/toolkit"
)

/** MRCP parser */
type MRCPParser struct {
	base            *toolkit.AptMessageParser // todo(toolkit.AptMessageParser 还没完成)
	ResourceFactory *resource.MRCPResourceFactory
	Resource        *resource.MRCPResource
}

/** Create MRCP stream parser */
func MRCPParserCreate(f *resource.MRCPResourceFactory, p *memory.AprPool) *MRCPParser {
	return &MRCPParser{}
}

/** Set resource by name to be used for parsing of MRCPv1 messages */
func (parser *MRCPParser) MRCPParserResourceSet(name string) {

}

/** Set verbose mode for the parser */
func (parser *MRCPParser) MRCPParserVerboseSet(verbose bool) {

}

/** Parse MRCP stream */
func (parser *MRCPParser) MRCPParserRun(stream *toolkit.AptTextStream, message []*message.MRCPMessage) toolkit.AptMessageStage {
	return 0
}

/** MRCP generator */
type MRCPGenerator struct {
	base            *toolkit.AptMessageGenerator // todo(toolkit.AptMessageGenerator 还没完成)
	ResourceFactory *resource.MRCPResourceFactory
}

/** Create MRCP stream generator */
func MRCPGeneratorCreate(f *resource.MRCPResourceFactory, p *memory.AprPool) *MRCPGenerator {
	return &MRCPGenerator{}
}

/** Set verbose mode for the generator */
func (g *MRCPGenerator) MRCPGeneratorVerboseSet(verbose bool) {

}

/** Generate MRCP stream */
func (g *MRCPGenerator) MRCPGeneratorRun(message *message.MRCPMessage, stream *toolkit.AptTextStream) toolkit.AptMessageStage {
	return 0
}

/** Generate MRCP message (excluding message body) */
func MRCPMessageGenerate(cf *resource.MRCPResourceFactory, message *message.MRCPMessage, stream *toolkit.AptTextStream) error {
	return nil
}
