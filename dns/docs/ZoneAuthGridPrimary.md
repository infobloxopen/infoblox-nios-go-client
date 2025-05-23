# ZoneAuthGridPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The grid member name. | [optional] 
**Stealth** | Pointer to **bool** | This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone. | [optional] 
**GridReplicate** | Pointer to **bool** | The flag represents DNS zone transfers if set to False, and ID Grid Replication if set to True. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone. | [optional] 
**Lead** | Pointer to **bool** | This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone. | [optional] 
**PreferredPrimaries** | Pointer to [**ZoneauthgridprimaryPreferredPrimaries**](ZoneauthgridprimaryPreferredPrimaries.md) |  | [optional] 
**EnablePreferredPrimaries** | Pointer to **bool** | This flag represents whether the preferred_primaries field values of this member are used. | [optional] 

## Methods

### NewZoneAuthGridPrimary

`func NewZoneAuthGridPrimary() *ZoneAuthGridPrimary`

NewZoneAuthGridPrimary instantiates a new ZoneAuthGridPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthGridPrimaryWithDefaults

`func NewZoneAuthGridPrimaryWithDefaults() *ZoneAuthGridPrimary`

NewZoneAuthGridPrimaryWithDefaults instantiates a new ZoneAuthGridPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZoneAuthGridPrimary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneAuthGridPrimary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneAuthGridPrimary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneAuthGridPrimary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStealth

`func (o *ZoneAuthGridPrimary) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *ZoneAuthGridPrimary) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *ZoneAuthGridPrimary) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *ZoneAuthGridPrimary) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetGridReplicate

`func (o *ZoneAuthGridPrimary) GetGridReplicate() bool`

GetGridReplicate returns the GridReplicate field if non-nil, zero value otherwise.

### GetGridReplicateOk

`func (o *ZoneAuthGridPrimary) GetGridReplicateOk() (*bool, bool)`

GetGridReplicateOk returns a tuple with the GridReplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridReplicate

`func (o *ZoneAuthGridPrimary) SetGridReplicate(v bool)`

SetGridReplicate sets GridReplicate field to given value.

### HasGridReplicate

`func (o *ZoneAuthGridPrimary) HasGridReplicate() bool`

HasGridReplicate returns a boolean if a field has been set.

### GetLead

`func (o *ZoneAuthGridPrimary) GetLead() bool`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ZoneAuthGridPrimary) GetLeadOk() (*bool, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ZoneAuthGridPrimary) SetLead(v bool)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ZoneAuthGridPrimary) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetPreferredPrimaries

`func (o *ZoneAuthGridPrimary) GetPreferredPrimaries() ZoneauthgridprimaryPreferredPrimaries`

GetPreferredPrimaries returns the PreferredPrimaries field if non-nil, zero value otherwise.

### GetPreferredPrimariesOk

`func (o *ZoneAuthGridPrimary) GetPreferredPrimariesOk() (*ZoneauthgridprimaryPreferredPrimaries, bool)`

GetPreferredPrimariesOk returns a tuple with the PreferredPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredPrimaries

`func (o *ZoneAuthGridPrimary) SetPreferredPrimaries(v ZoneauthgridprimaryPreferredPrimaries)`

SetPreferredPrimaries sets PreferredPrimaries field to given value.

### HasPreferredPrimaries

`func (o *ZoneAuthGridPrimary) HasPreferredPrimaries() bool`

HasPreferredPrimaries returns a boolean if a field has been set.

### GetEnablePreferredPrimaries

`func (o *ZoneAuthGridPrimary) GetEnablePreferredPrimaries() bool`

GetEnablePreferredPrimaries returns the EnablePreferredPrimaries field if non-nil, zero value otherwise.

### GetEnablePreferredPrimariesOk

`func (o *ZoneAuthGridPrimary) GetEnablePreferredPrimariesOk() (*bool, bool)`

GetEnablePreferredPrimariesOk returns a tuple with the EnablePreferredPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreferredPrimaries

`func (o *ZoneAuthGridPrimary) SetEnablePreferredPrimaries(v bool)`

SetEnablePreferredPrimaries sets EnablePreferredPrimaries field to given value.

### HasEnablePreferredPrimaries

`func (o *ZoneAuthGridPrimary) HasEnablePreferredPrimaries() bool`

HasEnablePreferredPrimaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


