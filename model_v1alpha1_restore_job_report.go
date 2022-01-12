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

type V1alpha1RestoreJobReport struct {
	BackedUpItems int32 `json:"backedUpItems,omitempty"`
	Details map[string][]string `json:"details,omitempty"`
	EndTime string `json:"endTime,omitempty"`
	RestoredItems int32 `json:"restoredItems,omitempty"`
	Results map[string]V1alpha1Result `json:"results,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	TotalItems int32 `json:"totalItems,omitempty"`
	TotalPVC int32 `json:"totalPVC,omitempty"`
}
