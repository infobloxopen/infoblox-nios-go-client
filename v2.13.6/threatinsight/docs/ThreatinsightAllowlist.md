# ThreatinsightAllowlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the threat insight allowlist. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the threat insight allowlist is disabled. | [optional] 
**Fqdn** | Pointer to **string** | The FQDN of the threat insight allowlist. | [optional] 
**Type** | Pointer to **string** | The type of the threat insight allowlist. | [optional] [readonly] 

## Methods

### NewThreatinsightAllowlist

`func NewThreatinsightAllowlist() *ThreatinsightAllowlist`

NewThreatinsightAllowlist instantiates a new ThreatinsightAllowlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatinsightAllowlistWithDefaults

`func NewThreatinsightAllowlistWithDefaults() *ThreatinsightAllowlist`

NewThreatinsightAllowlistWithDefaults instantiates a new ThreatinsightAllowlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatinsightAllowlist) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatinsightAllowlist) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatinsightAllowlist) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatinsightAllowlist) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *ThreatinsightAllowlist) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ThreatinsightAllowlist) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ThreatinsightAllowlist) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ThreatinsightAllowlist) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *ThreatinsightAllowlist) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ThreatinsightAllowlist) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ThreatinsightAllowlist) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ThreatinsightAllowlist) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetFqdn

`func (o *ThreatinsightAllowlist) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ThreatinsightAllowlist) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ThreatinsightAllowlist) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ThreatinsightAllowlist) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetType

`func (o *ThreatinsightAllowlist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreatinsightAllowlist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreatinsightAllowlist) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreatinsightAllowlist) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


