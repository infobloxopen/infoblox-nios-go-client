# GetThreatinsightModulesetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Version** | Pointer to **string** | The version number of the threat insight module set. | [optional] [readonly] 
**Result** | Pointer to [**ThreatinsightModuleset**](ThreatinsightModuleset.md) |  | [optional] 

## Methods

### NewGetThreatinsightModulesetResponse

`func NewGetThreatinsightModulesetResponse() *GetThreatinsightModulesetResponse`

NewGetThreatinsightModulesetResponse instantiates a new GetThreatinsightModulesetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatinsightModulesetResponseWithDefaults

`func NewGetThreatinsightModulesetResponseWithDefaults() *GetThreatinsightModulesetResponse`

NewGetThreatinsightModulesetResponseWithDefaults instantiates a new GetThreatinsightModulesetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatinsightModulesetResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatinsightModulesetResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatinsightModulesetResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatinsightModulesetResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetUuid

`func (o *GetThreatinsightModulesetResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetThreatinsightModulesetResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetThreatinsightModulesetResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetThreatinsightModulesetResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *GetThreatinsightModulesetResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetThreatinsightModulesetResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetThreatinsightModulesetResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetThreatinsightModulesetResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatinsightModulesetResponse) GetResult() ThreatinsightModuleset`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatinsightModulesetResponse) GetResultOk() (*ThreatinsightModuleset, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatinsightModulesetResponse) SetResult(v ThreatinsightModuleset)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatinsightModulesetResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


