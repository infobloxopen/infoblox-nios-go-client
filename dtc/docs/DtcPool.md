# DtcPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AutoConsolidatedMonitors** | Pointer to **bool** | Flag for enabling auto managing DTC Consolidated Monitors in DTC Pool. | [optional] 
**Availability** | Pointer to **string** | A resource in the pool is available if ANY, at least QUORUM, or ALL monitors for the pool say that it is up. | [optional] 
**Comment** | Pointer to **string** | The comment for the DTC Pool; maximum 256 characters. | [optional] 
**ConsolidatedMonitors** | Pointer to [**[]DtcPoolConsolidatedMonitors**](DtcPoolConsolidatedMonitors.md) | List of monitors and associated members statuses of which are shared across members and consolidated in server availability determination. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the DTC Pool is disabled or not. When this is set to False, the fixed address is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Health** | Pointer to [**DtcPoolHealth**](DtcPoolHealth.md) |  | [optional] 
**LbAlternateMethod** | Pointer to **string** | The alternate load balancing method. Use this to select a method type from the pool if the preferred method does not return any results. | [optional] 
**LbAlternateTopology** | Pointer to **string** | The alternate topology for load balancing. | [optional] 
**LbDynamicRatioAlternate** | Pointer to [**DtcPoolLbDynamicRatioAlternate**](DtcPoolLbDynamicRatioAlternate.md) |  | [optional] 
**LbDynamicRatioPreferred** | Pointer to [**DtcPoolLbDynamicRatioPreferred**](DtcPoolLbDynamicRatioPreferred.md) |  | [optional] 
**LbPreferredMethod** | Pointer to **string** | The preferred load balancing method. Use this to select a method type from the pool. | [optional] 
**LbPreferredTopology** | Pointer to **string** | The preferred topology for load balancing. | [optional] 
**Monitors** | Pointer to **[]string** | The monitors related to pool. | [optional] 
**Name** | Pointer to **string** | The DTC Pool display name. | [optional] 
**Quorum** | Pointer to **int64** | For availability mode QUORUM, at least this many monitors must report the resource as up for it to be available | [optional] 
**Servers** | Pointer to [**[]DtcPoolServers**](DtcPoolServers.md) | The servers related to the pool. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the DTC Pool. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewDtcPool

`func NewDtcPool() *DtcPool`

NewDtcPool instantiates a new DtcPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcPoolWithDefaults

`func NewDtcPoolWithDefaults() *DtcPool`

NewDtcPoolWithDefaults instantiates a new DtcPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcPool) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcPool) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcPool) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcPool) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoConsolidatedMonitors

`func (o *DtcPool) GetAutoConsolidatedMonitors() bool`

GetAutoConsolidatedMonitors returns the AutoConsolidatedMonitors field if non-nil, zero value otherwise.

### GetAutoConsolidatedMonitorsOk

`func (o *DtcPool) GetAutoConsolidatedMonitorsOk() (*bool, bool)`

GetAutoConsolidatedMonitorsOk returns a tuple with the AutoConsolidatedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidatedMonitors

`func (o *DtcPool) SetAutoConsolidatedMonitors(v bool)`

SetAutoConsolidatedMonitors sets AutoConsolidatedMonitors field to given value.

### HasAutoConsolidatedMonitors

`func (o *DtcPool) HasAutoConsolidatedMonitors() bool`

HasAutoConsolidatedMonitors returns a boolean if a field has been set.

### GetAvailability

`func (o *DtcPool) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *DtcPool) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *DtcPool) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *DtcPool) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetComment

`func (o *DtcPool) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcPool) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcPool) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcPool) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidatedMonitors

`func (o *DtcPool) GetConsolidatedMonitors() []DtcPoolConsolidatedMonitors`

GetConsolidatedMonitors returns the ConsolidatedMonitors field if non-nil, zero value otherwise.

### GetConsolidatedMonitorsOk

`func (o *DtcPool) GetConsolidatedMonitorsOk() (*[]DtcPoolConsolidatedMonitors, bool)`

GetConsolidatedMonitorsOk returns a tuple with the ConsolidatedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedMonitors

`func (o *DtcPool) SetConsolidatedMonitors(v []DtcPoolConsolidatedMonitors)`

SetConsolidatedMonitors sets ConsolidatedMonitors field to given value.

### HasConsolidatedMonitors

`func (o *DtcPool) HasConsolidatedMonitors() bool`

HasConsolidatedMonitors returns a boolean if a field has been set.

### GetDisable

`func (o *DtcPool) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcPool) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcPool) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcPool) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *DtcPool) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *DtcPool) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *DtcPool) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *DtcPool) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *DtcPool) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *DtcPool) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *DtcPool) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *DtcPool) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DtcPool) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DtcPool) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DtcPool) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DtcPool) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHealth

`func (o *DtcPool) GetHealth() DtcPoolHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DtcPool) GetHealthOk() (*DtcPoolHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DtcPool) SetHealth(v DtcPoolHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DtcPool) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLbAlternateMethod

`func (o *DtcPool) GetLbAlternateMethod() string`

GetLbAlternateMethod returns the LbAlternateMethod field if non-nil, zero value otherwise.

### GetLbAlternateMethodOk

`func (o *DtcPool) GetLbAlternateMethodOk() (*string, bool)`

GetLbAlternateMethodOk returns a tuple with the LbAlternateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbAlternateMethod

`func (o *DtcPool) SetLbAlternateMethod(v string)`

SetLbAlternateMethod sets LbAlternateMethod field to given value.

### HasLbAlternateMethod

`func (o *DtcPool) HasLbAlternateMethod() bool`

HasLbAlternateMethod returns a boolean if a field has been set.

### GetLbAlternateTopology

`func (o *DtcPool) GetLbAlternateTopology() string`

GetLbAlternateTopology returns the LbAlternateTopology field if non-nil, zero value otherwise.

### GetLbAlternateTopologyOk

`func (o *DtcPool) GetLbAlternateTopologyOk() (*string, bool)`

GetLbAlternateTopologyOk returns a tuple with the LbAlternateTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbAlternateTopology

`func (o *DtcPool) SetLbAlternateTopology(v string)`

SetLbAlternateTopology sets LbAlternateTopology field to given value.

### HasLbAlternateTopology

`func (o *DtcPool) HasLbAlternateTopology() bool`

HasLbAlternateTopology returns a boolean if a field has been set.

### GetLbDynamicRatioAlternate

`func (o *DtcPool) GetLbDynamicRatioAlternate() DtcPoolLbDynamicRatioAlternate`

GetLbDynamicRatioAlternate returns the LbDynamicRatioAlternate field if non-nil, zero value otherwise.

### GetLbDynamicRatioAlternateOk

`func (o *DtcPool) GetLbDynamicRatioAlternateOk() (*DtcPoolLbDynamicRatioAlternate, bool)`

GetLbDynamicRatioAlternateOk returns a tuple with the LbDynamicRatioAlternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbDynamicRatioAlternate

`func (o *DtcPool) SetLbDynamicRatioAlternate(v DtcPoolLbDynamicRatioAlternate)`

SetLbDynamicRatioAlternate sets LbDynamicRatioAlternate field to given value.

### HasLbDynamicRatioAlternate

`func (o *DtcPool) HasLbDynamicRatioAlternate() bool`

HasLbDynamicRatioAlternate returns a boolean if a field has been set.

### GetLbDynamicRatioPreferred

`func (o *DtcPool) GetLbDynamicRatioPreferred() DtcPoolLbDynamicRatioPreferred`

GetLbDynamicRatioPreferred returns the LbDynamicRatioPreferred field if non-nil, zero value otherwise.

### GetLbDynamicRatioPreferredOk

`func (o *DtcPool) GetLbDynamicRatioPreferredOk() (*DtcPoolLbDynamicRatioPreferred, bool)`

GetLbDynamicRatioPreferredOk returns a tuple with the LbDynamicRatioPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbDynamicRatioPreferred

`func (o *DtcPool) SetLbDynamicRatioPreferred(v DtcPoolLbDynamicRatioPreferred)`

SetLbDynamicRatioPreferred sets LbDynamicRatioPreferred field to given value.

### HasLbDynamicRatioPreferred

`func (o *DtcPool) HasLbDynamicRatioPreferred() bool`

HasLbDynamicRatioPreferred returns a boolean if a field has been set.

### GetLbPreferredMethod

`func (o *DtcPool) GetLbPreferredMethod() string`

GetLbPreferredMethod returns the LbPreferredMethod field if non-nil, zero value otherwise.

### GetLbPreferredMethodOk

`func (o *DtcPool) GetLbPreferredMethodOk() (*string, bool)`

GetLbPreferredMethodOk returns a tuple with the LbPreferredMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPreferredMethod

`func (o *DtcPool) SetLbPreferredMethod(v string)`

SetLbPreferredMethod sets LbPreferredMethod field to given value.

### HasLbPreferredMethod

`func (o *DtcPool) HasLbPreferredMethod() bool`

HasLbPreferredMethod returns a boolean if a field has been set.

### GetLbPreferredTopology

`func (o *DtcPool) GetLbPreferredTopology() string`

GetLbPreferredTopology returns the LbPreferredTopology field if non-nil, zero value otherwise.

### GetLbPreferredTopologyOk

`func (o *DtcPool) GetLbPreferredTopologyOk() (*string, bool)`

GetLbPreferredTopologyOk returns a tuple with the LbPreferredTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPreferredTopology

`func (o *DtcPool) SetLbPreferredTopology(v string)`

SetLbPreferredTopology sets LbPreferredTopology field to given value.

### HasLbPreferredTopology

`func (o *DtcPool) HasLbPreferredTopology() bool`

HasLbPreferredTopology returns a boolean if a field has been set.

### GetMonitors

`func (o *DtcPool) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *DtcPool) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *DtcPool) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *DtcPool) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetName

`func (o *DtcPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuorum

`func (o *DtcPool) GetQuorum() int64`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *DtcPool) GetQuorumOk() (*int64, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *DtcPool) SetQuorum(v int64)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *DtcPool) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetServers

`func (o *DtcPool) GetServers() []DtcPoolServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DtcPool) GetServersOk() (*[]DtcPoolServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DtcPool) SetServers(v []DtcPoolServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DtcPool) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTtl

`func (o *DtcPool) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcPool) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcPool) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcPool) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *DtcPool) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *DtcPool) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *DtcPool) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *DtcPool) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetUuid

`func (o *DtcPool) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DtcPool) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DtcPool) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DtcPool) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


