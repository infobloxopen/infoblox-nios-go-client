# GetSamlAuthserviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The descriptive comment for the SAML authentication service. | [optional] 
**Idp** | Pointer to [**SamlAuthserviceIdp**](SamlAuthserviceIdp.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the SAML authentication service. | [optional] 
**SessionTimeout** | Pointer to **int64** | The session timeout in seconds. | [optional] 
**Result** | Pointer to [**SamlAuthservice**](SamlAuthservice.md) |  | [optional] 

## Methods

### NewGetSamlAuthserviceResponse

`func NewGetSamlAuthserviceResponse() *GetSamlAuthserviceResponse`

NewGetSamlAuthserviceResponse instantiates a new GetSamlAuthserviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSamlAuthserviceResponseWithDefaults

`func NewGetSamlAuthserviceResponseWithDefaults() *GetSamlAuthserviceResponse`

NewGetSamlAuthserviceResponseWithDefaults instantiates a new GetSamlAuthserviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSamlAuthserviceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSamlAuthserviceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSamlAuthserviceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSamlAuthserviceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSamlAuthserviceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSamlAuthserviceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSamlAuthserviceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSamlAuthserviceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIdp

`func (o *GetSamlAuthserviceResponse) GetIdp() SamlAuthserviceIdp`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *GetSamlAuthserviceResponse) GetIdpOk() (*SamlAuthserviceIdp, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *GetSamlAuthserviceResponse) SetIdp(v SamlAuthserviceIdp)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *GetSamlAuthserviceResponse) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetName

`func (o *GetSamlAuthserviceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSamlAuthserviceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSamlAuthserviceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSamlAuthserviceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *GetSamlAuthserviceResponse) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *GetSamlAuthserviceResponse) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *GetSamlAuthserviceResponse) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *GetSamlAuthserviceResponse) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetResult

`func (o *GetSamlAuthserviceResponse) GetResult() SamlAuthservice`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSamlAuthserviceResponse) GetResultOk() (*SamlAuthservice, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSamlAuthserviceResponse) SetResult(v SamlAuthservice)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSamlAuthserviceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


