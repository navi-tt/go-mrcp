package control

import (
	apr_toolkit "github.com/navi-tt/go-mrcp/apr-toolkit"
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message"
)

/** MRCP parser */
type MRCPParser struct {
	base            *apr_toolkit.AptMessageParser // todo(apr_toolkit.AptMessageParser 还没完成)
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
func (parser *MRCPParser) MRCPParserRun(stream *apr_toolkit.AptTextStream, message []*message.MRCPMessage) apr_toolkit.AptMessageStage {
	return 0
}

/** MRCP generator */
type MRCPGenerator struct {
	base            *apr_toolkit.AptMessageGenerator // todo(apr_toolkit.AptMessageGenerator 还没完成)
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
func (g *MRCPGenerator) MRCPGeneratorRun(message *message.MRCPMessage, stream *apr_toolkit.AptTextStream) apr_toolkit.AptMessageStage {
	return 0
}

/** Generate MRCP message (excluding message body) */
func MRCPMessageGenerate(cf *resource.MRCPResourceFactory, message *message.MRCPMessage, stream *apr_toolkit.AptTextStream) error {
	return nil
}
