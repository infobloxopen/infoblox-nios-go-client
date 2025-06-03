# RirOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Id** | Pointer to **string** | The RIR organization identifier. | [optional] 
**Maintainer** | Pointer to **string** | The RIR organization maintainer. | [optional] 
**Name** | Pointer to **string** | The RIR organization name. | [optional] 
**Password** | Pointer to **string** | The password for the maintainer of RIR organization. | [optional] 
**Rir** | Pointer to **string** | The RIR associated with RIR organization. | [optional] 
**SenderEmail** | Pointer to **string** | The sender e-mail address for RIR organization. | [optional] 

## Methods

### NewRirOrganization

`func NewRirOrganization() *RirOrganization`

NewRirOrganization instantiates a new RirOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRirOrganizationWithDefaults

`func NewRirOrganizationWithDefaults() *RirOrganization`

NewRirOrganizationWithDefaults instantiates a new RirOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RirOrganization) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RirOrganization) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RirOrganization) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RirOrganization) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetExtattrs

`func (o *RirOrganization) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *RirOrganization) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *RirOrganization) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *RirOrganization) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetId

`func (o *RirOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RirOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RirOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RirOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaintainer

`func (o *RirOrganization) GetMaintainer() string`

GetMaintainer returns the Maintainer field if non-nil, zero value otherwise.

### GetMaintainerOk

`func (o *RirOrganization) GetMaintainerOk() (*string, bool)`

GetMaintainerOk returns a tuple with the Maintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainer

`func (o *RirOrganization) SetMaintainer(v string)`

SetMaintainer sets Maintainer field to given value.

### HasMaintainer

`func (o *RirOrganization) HasMaintainer() bool`

HasMaintainer returns a boolean if a field has been set.

### GetName

`func (o *RirOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RirOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RirOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RirOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *RirOrganization) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RirOrganization) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RirOrganization) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RirOrganization) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRir

`func (o *RirOrganization) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *RirOrganization) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *RirOrganization) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *RirOrganization) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetSenderEmail

`func (o *RirOrganization) GetSenderEmail() string`

GetSenderEmail returns the SenderEmail field if non-nil, zero value otherwise.

### GetSenderEmailOk

`func (o *RirOrganization) GetSenderEmailOk() (*string, bool)`

GetSenderEmailOk returns a tuple with the SenderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmail

`func (o *RirOrganization) SetSenderEmail(v string)`

SetSenderEmail sets SenderEmail field to given value.

### HasSenderEmail

`func (o *RirOrganization) HasSenderEmail() bool`

HasSenderEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


