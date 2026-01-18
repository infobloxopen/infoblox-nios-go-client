# GetThreatinsightAllowlistResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The descriptive comment for the threat insight allowlist. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the threat insight allowlist is disabled. | [optional] 
**Fqdn** | Pointer to **string** | The FQDN of the threat insight allowlist. | [optional] 
**Type** | Pointer to **string** | The type of the threat insight allowlist. | [optional] [readonly] 
**Result** | Pointer to [**ThreatinsightAllowlist**](ThreatinsightAllowlist.md) |  | [optional] 

## Methods

### NewGetThreatinsightAllowlistResponse

`func NewGetThreatinsightAllowlistResponse() *GetThreatinsightAllowlistResponse`

NewGetThreatinsightAllowlistResponse instantiates a new GetThreatinsightAllowlistResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatinsightAllowlistResponseWithDefaults

`func NewGetThreatinsightAllowlistResponseWithDefaults() *GetThreatinsightAllowlistResponse`

NewGetThreatinsightAllowlistResponseWithDefaults instantiates a new GetThreatinsightAllowlistResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatinsightAllowlistResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatinsightAllowlistResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatinsightAllowlistResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatinsightAllowlistResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetThreatinsightAllowlistResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetThreatinsightAllowlistResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetThreatinsightAllowlistResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetThreatinsightAllowlistResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetThreatinsightAllowlistResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetThreatinsightAllowlistResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetThreatinsightAllowlistResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetThreatinsightAllowlistResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetFqdn

`func (o *GetThreatinsightAllowlistResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GetThreatinsightAllowlistResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GetThreatinsightAllowlistResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GetThreatinsightAllowlistResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetType

`func (o *GetThreatinsightAllowlistResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetThreatinsightAllowlistResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetThreatinsightAllowlistResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetThreatinsightAllowlistResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatinsightAllowlistResponse) GetResult() ThreatinsightAllowlist`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatinsightAllowlistResponse) GetResultOk() (*ThreatinsightAllowlist, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatinsightAllowlistResponse) SetResult(v ThreatinsightAllowlist)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatinsightAllowlistResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


