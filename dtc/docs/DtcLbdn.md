# DtcLbdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AuthZones** | Pointer to **[]string** | List of linked auth zones. | [optional] 
**AutoConsolidatedMonitors** | Pointer to **bool** | Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools. | [optional] 
**Comment** | Pointer to **string** | Comment for the DTC LBDN; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Health** | Pointer to [**DtcLbdnHealth**](DtcLbdnHealth.md) |  | [optional] 
**LbMethod** | Pointer to **string** | The load balancing method. Used to select pool. | [optional] 
**Name** | Pointer to **string** | The display name of the DTC LBDN, not DNS related. | [optional] 
**Patterns** | Pointer to **[]string** | LBDN wildcards for pattern match. | [optional] 
**Persistence** | Pointer to **int64** | Maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching. | [optional] 
**Pools** | Pointer to [**[]DtcLbdnPools**](DtcLbdnPools.md) | The maximum time, in seconds, for which client specific LBDN responses will be cached. Zero specifies no caching. | [optional] 
**Priority** | Pointer to **int64** | The LBDN pattern match priority for \&quot;overlapping\&quot; DTC LBDN objects. LBDNs are \&quot;overlapping\&quot; if they are simultaneously assigned to a zone and have patterns that can match the same FQDN. The matching LBDN with highest priority (lowest ordinal) will be used. | [optional] 
**Topology** | Pointer to **string** | The topology rules for TOPOLOGY method. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the DTC LBDN. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**Types** | Pointer to **[]string** | The list of resource record types supported by LBDN. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 

## Methods

### NewDtcLbdn

`func NewDtcLbdn() *DtcLbdn`

NewDtcLbdn instantiates a new DtcLbdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcLbdnWithDefaults

`func NewDtcLbdnWithDefaults() *DtcLbdn`

NewDtcLbdnWithDefaults instantiates a new DtcLbdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcLbdn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcLbdn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcLbdn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcLbdn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthZones

`func (o *DtcLbdn) GetAuthZones() []string`

GetAuthZones returns the AuthZones field if non-nil, zero value otherwise.

### GetAuthZonesOk

`func (o *DtcLbdn) GetAuthZonesOk() (*[]string, bool)`

GetAuthZonesOk returns a tuple with the AuthZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthZones

`func (o *DtcLbdn) SetAuthZones(v []string)`

SetAuthZones sets AuthZones field to given value.

### HasAuthZones

`func (o *DtcLbdn) HasAuthZones() bool`

HasAuthZones returns a boolean if a field has been set.

### GetAutoConsolidatedMonitors

`func (o *DtcLbdn) GetAutoConsolidatedMonitors() bool`

GetAutoConsolidatedMonitors returns the AutoConsolidatedMonitors field if non-nil, zero value otherwise.

### GetAutoConsolidatedMonitorsOk

`func (o *DtcLbdn) GetAutoConsolidatedMonitorsOk() (*bool, bool)`

GetAutoConsolidatedMonitorsOk returns a tuple with the AutoConsolidatedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidatedMonitors

`func (o *DtcLbdn) SetAutoConsolidatedMonitors(v bool)`

SetAutoConsolidatedMonitors sets AutoConsolidatedMonitors field to given value.

### HasAutoConsolidatedMonitors

`func (o *DtcLbdn) HasAutoConsolidatedMonitors() bool`

HasAutoConsolidatedMonitors returns a boolean if a field has been set.

### GetComment

`func (o *DtcLbdn) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcLbdn) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcLbdn) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcLbdn) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcLbdn) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcLbdn) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcLbdn) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcLbdn) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *DtcLbdn) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *DtcLbdn) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *DtcLbdn) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *DtcLbdn) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *DtcLbdn) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *DtcLbdn) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *DtcLbdn) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *DtcLbdn) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DtcLbdn) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DtcLbdn) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DtcLbdn) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DtcLbdn) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHealth

`func (o *DtcLbdn) GetHealth() DtcLbdnHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DtcLbdn) GetHealthOk() (*DtcLbdnHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DtcLbdn) SetHealth(v DtcLbdnHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DtcLbdn) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLbMethod

`func (o *DtcLbdn) GetLbMethod() string`

GetLbMethod returns the LbMethod field if non-nil, zero value otherwise.

### GetLbMethodOk

`func (o *DtcLbdn) GetLbMethodOk() (*string, bool)`

GetLbMethodOk returns a tuple with the LbMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMethod

`func (o *DtcLbdn) SetLbMethod(v string)`

SetLbMethod sets LbMethod field to given value.

### HasLbMethod

`func (o *DtcLbdn) HasLbMethod() bool`

HasLbMethod returns a boolean if a field has been set.

### GetName

`func (o *DtcLbdn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcLbdn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcLbdn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcLbdn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *DtcLbdn) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *DtcLbdn) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *DtcLbdn) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *DtcLbdn) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetPersistence

`func (o *DtcLbdn) GetPersistence() int64`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *DtcLbdn) GetPersistenceOk() (*int64, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *DtcLbdn) SetPersistence(v int64)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *DtcLbdn) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### GetPools

`func (o *DtcLbdn) GetPools() []DtcLbdnPools`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *DtcLbdn) GetPoolsOk() (*[]DtcLbdnPools, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *DtcLbdn) SetPools(v []DtcLbdnPools)`

SetPools sets Pools field to given value.

### HasPools

`func (o *DtcLbdn) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetPriority

`func (o *DtcLbdn) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtcLbdn) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtcLbdn) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtcLbdn) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTopology

`func (o *DtcLbdn) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *DtcLbdn) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *DtcLbdn) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *DtcLbdn) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetTtl

`func (o *DtcLbdn) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcLbdn) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcLbdn) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcLbdn) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTypes

`func (o *DtcLbdn) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DtcLbdn) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DtcLbdn) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DtcLbdn) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUseTtl

`func (o *DtcLbdn) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *DtcLbdn) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *DtcLbdn) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *DtcLbdn) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


