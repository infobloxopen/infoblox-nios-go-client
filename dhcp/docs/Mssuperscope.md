# Mssuperscope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The superscope descriptive comment. | [optional] 
**DhcpUtilization** | Pointer to **int64** | The percentage of the total DHCP usage of the ranges in the superscope. | [optional] [readonly] 
**DhcpUtilizationStatus** | Pointer to **string** | Utilization level of the DHCP range objects that belong to the superscope. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether the superscope is disabled. | [optional] 
**DynamicHosts** | Pointer to **int64** | The total number of DHCP leases issued for the DHCP range objects that belong to the superscope. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**HighWaterMark** | Pointer to **int64** | The percentage value for DHCP range usage after which an alarm will be active. | [optional] [readonly] 
**HighWaterMarkReset** | Pointer to **int64** | The percentage value for DHCP range usage after which an alarm will be reset. | [optional] [readonly] 
**LowWaterMark** | Pointer to **int64** | The percentage value for DHCP range usage below which an alarm will be active. | [optional] [readonly] 
**LowWaterMarkReset** | Pointer to **int64** | The percentage value for DHCP range usage below which an alarm will be reset. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Microsoft DHCP superscope. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which the superscope resides. | [optional] 
**Ranges** | Pointer to **[]map[string]interface{}** | The list of DHCP ranges that are associated with the superscope. | [optional] 
**StaticHosts** | Pointer to **int64** | The number of static DHCP addresses configured in DHCP range objects that belong to the superscope. | [optional] [readonly] 
**TotalHosts** | Pointer to **int64** | The total number of DHCP addresses configured in DHCP range objects that belong to the superscope. | [optional] [readonly] 

## Methods

### NewMssuperscope

`func NewMssuperscope() *Mssuperscope`

NewMssuperscope instantiates a new Mssuperscope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssuperscopeWithDefaults

`func NewMssuperscopeWithDefaults() *Mssuperscope`

NewMssuperscopeWithDefaults instantiates a new Mssuperscope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Mssuperscope) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Mssuperscope) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Mssuperscope) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Mssuperscope) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Mssuperscope) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Mssuperscope) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Mssuperscope) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Mssuperscope) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDhcpUtilization

`func (o *Mssuperscope) GetDhcpUtilization() int64`

GetDhcpUtilization returns the DhcpUtilization field if non-nil, zero value otherwise.

### GetDhcpUtilizationOk

`func (o *Mssuperscope) GetDhcpUtilizationOk() (*int64, bool)`

GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilization

`func (o *Mssuperscope) SetDhcpUtilization(v int64)`

SetDhcpUtilization sets DhcpUtilization field to given value.

### HasDhcpUtilization

`func (o *Mssuperscope) HasDhcpUtilization() bool`

HasDhcpUtilization returns a boolean if a field has been set.

### GetDhcpUtilizationStatus

`func (o *Mssuperscope) GetDhcpUtilizationStatus() string`

GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field if non-nil, zero value otherwise.

### GetDhcpUtilizationStatusOk

`func (o *Mssuperscope) GetDhcpUtilizationStatusOk() (*string, bool)`

GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpUtilizationStatus

`func (o *Mssuperscope) SetDhcpUtilizationStatus(v string)`

SetDhcpUtilizationStatus sets DhcpUtilizationStatus field to given value.

### HasDhcpUtilizationStatus

`func (o *Mssuperscope) HasDhcpUtilizationStatus() bool`

HasDhcpUtilizationStatus returns a boolean if a field has been set.

### GetDisable

`func (o *Mssuperscope) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Mssuperscope) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Mssuperscope) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Mssuperscope) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDynamicHosts

`func (o *Mssuperscope) GetDynamicHosts() int64`

GetDynamicHosts returns the DynamicHosts field if non-nil, zero value otherwise.

### GetDynamicHostsOk

`func (o *Mssuperscope) GetDynamicHostsOk() (*int64, bool)`

GetDynamicHostsOk returns a tuple with the DynamicHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicHosts

`func (o *Mssuperscope) SetDynamicHosts(v int64)`

SetDynamicHosts sets DynamicHosts field to given value.

### HasDynamicHosts

`func (o *Mssuperscope) HasDynamicHosts() bool`

HasDynamicHosts returns a boolean if a field has been set.

### GetExtattrs

`func (o *Mssuperscope) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Mssuperscope) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Mssuperscope) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Mssuperscope) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetHighWaterMark

`func (o *Mssuperscope) GetHighWaterMark() int64`

GetHighWaterMark returns the HighWaterMark field if non-nil, zero value otherwise.

### GetHighWaterMarkOk

`func (o *Mssuperscope) GetHighWaterMarkOk() (*int64, bool)`

GetHighWaterMarkOk returns a tuple with the HighWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMark

`func (o *Mssuperscope) SetHighWaterMark(v int64)`

SetHighWaterMark sets HighWaterMark field to given value.

### HasHighWaterMark

`func (o *Mssuperscope) HasHighWaterMark() bool`

HasHighWaterMark returns a boolean if a field has been set.

### GetHighWaterMarkReset

`func (o *Mssuperscope) GetHighWaterMarkReset() int64`

GetHighWaterMarkReset returns the HighWaterMarkReset field if non-nil, zero value otherwise.

### GetHighWaterMarkResetOk

`func (o *Mssuperscope) GetHighWaterMarkResetOk() (*int64, bool)`

GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighWaterMarkReset

`func (o *Mssuperscope) SetHighWaterMarkReset(v int64)`

SetHighWaterMarkReset sets HighWaterMarkReset field to given value.

### HasHighWaterMarkReset

`func (o *Mssuperscope) HasHighWaterMarkReset() bool`

HasHighWaterMarkReset returns a boolean if a field has been set.

### GetLowWaterMark

`func (o *Mssuperscope) GetLowWaterMark() int64`

GetLowWaterMark returns the LowWaterMark field if non-nil, zero value otherwise.

### GetLowWaterMarkOk

`func (o *Mssuperscope) GetLowWaterMarkOk() (*int64, bool)`

GetLowWaterMarkOk returns a tuple with the LowWaterMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMark

`func (o *Mssuperscope) SetLowWaterMark(v int64)`

SetLowWaterMark sets LowWaterMark field to given value.

### HasLowWaterMark

`func (o *Mssuperscope) HasLowWaterMark() bool`

HasLowWaterMark returns a boolean if a field has been set.

### GetLowWaterMarkReset

`func (o *Mssuperscope) GetLowWaterMarkReset() int64`

GetLowWaterMarkReset returns the LowWaterMarkReset field if non-nil, zero value otherwise.

### GetLowWaterMarkResetOk

`func (o *Mssuperscope) GetLowWaterMarkResetOk() (*int64, bool)`

GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowWaterMarkReset

`func (o *Mssuperscope) SetLowWaterMarkReset(v int64)`

SetLowWaterMarkReset sets LowWaterMarkReset field to given value.

### HasLowWaterMarkReset

`func (o *Mssuperscope) HasLowWaterMarkReset() bool`

HasLowWaterMarkReset returns a boolean if a field has been set.

### GetName

`func (o *Mssuperscope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mssuperscope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mssuperscope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mssuperscope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *Mssuperscope) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Mssuperscope) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Mssuperscope) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Mssuperscope) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetRanges

`func (o *Mssuperscope) GetRanges() []map[string]interface{}`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *Mssuperscope) GetRangesOk() (*[]map[string]interface{}, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *Mssuperscope) SetRanges(v []map[string]interface{})`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *Mssuperscope) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetStaticHosts

`func (o *Mssuperscope) GetStaticHosts() int64`

GetStaticHosts returns the StaticHosts field if non-nil, zero value otherwise.

### GetStaticHostsOk

`func (o *Mssuperscope) GetStaticHostsOk() (*int64, bool)`

GetStaticHostsOk returns a tuple with the StaticHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticHosts

`func (o *Mssuperscope) SetStaticHosts(v int64)`

SetStaticHosts sets StaticHosts field to given value.

### HasStaticHosts

`func (o *Mssuperscope) HasStaticHosts() bool`

HasStaticHosts returns a boolean if a field has been set.

### GetTotalHosts

`func (o *Mssuperscope) GetTotalHosts() int64`

GetTotalHosts returns the TotalHosts field if non-nil, zero value otherwise.

### GetTotalHostsOk

`func (o *Mssuperscope) GetTotalHostsOk() (*int64, bool)`

GetTotalHostsOk returns a tuple with the TotalHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHosts

`func (o *Mssuperscope) SetTotalHosts(v int64)`

SetTotalHosts sets TotalHosts field to given value.

### HasTotalHosts

`func (o *Mssuperscope) HasTotalHosts() bool`

HasTotalHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


