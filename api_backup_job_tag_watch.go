package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

/*
BackupJobTagApiService Watch backupjobs of a tenant
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenant tenant id

@return V1alpha1BackupJobList
*/

func (a *BackupJobTagApiService) WatchBackupJobs(ctx context.Context, tenant string) (WatchInterface, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupjobs"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", fmt.Sprintf("%v", tenant), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("watch", "true")

	// set Accept header
	localVarHeaderParams["Accept"] = "application/json"
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil {
		return nil, err
	}
	if localVarHttpResponse == nil {
		return nil, fmt.Errorf("response is nil")
	}

	return NewStreamWatcher(localVarHttpResponse.Body, WatchDecoder(BackupJobWatchDecoder)), nil
}

func BackupJobWatchDecoder(b []byte) (interface{}, error) {
	backupjob := &V1alpha1BackupJob{}
	if err := json.Unmarshal(b, backupjob); err != nil {
		return nil, err
	}
	return *backupjob, nil
}
