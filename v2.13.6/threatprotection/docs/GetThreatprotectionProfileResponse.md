# GetThreatprotectionProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for the Threat Protection profile. | [optional] 
**CurrentRuleset** | Pointer to **string** | The current Threat Protection profile ruleset. | [optional] 
**DisableMultipleDnsTcpRequest** | Pointer to **bool** | Determines if multiple BIND responses via TCP connection are disabled. | [optional] 
**EventsPerSecondPerRule** | Pointer to **int64** | The number of events logged per second per rule. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Members** | Pointer to **[]string** | The list of members that are associated with the profile. | [optional] 
**Name** | Pointer to **string** | The name of the Threat Protection profile. | [optional] 
**SourceMember** | Pointer to **string** | The source member. It can be used only during the create operation for cloning a profile from an existing member. | [optional] 
**SourceProfile** | Pointer to **string** | The source profile. It can be used only during the create operation for cloning a profile from an existing profile. | [optional] 
**UseCurrentRuleset** | Pointer to **bool** | Use flag for: current_ruleset | [optional] 
**UseDisableMultipleDnsTcpRequest** | Pointer to **bool** | Use flag for: disable_multiple_dns_tcp_request | [optional] 
**UseEventsPerSecondPerRule** | Pointer to **bool** | Use flag for: events_per_second_per_rule | [optional] 
**Result** | Pointer to [**ThreatprotectionProfile**](ThreatprotectionProfile.md) |  | [optional] 

## Methods

### NewGetThreatprotectionProfileResponse

`func NewGetThreatprotectionProfileResponse() *GetThreatprotectionProfileResponse`

NewGetThreatprotectionProfileResponse instantiates a new GetThreatprotectionProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionProfileResponseWithDefaults

`func NewGetThreatprotectionProfileResponseWithDefaults() *GetThreatprotectionProfileResponse`

NewGetThreatprotectionProfileResponseWithDefaults instantiates a new GetThreatprotectionProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionProfileResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionProfileResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionProfileResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionProfileResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetThreatprotectionProfileResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetThreatprotectionProfileResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetThreatprotectionProfileResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetThreatprotectionProfileResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentRuleset

`func (o *GetThreatprotectionProfileResponse) GetCurrentRuleset() string`

GetCurrentRuleset returns the CurrentRuleset field if non-nil, zero value otherwise.

### GetCurrentRulesetOk

`func (o *GetThreatprotectionProfileResponse) GetCurrentRulesetOk() (*string, bool)`

GetCurrentRulesetOk returns a tuple with the CurrentRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRuleset

`func (o *GetThreatprotectionProfileResponse) SetCurrentRuleset(v string)`

SetCurrentRuleset sets CurrentRuleset field to given value.

### HasCurrentRuleset

`func (o *GetThreatprotectionProfileResponse) HasCurrentRuleset() bool`

HasCurrentRuleset returns a boolean if a field has been set.

### GetDisableMultipleDnsTcpRequest

`func (o *GetThreatprotectionProfileResponse) GetDisableMultipleDnsTcpRequest() bool`

GetDisableMultipleDnsTcpRequest returns the DisableMultipleDnsTcpRequest field if non-nil, zero value otherwise.

### GetDisableMultipleDnsTcpRequestOk

`func (o *GetThreatprotectionProfileResponse) GetDisableMultipleDnsTcpRequestOk() (*bool, bool)`

GetDisableMultipleDnsTcpRequestOk returns a tuple with the DisableMultipleDnsTcpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMultipleDnsTcpRequest

`func (o *GetThreatprotectionProfileResponse) SetDisableMultipleDnsTcpRequest(v bool)`

SetDisableMultipleDnsTcpRequest sets DisableMultipleDnsTcpRequest field to given value.

### HasDisableMultipleDnsTcpRequest

`func (o *GetThreatprotectionProfileResponse) HasDisableMultipleDnsTcpRequest() bool`

HasDisableMultipleDnsTcpRequest returns a boolean if a field has been set.

### GetEventsPerSecondPerRule

`func (o *GetThreatprotectionProfileResponse) GetEventsPerSecondPerRule() int64`

GetEventsPerSecondPerRule returns the EventsPerSecondPerRule field if non-nil, zero value otherwise.

### GetEventsPerSecondPerRuleOk

`func (o *GetThreatprotectionProfileResponse) GetEventsPerSecondPerRuleOk() (*int64, bool)`

GetEventsPerSecondPerRuleOk returns a tuple with the EventsPerSecondPerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsPerSecondPerRule

`func (o *GetThreatprotectionProfileResponse) SetEventsPerSecondPerRule(v int64)`

SetEventsPerSecondPerRule sets EventsPerSecondPerRule field to given value.

### HasEventsPerSecondPerRule

`func (o *GetThreatprotectionProfileResponse) HasEventsPerSecondPerRule() bool`

HasEventsPerSecondPerRule returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetThreatprotectionProfileResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetThreatprotectionProfileResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetThreatprotectionProfileResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetThreatprotectionProfileResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetMembers

`func (o *GetThreatprotectionProfileResponse) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetThreatprotectionProfileResponse) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetThreatprotectionProfileResponse) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetThreatprotectionProfileResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *GetThreatprotectionProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetThreatprotectionProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetThreatprotectionProfileResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetThreatprotectionProfileResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceMember

`func (o *GetThreatprotectionProfileResponse) GetSourceMember() string`

GetSourceMember returns the SourceMember field if non-nil, zero value otherwise.

### GetSourceMemberOk

`func (o *GetThreatprotectionProfileResponse) GetSourceMemberOk() (*string, bool)`

GetSourceMemberOk returns a tuple with the SourceMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMember

`func (o *GetThreatprotectionProfileResponse) SetSourceMember(v string)`

SetSourceMember sets SourceMember field to given value.

### HasSourceMember

`func (o *GetThreatprotectionProfileResponse) HasSourceMember() bool`

HasSourceMember returns a boolean if a field has been set.

### GetSourceProfile

`func (o *GetThreatprotectionProfileResponse) GetSourceProfile() string`

GetSourceProfile returns the SourceProfile field if non-nil, zero value otherwise.

### GetSourceProfileOk

`func (o *GetThreatprotectionProfileResponse) GetSourceProfileOk() (*string, bool)`

GetSourceProfileOk returns a tuple with the SourceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProfile

`func (o *GetThreatprotectionProfileResponse) SetSourceProfile(v string)`

SetSourceProfile sets SourceProfile field to given value.

### HasSourceProfile

`func (o *GetThreatprotectionProfileResponse) HasSourceProfile() bool`

HasSourceProfile returns a boolean if a field has been set.

### GetUseCurrentRuleset

`func (o *GetThreatprotectionProfileResponse) GetUseCurrentRuleset() bool`

GetUseCurrentRuleset returns the UseCurrentRuleset field if non-nil, zero value otherwise.

### GetUseCurrentRulesetOk

`func (o *GetThreatprotectionProfileResponse) GetUseCurrentRulesetOk() (*bool, bool)`

GetUseCurrentRulesetOk returns a tuple with the UseCurrentRuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCurrentRuleset

`func (o *GetThreatprotectionProfileResponse) SetUseCurrentRuleset(v bool)`

SetUseCurrentRuleset sets UseCurrentRuleset field to given value.

### HasUseCurrentRuleset

`func (o *GetThreatprotectionProfileResponse) HasUseCurrentRuleset() bool`

HasUseCurrentRuleset returns a boolean if a field has been set.

### GetUseDisableMultipleDnsTcpRequest

`func (o *GetThreatprotectionProfileResponse) GetUseDisableMultipleDnsTcpRequest() bool`

GetUseDisableMultipleDnsTcpRequest returns the UseDisableMultipleDnsTcpRequest field if non-nil, zero value otherwise.

### GetUseDisableMultipleDnsTcpRequestOk

`func (o *GetThreatprotectionProfileResponse) GetUseDisableMultipleDnsTcpRequestOk() (*bool, bool)`

GetUseDisableMultipleDnsTcpRequestOk returns a tuple with the UseDisableMultipleDnsTcpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisableMultipleDnsTcpRequest

`func (o *GetThreatprotectionProfileResponse) SetUseDisableMultipleDnsTcpRequest(v bool)`

SetUseDisableMultipleDnsTcpRequest sets UseDisableMultipleDnsTcpRequest field to given value.

### HasUseDisableMultipleDnsTcpRequest

`func (o *GetThreatprotectionProfileResponse) HasUseDisableMultipleDnsTcpRequest() bool`

HasUseDisableMultipleDnsTcpRequest returns a boolean if a field has been set.

### GetUseEventsPerSecondPerRule

`func (o *GetThreatprotectionProfileResponse) GetUseEventsPerSecondPerRule() bool`

GetUseEventsPerSecondPerRule returns the UseEventsPerSecondPerRule field if non-nil, zero value otherwise.

### GetUseEventsPerSecondPerRuleOk

`func (o *GetThreatprotectionProfileResponse) GetUseEventsPerSecondPerRuleOk() (*bool, bool)`

GetUseEventsPerSecondPerRuleOk returns a tuple with the UseEventsPerSecondPerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEventsPerSecondPerRule

`func (o *GetThreatprotectionProfileResponse) SetUseEventsPerSecondPerRule(v bool)`

SetUseEventsPerSecondPerRule sets UseEventsPerSecondPerRule field to given value.

### HasUseEventsPerSecondPerRule

`func (o *GetThreatprotectionProfileResponse) HasUseEventsPerSecondPerRule() bool`

HasUseEventsPerSecondPerRule returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionProfileResponse) GetResult() ThreatprotectionProfile`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionProfileResponse) GetResultOk() (*ThreatprotectionProfile, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionProfileResponse) SetResult(v ThreatprotectionProfile)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionProfileResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


