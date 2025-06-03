# ZoneRpMemberSoaSerials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GridPrimary** | Pointer to **string** | The grid primary server for the zone. Only one of \&quot;grid_primary\&quot; or \&quot;ms_server_primary\&quot; will be set when the object is retrieved from the server. | [optional] [readonly] 
**MsServerPrimary** | Pointer to **string** | The primary MS server for the zone. Only one of \&quot;grid_primary\&quot; or \&quot;ms_server_primary\&quot; will be set when the object is retrieved from the server. | [optional] [readonly] 
**Serial** | Pointer to **int64** | The SOA serial number. | [optional] [readonly] 

## Methods

### NewZoneRpMemberSoaSerials

`func NewZoneRpMemberSoaSerials() *ZoneRpMemberSoaSerials`

NewZoneRpMemberSoaSerials instantiates a new ZoneRpMemberSoaSerials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneRpMemberSoaSerialsWithDefaults

`func NewZoneRpMemberSoaSerialsWithDefaults() *ZoneRpMemberSoaSerials`

NewZoneRpMemberSoaSerialsWithDefaults instantiates a new ZoneRpMemberSoaSerials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGridPrimary

`func (o *ZoneRpMemberSoaSerials) GetGridPrimary() string`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *ZoneRpMemberSoaSerials) GetGridPrimaryOk() (*string, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *ZoneRpMemberSoaSerials) SetGridPrimary(v string)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *ZoneRpMemberSoaSerials) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetMsServerPrimary

`func (o *ZoneRpMemberSoaSerials) GetMsServerPrimary() string`

GetMsServerPrimary returns the MsServerPrimary field if non-nil, zero value otherwise.

### GetMsServerPrimaryOk

`func (o *ZoneRpMemberSoaSerials) GetMsServerPrimaryOk() (*string, bool)`

GetMsServerPrimaryOk returns a tuple with the MsServerPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServerPrimary

`func (o *ZoneRpMemberSoaSerials) SetMsServerPrimary(v string)`

SetMsServerPrimary sets MsServerPrimary field to given value.

### HasMsServerPrimary

`func (o *ZoneRpMemberSoaSerials) HasMsServerPrimary() bool`

HasMsServerPrimary returns a boolean if a field has been set.

### GetSerial

`func (o *ZoneRpMemberSoaSerials) GetSerial() int64`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ZoneRpMemberSoaSerials) GetSerialOk() (*int64, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ZoneRpMemberSoaSerials) SetSerial(v int64)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ZoneRpMemberSoaSerials) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


