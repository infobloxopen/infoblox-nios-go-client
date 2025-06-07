# Filteroption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ApplyAsClass** | Pointer to **bool** | Determines if apply as class is enabled or not. If this flag is set to \&quot;true\&quot; the filter is treated as global DHCP class, e.g it is written to dhcpd config file even if it is not present in any DHCP range. | [optional] 
**Bootfile** | Pointer to **string** | A name of boot file of a DHCP filter option object. | [optional] 
**Bootserver** | Pointer to **string** | Determines the boot server of a DHCP filter option object. You can specify the name and/or IP address of the boot server that host needs to boot. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP filter option object. | [optional] 
**Expression** | Pointer to **string** | The conditional expression of a DHCP filter option object. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | Determines the lease time of a DHCP filter option object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP option filter object. | [optional] 
**NextServer** | Pointer to **string** | Determines the next server of a DHCP filter option object. You can specify the name and/or IP address of the next server that the host needs to boot. | [optional] 
**OptionList** | Pointer to [**[]FilteroptionOptionList**](FilteroptionOptionList.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**OptionSpace** | Pointer to **string** | The option space of a DHCP filter option object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | Determines the PXE (Preboot Execution Environment) lease time of a DHCP filter option object. To specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. | [optional] 

## Methods

### NewFilteroption

`func NewFilteroption() *Filteroption`

NewFilteroption instantiates a new Filteroption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilteroptionWithDefaults

`func NewFilteroptionWithDefaults() *Filteroption`

NewFilteroptionWithDefaults instantiates a new Filteroption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Filteroption) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Filteroption) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Filteroption) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Filteroption) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetApplyAsClass

`func (o *Filteroption) GetApplyAsClass() bool`

GetApplyAsClass returns the ApplyAsClass field if non-nil, zero value otherwise.

### GetApplyAsClassOk

`func (o *Filteroption) GetApplyAsClassOk() (*bool, bool)`

GetApplyAsClassOk returns a tuple with the ApplyAsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAsClass

`func (o *Filteroption) SetApplyAsClass(v bool)`

SetApplyAsClass sets ApplyAsClass field to given value.

### HasApplyAsClass

`func (o *Filteroption) HasApplyAsClass() bool`

HasApplyAsClass returns a boolean if a field has been set.

### GetBootfile

`func (o *Filteroption) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *Filteroption) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *Filteroption) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *Filteroption) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *Filteroption) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *Filteroption) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *Filteroption) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *Filteroption) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetComment

`func (o *Filteroption) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Filteroption) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Filteroption) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Filteroption) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpression

`func (o *Filteroption) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Filteroption) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Filteroption) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Filteroption) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExtattrs

`func (o *Filteroption) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Filteroption) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Filteroption) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Filteroption) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLeaseTime

`func (o *Filteroption) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *Filteroption) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *Filteroption) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *Filteroption) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *Filteroption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filteroption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filteroption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Filteroption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextServer

`func (o *Filteroption) GetNextServer() string`

GetNextServer returns the NextServer field if non-nil, zero value otherwise.

### GetNextServerOk

`func (o *Filteroption) GetNextServerOk() (*string, bool)`

GetNextServerOk returns a tuple with the NextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextServer

`func (o *Filteroption) SetNextServer(v string)`

SetNextServer sets NextServer field to given value.

### HasNextServer

`func (o *Filteroption) HasNextServer() bool`

HasNextServer returns a boolean if a field has been set.

### GetOptionList

`func (o *Filteroption) GetOptionList() []FilteroptionOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *Filteroption) GetOptionListOk() (*[]FilteroptionOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *Filteroption) SetOptionList(v []FilteroptionOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *Filteroption) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetOptionSpace

`func (o *Filteroption) GetOptionSpace() string`

GetOptionSpace returns the OptionSpace field if non-nil, zero value otherwise.

### GetOptionSpaceOk

`func (o *Filteroption) GetOptionSpaceOk() (*string, bool)`

GetOptionSpaceOk returns a tuple with the OptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSpace

`func (o *Filteroption) SetOptionSpace(v string)`

SetOptionSpace sets OptionSpace field to given value.

### HasOptionSpace

`func (o *Filteroption) HasOptionSpace() bool`

HasOptionSpace returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *Filteroption) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *Filteroption) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *Filteroption) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *Filteroption) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


