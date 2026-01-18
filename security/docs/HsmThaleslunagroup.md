# HsmThaleslunagroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The HSM Thales Luna group comment. | [optional] 
**GroupSn** | Pointer to **string** | The HSM Thales Luna group serial number. | [optional] [readonly] 
**HsmVersion** | Pointer to **string** | The HSM Thales Luna version. | [optional] 
**Name** | Pointer to **string** | The HSM Thales Luna group name. | [optional] 
**PassPhrase** | Pointer to **string** | The pass phrase used to unlock the HSM Thales Luna keystore. | [optional] 
**Status** | Pointer to **string** | The status of all HSM Thales Luna devices in the group. | [optional] [readonly] 
**Thalesluna** | Pointer to [**[]HsmThaleslunagroupThalesluna**](HsmThaleslunagroupThalesluna.md) | The list of HSM Thales Luna devices. | [optional] 

## Methods

### NewHsmThaleslunagroup

`func NewHsmThaleslunagroup() *HsmThaleslunagroup`

NewHsmThaleslunagroup instantiates a new HsmThaleslunagroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmThaleslunagroupWithDefaults

`func NewHsmThaleslunagroupWithDefaults() *HsmThaleslunagroup`

NewHsmThaleslunagroupWithDefaults instantiates a new HsmThaleslunagroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *HsmThaleslunagroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *HsmThaleslunagroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *HsmThaleslunagroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *HsmThaleslunagroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *HsmThaleslunagroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *HsmThaleslunagroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *HsmThaleslunagroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *HsmThaleslunagroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupSn

`func (o *HsmThaleslunagroup) GetGroupSn() string`

GetGroupSn returns the GroupSn field if non-nil, zero value otherwise.

### GetGroupSnOk

`func (o *HsmThaleslunagroup) GetGroupSnOk() (*string, bool)`

GetGroupSnOk returns a tuple with the GroupSn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSn

`func (o *HsmThaleslunagroup) SetGroupSn(v string)`

SetGroupSn sets GroupSn field to given value.

### HasGroupSn

`func (o *HsmThaleslunagroup) HasGroupSn() bool`

HasGroupSn returns a boolean if a field has been set.

### GetHsmVersion

`func (o *HsmThaleslunagroup) GetHsmVersion() string`

GetHsmVersion returns the HsmVersion field if non-nil, zero value otherwise.

### GetHsmVersionOk

`func (o *HsmThaleslunagroup) GetHsmVersionOk() (*string, bool)`

GetHsmVersionOk returns a tuple with the HsmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmVersion

`func (o *HsmThaleslunagroup) SetHsmVersion(v string)`

SetHsmVersion sets HsmVersion field to given value.

### HasHsmVersion

`func (o *HsmThaleslunagroup) HasHsmVersion() bool`

HasHsmVersion returns a boolean if a field has been set.

### GetName

`func (o *HsmThaleslunagroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HsmThaleslunagroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HsmThaleslunagroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HsmThaleslunagroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassPhrase

`func (o *HsmThaleslunagroup) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *HsmThaleslunagroup) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *HsmThaleslunagroup) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *HsmThaleslunagroup) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetStatus

`func (o *HsmThaleslunagroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HsmThaleslunagroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HsmThaleslunagroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HsmThaleslunagroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThalesluna

`func (o *HsmThaleslunagroup) GetThalesluna() []HsmThaleslunagroupThalesluna`

GetThalesluna returns the Thalesluna field if non-nil, zero value otherwise.

### GetThaleslunaOk

`func (o *HsmThaleslunagroup) GetThaleslunaOk() (*[]HsmThaleslunagroupThalesluna, bool)`

GetThaleslunaOk returns a tuple with the Thalesluna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThalesluna

`func (o *HsmThaleslunagroup) SetThalesluna(v []HsmThaleslunagroupThalesluna)`

SetThalesluna sets Thalesluna field to given value.

### HasThalesluna

`func (o *HsmThaleslunagroup) HasThalesluna() bool`

HasThalesluna returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


