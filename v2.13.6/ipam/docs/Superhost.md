# Superhost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for Super Host. | [optional] 
**DeleteAssociatedObjects** | Pointer to **bool** | True if we have to delete all DNS/DHCP associated objects with Super Host, false by default. | [optional] 
**DhcpAssociatedObjects** | Pointer to **[]string** | A list of DHCP objects refs which are associated with Super Host. | [optional] 
**Disabled** | Pointer to **bool** | Disable all DNS/DHCP associated objects with Super Host if True, False by default. | [optional] 
**DnsAssociatedObjects** | Pointer to **[]string** | A list of object refs of the DNS resource records which are associated with Super Host. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Name of the Superhost. | [optional] 

## Methods

### NewSuperhost

`func NewSuperhost() *Superhost`

NewSuperhost instantiates a new Superhost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperhostWithDefaults

`func NewSuperhostWithDefaults() *Superhost`

NewSuperhostWithDefaults instantiates a new Superhost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Superhost) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Superhost) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Superhost) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Superhost) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Superhost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Superhost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Superhost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Superhost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeleteAssociatedObjects

`func (o *Superhost) GetDeleteAssociatedObjects() bool`

GetDeleteAssociatedObjects returns the DeleteAssociatedObjects field if non-nil, zero value otherwise.

### GetDeleteAssociatedObjectsOk

`func (o *Superhost) GetDeleteAssociatedObjectsOk() (*bool, bool)`

GetDeleteAssociatedObjectsOk returns a tuple with the DeleteAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAssociatedObjects

`func (o *Superhost) SetDeleteAssociatedObjects(v bool)`

SetDeleteAssociatedObjects sets DeleteAssociatedObjects field to given value.

### HasDeleteAssociatedObjects

`func (o *Superhost) HasDeleteAssociatedObjects() bool`

HasDeleteAssociatedObjects returns a boolean if a field has been set.

### GetDhcpAssociatedObjects

`func (o *Superhost) GetDhcpAssociatedObjects() []string`

GetDhcpAssociatedObjects returns the DhcpAssociatedObjects field if non-nil, zero value otherwise.

### GetDhcpAssociatedObjectsOk

`func (o *Superhost) GetDhcpAssociatedObjectsOk() (*[]string, bool)`

GetDhcpAssociatedObjectsOk returns a tuple with the DhcpAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpAssociatedObjects

`func (o *Superhost) SetDhcpAssociatedObjects(v []string)`

SetDhcpAssociatedObjects sets DhcpAssociatedObjects field to given value.

### HasDhcpAssociatedObjects

`func (o *Superhost) HasDhcpAssociatedObjects() bool`

HasDhcpAssociatedObjects returns a boolean if a field has been set.

### GetDisabled

`func (o *Superhost) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Superhost) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Superhost) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Superhost) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnsAssociatedObjects

`func (o *Superhost) GetDnsAssociatedObjects() []string`

GetDnsAssociatedObjects returns the DnsAssociatedObjects field if non-nil, zero value otherwise.

### GetDnsAssociatedObjectsOk

`func (o *Superhost) GetDnsAssociatedObjectsOk() (*[]string, bool)`

GetDnsAssociatedObjectsOk returns a tuple with the DnsAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAssociatedObjects

`func (o *Superhost) SetDnsAssociatedObjects(v []string)`

SetDnsAssociatedObjects sets DnsAssociatedObjects field to given value.

### HasDnsAssociatedObjects

`func (o *Superhost) HasDnsAssociatedObjects() bool`

HasDnsAssociatedObjects returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Superhost) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Superhost) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Superhost) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Superhost) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Superhost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Superhost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Superhost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Superhost) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


