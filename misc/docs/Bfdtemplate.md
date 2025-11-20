# Bfdtemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AuthenticationKey** | Pointer to **string** | The authentication key for BFD protocol message-digest authentication. | [optional] 
**AuthenticationKeyId** | Pointer to **int64** | The authentication key identifier for BFD protocol authentication. Valid values are between 1 and 255. | [optional] 
**AuthenticationType** | Pointer to **string** | The authentication type for BFD protocol. | [optional] 
**DetectionMultiplier** | Pointer to **int64** | The detection time multiplier value for BFD protocol. The negotiated transmit interval, multiplied by this value, provides the detection time for the receiving system in asynchronous BFD mode. Valid values are between 3 and 50. | [optional] 
**MinRxInterval** | Pointer to **int64** | The minimum receive time (in seconds) for BFD protocol. Valid values are between 50 and 9999. | [optional] 
**MinTxInterval** | Pointer to **int64** | The minimum transmission time (in seconds) for BFD protocol. Valid values are between 50 and 9999. | [optional] 
**Name** | Pointer to **string** | The name of the BFD template object. | [optional] 

## Methods

### NewBfdtemplate

`func NewBfdtemplate() *Bfdtemplate`

NewBfdtemplate instantiates a new Bfdtemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBfdtemplateWithDefaults

`func NewBfdtemplateWithDefaults() *Bfdtemplate`

NewBfdtemplateWithDefaults instantiates a new Bfdtemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Bfdtemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Bfdtemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Bfdtemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Bfdtemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *Bfdtemplate) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *Bfdtemplate) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *Bfdtemplate) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *Bfdtemplate) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetAuthenticationKeyId

`func (o *Bfdtemplate) GetAuthenticationKeyId() int64`

GetAuthenticationKeyId returns the AuthenticationKeyId field if non-nil, zero value otherwise.

### GetAuthenticationKeyIdOk

`func (o *Bfdtemplate) GetAuthenticationKeyIdOk() (*int64, bool)`

GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKeyId

`func (o *Bfdtemplate) SetAuthenticationKeyId(v int64)`

SetAuthenticationKeyId sets AuthenticationKeyId field to given value.

### HasAuthenticationKeyId

`func (o *Bfdtemplate) HasAuthenticationKeyId() bool`

HasAuthenticationKeyId returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *Bfdtemplate) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *Bfdtemplate) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *Bfdtemplate) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *Bfdtemplate) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetDetectionMultiplier

`func (o *Bfdtemplate) GetDetectionMultiplier() int64`

GetDetectionMultiplier returns the DetectionMultiplier field if non-nil, zero value otherwise.

### GetDetectionMultiplierOk

`func (o *Bfdtemplate) GetDetectionMultiplierOk() (*int64, bool)`

GetDetectionMultiplierOk returns a tuple with the DetectionMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMultiplier

`func (o *Bfdtemplate) SetDetectionMultiplier(v int64)`

SetDetectionMultiplier sets DetectionMultiplier field to given value.

### HasDetectionMultiplier

`func (o *Bfdtemplate) HasDetectionMultiplier() bool`

HasDetectionMultiplier returns a boolean if a field has been set.

### GetMinRxInterval

`func (o *Bfdtemplate) GetMinRxInterval() int64`

GetMinRxInterval returns the MinRxInterval field if non-nil, zero value otherwise.

### GetMinRxIntervalOk

`func (o *Bfdtemplate) GetMinRxIntervalOk() (*int64, bool)`

GetMinRxIntervalOk returns a tuple with the MinRxInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRxInterval

`func (o *Bfdtemplate) SetMinRxInterval(v int64)`

SetMinRxInterval sets MinRxInterval field to given value.

### HasMinRxInterval

`func (o *Bfdtemplate) HasMinRxInterval() bool`

HasMinRxInterval returns a boolean if a field has been set.

### GetMinTxInterval

`func (o *Bfdtemplate) GetMinTxInterval() int64`

GetMinTxInterval returns the MinTxInterval field if non-nil, zero value otherwise.

### GetMinTxIntervalOk

`func (o *Bfdtemplate) GetMinTxIntervalOk() (*int64, bool)`

GetMinTxIntervalOk returns a tuple with the MinTxInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTxInterval

`func (o *Bfdtemplate) SetMinTxInterval(v int64)`

SetMinTxInterval sets MinTxInterval field to given value.

### HasMinTxInterval

`func (o *Bfdtemplate) HasMinTxInterval() bool`

HasMinTxInterval returns a boolean if a field has been set.

### GetName

`func (o *Bfdtemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bfdtemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bfdtemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bfdtemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


