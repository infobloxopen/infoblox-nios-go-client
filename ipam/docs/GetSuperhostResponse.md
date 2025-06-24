# GetSuperhostResponse

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
**Result** | Pointer to [**Superhost**](Superhost.md) |  | [optional] 

## Methods

### NewGetSuperhostResponse

`func NewGetSuperhostResponse() *GetSuperhostResponse`

NewGetSuperhostResponse instantiates a new GetSuperhostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSuperhostResponseWithDefaults

`func NewGetSuperhostResponseWithDefaults() *GetSuperhostResponse`

NewGetSuperhostResponseWithDefaults instantiates a new GetSuperhostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSuperhostResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSuperhostResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSuperhostResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSuperhostResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSuperhostResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSuperhostResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSuperhostResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSuperhostResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeleteAssociatedObjects

`func (o *GetSuperhostResponse) GetDeleteAssociatedObjects() bool`

GetDeleteAssociatedObjects returns the DeleteAssociatedObjects field if non-nil, zero value otherwise.

### GetDeleteAssociatedObjectsOk

`func (o *GetSuperhostResponse) GetDeleteAssociatedObjectsOk() (*bool, bool)`

GetDeleteAssociatedObjectsOk returns a tuple with the DeleteAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAssociatedObjects

`func (o *GetSuperhostResponse) SetDeleteAssociatedObjects(v bool)`

SetDeleteAssociatedObjects sets DeleteAssociatedObjects field to given value.

### HasDeleteAssociatedObjects

`func (o *GetSuperhostResponse) HasDeleteAssociatedObjects() bool`

HasDeleteAssociatedObjects returns a boolean if a field has been set.

### GetDhcpAssociatedObjects

`func (o *GetSuperhostResponse) GetDhcpAssociatedObjects() []string`

GetDhcpAssociatedObjects returns the DhcpAssociatedObjects field if non-nil, zero value otherwise.

### GetDhcpAssociatedObjectsOk

`func (o *GetSuperhostResponse) GetDhcpAssociatedObjectsOk() (*[]string, bool)`

GetDhcpAssociatedObjectsOk returns a tuple with the DhcpAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpAssociatedObjects

`func (o *GetSuperhostResponse) SetDhcpAssociatedObjects(v []string)`

SetDhcpAssociatedObjects sets DhcpAssociatedObjects field to given value.

### HasDhcpAssociatedObjects

`func (o *GetSuperhostResponse) HasDhcpAssociatedObjects() bool`

HasDhcpAssociatedObjects returns a boolean if a field has been set.

### GetDisabled

`func (o *GetSuperhostResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetSuperhostResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetSuperhostResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetSuperhostResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnsAssociatedObjects

`func (o *GetSuperhostResponse) GetDnsAssociatedObjects() []string`

GetDnsAssociatedObjects returns the DnsAssociatedObjects field if non-nil, zero value otherwise.

### GetDnsAssociatedObjectsOk

`func (o *GetSuperhostResponse) GetDnsAssociatedObjectsOk() (*[]string, bool)`

GetDnsAssociatedObjectsOk returns a tuple with the DnsAssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAssociatedObjects

`func (o *GetSuperhostResponse) SetDnsAssociatedObjects(v []string)`

SetDnsAssociatedObjects sets DnsAssociatedObjects field to given value.

### HasDnsAssociatedObjects

`func (o *GetSuperhostResponse) HasDnsAssociatedObjects() bool`

HasDnsAssociatedObjects returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetSuperhostResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetSuperhostResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetSuperhostResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetSuperhostResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GetSuperhostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSuperhostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSuperhostResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSuperhostResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetSuperhostResponse) GetResult() Superhost`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSuperhostResponse) GetResultOk() (*Superhost, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSuperhostResponse) SetResult(v Superhost)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSuperhostResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


