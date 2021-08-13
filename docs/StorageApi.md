# \StorageApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorage**](StorageApi.md#CreateStorage) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages | Create a storage.
[**CreateStorageCredentials**](StorageApi.md#CreateStorageCredentials) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage}/credentials | Create a credential for a storage.
[**DeleteStorage**](StorageApi.md#DeleteStorage) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage} | Delete a storage.
[**GetStorage**](StorageApi.md#GetStorage) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage} | Retrieve storage details.
[**ListStorages**](StorageApi.md#ListStorages) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages | List all storages of a tenant.


# **CreateStorage**
> V1alpha1Storage CreateStorage(ctx, tenant, body)
Create a storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1Tenant**](V1alpha1Tenant.md)|  | 

### Return type

[**V1alpha1Storage**](v1alpha1.Storage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStorageCredentials**
> YsapiError CreateStorageCredentials(ctx, tenant, storage, body)
Create a credential for a storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **storage** | **string**| storage name | 
  **body** | [**V1Secret**](V1Secret.md)|  | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorage**
> YsapiError DeleteStorage(ctx, tenant, storage)
Delete a storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **storage** | **string**| storage name | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorage**
> V1alpha1Storage GetStorage(ctx, tenant, storage)
Retrieve storage details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **storage** | **string**| storage name | 

### Return type

[**V1alpha1Storage**](v1alpha1.Storage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStorages**
> V1alpha1StorageList ListStorages(ctx, tenant)
List all storages of a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 

### Return type

[**V1alpha1StorageList**](v1alpha1.StorageList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

