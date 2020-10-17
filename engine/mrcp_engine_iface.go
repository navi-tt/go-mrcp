package engine

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 5:47 下午
 */

/** Destroy engine */
func MRCPEngineVirtualDestroy(e *MRCPEngine) error {
	return nil
}

/** Open engine */
func MRCPEngineVirtualOpen(e *MRCPEngine) error {
	return nil
}

/** Response to open engine request */
func MRCPEngineOnOpen(e *MRCPEngine, status bool) {

}

/** Close engine */
func MRCPEngineVirtualClose(e *MRCPEngine) error {
	return nil
}
