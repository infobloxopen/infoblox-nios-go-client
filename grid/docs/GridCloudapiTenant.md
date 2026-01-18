# GridCloudapiTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CloudInfo** | Pointer to [**GridCloudapiTenantCloudInfo**](GridCloudapiTenantCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the Grid Cloud API Tenant object; maximum 256 characters. | [optional] 
**CreatedTs** | Pointer to **int64** | The timestamp when the tenant was first created in the system. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID associated with the tenant. This is set only when the tenant is first created. | [optional] [readonly] 
**LastEventTs** | Pointer to **int64** | The timestamp when the last event associated with the tenant happened. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the tenant. | [optional] 
**NetworkCount** | Pointer to **int64** | Number of Networks associated with the tenant. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | Number of VMs associated with the tenant. | [optional] [readonly] 

## Methods

### NewGridCloudapiTenant

`func NewGridCloudapiTenant() *GridCloudapiTenant`

NewGridCloudapiTenant instantiates a new GridCloudapiTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCloudapiTenantWithDefaults

`func NewGridCloudapiTenantWithDefaults() *GridCloudapiTenant`

NewGridCloudapiTenantWithDefaults instantiates a new GridCloudapiTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridCloudapiTenant) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridCloudapiTenant) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridCloudapiTenant) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridCloudapiTenant) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GridCloudapiTenant) GetCloudInfo() GridCloudapiTenantCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GridCloudapiTenant) GetCloudInfoOk() (*GridCloudapiTenantCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GridCloudapiTenant) SetCloudInfo(v GridCloudapiTenantCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GridCloudapiTenant) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GridCloudapiTenant) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GridCloudapiTenant) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GridCloudapiTenant) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GridCloudapiTenant) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedTs

`func (o *GridCloudapiTenant) GetCreatedTs() int64`

GetCreatedTs returns the CreatedTs field if non-nil, zero value otherwise.

### GetCreatedTsOk

`func (o *GridCloudapiTenant) GetCreatedTsOk() (*int64, bool)`

GetCreatedTsOk returns a tuple with the CreatedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTs

`func (o *GridCloudapiTenant) SetCreatedTs(v int64)`

SetCreatedTs sets CreatedTs field to given value.

### HasCreatedTs

`func (o *GridCloudapiTenant) HasCreatedTs() bool`

HasCreatedTs returns a boolean if a field has been set.

### GetId

`func (o *GridCloudapiTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GridCloudapiTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GridCloudapiTenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GridCloudapiTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEventTs

`func (o *GridCloudapiTenant) GetLastEventTs() int64`

GetLastEventTs returns the LastEventTs field if non-nil, zero value otherwise.

### GetLastEventTsOk

`func (o *GridCloudapiTenant) GetLastEventTsOk() (*int64, bool)`

GetLastEventTsOk returns a tuple with the LastEventTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTs

`func (o *GridCloudapiTenant) SetLastEventTs(v int64)`

SetLastEventTs sets LastEventTs field to given value.

### HasLastEventTs

`func (o *GridCloudapiTenant) HasLastEventTs() bool`

HasLastEventTs returns a boolean if a field has been set.

### GetName

`func (o *GridCloudapiTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridCloudapiTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridCloudapiTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridCloudapiTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCount

`func (o *GridCloudapiTenant) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *GridCloudapiTenant) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *GridCloudapiTenant) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *GridCloudapiTenant) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetVmCount

`func (o *GridCloudapiTenant) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *GridCloudapiTenant) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *GridCloudapiTenant) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *GridCloudapiTenant) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


