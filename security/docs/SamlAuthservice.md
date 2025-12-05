# SamlAuthservice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The descriptive comment for the SAML authentication service. | [optional] 
**Idp** | Pointer to [**SamlAuthserviceIdp**](SamlAuthserviceIdp.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the SAML authentication service. | [optional] 
**SessionTimeout** | Pointer to **int64** | The session timeout in seconds. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewSamlAuthservice

`func NewSamlAuthservice() *SamlAuthservice`

NewSamlAuthservice instantiates a new SamlAuthservice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAuthserviceWithDefaults

`func NewSamlAuthserviceWithDefaults() *SamlAuthservice`

NewSamlAuthserviceWithDefaults instantiates a new SamlAuthservice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SamlAuthservice) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SamlAuthservice) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SamlAuthservice) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SamlAuthservice) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SamlAuthservice) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SamlAuthservice) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SamlAuthservice) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SamlAuthservice) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIdp

`func (o *SamlAuthservice) GetIdp() SamlAuthserviceIdp`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *SamlAuthservice) GetIdpOk() (*SamlAuthserviceIdp, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *SamlAuthservice) SetIdp(v SamlAuthserviceIdp)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *SamlAuthservice) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetName

`func (o *SamlAuthservice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlAuthservice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlAuthservice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlAuthservice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *SamlAuthservice) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SamlAuthservice) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SamlAuthservice) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SamlAuthservice) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetUuid

`func (o *SamlAuthservice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SamlAuthservice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SamlAuthservice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SamlAuthservice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


