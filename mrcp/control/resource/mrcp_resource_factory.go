package resource

import (
	"github.com/navi-tt/go-mrcp/mrcp"
)

/** Resource factory definition (aggregation of resources) */
type MRCPResourceFactory struct {
	ResourceArray []*MRCPResource          // Array of MRCP resources (reference by id)
	resourceCount int64                    // Number of MRCP resources
	ResourceHash  map[string]*MRCPResource // Hash of MRCP resources (reference by name) todo(apt/tables/apr_hash.c/apr_hash_t)
}

/** Create MRCP resource factory */
func MRCPResourceFactoryCreate(resourceCount int64) *MRCPResourceFactory {
	return &MRCPResourceFactory{}
}

/** Destroy MRCP resource factory */
func MRCPResourceFactoryDestroy(factory *MRCPResourceFactory) error {
	return nil
}

/** Register MRCP resource */
func MRCPResourceRegister(factory *MRCPResourceFactory, resource *MRCPResource) error {
	return nil
}

/** Get MRCP resource by resource id */
func MRCPResourceGet(factory *MRCPResourceFactory, rid mrcp.MRCPResourceId) (*MRCPResource, error) {
	return nil, nil
}

/** Find MRCP resource by resource name */
func MRCPResourceFind(factory *MRCPResourceFactory, name string) (*MRCPResource, error) {
	return nil, nil
}
