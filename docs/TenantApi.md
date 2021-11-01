# \TenantApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantApi.md#CreateTenant) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants | Create a tenant.
[**DeleteTenant**](TenantApi.md#DeleteTenant) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant} | Delete a tenant.
[**GetTenant**](TenantApi.md#GetTenant) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant} | Retrieve tenant details.
[**ListTenants**](TenantApi.md#ListTenants) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants | List all tenants
[**UpdateTenant**](TenantApi.md#UpdateTenant) | **Put** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant} | Update a tenant


# **CreateTenant**
> V1alpha1Tenant CreateTenant(ctx, body)
Create a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1Tenant**](V1alpha1Tenant.md)|  | 

### Return type

[**V1alpha1Tenant**](v1alpha1.Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTenant**
> YsapiError DeleteTenant(ctx, tenant)
Delete a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant ID | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenant**
> V1alpha1Tenant GetTenant(ctx, tenant)
Retrieve tenant details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant ID | 

### Return type

[**V1alpha1Tenant**](v1alpha1.Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTenants**
> V1alpha1TenantList ListTenants(ctx, optional)
List all tenants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantApiListTenantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantApiListTenantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1TenantList**](v1alpha1.TenantList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTenant**
> V1alpha1Tenant UpdateTenant(ctx, tenant, body)
Update a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant ID | 
  **body** | [**V1alpha1Tenant**](V1alpha1Tenant.md)|  | 

### Return type

[**V1alpha1Tenant**](v1alpha1.Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

