# \BackupPlanTagApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackupPlan**](BackupPlanTagApi.md#CreateBackupPlan) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupplans | Create a backupplan.
[**DeleteBackupPlan**](BackupPlanTagApi.md#DeleteBackupPlan) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupplans/{backupplan} | Delete a backupplan.
[**GetBackupPlan**](BackupPlanTagApi.md#GetBackupPlan) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupplans/{backupplan} | Retrieve backupplan details.
[**ListAllBackupPlans**](BackupPlanTagApi.md#ListAllBackupPlans) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/backupplans | List all backupplans
[**ListBackupPlans**](BackupPlanTagApi.md#ListBackupPlans) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupplans | List all backupplans of a tenant
[**UpdateBackupPlan**](BackupPlanTagApi.md#UpdateBackupPlan) | **Put** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/backupplans/{backupplan} | Update backupplan details.


# **CreateBackupPlan**
> V1alpha1BackupPlan CreateBackupPlan(ctx, tenant, body)
Create a backupplan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1BackupPlan**](V1alpha1BackupPlan.md)|  | 

### Return type

[**V1alpha1BackupPlan**](v1alpha1.BackupPlan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackupPlan**
> YsapiError DeleteBackupPlan(ctx, tenant, backupplan)
Delete a backupplan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **backupplan** | **string**| backupplan name | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupPlan**
> V1alpha1BackupPlan GetBackupPlan(ctx, tenant, backupplan)
Retrieve backupplan details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **backupplan** | **string**| backupplan name | 

### Return type

[**V1alpha1BackupPlan**](v1alpha1.BackupPlan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllBackupPlans**
> V1alpha1BackupPlanList ListAllBackupPlans(ctx, optional)
List all backupplans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BackupPlanTagApiListAllBackupPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupPlanTagApiListAllBackupPlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1BackupPlanList**](v1alpha1.BackupPlanList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackupPlans**
> V1alpha1BackupPlanList ListBackupPlans(ctx, tenant, optional)
List all backupplans of a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
 **optional** | ***BackupPlanTagApiListBackupPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupPlanTagApiListBackupPlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| backup plan name | 
 **watch** | **optional.String**| watch | 
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1BackupPlanList**](v1alpha1.BackupPlanList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBackupPlan**
> V1alpha1BackupPlan UpdateBackupPlan(ctx, tenant, backupplan)
Update backupplan details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **backupplan** | **string**| backupplan name | 

### Return type

[**V1alpha1BackupPlan**](v1alpha1.BackupPlan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

