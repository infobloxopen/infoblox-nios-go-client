# GetDiscoveryCredentialgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Credential group. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryCredentialgroup**](DiscoveryCredentialgroup.md) |  | [optional] 

## Methods

### NewGetDiscoveryCredentialgroupResponse

`func NewGetDiscoveryCredentialgroupResponse() *GetDiscoveryCredentialgroupResponse`

NewGetDiscoveryCredentialgroupResponse instantiates a new GetDiscoveryCredentialgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryCredentialgroupResponseWithDefaults

`func NewGetDiscoveryCredentialgroupResponseWithDefaults() *GetDiscoveryCredentialgroupResponse`

NewGetDiscoveryCredentialgroupResponseWithDefaults instantiates a new GetDiscoveryCredentialgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryCredentialgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryCredentialgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryCredentialgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryCredentialgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetName

`func (o *GetDiscoveryCredentialgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoveryCredentialgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoveryCredentialgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoveryCredentialgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *GetDiscoveryCredentialgroupResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDiscoveryCredentialgroupResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDiscoveryCredentialgroupResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDiscoveryCredentialgroupResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryCredentialgroupResponse) GetResult() DiscoveryCredentialgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryCredentialgroupResponse) GetResultOk() (*DiscoveryCredentialgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryCredentialgroupResponse) SetResult(v DiscoveryCredentialgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryCredentialgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


