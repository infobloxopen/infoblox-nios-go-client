# HsmEntrustnshieldgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CardName** | Pointer to **string** | The Entrust nShield HSM softcard name. | [optional] 
**Comment** | Pointer to **string** | The Entrust nShield HSM group comment. | [optional] 
**EntrustnshieldHsm** | Pointer to [**[]HsmEntrustnshieldgroupEntrustnshieldHsm**](HsmEntrustnshieldgroupEntrustnshieldHsm.md) | The list of Entrust nShield HSM devices. | [optional] 
**KeyServerIp** | Pointer to **string** | The remote file server (RFS) IPv4 Address. | [optional] 
**KeyServerPort** | Pointer to **int64** | The remote file server (RFS) port. | [optional] 
**Name** | Pointer to **string** | The Entrust nShield HSM group name. | [optional] 
**PassPhrase** | Pointer to **string** | The password phrase used to unlock the Entrust nShield HSM keystore. | [optional] 
**Protection** | Pointer to **string** | The level of protection that the HSM group uses for the DNSSEC key data. | [optional] 
**Status** | Pointer to **string** | The status of all Entrust nShield HSM devices in the group. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewHsmEntrustnshieldgroup

`func NewHsmEntrustnshieldgroup() *HsmEntrustnshieldgroup`

NewHsmEntrustnshieldgroup instantiates a new HsmEntrustnshieldgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmEntrustnshieldgroupWithDefaults

`func NewHsmEntrustnshieldgroupWithDefaults() *HsmEntrustnshieldgroup`

NewHsmEntrustnshieldgroupWithDefaults instantiates a new HsmEntrustnshieldgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *HsmEntrustnshieldgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *HsmEntrustnshieldgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *HsmEntrustnshieldgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *HsmEntrustnshieldgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCardName

`func (o *HsmEntrustnshieldgroup) GetCardName() string`

GetCardName returns the CardName field if non-nil, zero value otherwise.

### GetCardNameOk

`func (o *HsmEntrustnshieldgroup) GetCardNameOk() (*string, bool)`

GetCardNameOk returns a tuple with the CardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardName

`func (o *HsmEntrustnshieldgroup) SetCardName(v string)`

SetCardName sets CardName field to given value.

### HasCardName

`func (o *HsmEntrustnshieldgroup) HasCardName() bool`

HasCardName returns a boolean if a field has been set.

### GetComment

`func (o *HsmEntrustnshieldgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *HsmEntrustnshieldgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *HsmEntrustnshieldgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *HsmEntrustnshieldgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEntrustnshieldHsm

`func (o *HsmEntrustnshieldgroup) GetEntrustnshieldHsm() []HsmEntrustnshieldgroupEntrustnshieldHsm`

GetEntrustnshieldHsm returns the EntrustnshieldHsm field if non-nil, zero value otherwise.

### GetEntrustnshieldHsmOk

`func (o *HsmEntrustnshieldgroup) GetEntrustnshieldHsmOk() (*[]HsmEntrustnshieldgroupEntrustnshieldHsm, bool)`

GetEntrustnshieldHsmOk returns a tuple with the EntrustnshieldHsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrustnshieldHsm

`func (o *HsmEntrustnshieldgroup) SetEntrustnshieldHsm(v []HsmEntrustnshieldgroupEntrustnshieldHsm)`

SetEntrustnshieldHsm sets EntrustnshieldHsm field to given value.

### HasEntrustnshieldHsm

`func (o *HsmEntrustnshieldgroup) HasEntrustnshieldHsm() bool`

HasEntrustnshieldHsm returns a boolean if a field has been set.

### GetKeyServerIp

`func (o *HsmEntrustnshieldgroup) GetKeyServerIp() string`

GetKeyServerIp returns the KeyServerIp field if non-nil, zero value otherwise.

### GetKeyServerIpOk

`func (o *HsmEntrustnshieldgroup) GetKeyServerIpOk() (*string, bool)`

GetKeyServerIpOk returns a tuple with the KeyServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerIp

`func (o *HsmEntrustnshieldgroup) SetKeyServerIp(v string)`

SetKeyServerIp sets KeyServerIp field to given value.

### HasKeyServerIp

`func (o *HsmEntrustnshieldgroup) HasKeyServerIp() bool`

HasKeyServerIp returns a boolean if a field has been set.

### GetKeyServerPort

`func (o *HsmEntrustnshieldgroup) GetKeyServerPort() int64`

GetKeyServerPort returns the KeyServerPort field if non-nil, zero value otherwise.

### GetKeyServerPortOk

`func (o *HsmEntrustnshieldgroup) GetKeyServerPortOk() (*int64, bool)`

GetKeyServerPortOk returns a tuple with the KeyServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerPort

`func (o *HsmEntrustnshieldgroup) SetKeyServerPort(v int64)`

SetKeyServerPort sets KeyServerPort field to given value.

### HasKeyServerPort

`func (o *HsmEntrustnshieldgroup) HasKeyServerPort() bool`

HasKeyServerPort returns a boolean if a field has been set.

### GetName

`func (o *HsmEntrustnshieldgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HsmEntrustnshieldgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HsmEntrustnshieldgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HsmEntrustnshieldgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassPhrase

`func (o *HsmEntrustnshieldgroup) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *HsmEntrustnshieldgroup) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *HsmEntrustnshieldgroup) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *HsmEntrustnshieldgroup) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetProtection

`func (o *HsmEntrustnshieldgroup) GetProtection() string`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *HsmEntrustnshieldgroup) GetProtectionOk() (*string, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *HsmEntrustnshieldgroup) SetProtection(v string)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *HsmEntrustnshieldgroup) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetStatus

`func (o *HsmEntrustnshieldgroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HsmEntrustnshieldgroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HsmEntrustnshieldgroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HsmEntrustnshieldgroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *HsmEntrustnshieldgroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HsmEntrustnshieldgroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HsmEntrustnshieldgroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HsmEntrustnshieldgroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


