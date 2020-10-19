package engine

import (
	"github.com/navi-tt/go-mrcp/mrcp"
)

/** Termorary define legacy mrcp_resource_engine_t as mrcp_engine_t */
type MRCPResourceEngine = MRCPEngine

/**
 * Create resource engine
 * @deprecated @see mrcp_engine_create
 */
func MRCPResourceEngineCreate(rid mrcp.MRCPResourceId, obj interface{}, vTable *MRCPEngineMethodVTable) *MRCPEngine {
	return MRCPEngineCreate(rid, obj, vTable)
}
