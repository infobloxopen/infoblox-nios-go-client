# GetDtcLbdnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AuthZones** | Pointer to **[]map[string]interface{}** | List of linked auth zones. | [optional] 
**AutoConsolidatedMonitors** | Pointer to **bool** | Flag for enabling auto managing DTC Consolidated Monitors on related DTC Pools. | [optional] 
**Comment** | Pointer to **string** | Comment for the DTC LBDN; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the DTC LBDN is disabled or not. When this is set to False, the fixed address is enabled. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**DtcLbdn**](DtcLbdn.md) |  | [optional] 

## Methods

### NewGetDtcLbdnResponse

`func NewGetDtcLbdnResponse() *GetDtcLbdnResponse`

NewGetDtcLbdnResponse instantiates a new GetDtcLbdnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcLbdnResponseWithDefaults

`func NewGetDtcLbdnResponseWithDefaults() *GetDtcLbdnResponse`

NewGetDtcLbdnResponseWithDefaults instantiates a new GetDtcLbdnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcLbdnResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcLbdnResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcLbdnResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcLbdnResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthZones

`func (o *GetDtcLbdnResponse) GetAuthZones() []map[string]interface{}`

GetAuthZones returns the AuthZones field if non-nil, zero value otherwise.

### GetAuthZonesOk

`func (o *GetDtcLbdnResponse) GetAuthZonesOk() (*[]map[string]interface{}, bool)`

GetAuthZonesOk returns a tuple with the AuthZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthZones

`func (o *GetDtcLbdnResponse) SetAuthZones(v []map[string]interface{})`

SetAuthZones sets AuthZones field to given value.

### HasAuthZones

`func (o *GetDtcLbdnResponse) HasAuthZones() bool`

HasAuthZones returns a boolean if a field has been set.

### GetAutoConsolidatedMonitors

`func (o *GetDtcLbdnResponse) GetAutoConsolidatedMonitors() bool`

GetAutoConsolidatedMonitors returns the AutoConsolidatedMonitors field if non-nil, zero value otherwise.

### GetAutoConsolidatedMonitorsOk

`func (o *GetDtcLbdnResponse) GetAutoConsolidatedMonitorsOk() (*bool, bool)`

GetAutoConsolidatedMonitorsOk returns a tuple with the AutoConsolidatedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidatedMonitors

`func (o *GetDtcLbdnResponse) SetAutoConsolidatedMonitors(v bool)`

SetAutoConsolidatedMonitors sets AutoConsolidatedMonitors field to given value.

### HasAutoConsolidatedMonitors

`func (o *GetDtcLbdnResponse) HasAutoConsolidatedMonitors() bool`

HasAutoConsolidatedMonitors returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcLbdnResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcLbdnResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcLbdnResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcLbdnResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDtcLbdnResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDtcLbdnResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDtcLbdnResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDtcLbdnResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetDtcLbdnResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetDtcLbdnResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetDtcLbdnResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetDtcLbdnResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetHealth

`func (o *GetDtcLbdnResponse) GetHealth() DtcLbdnHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GetDtcLbdnResponse) GetHealthOk() (*DtcLbdnHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GetDtcLbdnResponse) SetHealth(v DtcLbdnHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GetDtcLbdnResponse) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetLbMethod

`func (o *GetDtcLbdnResponse) GetLbMethod() string`

GetLbMethod returns the LbMethod field if non-nil, zero value otherwise.

### GetLbMethodOk

`func (o *GetDtcLbdnResponse) GetLbMethodOk() (*string, bool)`

GetLbMethodOk returns a tuple with the LbMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMethod

`func (o *GetDtcLbdnResponse) SetLbMethod(v string)`

SetLbMethod sets LbMethod field to given value.

### HasLbMethod

`func (o *GetDtcLbdnResponse) HasLbMethod() bool`

HasLbMethod returns a boolean if a field has been set.

### GetName

`func (o *GetDtcLbdnResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcLbdnResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcLbdnResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcLbdnResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *GetDtcLbdnResponse) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *GetDtcLbdnResponse) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *GetDtcLbdnResponse) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *GetDtcLbdnResponse) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetPersistence

`func (o *GetDtcLbdnResponse) GetPersistence() int64`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *GetDtcLbdnResponse) GetPersistenceOk() (*int64, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *GetDtcLbdnResponse) SetPersistence(v int64)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *GetDtcLbdnResponse) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### GetPools

`func (o *GetDtcLbdnResponse) GetPools() []DtcLbdnPools`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *GetDtcLbdnResponse) GetPoolsOk() (*[]DtcLbdnPools, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *GetDtcLbdnResponse) SetPools(v []DtcLbdnPools)`

SetPools sets Pools field to given value.

### HasPools

`func (o *GetDtcLbdnResponse) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetPriority

`func (o *GetDtcLbdnResponse) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetDtcLbdnResponse) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetDtcLbdnResponse) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetDtcLbdnResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTopology

`func (o *GetDtcLbdnResponse) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *GetDtcLbdnResponse) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *GetDtcLbdnResponse) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *GetDtcLbdnResponse) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetTtl

`func (o *GetDtcLbdnResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetDtcLbdnResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetDtcLbdnResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetDtcLbdnResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTypes

`func (o *GetDtcLbdnResponse) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GetDtcLbdnResponse) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GetDtcLbdnResponse) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *GetDtcLbdnResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetDtcLbdnResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetDtcLbdnResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetDtcLbdnResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetDtcLbdnResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcLbdnResponse) GetResult() DtcLbdn`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcLbdnResponse) GetResultOk() (*DtcLbdn, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcLbdnResponse) SetResult(v DtcLbdn)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcLbdnResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


