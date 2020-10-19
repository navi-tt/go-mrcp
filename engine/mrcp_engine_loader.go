package engine

import (
	"container/list"
)

/** Engine loader declaration */
type MRCPEngineLoader struct {
	/** Table of plugins (apr_dso_handle_t*) */
	pluginsList *list.List
	pluginsMap  map[string]*list.Element
	//pool        *memory.AprPool
}

/** Create engine loader */
func MRCPEngineLoaderCreate() *MRCPEngineLoader {
	return &MRCPEngineLoader{}
}

/** Destroy engine loader */
func MRCPEngineLoaderDestroy(loader *MRCPEngineLoader) error {
	return nil
}

/** Unload loaded plugins */
func (loader *MRCPEngineLoader) MRCPEngineLoaderPluginsUnload() error {
	return nil
}

/** Load engine plugin */
func (loader *MRCPEngineLoader) MRCPEngineLoaderPluginLoad(id string, path string, config *MRCPEngineConfig) *MRCPEngine {
	return nil
}
