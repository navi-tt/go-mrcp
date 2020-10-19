package engine

//
//import (
//	"github.com/navi-tt/go-mrcp/apr/memory"
//	"log"
//)
//
//const (
//	/** [REQUIRED] Symbol name of the main entry point in plugin DSO */
//	MRCP_PLUGIN_ENGINE_SYM_NAME = "mrcp_plugin_create"
//	/** [REQUIRED] Symbol name of the vesrion number entry point in plugin DSO */
//	MRCP_PLUGIN_VERSION_SYM_NAME = "mrcp_plugin_version"
//	/** [IMPLIED] Symbol name of the log accessor entry point in plugin DSO */
//	MRCP_PLUGIN_LOGGER_SYM_NAME = "mrcp_plugin_logger_set"
//	/** [IMPLIED] Symbol name of the log source accessor entry point in plugin DSO */
//	MRCP_PLUGIN_LOG_SOURCE_SYM_NAME = "mrcp_plugin_log_source_set"
//)
//
//const (
//	/** major version
//	 * Major API changes that could cause compatibility problems for older
//	 * plugins such as structure size changes.  No binary compatibility is
//	 * possible across a change in the major version.
//	 */
//	PLUGIN_MAJOR_VERSION = 1
//	/** minor version
//	 * Minor API changes that do not cause binary compatibility problems.
//	 * Reset to 0 when upgrading PLUGIN_MAJOR_VERSION
//	 */
//	PLUGIN_MINOR_VERSION = 5
//	/** patch level
//	 * The Patch Level never includes API changes, simply bug fixes.
//	 * Reset to 0 when upgrading PLUGIN_MINOR_VERSION
//	 */
//	PLUGIN_PATCH_VERSION = 0
//)
//
///** Prototype of engine creator (entry point of plugin DSO) */
//type MRCPPluginCreatorFunc func ()  *MRCPEngine
//
///** Prototype of log accessor (entry point of plugin DSO) */
//type  MRCPPluginLogAccessorFunc func(logger log.Logger)error
//
///** Prototype of log source accessor (entry point of plugin DSO) */
////typedef apt_bool_t (*mrcp_plugin_log_source_accessor_f)(apt_log_source_t *log_source);
