# \ClusterApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClusterApi.md#CreateCluster) | **Post** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/clusters | Create a cluster.
[**DeleteCluster**](ClusterApi.md#DeleteCluster) | **Delete** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/clusters/{cluster} | Delete a cluster.
[**GetCluster**](ClusterApi.md#GetCluster) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/clusters/{cluster} | Retrieve storage details.
[**GetNamespaces**](ClusterApi.md#GetNamespaces) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/clusters/{cluster}/resources/namespaces | Get namespaces in a cluster
[**ListClusters**](ClusterApi.md#ListClusters) | **Get** /jibuapis/ys.jibudata.com/v1alpha1/tenants/{tenant}/clusters | List all clusters of a tenant.


# **CreateCluster**
> V1alpha1Cluster CreateCluster(ctx, tenant, body)
Create a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **body** | [**V1alpha1Cluster**](V1alpha1Cluster.md)|  | 

### Return type

[**V1alpha1Cluster**](v1alpha1.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCluster**
> YsapiError DeleteCluster(ctx, tenant, cluster)
Delete a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **cluster** | **string**| cluster name | 

### Return type

[**YsapiError**](ysapi.Error.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCluster**
> V1alpha1Cluster GetCluster(ctx, tenant, cluster)
Retrieve storage details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **cluster** | **string**| cluster name | 

### Return type

[**V1alpha1Cluster**](v1alpha1.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespaces**
> V1NamespaceList GetNamespaces(ctx, tenant, cluster)
Get namespaces in a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 
  **cluster** | **string**| cluster name | 

### Return type

[**V1NamespaceList**](v1.NamespaceList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusters**
> V1alpha1ClusterList ListClusters(ctx, tenant)
List all clusters of a tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenant** | **string**| tenant id | 

### Return type

[**V1alpha1ClusterList**](v1alpha1.ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

