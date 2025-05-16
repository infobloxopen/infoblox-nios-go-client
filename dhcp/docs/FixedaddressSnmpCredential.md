# FixedaddressSnmpCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunityString** | Pointer to **string** | The public community string. | [optional] 
**Comment** | Pointer to **string** | Comments for the SNMPv1 and SNMPv2 users. | [optional] 
**CredentialGroup** | Pointer to **string** | Group for the SNMPv1 and SNMPv2 credential. | [optional] 

## Methods

### NewFixedaddressSnmpCredential

`func NewFixedaddressSnmpCredential() *FixedaddressSnmpCredential`

NewFixedaddressSnmpCredential instantiates a new FixedaddressSnmpCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedaddressSnmpCredentialWithDefaults

`func NewFixedaddressSnmpCredentialWithDefaults() *FixedaddressSnmpCredential`

NewFixedaddressSnmpCredentialWithDefaults instantiates a new FixedaddressSnmpCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunityString

`func (o *FixedaddressSnmpCredential) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *FixedaddressSnmpCredential) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *FixedaddressSnmpCredential) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *FixedaddressSnmpCredential) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetComment

`func (o *FixedaddressSnmpCredential) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FixedaddressSnmpCredential) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FixedaddressSnmpCredential) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FixedaddressSnmpCredential) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialGroup

`func (o *FixedaddressSnmpCredential) GetCredentialGroup() string`

GetCredentialGroup returns the CredentialGroup field if non-nil, zero value otherwise.

### GetCredentialGroupOk

`func (o *FixedaddressSnmpCredential) GetCredentialGroupOk() (*string, bool)`

GetCredentialGroupOk returns a tuple with the CredentialGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialGroup

`func (o *FixedaddressSnmpCredential) SetCredentialGroup(v string)`

SetCredentialGroup sets CredentialGroup field to given value.

### HasCredentialGroup

`func (o *FixedaddressSnmpCredential) HasCredentialGroup() bool`

HasCredentialGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


