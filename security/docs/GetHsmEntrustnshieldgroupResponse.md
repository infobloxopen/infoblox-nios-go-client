# GetHsmEntrustnshieldgroupResponse

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
**Result** | Pointer to [**HsmEntrustnshieldgroup**](HsmEntrustnshieldgroup.md) |  | [optional] 

## Methods

### NewGetHsmEntrustnshieldgroupResponse

`func NewGetHsmEntrustnshieldgroupResponse() *GetHsmEntrustnshieldgroupResponse`

NewGetHsmEntrustnshieldgroupResponse instantiates a new GetHsmEntrustnshieldgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHsmEntrustnshieldgroupResponseWithDefaults

`func NewGetHsmEntrustnshieldgroupResponseWithDefaults() *GetHsmEntrustnshieldgroupResponse`

NewGetHsmEntrustnshieldgroupResponseWithDefaults instantiates a new GetHsmEntrustnshieldgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetHsmEntrustnshieldgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetHsmEntrustnshieldgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetHsmEntrustnshieldgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCardName

`func (o *GetHsmEntrustnshieldgroupResponse) GetCardName() string`

GetCardName returns the CardName field if non-nil, zero value otherwise.

### GetCardNameOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetCardNameOk() (*string, bool)`

GetCardNameOk returns a tuple with the CardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardName

`func (o *GetHsmEntrustnshieldgroupResponse) SetCardName(v string)`

SetCardName sets CardName field to given value.

### HasCardName

`func (o *GetHsmEntrustnshieldgroupResponse) HasCardName() bool`

HasCardName returns a boolean if a field has been set.

### GetComment

`func (o *GetHsmEntrustnshieldgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetHsmEntrustnshieldgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetHsmEntrustnshieldgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEntrustnshieldHsm

`func (o *GetHsmEntrustnshieldgroupResponse) GetEntrustnshieldHsm() []HsmEntrustnshieldgroupEntrustnshieldHsm`

GetEntrustnshieldHsm returns the EntrustnshieldHsm field if non-nil, zero value otherwise.

### GetEntrustnshieldHsmOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetEntrustnshieldHsmOk() (*[]HsmEntrustnshieldgroupEntrustnshieldHsm, bool)`

GetEntrustnshieldHsmOk returns a tuple with the EntrustnshieldHsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrustnshieldHsm

`func (o *GetHsmEntrustnshieldgroupResponse) SetEntrustnshieldHsm(v []HsmEntrustnshieldgroupEntrustnshieldHsm)`

SetEntrustnshieldHsm sets EntrustnshieldHsm field to given value.

### HasEntrustnshieldHsm

`func (o *GetHsmEntrustnshieldgroupResponse) HasEntrustnshieldHsm() bool`

HasEntrustnshieldHsm returns a boolean if a field has been set.

### GetKeyServerIp

`func (o *GetHsmEntrustnshieldgroupResponse) GetKeyServerIp() string`

GetKeyServerIp returns the KeyServerIp field if non-nil, zero value otherwise.

### GetKeyServerIpOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetKeyServerIpOk() (*string, bool)`

GetKeyServerIpOk returns a tuple with the KeyServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerIp

`func (o *GetHsmEntrustnshieldgroupResponse) SetKeyServerIp(v string)`

SetKeyServerIp sets KeyServerIp field to given value.

### HasKeyServerIp

`func (o *GetHsmEntrustnshieldgroupResponse) HasKeyServerIp() bool`

HasKeyServerIp returns a boolean if a field has been set.

### GetKeyServerPort

`func (o *GetHsmEntrustnshieldgroupResponse) GetKeyServerPort() int64`

GetKeyServerPort returns the KeyServerPort field if non-nil, zero value otherwise.

### GetKeyServerPortOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetKeyServerPortOk() (*int64, bool)`

GetKeyServerPortOk returns a tuple with the KeyServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerPort

`func (o *GetHsmEntrustnshieldgroupResponse) SetKeyServerPort(v int64)`

SetKeyServerPort sets KeyServerPort field to given value.

### HasKeyServerPort

`func (o *GetHsmEntrustnshieldgroupResponse) HasKeyServerPort() bool`

HasKeyServerPort returns a boolean if a field has been set.

### GetName

`func (o *GetHsmEntrustnshieldgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHsmEntrustnshieldgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHsmEntrustnshieldgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassPhrase

`func (o *GetHsmEntrustnshieldgroupResponse) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *GetHsmEntrustnshieldgroupResponse) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *GetHsmEntrustnshieldgroupResponse) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetProtection

`func (o *GetHsmEntrustnshieldgroupResponse) GetProtection() string`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetProtectionOk() (*string, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *GetHsmEntrustnshieldgroupResponse) SetProtection(v string)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *GetHsmEntrustnshieldgroupResponse) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetStatus

`func (o *GetHsmEntrustnshieldgroupResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHsmEntrustnshieldgroupResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetHsmEntrustnshieldgroupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetHsmEntrustnshieldgroupResponse) GetResult() HsmEntrustnshieldgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetHsmEntrustnshieldgroupResponse) GetResultOk() (*HsmEntrustnshieldgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetHsmEntrustnshieldgroupResponse) SetResult(v HsmEntrustnshieldgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetHsmEntrustnshieldgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


