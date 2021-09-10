# \RestoreJobTagApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRestoreJob**](RestoreJobTagApi.md#CreateRestoreJob) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restorejobs | Create a restorejob for a tenant.
[**DeleteRestoreJob**](RestoreJobTagApi.md#DeleteRestoreJob) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restorejobs/{restorejob} | Delete a restorejob for a tenant.
[**GetRestoreJob**](RestoreJobTagApi.md#GetRestoreJob) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restorejobs/{restorejob} | Retrieve restorejob details.
[**ListRestoreJobs**](RestoreJobTagApi.md#ListRestoreJobs) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restorejobs | List all restorejobs of a tenant
[**UpdateRestoreJob**](RestoreJobTagApi.md#UpdateRestoreJob) | **Put** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restorejobs/{restorejob} | Update a restorejob for a tenant.


# **CreateRestoreJob**
> V1alpha1RestoreJob CreateRestoreJob(ctx, tenant, body)
Create a restorejob for a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1RestoreJob**](V1alpha1RestoreJob.md)|  | 

### Return type

[**V1alpha1RestoreJob**](v1alpha1.RestoreJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRestoreJob**
> YsapiError DeleteRestoreJob(ctx, tenant, restorejob)
Delete a restorejob for a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **restorejob** | **string**| restorejob name | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRestoreJob**
> V1alpha1RestoreJob GetRestoreJob(ctx, tenant, restorejob)
Retrieve restorejob details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **restorejob** | **string**| restorejob name | 

### Return type

[**V1alpha1RestoreJob**](v1alpha1.RestoreJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRestoreJobs**
> V1alpha1RestoreJobList ListRestoreJobs(ctx, tenant)
List all restorejobs of a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 

### Return type

[**V1alpha1RestoreJobList**](v1alpha1.RestoreJobList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRestoreJob**
> V1alpha1RestoreJob UpdateRestoreJob(ctx, tenant, restorejob, body)
Update a restorejob for a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **restorejob** | **string**| restorejob name | 
  **body** | [**V1alpha1RestoreJob**](V1alpha1RestoreJob.md)|  | 

### Return type

[**V1alpha1RestoreJob**](v1alpha1.RestoreJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

