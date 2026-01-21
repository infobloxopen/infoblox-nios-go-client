# GetDtcPoolResponse

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
**Result** | Pointer to [**DtcPool**](DtcPool.md) |  | [optional] 

## Methods

### NewGetDtcPoolResponse

`func NewGetDtcPoolResponse() *GetDtcPoolResponse`

NewGetDtcPoolResponse instantiates a new GetDtcPoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcPoolResponseWithDefaults

`func NewGetDtcPoolResponseWithDefaults() *GetDtcPoolResponse`

NewGetDtcPoolResponseWithDefaults instantiates a new GetDtcPoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcPoolResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcPoolResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcPoolResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcPoolResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoConsolidatedMonitors

`func (o *GetDtcPoolResponse) GetAutoConsolidatedMonitors() bool`

GetAutoConsolidatedMonitors returns the AutoConsolidatedMonitors field if non-nil, zero value otherwise.

### GetAutoConsolidatedMonitorsOk

`func (o *GetDtcPoolResponse) GetAutoConsolidatedMonitorsOk() (*bool, bool)`

GetAutoConsolidatedMonitorsOk returns a tuple with the AutoConsolidatedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidatedMonitors

`func (o *GetDtcPoolResponse) SetAutoConsolidatedMonitors(v bool)`

SetAutoConsolidatedMonitors sets AutoConsolidatedMonitors field to given value.

### HasAutoConsolidatedMonitors

`func (o *GetDtcPoolResponse) HasAutoConsolidatedMonitors() bool`

HasAutoConsolidatedMonitors returns a boolean if a field has been set.

### GetAvailability

`func (o *GetDtcPoolResponse) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetDtcPoolResponse) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetDtcPoolResponse) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetDtcPoolResponse) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcPoolResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcPoolResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcPoolResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcPoolResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConsolidatedMonitors

`func (o *GetDtcPoolResponse) GetConsolidatedMonitors() []DtcPoolConsolidatedMonitors`

GetConsolidatedMonitors returns the ConsolidatedMonitors field if non-nil, zero value otherwise.

### GetConsolidatedMonitorsOk

`func (o *GetDtcPoolResponse) GetConsolidatedMonitorsOk() (*[]DtcPoolConsolidatedMonitors, bool)`

GetConsolidatedMonitorsOk returns a tuple with the ConsolidatedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidatedMonitors

`func (o *GetDtcPoolResponse) SetConsolidatedMonitors(v []DtcPoolConsolidatedMonitors)`

SetConsolidatedMonitors sets ConsolidatedMonitors field to given value.

### HasConsolidatedMonitors

`func (o *GetDtcPoolResponse) HasConsolidatedMonitors() bool`

HasConsolidatedMonitors returns a boolean if a field has been set.

### GetDisable

`func (o *GetDtcPoolResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDtcPoolResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDtcPoolResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDtcPoolResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetDtcPoolResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetDtcPoolResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetDtcPoolResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetDtcPoolResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetDtcPoolResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetDtcPoolResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetDtcPoolResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetDtcPoolResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDtcPoolResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDtcPoolResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDtcPoolResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDtcPoolResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHealth

`func (o *GetDtcPoolResponse) GetHealth() DtcPoolHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetDtcPoolResponse) GetHealthOk() (*DtcPoolHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetDtcPoolResponse) SetHealth(v DtcPoolHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetDtcPoolResponse) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLbAlternateMethod

`func (o *GetDtcPoolResponse) GetLbAlternateMethod() string`

GetLbAlternateMethod returns the LbAlternateMethod field if non-nil, zero value otherwise.

### GetLbAlternateMethodOk

`func (o *GetDtcPoolResponse) GetLbAlternateMethodOk() (*string, bool)`

GetLbAlternateMethodOk returns a tuple with the LbAlternateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbAlternateMethod

`func (o *GetDtcPoolResponse) SetLbAlternateMethod(v string)`

SetLbAlternateMethod sets LbAlternateMethod field to given value.

### HasLbAlternateMethod

`func (o *GetDtcPoolResponse) HasLbAlternateMethod() bool`

HasLbAlternateMethod returns a boolean if a field has been set.

### GetLbAlternateTopology

`func (o *GetDtcPoolResponse) GetLbAlternateTopology() string`

GetLbAlternateTopology returns the LbAlternateTopology field if non-nil, zero value otherwise.

### GetLbAlternateTopologyOk

`func (o *GetDtcPoolResponse) GetLbAlternateTopologyOk() (*string, bool)`

GetLbAlternateTopologyOk returns a tuple with the LbAlternateTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbAlternateTopology

`func (o *GetDtcPoolResponse) SetLbAlternateTopology(v string)`

SetLbAlternateTopology sets LbAlternateTopology field to given value.

### HasLbAlternateTopology

`func (o *GetDtcPoolResponse) HasLbAlternateTopology() bool`

HasLbAlternateTopology returns a boolean if a field has been set.

### GetLbDynamicRatioAlternate

`func (o *GetDtcPoolResponse) GetLbDynamicRatioAlternate() DtcPoolLbDynamicRatioAlternate`

GetLbDynamicRatioAlternate returns the LbDynamicRatioAlternate field if non-nil, zero value otherwise.

### GetLbDynamicRatioAlternateOk

`func (o *GetDtcPoolResponse) GetLbDynamicRatioAlternateOk() (*DtcPoolLbDynamicRatioAlternate, bool)`

GetLbDynamicRatioAlternateOk returns a tuple with the LbDynamicRatioAlternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbDynamicRatioAlternate

`func (o *GetDtcPoolResponse) SetLbDynamicRatioAlternate(v DtcPoolLbDynamicRatioAlternate)`

SetLbDynamicRatioAlternate sets LbDynamicRatioAlternate field to given value.

### HasLbDynamicRatioAlternate

`func (o *GetDtcPoolResponse) HasLbDynamicRatioAlternate() bool`

HasLbDynamicRatioAlternate returns a boolean if a field has been set.

### GetLbDynamicRatioPreferred

`func (o *GetDtcPoolResponse) GetLbDynamicRatioPreferred() DtcPoolLbDynamicRatioPreferred`

GetLbDynamicRatioPreferred returns the LbDynamicRatioPreferred field if non-nil, zero value otherwise.

### GetLbDynamicRatioPreferredOk

`func (o *GetDtcPoolResponse) GetLbDynamicRatioPreferredOk() (*DtcPoolLbDynamicRatioPreferred, bool)`

GetLbDynamicRatioPreferredOk returns a tuple with the LbDynamicRatioPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbDynamicRatioPreferred

`func (o *GetDtcPoolResponse) SetLbDynamicRatioPreferred(v DtcPoolLbDynamicRatioPreferred)`

SetLbDynamicRatioPreferred sets LbDynamicRatioPreferred field to given value.

### HasLbDynamicRatioPreferred

`func (o *GetDtcPoolResponse) HasLbDynamicRatioPreferred() bool`

HasLbDynamicRatioPreferred returns a boolean if a field has been set.

### GetLbPreferredMethod

`func (o *GetDtcPoolResponse) GetLbPreferredMethod() string`

GetLbPreferredMethod returns the LbPreferredMethod field if non-nil, zero value otherwise.

### GetLbPreferredMethodOk

`func (o *GetDtcPoolResponse) GetLbPreferredMethodOk() (*string, bool)`

GetLbPreferredMethodOk returns a tuple with the LbPreferredMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPreferredMethod

`func (o *GetDtcPoolResponse) SetLbPreferredMethod(v string)`

SetLbPreferredMethod sets LbPreferredMethod field to given value.

### HasLbPreferredMethod

`func (o *GetDtcPoolResponse) HasLbPreferredMethod() bool`

HasLbPreferredMethod returns a boolean if a field has been set.

### GetLbPreferredTopology

`func (o *GetDtcPoolResponse) GetLbPreferredTopology() string`

GetLbPreferredTopology returns the LbPreferredTopology field if non-nil, zero value otherwise.

### GetLbPreferredTopologyOk

`func (o *GetDtcPoolResponse) GetLbPreferredTopologyOk() (*string, bool)`

GetLbPreferredTopologyOk returns a tuple with the LbPreferredTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPreferredTopology

`func (o *GetDtcPoolResponse) SetLbPreferredTopology(v string)`

SetLbPreferredTopology sets LbPreferredTopology field to given value.

### HasLbPreferredTopology

`func (o *GetDtcPoolResponse) HasLbPreferredTopology() bool`

HasLbPreferredTopology returns a boolean if a field has been set.

### GetMonitors

`func (o *GetDtcPoolResponse) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *GetDtcPoolResponse) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *GetDtcPoolResponse) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *GetDtcPoolResponse) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetName

`func (o *GetDtcPoolResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcPoolResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcPoolResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcPoolResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuorum

`func (o *GetDtcPoolResponse) GetQuorum() int64`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *GetDtcPoolResponse) GetQuorumOk() (*int64, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *GetDtcPoolResponse) SetQuorum(v int64)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *GetDtcPoolResponse) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetServers

`func (o *GetDtcPoolResponse) GetServers() []DtcPoolServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetDtcPoolResponse) GetServersOk() (*[]DtcPoolServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetDtcPoolResponse) SetServers(v []DtcPoolServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetDtcPoolResponse) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTtl

`func (o *GetDtcPoolResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetDtcPoolResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetDtcPoolResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetDtcPoolResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetDtcPoolResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetDtcPoolResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetDtcPoolResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetDtcPoolResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetUuid

`func (o *GetDtcPoolResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDtcPoolResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDtcPoolResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDtcPoolResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcPoolResponse) GetResult() DtcPool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcPoolResponse) GetResultOk() (*DtcPool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcPoolResponse) SetResult(v DtcPool)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcPoolResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


