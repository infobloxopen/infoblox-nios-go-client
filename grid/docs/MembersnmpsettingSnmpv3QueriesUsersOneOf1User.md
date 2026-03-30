# MembersnmpsettingSnmpv3QueriesUsersOneOf1User

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

### NewMembersnmpsettingSnmpv3QueriesUsersOneOf1User

`func NewMembersnmpsettingSnmpv3QueriesUsersOneOf1User() *MembersnmpsettingSnmpv3QueriesUsersOneOf1User`

NewMembersnmpsettingSnmpv3QueriesUsersOneOf1User instantiates a new MembersnmpsettingSnmpv3QueriesUsersOneOf1User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembersnmpsettingSnmpv3QueriesUsersOneOf1UserWithDefaults

`func NewMembersnmpsettingSnmpv3QueriesUsersOneOf1UserWithDefaults() *MembersnmpsettingSnmpv3QueriesUsersOneOf1User`

NewMembersnmpsettingSnmpv3QueriesUsersOneOf1UserWithDefaults instantiates a new MembersnmpsettingSnmpv3QueriesUsersOneOf1User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetComment

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrs

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *MembersnmpsettingSnmpv3QueriesUsersOneOf1User) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


