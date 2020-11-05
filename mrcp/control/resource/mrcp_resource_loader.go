package resource

import (
	"github.com/navi-tt/go-mrcp/mrcp"
)

/** Resource loader */
type MRCPResourceLoader struct {
	factory *MRCPResourceFactory

}

/** Create MRCP resource loader */
func MRCPResourceLoaderCreate(loadAll bool) *MRCPResourceLoader {
	return nil
}

/** Load all MRCP resources */
func (loader *MRCPResourceLoader) MRCPResourcesLoad() error {
	return nil
}

/** Load MRCP resource by resource name */
func (loader *MRCPResourceLoader) MRCPResourceLoad(name string) error {
	return nil
}

/** Load MRCP resource by resource identifier */
func (loader *MRCPResourceLoader) MRCPResourceLoadById(id mrcp.MRCPResourceId) error {
	return nil
}

/** Get MRCP resource factory */
func (loader *MRCPResourceLoader) MRCPResourceFactoryGet() (*MRCPResourceFactory, error) {
	return nil, nil
}
