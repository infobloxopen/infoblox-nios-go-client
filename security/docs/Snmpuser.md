# Snmpuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AuthenticationPassword** | Pointer to **string** | Determines an authentication password for the user. This is a write-only attribute. | [optional] 
**AuthenticationProtocol** | Pointer to **string** | The authentication protocol to be used for this user. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment for the SNMPv3 User. | [optional] 
**Disable** | Pointer to **bool** | Determines if SNMPv3 user is disabled or not. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**PrivacyPassword** | Pointer to **string** | Determines a password for the privacy protocol. | [optional] 
**PrivacyProtocol** | Pointer to **string** | The privacy protocol to be used for this user. | [optional] 

## Methods

### NewSnmpuser

`func NewSnmpuser() *Snmpuser`

NewSnmpuser instantiates a new Snmpuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpuserWithDefaults

`func NewSnmpuserWithDefaults() *Snmpuser`

NewSnmpuserWithDefaults instantiates a new Snmpuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Snmpuser) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Snmpuser) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Snmpuser) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Snmpuser) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthenticationPassword

`func (o *Snmpuser) GetAuthenticationPassword() string`

GetAuthenticationPassword returns the AuthenticationPassword field if non-nil, zero value otherwise.

### GetAuthenticationPasswordOk

`func (o *Snmpuser) GetAuthenticationPasswordOk() (*string, bool)`

GetAuthenticationPasswordOk returns a tuple with the AuthenticationPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPassword

`func (o *Snmpuser) SetAuthenticationPassword(v string)`

SetAuthenticationPassword sets AuthenticationPassword field to given value.

### HasAuthenticationPassword

`func (o *Snmpuser) HasAuthenticationPassword() bool`

HasAuthenticationPassword returns a boolean if a field has been set.

### GetAuthenticationProtocol

`func (o *Snmpuser) GetAuthenticationProtocol() string`

GetAuthenticationProtocol returns the AuthenticationProtocol field if non-nil, zero value otherwise.

### GetAuthenticationProtocolOk

`func (o *Snmpuser) GetAuthenticationProtocolOk() (*string, bool)`

GetAuthenticationProtocolOk returns a tuple with the AuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProtocol

`func (o *Snmpuser) SetAuthenticationProtocol(v string)`

SetAuthenticationProtocol sets AuthenticationProtocol field to given value.

### HasAuthenticationProtocol

`func (o *Snmpuser) HasAuthenticationProtocol() bool`

HasAuthenticationProtocol returns a boolean if a field has been set.

### GetComment

`func (o *Snmpuser) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Snmpuser) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Snmpuser) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Snmpuser) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *Snmpuser) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Snmpuser) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Snmpuser) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Snmpuser) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Snmpuser) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Snmpuser) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Snmpuser) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Snmpuser) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Snmpuser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snmpuser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snmpuser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snmpuser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *Snmpuser) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *Snmpuser) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *Snmpuser) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *Snmpuser) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetPrivacyProtocol

`func (o *Snmpuser) GetPrivacyProtocol() string`

GetPrivacyProtocol returns the PrivacyProtocol field if non-nil, zero value otherwise.

### GetPrivacyProtocolOk

`func (o *Snmpuser) GetPrivacyProtocolOk() (*string, bool)`

GetPrivacyProtocolOk returns a tuple with the PrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyProtocol

`func (o *Snmpuser) SetPrivacyProtocol(v string)`

SetPrivacyProtocol sets PrivacyProtocol field to given value.

### HasPrivacyProtocol

`func (o *Snmpuser) HasPrivacyProtocol() bool`

HasPrivacyProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


