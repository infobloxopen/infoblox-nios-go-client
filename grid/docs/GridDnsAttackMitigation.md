# GridDnsAttackMitigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectChr** | Pointer to [**GriddnsattackmitigationDetectChr**](GriddnsattackmitigationDetectChr.md) |  | [optional] 
**DetectChrGrace** | Pointer to **int64** | The cache utilization (in percentage) when Cache Hit Ratio (CHR) starts. | [optional] 
**DetectNxdomainResponses** | Pointer to [**GriddnsattackmitigationDetectNxdomainResponses**](GriddnsattackmitigationDetectNxdomainResponses.md) |  | [optional] 
**DetectUdpDrop** | Pointer to [**GriddnsattackmitigationDetectUdpDrop**](GriddnsattackmitigationDetectUdpDrop.md) |  | [optional] 
**Interval** | Pointer to **int64** | The minimum time interval (in seconds) between changes in attack status. | [optional] 

## Methods

### NewGridDnsAttackMitigation

`func NewGridDnsAttackMitigation() *GridDnsAttackMitigation`

NewGridDnsAttackMitigation instantiates a new GridDnsAttackMitigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsAttackMitigationWithDefaults

`func NewGridDnsAttackMitigationWithDefaults() *GridDnsAttackMitigation`

NewGridDnsAttackMitigationWithDefaults instantiates a new GridDnsAttackMitigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectChr

`func (o *GridDnsAttackMitigation) GetDetectChr() GriddnsattackmitigationDetectChr`

GetDetectChr returns the DetectChr field if non-nil, zero value otherwise.

### GetDetectChrOk

`func (o *GridDnsAttackMitigation) GetDetectChrOk() (*GriddnsattackmitigationDetectChr, bool)`

GetDetectChrOk returns a tuple with the DetectChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectChr

`func (o *GridDnsAttackMitigation) SetDetectChr(v GriddnsattackmitigationDetectChr)`

SetDetectChr sets DetectChr field to given value.

### HasDetectChr

`func (o *GridDnsAttackMitigation) HasDetectChr() bool`

HasDetectChr returns a boolean if a field has been set.

### GetDetectChrGrace

`func (o *GridDnsAttackMitigation) GetDetectChrGrace() int64`

GetDetectChrGrace returns the DetectChrGrace field if non-nil, zero value otherwise.

### GetDetectChrGraceOk

`func (o *GridDnsAttackMitigation) GetDetectChrGraceOk() (*int64, bool)`

GetDetectChrGraceOk returns a tuple with the DetectChrGrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectChrGrace

`func (o *GridDnsAttackMitigation) SetDetectChrGrace(v int64)`

SetDetectChrGrace sets DetectChrGrace field to given value.

### HasDetectChrGrace

`func (o *GridDnsAttackMitigation) HasDetectChrGrace() bool`

HasDetectChrGrace returns a boolean if a field has been set.

### GetDetectNxdomainResponses

`func (o *GridDnsAttackMitigation) GetDetectNxdomainResponses() GriddnsattackmitigationDetectNxdomainResponses`

GetDetectNxdomainResponses returns the DetectNxdomainResponses field if non-nil, zero value otherwise.

### GetDetectNxdomainResponsesOk

`func (o *GridDnsAttackMitigation) GetDetectNxdomainResponsesOk() (*GriddnsattackmitigationDetectNxdomainResponses, bool)`

GetDetectNxdomainResponsesOk returns a tuple with the DetectNxdomainResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectNxdomainResponses

`func (o *GridDnsAttackMitigation) SetDetectNxdomainResponses(v GriddnsattackmitigationDetectNxdomainResponses)`

SetDetectNxdomainResponses sets DetectNxdomainResponses field to given value.

### HasDetectNxdomainResponses

`func (o *GridDnsAttackMitigation) HasDetectNxdomainResponses() bool`

HasDetectNxdomainResponses returns a boolean if a field has been set.

### GetDetectUdpDrop

`func (o *GridDnsAttackMitigation) GetDetectUdpDrop() GriddnsattackmitigationDetectUdpDrop`

GetDetectUdpDrop returns the DetectUdpDrop field if non-nil, zero value otherwise.

### GetDetectUdpDropOk

`func (o *GridDnsAttackMitigation) GetDetectUdpDropOk() (*GriddnsattackmitigationDetectUdpDrop, bool)`

GetDetectUdpDropOk returns a tuple with the DetectUdpDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectUdpDrop

`func (o *GridDnsAttackMitigation) SetDetectUdpDrop(v GriddnsattackmitigationDetectUdpDrop)`

SetDetectUdpDrop sets DetectUdpDrop field to given value.

### HasDetectUdpDrop

`func (o *GridDnsAttackMitigation) HasDetectUdpDrop() bool`

HasDetectUdpDrop returns a boolean if a field has been set.

### GetInterval

`func (o *GridDnsAttackMitigation) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GridDnsAttackMitigation) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GridDnsAttackMitigation) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GridDnsAttackMitigation) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


