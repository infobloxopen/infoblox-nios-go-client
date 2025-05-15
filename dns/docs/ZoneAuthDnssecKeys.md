# ZoneAuthDnssecKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **int64** | The tag of the key for the zone. | [optional] 
**Status** | Pointer to **string** | The status of the key for the zone. | [optional] [readonly] 
**NextEventDate** | Pointer to **int64** | The next event date for the key, the rollover date for an active key or the removal date for an already rolled one. | [optional] [readonly] 
**Type** | Pointer to **string** | The key type. | [optional] [readonly] 
**Algorithm** | Pointer to **string** | The public-key encryption algorithm. Values 1, 3 and 6 are deprecated from NIOS 9.0. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The Base-64 encoding of the public key. | [optional] [readonly] 

## Methods

### NewZoneAuthDnssecKeys

`func NewZoneAuthDnssecKeys() *ZoneAuthDnssecKeys`

NewZoneAuthDnssecKeys instantiates a new ZoneAuthDnssecKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthDnssecKeysWithDefaults

`func NewZoneAuthDnssecKeysWithDefaults() *ZoneAuthDnssecKeys`

NewZoneAuthDnssecKeysWithDefaults instantiates a new ZoneAuthDnssecKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ZoneAuthDnssecKeys) GetTag() int64`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ZoneAuthDnssecKeys) GetTagOk() (*int64, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ZoneAuthDnssecKeys) SetTag(v int64)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ZoneAuthDnssecKeys) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetStatus

`func (o *ZoneAuthDnssecKeys) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneAuthDnssecKeys) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneAuthDnssecKeys) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZoneAuthDnssecKeys) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNextEventDate

`func (o *ZoneAuthDnssecKeys) GetNextEventDate() int64`

GetNextEventDate returns the NextEventDate field if non-nil, zero value otherwise.

### GetNextEventDateOk

`func (o *ZoneAuthDnssecKeys) GetNextEventDateOk() (*int64, bool)`

GetNextEventDateOk returns a tuple with the NextEventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEventDate

`func (o *ZoneAuthDnssecKeys) SetNextEventDate(v int64)`

SetNextEventDate sets NextEventDate field to given value.

### HasNextEventDate

`func (o *ZoneAuthDnssecKeys) HasNextEventDate() bool`

HasNextEventDate returns a boolean if a field has been set.

### GetType

`func (o *ZoneAuthDnssecKeys) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneAuthDnssecKeys) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneAuthDnssecKeys) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneAuthDnssecKeys) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlgorithm

`func (o *ZoneAuthDnssecKeys) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ZoneAuthDnssecKeys) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ZoneAuthDnssecKeys) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ZoneAuthDnssecKeys) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetPublicKey

`func (o *ZoneAuthDnssecKeys) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ZoneAuthDnssecKeys) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ZoneAuthDnssecKeys) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ZoneAuthDnssecKeys) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


