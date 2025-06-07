# GetGridCloudapiTenantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CloudInfo** | Pointer to [**GridCloudapiTenantCloudInfo**](GridCloudapiTenantCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the Grid Cloud API Tenant object; maximum 256 characters. | [optional] 
**CreatedTs** | Pointer to **int64** | The timestamp when the tenant was first created in the system. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID associated with the tenant. This is set only when the tenant is first created. | [optional] [readonly] 
**LastEventTs** | Pointer to **int64** | The timestamp when the last event associated with the tenant happened. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the tenant. | [optional] 
**NetworkCount** | Pointer to **int64** | Number of Networks associated with the tenant. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | Number of VMs associated with the tenant. | [optional] [readonly] 
**Result** | Pointer to [**GridCloudapiTenant**](GridCloudapiTenant.md) |  | [optional] 

## Methods

### NewGetGridCloudapiTenantResponse

`func NewGetGridCloudapiTenantResponse() *GetGridCloudapiTenantResponse`

NewGetGridCloudapiTenantResponse instantiates a new GetGridCloudapiTenantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridCloudapiTenantResponseWithDefaults

`func NewGetGridCloudapiTenantResponseWithDefaults() *GetGridCloudapiTenantResponse`

NewGetGridCloudapiTenantResponseWithDefaults instantiates a new GetGridCloudapiTenantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridCloudapiTenantResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridCloudapiTenantResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridCloudapiTenantResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridCloudapiTenantResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetGridCloudapiTenantResponse) GetCloudInfo() GridCloudapiTenantCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetGridCloudapiTenantResponse) GetCloudInfoOk() (*GridCloudapiTenantCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetGridCloudapiTenantResponse) SetCloudInfo(v GridCloudapiTenantCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetGridCloudapiTenantResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetGridCloudapiTenantResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetGridCloudapiTenantResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetGridCloudapiTenantResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetGridCloudapiTenantResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedTs

`func (o *GetGridCloudapiTenantResponse) GetCreatedTs() int64`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *GetGridCloudapiTenantResponse) GetCreatedTsOk() (*int64, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *GetGridCloudapiTenantResponse) SetCreatedTs(v int64)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *GetGridCloudapiTenantResponse) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetId

`func (o *GetGridCloudapiTenantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGridCloudapiTenantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGridCloudapiTenantResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetGridCloudapiTenantResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEventTs

`func (o *GetGridCloudapiTenantResponse) GetLastEventTs() int64`

GetLastEventTs returns the LastEventTs field if non-nil, zero value otherwise.

### GetLastEventTsOk

`func (o *GetGridCloudapiTenantResponse) GetLastEventTsOk() (*int64, bool)`

GetLastEventTsOk returns a tuple with the LastEventTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTs

`func (o *GetGridCloudapiTenantResponse) SetLastEventTs(v int64)`

SetLastEventTs sets LastEventTs field to given value.

### HasLastEventTs

`func (o *GetGridCloudapiTenantResponse) HasLastEventTs() bool`

HasLastEventTs returns a boolean if a field has been set.

### GetName

`func (o *GetGridCloudapiTenantResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGridCloudapiTenantResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGridCloudapiTenantResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGridCloudapiTenantResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCount

`func (o *GetGridCloudapiTenantResponse) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *GetGridCloudapiTenantResponse) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *GetGridCloudapiTenantResponse) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *GetGridCloudapiTenantResponse) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetVmCount

`func (o *GetGridCloudapiTenantResponse) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *GetGridCloudapiTenantResponse) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *GetGridCloudapiTenantResponse) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *GetGridCloudapiTenantResponse) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetResult

`func (o *GetGridCloudapiTenantResponse) GetResult() GridCloudapiTenant`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridCloudapiTenantResponse) GetResultOk() (*GridCloudapiTenant, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridCloudapiTenantResponse) SetResult(v GridCloudapiTenant)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridCloudapiTenantResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


