# GetMssuperscopeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The superscope descriptive comment. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP usage of the ranges in the superscope. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | Utilization level of the DHCP range objects that belong to the superscope. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether the superscope is disabled. | [optional] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the DHCP range objects that belong to the superscope. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage value for DHCP range usage after which an alarm will be active. | [optional] [readonly] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage value for DHCP range usage after which an alarm will be reset. | [optional] [readonly] 
**LowWaterMark** | Pointer to **int64** | The percentage value for DHCP range usage below which an alarm will be active. | [optional] [readonly] 
**LowWaterMarkReset** | Pointer to **int64** | The percentage value for DHCP range usage below which an alarm will be reset. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Microsoft DHCP superscope. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which the superscope resides. | [optional] 
**Ranges** | Pointer to **[]string** | The list of DHCP ranges that are associated with the superscope. | [optional] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in DHCP range objects that belong to the superscope. | [optional] [readonly] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in DHCP range objects that belong to the superscope. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Mssuperscope**](Mssuperscope.md) |  | [optional] 

## Methods

### NewGetMssuperscopeResponse

`func NewGetMssuperscopeResponse() *GetMssuperscopeResponse`

NewGetMssuperscopeResponse instantiates a new GetMssuperscopeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMssuperscopeResponseWithDefaults

`func NewGetMssuperscopeResponseWithDefaults() *GetMssuperscopeResponse`

NewGetMssuperscopeResponseWithDefaults instantiates a new GetMssuperscopeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMssuperscopeResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMssuperscopeResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMssuperscopeResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMssuperscopeResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetMssuperscopeResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetMssuperscopeResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetMssuperscopeResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetMssuperscopeResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *GetMssuperscopeResponse) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *GetMssuperscopeResponse) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *GetMssuperscopeResponse) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *GetMssuperscopeResponse) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *GetMssuperscopeResponse) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *GetMssuperscopeResponse) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *GetMssuperscopeResponse) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *GetMssuperscopeResponse) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *GetMssuperscopeResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetMssuperscopeResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetMssuperscopeResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetMssuperscopeResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *GetMssuperscopeResponse) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *GetMssuperscopeResponse) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *GetMssuperscopeResponse) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *GetMssuperscopeResponse) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetMssuperscopeResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetMssuperscopeResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetMssuperscopeResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetMssuperscopeResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetMssuperscopeResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetMssuperscopeResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetMssuperscopeResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetMssuperscopeResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetMssuperscopeResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetMssuperscopeResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetMssuperscopeResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetMssuperscopeResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *GetMssuperscopeResponse) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *GetMssuperscopeResponse) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *GetMssuperscopeResponse) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *GetMssuperscopeResponse) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *GetMssuperscopeResponse) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *GetMssuperscopeResponse) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *GetMssuperscopeResponse) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *GetMssuperscopeResponse) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *GetMssuperscopeResponse) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *GetMssuperscopeResponse) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *GetMssuperscopeResponse) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *GetMssuperscopeResponse) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *GetMssuperscopeResponse) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *GetMssuperscopeResponse) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *GetMssuperscopeResponse) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *GetMssuperscopeResponse) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetName

`func (o *GetMssuperscopeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMssuperscopeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMssuperscopeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMssuperscopeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetMssuperscopeResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetMssuperscopeResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetMssuperscopeResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetMssuperscopeResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetRanges

`func (o *GetMssuperscopeResponse) GetRanges() []string`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GetMssuperscopeResponse) GetRangesOk() (*[]string, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GetMssuperscopeResponse) SetRanges(v []string)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *GetMssuperscopeResponse) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetStaticHosts

`func (o *GetMssuperscopeResponse) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *GetMssuperscopeResponse) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *GetMssuperscopeResponse) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *GetMssuperscopeResponse) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetTotalHosts

`func (o *GetMssuperscopeResponse) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *GetMssuperscopeResponse) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *GetMssuperscopeResponse) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *GetMssuperscopeResponse) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.

### GetUuid

`func (o *GetMssuperscopeResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetMssuperscopeResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetMssuperscopeResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetMssuperscopeResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetMssuperscopeResponse) GetResult() Mssuperscope`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMssuperscopeResponse) GetResultOk() (*Mssuperscope, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMssuperscopeResponse) SetResult(v Mssuperscope)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMssuperscopeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


