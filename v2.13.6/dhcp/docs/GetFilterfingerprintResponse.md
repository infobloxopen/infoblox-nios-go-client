# GetFilterfingerprintResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Fingerprint** | Pointer to **[]string** | The list of DHCP Fingerprint objects. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP Fingerprint Filter object. | [optional] 
**Result** | Pointer to [**Filterfingerprint**](Filterfingerprint.md) |  | [optional] 

## Methods

### NewGetFilterfingerprintResponse

`func NewGetFilterfingerprintResponse() *GetFilterfingerprintResponse`

NewGetFilterfingerprintResponse instantiates a new GetFilterfingerprintResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFilterfingerprintResponseWithDefaults

`func NewGetFilterfingerprintResponseWithDefaults() *GetFilterfingerprintResponse`

NewGetFilterfingerprintResponseWithDefaults instantiates a new GetFilterfingerprintResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFilterfingerprintResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFilterfingerprintResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFilterfingerprintResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFilterfingerprintResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetFilterfingerprintResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFilterfingerprintResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFilterfingerprintResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFilterfingerprintResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetFilterfingerprintResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetFilterfingerprintResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetFilterfingerprintResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetFilterfingerprintResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFingerprint

`func (o *GetFilterfingerprintResponse) GetFingerprint() []string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GetFilterfingerprintResponse) GetFingerprintOk() (*[]string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GetFilterfingerprintResponse) SetFingerprint(v []string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GetFilterfingerprintResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetName

`func (o *GetFilterfingerprintResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFilterfingerprintResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFilterfingerprintResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFilterfingerprintResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetFilterfingerprintResponse) GetResult() Filterfingerprint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFilterfingerprintResponse) GetResultOk() (*Filterfingerprint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFilterfingerprintResponse) SetResult(v Filterfingerprint)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFilterfingerprintResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


