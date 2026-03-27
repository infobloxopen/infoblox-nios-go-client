# GridsnmpsettingSnmpv3QueriesUsersOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the SNMPv3 user object | [optional] [readonly] 
**AuthenticationProtocol** | Pointer to **string** | The authentication protocol to be used for this user. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment for the SNMPv3 User. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines if SNMPv3 user is disabled or not. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the user. | [optional] [readonly] 
**PrivacyProtocol** | Pointer to **string** | The privacy protocol to be used for this user. | [optional] [readonly] 

## Methods

### NewGridsnmpsettingSnmpv3QueriesUsersOneOf

`func NewGridsnmpsettingSnmpv3QueriesUsersOneOf() *GridsnmpsettingSnmpv3QueriesUsersOneOf`

NewGridsnmpsettingSnmpv3QueriesUsersOneOf instantiates a new GridsnmpsettingSnmpv3QueriesUsersOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridsnmpsettingSnmpv3QueriesUsersOneOfWithDefaults

`func NewGridsnmpsettingSnmpv3QueriesUsersOneOfWithDefaults() *GridsnmpsettingSnmpv3QueriesUsersOneOf`

NewGridsnmpsettingSnmpv3QueriesUsersOneOfWithDefaults instantiates a new GridsnmpsettingSnmpv3QueriesUsersOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetComment

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *GridsnmpsettingSnmpv3QueriesUsersOneOf) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


