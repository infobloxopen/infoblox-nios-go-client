# GetHsmAllgroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Groups** | Pointer to **[]string** | The list of HSM groups configured on the appliance. | [optional] 
**Result** | Pointer to [**HsmAllgroups**](HsmAllgroups.md) |  | [optional] 

## Methods

### NewGetHsmAllgroupsResponse

`func NewGetHsmAllgroupsResponse() *GetHsmAllgroupsResponse`

NewGetHsmAllgroupsResponse instantiates a new GetHsmAllgroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHsmAllgroupsResponseWithDefaults

`func NewGetHsmAllgroupsResponseWithDefaults() *GetHsmAllgroupsResponse`

NewGetHsmAllgroupsResponseWithDefaults instantiates a new GetHsmAllgroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetHsmAllgroupsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetHsmAllgroupsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetHsmAllgroupsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetHsmAllgroupsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetGroups

`func (o *GetHsmAllgroupsResponse) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetHsmAllgroupsResponse) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetHsmAllgroupsResponse) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetHsmAllgroupsResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetResult

`func (o *GetHsmAllgroupsResponse) GetResult() HsmAllgroups`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetHsmAllgroupsResponse) GetResultOk() (*HsmAllgroups, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetHsmAllgroupsResponse) SetResult(v HsmAllgroups)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetHsmAllgroupsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


