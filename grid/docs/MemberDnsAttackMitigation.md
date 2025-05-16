# MemberDnsAttackMitigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectChr** | Pointer to [**MemberdnsattackmitigationDetectChr**](MemberdnsattackmitigationDetectChr.md) |  | [optional] 
**DetectChrGrace** | Pointer to **int64** | The cache utilization (in percentage) when Cache Hit Ratio (CHR) starts. | [optional] 
**DetectNxdomainResponses** | Pointer to [**MemberdnsattackmitigationDetectNxdomainResponses**](MemberdnsattackmitigationDetectNxdomainResponses.md) |  | [optional] 
**DetectUdpDrop** | Pointer to [**MemberdnsattackmitigationDetectUdpDrop**](MemberdnsattackmitigationDetectUdpDrop.md) |  | [optional] 
**Interval** | Pointer to **int64** | The minimum time interval (in seconds) between changes in attack status. | [optional] 

## Methods

### NewMemberDnsAttackMitigation

`func NewMemberDnsAttackMitigation() *MemberDnsAttackMitigation`

NewMemberDnsAttackMitigation instantiates a new MemberDnsAttackMitigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsAttackMitigationWithDefaults

`func NewMemberDnsAttackMitigationWithDefaults() *MemberDnsAttackMitigation`

NewMemberDnsAttackMitigationWithDefaults instantiates a new MemberDnsAttackMitigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectChr

`func (o *MemberDnsAttackMitigation) GetDetectChr() MemberdnsattackmitigationDetectChr`

GetDetectChr returns the DetectChr field if non-nil, zero value otherwise.

### GetDetectChrOk

`func (o *MemberDnsAttackMitigation) GetDetectChrOk() (*MemberdnsattackmitigationDetectChr, bool)`

GetDetectChrOk returns a tuple with the DetectChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectChr

`func (o *MemberDnsAttackMitigation) SetDetectChr(v MemberdnsattackmitigationDetectChr)`

SetDetectChr sets DetectChr field to given value.

### HasDetectChr

`func (o *MemberDnsAttackMitigation) HasDetectChr() bool`

HasDetectChr returns a boolean if a field has been set.

### GetDetectChrGrace

`func (o *MemberDnsAttackMitigation) GetDetectChrGrace() int64`

GetDetectChrGrace returns the DetectChrGrace field if non-nil, zero value otherwise.

### GetDetectChrGraceOk

`func (o *MemberDnsAttackMitigation) GetDetectChrGraceOk() (*int64, bool)`

GetDetectChrGraceOk returns a tuple with the DetectChrGrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectChrGrace

`func (o *MemberDnsAttackMitigation) SetDetectChrGrace(v int64)`

SetDetectChrGrace sets DetectChrGrace field to given value.

### HasDetectChrGrace

`func (o *MemberDnsAttackMitigation) HasDetectChrGrace() bool`

HasDetectChrGrace returns a boolean if a field has been set.

### GetDetectNxdomainResponses

`func (o *MemberDnsAttackMitigation) GetDetectNxdomainResponses() MemberdnsattackmitigationDetectNxdomainResponses`

GetDetectNxdomainResponses returns the DetectNxdomainResponses field if non-nil, zero value otherwise.

### GetDetectNxdomainResponsesOk

`func (o *MemberDnsAttackMitigation) GetDetectNxdomainResponsesOk() (*MemberdnsattackmitigationDetectNxdomainResponses, bool)`

GetDetectNxdomainResponsesOk returns a tuple with the DetectNxdomainResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectNxdomainResponses

`func (o *MemberDnsAttackMitigation) SetDetectNxdomainResponses(v MemberdnsattackmitigationDetectNxdomainResponses)`

SetDetectNxdomainResponses sets DetectNxdomainResponses field to given value.

### HasDetectNxdomainResponses

`func (o *MemberDnsAttackMitigation) HasDetectNxdomainResponses() bool`

HasDetectNxdomainResponses returns a boolean if a field has been set.

### GetDetectUdpDrop

`func (o *MemberDnsAttackMitigation) GetDetectUdpDrop() MemberdnsattackmitigationDetectUdpDrop`

GetDetectUdpDrop returns the DetectUdpDrop field if non-nil, zero value otherwise.

### GetDetectUdpDropOk

`func (o *MemberDnsAttackMitigation) GetDetectUdpDropOk() (*MemberdnsattackmitigationDetectUdpDrop, bool)`

GetDetectUdpDropOk returns a tuple with the DetectUdpDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectUdpDrop

`func (o *MemberDnsAttackMitigation) SetDetectUdpDrop(v MemberdnsattackmitigationDetectUdpDrop)`

SetDetectUdpDrop sets DetectUdpDrop field to given value.

### HasDetectUdpDrop

`func (o *MemberDnsAttackMitigation) HasDetectUdpDrop() bool`

HasDetectUdpDrop returns a boolean if a field has been set.

### GetInterval

`func (o *MemberDnsAttackMitigation) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MemberDnsAttackMitigation) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MemberDnsAttackMitigation) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MemberDnsAttackMitigation) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


