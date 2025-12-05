# GetFilteroptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ApplyAsClass** | Pointer to **bool** | Determines if apply as class is enabled or not. If this flag is set to \&quot;true\&quot; the filter is treated as global DHCP class, e.g it is written to dhcpd config file even if it is not present in any DHCP range. | [optional] 
**Bootfile** | Pointer to **string** | A name of boot file of a DHCP filter option object. | [optional] 
**Bootserver** | Pointer to **string** | Determines the boot server of a DHCP filter option object. You can specify the name and/or IP address of the boot server that host needs to boot. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP filter option object. | [optional] 
**Expression** | Pointer to **string** | The conditional expression of a DHCP filter option object. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | Determines the lease time of a DHCP filter option object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP option filter object. | [optional] 
**NextServer** | Pointer to **string** | Determines the next server of a DHCP filter option object. You can specify the name and/or IP address of the next server that the host needs to boot. | [optional] 
**OptionList** | Pointer to [**[]FilteroptionOptionList**](FilteroptionOptionList.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**OptionSpace** | Pointer to **string** | The option space of a DHCP filter option object. | [optional] 
**PxeLeaseTime** | Pointer to **int64** | Determines the PXE (Preboot Execution Environment) lease time of a DHCP filter option object. To specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Filteroption**](Filteroption.md) |  | [optional] 

## Methods

### NewGetFilteroptionResponse

`func NewGetFilteroptionResponse() *GetFilteroptionResponse`

NewGetFilteroptionResponse instantiates a new GetFilteroptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFilteroptionResponseWithDefaults

`func NewGetFilteroptionResponseWithDefaults() *GetFilteroptionResponse`

NewGetFilteroptionResponseWithDefaults instantiates a new GetFilteroptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFilteroptionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFilteroptionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFilteroptionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFilteroptionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetApplyAsClass

`func (o *GetFilteroptionResponse) GetApplyAsClass() bool`

GetApplyAsClass returns the ApplyAsClass field if non-nil, zero value otherwise.

### GetApplyAsClassOk

`func (o *GetFilteroptionResponse) GetApplyAsClassOk() (*bool, bool)`

GetApplyAsClassOk returns a tuple with the ApplyAsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAsClass

`func (o *GetFilteroptionResponse) SetApplyAsClass(v bool)`

SetApplyAsClass sets ApplyAsClass field to given value.

### HasApplyAsClass

`func (o *GetFilteroptionResponse) HasApplyAsClass() bool`

HasApplyAsClass returns a boolean if a field has been set.

### GetBootfile

`func (o *GetFilteroptionResponse) GetBootfile() string`

GetBootfile returns the Bootfile field if non-nil, zero value otherwise.

### GetBootfileOk

`func (o *GetFilteroptionResponse) GetBootfileOk() (*string, bool)`

GetBootfileOk returns a tuple with the Bootfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootfile

`func (o *GetFilteroptionResponse) SetBootfile(v string)`

SetBootfile sets Bootfile field to given value.

### HasBootfile

`func (o *GetFilteroptionResponse) HasBootfile() bool`

HasBootfile returns a boolean if a field has been set.

### GetBootserver

`func (o *GetFilteroptionResponse) GetBootserver() string`

GetBootserver returns the Bootserver field if non-nil, zero value otherwise.

### GetBootserverOk

`func (o *GetFilteroptionResponse) GetBootserverOk() (*string, bool)`

GetBootserverOk returns a tuple with the Bootserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootserver

`func (o *GetFilteroptionResponse) SetBootserver(v string)`

SetBootserver sets Bootserver field to given value.

### HasBootserver

`func (o *GetFilteroptionResponse) HasBootserver() bool`

HasBootserver returns a boolean if a field has been set.

### GetComment

`func (o *GetFilteroptionResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFilteroptionResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFilteroptionResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFilteroptionResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpression

`func (o *GetFilteroptionResponse) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *GetFilteroptionResponse) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *GetFilteroptionResponse) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *GetFilteroptionResponse) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetFilteroptionResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetFilteroptionResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetFilteroptionResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetFilteroptionResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetFilteroptionResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetFilteroptionResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetFilteroptionResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetFilteroptionResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetFilteroptionResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetFilteroptionResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetFilteroptionResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetFilteroptionResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLeaseTime

`func (o *GetFilteroptionResponse) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *GetFilteroptionResponse) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *GetFilteroptionResponse) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *GetFilteroptionResponse) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *GetFilteroptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFilteroptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFilteroptionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFilteroptionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextServer

`func (o *GetFilteroptionResponse) GetNextServer() string`

GetNextServer returns the NextServer field if non-nil, zero value otherwise.

### GetNextServerOk

`func (o *GetFilteroptionResponse) GetNextServerOk() (*string, bool)`

GetNextServerOk returns a tuple with the NextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextServer

`func (o *GetFilteroptionResponse) SetNextServer(v string)`

SetNextServer sets NextServer field to given value.

### HasNextServer

`func (o *GetFilteroptionResponse) HasNextServer() bool`

HasNextServer returns a boolean if a field has been set.

### GetOptionList

`func (o *GetFilteroptionResponse) GetOptionList() []FilteroptionOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *GetFilteroptionResponse) GetOptionListOk() (*[]FilteroptionOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *GetFilteroptionResponse) SetOptionList(v []FilteroptionOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *GetFilteroptionResponse) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetOptionSpace

`func (o *GetFilteroptionResponse) GetOptionSpace() string`

GetOptionSpace returns the OptionSpace field if non-nil, zero value otherwise.

### GetOptionSpaceOk

`func (o *GetFilteroptionResponse) GetOptionSpaceOk() (*string, bool)`

GetOptionSpaceOk returns a tuple with the OptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSpace

`func (o *GetFilteroptionResponse) SetOptionSpace(v string)`

SetOptionSpace sets OptionSpace field to given value.

### HasOptionSpace

`func (o *GetFilteroptionResponse) HasOptionSpace() bool`

HasOptionSpace returns a boolean if a field has been set.

### GetPxeLeaseTime

`func (o *GetFilteroptionResponse) GetPxeLeaseTime() int64`

GetPxeLeaseTime returns the PxeLeaseTime field if non-nil, zero value otherwise.

### GetPxeLeaseTimeOk

`func (o *GetFilteroptionResponse) GetPxeLeaseTimeOk() (*int64, bool)`

GetPxeLeaseTimeOk returns a tuple with the PxeLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeLeaseTime

`func (o *GetFilteroptionResponse) SetPxeLeaseTime(v int64)`

SetPxeLeaseTime sets PxeLeaseTime field to given value.

### HasPxeLeaseTime

`func (o *GetFilteroptionResponse) HasPxeLeaseTime() bool`

HasPxeLeaseTime returns a boolean if a field has been set.

### GetUuid

`func (o *GetFilteroptionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetFilteroptionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetFilteroptionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetFilteroptionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetFilteroptionResponse) GetResult() Filteroption`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFilteroptionResponse) GetResultOk() (*Filteroption, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFilteroptionResponse) SetResult(v Filteroption)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFilteroptionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


