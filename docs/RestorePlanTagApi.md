# \RestorePlanTagApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRestorePlan**](RestorePlanTagApi.md#CreateRestorePlan) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restoreplans | Create a restoreplan.
[**DeleteRestorePlan**](RestorePlanTagApi.md#DeleteRestorePlan) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restoreplans/{restoreplan} | Delete a restoreplan.
[**GetRestorePlan**](RestorePlanTagApi.md#GetRestorePlan) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restoreplans/{restoreplan} | Retrieve restoreplan details.
[**ListAllRestorePlans**](RestorePlanTagApi.md#ListAllRestorePlans) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/restoreplans | List all restoreplans.
[**ListRestorePlans**](RestorePlanTagApi.md#ListRestorePlans) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restoreplans | List all restoreplans of a tenant
[**UpdateRestorePlan**](RestorePlanTagApi.md#UpdateRestorePlan) | **Put** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/restoreplans/{restoreplan} | Update restoreplan details.


# **CreateRestorePlan**
> V1alpha1RestorePlan CreateRestorePlan(ctx, tenant, body)
Create a restoreplan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1RestorePlan**](V1alpha1RestorePlan.md)|  | 

### Return type

[**V1alpha1RestorePlan**](v1alpha1.RestorePlan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRestorePlan**
> YsapiError DeleteRestorePlan(ctx, tenant, restoreplan)
Delete a restoreplan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **restoreplan** | **string**| restoreplan name | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRestorePlan**
> V1alpha1RestorePlan GetRestorePlan(ctx, tenant, restoreplan)
Retrieve restoreplan details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **restoreplan** | **string**| restoreplan name | 

### Return type

[**V1alpha1RestorePlan**](v1alpha1.RestorePlan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllRestorePlans**
> V1alpha1RestorePlanList ListAllRestorePlans(ctx, optional)
List all restoreplans.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RestorePlanTagApiListAllRestorePlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestorePlanTagApiListAllRestorePlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1RestorePlanList**](v1alpha1.RestorePlanList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRestorePlans**
> V1alpha1RestorePlanList ListRestorePlans(ctx, tenant, optional)
List all restoreplans of a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
 **optional** | ***RestorePlanTagApiListRestorePlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RestorePlanTagApiListRestorePlansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| restore plan name | 
 **watch** | **optional.String**| watch | 
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1RestorePlanList**](v1alpha1.RestorePlanList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRestorePlan**
> V1alpha1RestorePlan UpdateRestorePlan(ctx, tenant, restoreplan)
Update restoreplan details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **restoreplan** | **string**| restoreplan name | 

### Return type

[**V1alpha1RestorePlan**](v1alpha1.RestorePlan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

