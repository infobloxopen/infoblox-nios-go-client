# ZoneAuthMemberSoaMnames

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GridPrimary** | Pointer to **string** | The grid primary server for the zone. Only one of \&quot;grid_primary\&quot; or \&quot;ms_server_primary\&quot; should be set when modifying or creating the object. | [optional] 
**MsServerPrimary** | Pointer to **string** | The primary MS server for the zone. Only one of \&quot;grid_primary\&quot; or \&quot;ms_server_primary\&quot; should be set when modifying or creating the object. | [optional] 
**Mname** | Pointer to **string** | Master&#39;s SOA MNAME. This value can be in unicode format. | [optional] 
**DnsMname** | Pointer to **string** | Master&#39;s SOA MNAME in punycode format. | [optional] [readonly] 

## Methods

### NewZoneAuthMemberSoaMnames

`func NewZoneAuthMemberSoaMnames() *ZoneAuthMemberSoaMnames`

NewZoneAuthMemberSoaMnames instantiates a new ZoneAuthMemberSoaMnames object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthMemberSoaMnamesWithDefaults

`func NewZoneAuthMemberSoaMnamesWithDefaults() *ZoneAuthMemberSoaMnames`

NewZoneAuthMemberSoaMnamesWithDefaults instantiates a new ZoneAuthMemberSoaMnames object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGridPrimary

`func (o *ZoneAuthMemberSoaMnames) GetGridPrimary() string`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *ZoneAuthMemberSoaMnames) GetGridPrimaryOk() (*string, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *ZoneAuthMemberSoaMnames) SetGridPrimary(v string)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *ZoneAuthMemberSoaMnames) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetMsServerPrimary

`func (o *ZoneAuthMemberSoaMnames) GetMsServerPrimary() string`

GetMsServerPrimary returns the MsServerPrimary field if non-nil, zero value otherwise.

### GetMsServerPrimaryOk

`func (o *ZoneAuthMemberSoaMnames) GetMsServerPrimaryOk() (*string, bool)`

GetMsServerPrimaryOk returns a tuple with the MsServerPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServerPrimary

`func (o *ZoneAuthMemberSoaMnames) SetMsServerPrimary(v string)`

SetMsServerPrimary sets MsServerPrimary field to given value.

### HasMsServerPrimary

`func (o *ZoneAuthMemberSoaMnames) HasMsServerPrimary() bool`

HasMsServerPrimary returns a boolean if a field has been set.

### GetMname

`func (o *ZoneAuthMemberSoaMnames) GetMname() string`

GetMname returns the Mname field if non-nil, zero value otherwise.

### GetMnameOk

`func (o *ZoneAuthMemberSoaMnames) GetMnameOk() (*string, bool)`

GetMnameOk returns a tuple with the Mname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMname

`func (o *ZoneAuthMemberSoaMnames) SetMname(v string)`

SetMname sets Mname field to given value.

### HasMname

`func (o *ZoneAuthMemberSoaMnames) HasMname() bool`

HasMname returns a boolean if a field has been set.

### GetDnsMname

`func (o *ZoneAuthMemberSoaMnames) GetDnsMname() string`

GetDnsMname returns the DnsMname field if non-nil, zero value otherwise.

### GetDnsMnameOk

`func (o *ZoneAuthMemberSoaMnames) GetDnsMnameOk() (*string, bool)`

GetDnsMnameOk returns a tuple with the DnsMname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMname

`func (o *ZoneAuthMemberSoaMnames) SetDnsMname(v string)`

SetDnsMname sets DnsMname field to given value.

### HasDnsMname

`func (o *ZoneAuthMemberSoaMnames) HasDnsMname() bool`

HasDnsMname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


