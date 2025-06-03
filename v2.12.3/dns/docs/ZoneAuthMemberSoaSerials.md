# ZoneAuthMemberSoaSerials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GridPrimary** | Pointer to **string** | The grid primary server for the zone. Only one of \&quot;grid_primary\&quot; or \&quot;ms_server_primary\&quot; will be set when the object is retrieved from the server. | [optional] [readonly] 
**MsServerPrimary** | Pointer to **string** | The primary MS server for the zone. Only one of \&quot;grid_primary\&quot; or \&quot;ms_server_primary\&quot; will be set when the object is retrieved from the server. | [optional] [readonly] 
**Serial** | Pointer to **int64** | The SOA serial number. | [optional] [readonly] 

## Methods

### NewZoneAuthMemberSoaSerials

`func NewZoneAuthMemberSoaSerials() *ZoneAuthMemberSoaSerials`

NewZoneAuthMemberSoaSerials instantiates a new ZoneAuthMemberSoaSerials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthMemberSoaSerialsWithDefaults

`func NewZoneAuthMemberSoaSerialsWithDefaults() *ZoneAuthMemberSoaSerials`

NewZoneAuthMemberSoaSerialsWithDefaults instantiates a new ZoneAuthMemberSoaSerials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGridPrimary

`func (o *ZoneAuthMemberSoaSerials) GetGridPrimary() string`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *ZoneAuthMemberSoaSerials) GetGridPrimaryOk() (*string, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *ZoneAuthMemberSoaSerials) SetGridPrimary(v string)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *ZoneAuthMemberSoaSerials) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetMsServerPrimary

`func (o *ZoneAuthMemberSoaSerials) GetMsServerPrimary() string`

GetMsServerPrimary returns the MsServerPrimary field if non-nil, zero value otherwise.

### GetMsServerPrimaryOk

`func (o *ZoneAuthMemberSoaSerials) GetMsServerPrimaryOk() (*string, bool)`

GetMsServerPrimaryOk returns a tuple with the MsServerPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServerPrimary

`func (o *ZoneAuthMemberSoaSerials) SetMsServerPrimary(v string)`

SetMsServerPrimary sets MsServerPrimary field to given value.

### HasMsServerPrimary

`func (o *ZoneAuthMemberSoaSerials) HasMsServerPrimary() bool`

HasMsServerPrimary returns a boolean if a field has been set.

### GetSerial

`func (o *ZoneAuthMemberSoaSerials) GetSerial() int64`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ZoneAuthMemberSoaSerials) GetSerialOk() (*int64, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ZoneAuthMemberSoaSerials) SetSerial(v int64)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ZoneAuthMemberSoaSerials) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


