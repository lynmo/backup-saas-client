# \StorageApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorage**](StorageApi.md#CreateStorage) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages | Create a storage.
[**CreateStorageCredentials**](StorageApi.md#CreateStorageCredentials) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage}/credentials | Create a credential for a storage.
[**DeleteStorage**](StorageApi.md#DeleteStorage) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage} | Delete a storage.
[**GetStorage**](StorageApi.md#GetStorage) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage} | Retrieve storage details.
[**ListAllStorages**](StorageApi.md#ListAllStorages) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/storages | List all storages.
[**ListStorages**](StorageApi.md#ListStorages) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages | List all storages of a tenant.
[**UpdateStorage**](StorageApi.md#UpdateStorage) | **Put** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/storages/{storage} | Update a storage


# **CreateStorage**
> V1alpha1Storage CreateStorage(ctx, tenant, body)
Create a storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1Storage**](V1alpha1Storage.md)|  | 

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
> V1alpha1Storage GetStorage(ctx, tenant, storage, optional)
Retrieve storage details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **storage** | **string**| storage name | 
 **optional** | ***StorageApiGetStorageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiGetStorageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSecrets** | **optional.String**| include secrets | 

### Return type

[**V1alpha1Storage**](v1alpha1.Storage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllStorages**
> V1alpha1StorageList ListAllStorages(ctx, optional)
List all storages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiListAllStoragesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiListAllStoragesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSecrets** | **optional.String**| include secrets | 
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1StorageList**](v1alpha1.StorageList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStorages**
> V1alpha1StorageList ListStorages(ctx, tenant, optional)
List all storages of a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
 **optional** | ***StorageApiListStoragesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiListStoragesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSecrets** | **optional.String**| include secrets | 
 **page** | **optional.String**| page | [default to page&#x3D;1]
 **limit** | **optional.String**| limit | 
 **ascending** | **optional.String**| sort parameters, e.g. reverse&#x3D;true | [default to ascending&#x3D;false]
 **sortBy** | **optional.String**| sort parameters, e.g. orderBy&#x3D;createTime | 

### Return type

[**V1alpha1StorageList**](v1alpha1.StorageList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStorage**
> V1alpha1Storage UpdateStorage(ctx, tenant, storage, body)
Update a storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **storage** | **string**| storage name | 
  **body** | [**V1alpha1Storage**](V1alpha1Storage.md)|  | 

### Return type

[**V1alpha1Storage**](v1alpha1.Storage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

