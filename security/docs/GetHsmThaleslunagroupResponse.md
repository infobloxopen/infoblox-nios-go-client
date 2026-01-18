# GetHsmThaleslunagroupResponse

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
**Result** | Pointer to [**HsmThaleslunagroup**](HsmThaleslunagroup.md) |  | [optional] 

## Methods

### NewGetHsmThaleslunagroupResponse

`func NewGetHsmThaleslunagroupResponse() *GetHsmThaleslunagroupResponse`

NewGetHsmThaleslunagroupResponse instantiates a new GetHsmThaleslunagroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHsmThaleslunagroupResponseWithDefaults

`func NewGetHsmThaleslunagroupResponseWithDefaults() *GetHsmThaleslunagroupResponse`

NewGetHsmThaleslunagroupResponseWithDefaults instantiates a new GetHsmThaleslunagroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetHsmThaleslunagroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetHsmThaleslunagroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetHsmThaleslunagroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetHsmThaleslunagroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetHsmThaleslunagroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetHsmThaleslunagroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetHsmThaleslunagroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetHsmThaleslunagroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupSn

`func (o *GetHsmThaleslunagroupResponse) GetGroupSn() string`

GetGroupSn returns the GroupSn field if non-nil, zero value otherwise.

### GetGroupSnOk

`func (o *GetHsmThaleslunagroupResponse) GetGroupSnOk() (*string, bool)`

GetGroupSnOk returns a tuple with the GroupSn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSn

`func (o *GetHsmThaleslunagroupResponse) SetGroupSn(v string)`

SetGroupSn sets GroupSn field to given value.

### HasGroupSn

`func (o *GetHsmThaleslunagroupResponse) HasGroupSn() bool`

HasGroupSn returns a boolean if a field has been set.

### GetHsmVersion

`func (o *GetHsmThaleslunagroupResponse) GetHsmVersion() string`

GetHsmVersion returns the HsmVersion field if non-nil, zero value otherwise.

### GetHsmVersionOk

`func (o *GetHsmThaleslunagroupResponse) GetHsmVersionOk() (*string, bool)`

GetHsmVersionOk returns a tuple with the HsmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmVersion

`func (o *GetHsmThaleslunagroupResponse) SetHsmVersion(v string)`

SetHsmVersion sets HsmVersion field to given value.

### HasHsmVersion

`func (o *GetHsmThaleslunagroupResponse) HasHsmVersion() bool`

HasHsmVersion returns a boolean if a field has been set.

### GetName

`func (o *GetHsmThaleslunagroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHsmThaleslunagroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHsmThaleslunagroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHsmThaleslunagroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassPhrase

`func (o *GetHsmThaleslunagroupResponse) GetPassPhrase() string`

GetPassPhrase returns the PassPhrase field if non-nil, zero value otherwise.

### GetPassPhraseOk

`func (o *GetHsmThaleslunagroupResponse) GetPassPhraseOk() (*string, bool)`

GetPassPhraseOk returns a tuple with the PassPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPhrase

`func (o *GetHsmThaleslunagroupResponse) SetPassPhrase(v string)`

SetPassPhrase sets PassPhrase field to given value.

### HasPassPhrase

`func (o *GetHsmThaleslunagroupResponse) HasPassPhrase() bool`

HasPassPhrase returns a boolean if a field has been set.

### GetStatus

`func (o *GetHsmThaleslunagroupResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHsmThaleslunagroupResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHsmThaleslunagroupResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetHsmThaleslunagroupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThalesluna

`func (o *GetHsmThaleslunagroupResponse) GetThalesluna() []HsmThaleslunagroupThalesluna`

GetThalesluna returns the Thalesluna field if non-nil, zero value otherwise.

### GetThaleslunaOk

`func (o *GetHsmThaleslunagroupResponse) GetThaleslunaOk() (*[]HsmThaleslunagroupThalesluna, bool)`

GetThaleslunaOk returns a tuple with the Thalesluna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThalesluna

`func (o *GetHsmThaleslunagroupResponse) SetThalesluna(v []HsmThaleslunagroupThalesluna)`

SetThalesluna sets Thalesluna field to given value.

### HasThalesluna

`func (o *GetHsmThaleslunagroupResponse) HasThalesluna() bool`

HasThalesluna returns a boolean if a field has been set.

### GetResult

`func (o *GetHsmThaleslunagroupResponse) GetResult() HsmThaleslunagroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetHsmThaleslunagroupResponse) GetResultOk() (*HsmThaleslunagroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetHsmThaleslunagroupResponse) SetResult(v HsmThaleslunagroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetHsmThaleslunagroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


