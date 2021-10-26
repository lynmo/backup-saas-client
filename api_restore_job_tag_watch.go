package swagger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

/*
RestoreJobTagApiService Watch restorejobs of a tenant
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenant tenant id

@return V1alpha1RestoreJobList
*/

func (a *RestoreJobTagApiService) WatchRestoreJobs(ctx context.Context, tenant string) (WatchInterface, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restorejobs"
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

	return NewStreamWatcher(localVarHttpResponse.Body, WatchDecoder(RestoreJobWatchDecoder)), nil
}

func RestoreJobWatchDecoder(b []byte) (interface{}, error) {
	restorejob := &V1alpha1RestoreJob{}
	if err := json.Unmarshal(b, restorejob); err != nil {
		return nil, err
	}
	return *restorejob, nil
}
