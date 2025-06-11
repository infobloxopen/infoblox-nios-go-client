# Allendpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The Grid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN). | [optional] [readonly] 
**Comment** | Pointer to **string** | The Grid endpoint descriptive comment. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether a Grid endpoint is disabled or not. When this is set to False, the Grid endpoint is enabled. | [optional] [readonly] 
**SubscribingMember** | Pointer to **string** | The name of the Grid Member object that is serving Grid endpoint. | [optional] [readonly] 
**Type** | Pointer to **string** | The Grid endpoint type. | [optional] [readonly] 
**Version** | Pointer to **string** | The Grid endpoint version. | [optional] [readonly] 

## Methods

### NewAllendpoints

`func NewAllendpoints() *Allendpoints`

NewAllendpoints instantiates a new Allendpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllendpointsWithDefaults

`func NewAllendpointsWithDefaults() *Allendpoints`

NewAllendpointsWithDefaults instantiates a new Allendpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Allendpoints) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Allendpoints) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Allendpoints) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Allendpoints) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *Allendpoints) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Allendpoints) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Allendpoints) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Allendpoints) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *Allendpoints) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Allendpoints) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Allendpoints) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Allendpoints) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *Allendpoints) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Allendpoints) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Allendpoints) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Allendpoints) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetSubscribingMember

`func (o *Allendpoints) GetSubscribingMember() string`

GetSubscribingMember returns the SubscribingMember field if non-nil, zero value otherwise.

### GetSubscribingMemberOk

`func (o *Allendpoints) GetSubscribingMemberOk() (*string, bool)`

GetSubscribingMemberOk returns a tuple with the SubscribingMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribingMember

`func (o *Allendpoints) SetSubscribingMember(v string)`

SetSubscribingMember sets SubscribingMember field to given value.

### HasSubscribingMember

`func (o *Allendpoints) HasSubscribingMember() bool`

HasSubscribingMember returns a boolean if a field has been set.

### GetType

`func (o *Allendpoints) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Allendpoints) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Allendpoints) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Allendpoints) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *Allendpoints) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Allendpoints) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Allendpoints) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Allendpoints) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


