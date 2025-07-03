# ZoneStubStubMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The grid member name. | [optional] 
**Stealth** | Pointer to **bool** | This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone. | [optional] 
**GridReplicate** | Pointer to **bool** | The flag represents DNS zone transfers if set to False, and ID Grid Replication if set to True. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone. | [optional] 
**Lead** | Pointer to **bool** | This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone. | [optional] 
**PreferredPrimaries** | Pointer to [**[]ZonestubstubmembersPreferredPrimaries**](ZonestubstubmembersPreferredPrimaries.md) | The primary preference list with Grid member names and\\or External Server extserver structs for this member. | [optional] 
**EnablePreferredPrimaries** | Pointer to **bool** | This flag represents whether the preferred_primaries field values of this member are used. | [optional] 

## Methods

### NewZoneStubStubMembers

`func NewZoneStubStubMembers() *ZoneStubStubMembers`

NewZoneStubStubMembers instantiates a new ZoneStubStubMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneStubStubMembersWithDefaults

`func NewZoneStubStubMembersWithDefaults() *ZoneStubStubMembers`

NewZoneStubStubMembersWithDefaults instantiates a new ZoneStubStubMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZoneStubStubMembers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneStubStubMembers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneStubStubMembers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneStubStubMembers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStealth

`func (o *ZoneStubStubMembers) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *ZoneStubStubMembers) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *ZoneStubStubMembers) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *ZoneStubStubMembers) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetGridReplicate

`func (o *ZoneStubStubMembers) GetGridReplicate() bool`

GetGridReplicate returns the GridReplicate field if non-nil, zero value otherwise.

### GetGridReplicateOk

`func (o *ZoneStubStubMembers) GetGridReplicateOk() (*bool, bool)`

GetGridReplicateOk returns a tuple with the GridReplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridReplicate

`func (o *ZoneStubStubMembers) SetGridReplicate(v bool)`

SetGridReplicate sets GridReplicate field to given value.

### HasGridReplicate

`func (o *ZoneStubStubMembers) HasGridReplicate() bool`

HasGridReplicate returns a boolean if a field has been set.

### GetLead

`func (o *ZoneStubStubMembers) GetLead() bool`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ZoneStubStubMembers) GetLeadOk() (*bool, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ZoneStubStubMembers) SetLead(v bool)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ZoneStubStubMembers) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetPreferredPrimaries

`func (o *ZoneStubStubMembers) GetPreferredPrimaries() []ZonestubstubmembersPreferredPrimaries`

GetPreferredPrimaries returns the PreferredPrimaries field if non-nil, zero value otherwise.

### GetPreferredPrimariesOk

`func (o *ZoneStubStubMembers) GetPreferredPrimariesOk() (*[]ZonestubstubmembersPreferredPrimaries, bool)`

GetPreferredPrimariesOk returns a tuple with the PreferredPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredPrimaries

`func (o *ZoneStubStubMembers) SetPreferredPrimaries(v []ZonestubstubmembersPreferredPrimaries)`

SetPreferredPrimaries sets PreferredPrimaries field to given value.

### HasPreferredPrimaries

`func (o *ZoneStubStubMembers) HasPreferredPrimaries() bool`

HasPreferredPrimaries returns a boolean if a field has been set.

### GetEnablePreferredPrimaries

`func (o *ZoneStubStubMembers) GetEnablePreferredPrimaries() bool`

GetEnablePreferredPrimaries returns the EnablePreferredPrimaries field if non-nil, zero value otherwise.

### GetEnablePreferredPrimariesOk

`func (o *ZoneStubStubMembers) GetEnablePreferredPrimariesOk() (*bool, bool)`

GetEnablePreferredPrimariesOk returns a tuple with the EnablePreferredPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreferredPrimaries

`func (o *ZoneStubStubMembers) SetEnablePreferredPrimaries(v bool)`

SetEnablePreferredPrimaries sets EnablePreferredPrimaries field to given value.

### HasEnablePreferredPrimaries

`func (o *ZoneStubStubMembers) HasEnablePreferredPrimaries() bool`

HasEnablePreferredPrimaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


