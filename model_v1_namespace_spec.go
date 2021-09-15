/*
 * Backup Service
 *
 * Backup Service OpenAPI
 *
 * API version: v0.0.0-master+$Format:%h$
 * Contact: shaofeng@jibudata.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// NamespaceSpec describes the attributes on a Namespace.
type V1NamespaceSpec struct {
	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
	Finalizers []string `json:"finalizers,omitempty"`
}
