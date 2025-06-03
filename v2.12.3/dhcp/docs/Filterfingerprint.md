# Filterfingerprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Fingerprint** | Pointer to **[]string** | The list of DHCP Fingerprint objects. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP Fingerprint Filter object. | [optional] 

## Methods

### NewFilterfingerprint

`func NewFilterfingerprint() *Filterfingerprint`

NewFilterfingerprint instantiates a new Filterfingerprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterfingerprintWithDefaults

`func NewFilterfingerprintWithDefaults() *Filterfingerprint`

NewFilterfingerprintWithDefaults instantiates a new Filterfingerprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Filterfingerprint) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Filterfingerprint) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Filterfingerprint) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Filterfingerprint) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Filterfingerprint) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Filterfingerprint) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Filterfingerprint) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Filterfingerprint) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *Filterfingerprint) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Filterfingerprint) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Filterfingerprint) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Filterfingerprint) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFingerprint

`func (o *Filterfingerprint) GetFingerprint() []string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Filterfingerprint) GetFingerprintOk() (*[]string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Filterfingerprint) SetFingerprint(v []string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Filterfingerprint) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetName

`func (o *Filterfingerprint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filterfingerprint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filterfingerprint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Filterfingerprint) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


