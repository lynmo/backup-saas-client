# \BackupJobTagApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupJob**](BackupJobTagApi.md#CreateBackupJob) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupjobs | Create a backupjob for a tenant.
[**DeleteBackupJob**](BackupJobTagApi.md#DeleteBackupJob) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupjobs/{backupjob} | Delete a backupjob for a tenant.
[**GetBackupJob**](BackupJobTagApi.md#GetBackupJob) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupjobs/{backupjob} | Retrieve backupjob details.
[**ListBackupJobs**](BackupJobTagApi.md#ListBackupJobs) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupjobs | List all backupjobs of a tenant


# **CreateBackupJob**
> V1alpha1BackupJob CreateBackupJob(ctx, tenant, body)
Create a backupjob for a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1BackupJob**](V1alpha1BackupJob.md)|  | 

### Return type

[**V1alpha1BackupJob**](v1alpha1.BackupJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupJob**
> YsapiError DeleteBackupJob(ctx, tenant, backupjob)
Delete a backupjob for a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **backupjob** | **string**| backupjob name | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupJob**
> V1alpha1BackupJob GetBackupJob(ctx, tenant, backupjob)
Retrieve backupjob details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **backupjob** | **string**| backupjob name | 

### Return type

[**V1alpha1BackupJob**](v1alpha1.BackupJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupJobs**
> V1alpha1BackupJobList ListBackupJobs(ctx, tenant)
List all backupjobs of a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 

### Return type

[**V1alpha1BackupJobList**](v1alpha1.BackupJobList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

