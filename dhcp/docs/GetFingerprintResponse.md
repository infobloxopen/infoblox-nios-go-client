# GetFingerprintResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the Fingerprint; maximum 256 characters. | [optional] 
**DeviceClass** | Pointer to **string** | A class of DHCP Fingerprint object; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the DHCP Fingerprint object is disabled or not. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv6OptionSequence** | Pointer to **[]string** | A list (comma separated list) of IPv6 option number sequences of the device or operating system. | [optional] 
**Name** | Pointer to **string** | Name of the DHCP Fingerprint object. | [optional] 
**OptionSequence** | Pointer to **[]string** | A list (comma separated list) of IPv4 option number sequences of the device or operating system. | [optional] 
**Type** | Pointer to **string** | The type of the DHCP Fingerprint object. | [optional] 
**VendorId** | Pointer to **[]string** | A list of vendor IDs of the device or operating system. | [optional] 
**Result** | Pointer to [**Fingerprint**](Fingerprint.md) |  | [optional] 

## Methods

### NewGetFingerprintResponse

`func NewGetFingerprintResponse() *GetFingerprintResponse`

NewGetFingerprintResponse instantiates a new GetFingerprintResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFingerprintResponseWithDefaults

`func NewGetFingerprintResponseWithDefaults() *GetFingerprintResponse`

NewGetFingerprintResponseWithDefaults instantiates a new GetFingerprintResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFingerprintResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFingerprintResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFingerprintResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFingerprintResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetFingerprintResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFingerprintResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFingerprintResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFingerprintResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeviceClass

`func (o *GetFingerprintResponse) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *GetFingerprintResponse) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *GetFingerprintResponse) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.

### HasDeviceClass

`func (o *GetFingerprintResponse) HasDeviceClass() bool`

HasDeviceClass returns a boolean if a field has been set.

### GetDisable

`func (o *GetFingerprintResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetFingerprintResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetFingerprintResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetFingerprintResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetFingerprintResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetFingerprintResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetFingerprintResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetFingerprintResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetFingerprintResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetFingerprintResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetFingerprintResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetFingerprintResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetFingerprintResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetFingerprintResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetFingerprintResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetFingerprintResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv6OptionSequence

`func (o *GetFingerprintResponse) GetIpv6OptionSequence() []string`

GetIpv6OptionSequence returns the Ipv6OptionSequence field if non-nil, zero value otherwise.

### GetIpv6OptionSequenceOk

`func (o *GetFingerprintResponse) GetIpv6OptionSequenceOk() (*[]string, bool)`

GetIpv6OptionSequenceOk returns a tuple with the Ipv6OptionSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6OptionSequence

`func (o *GetFingerprintResponse) SetIpv6OptionSequence(v []string)`

SetIpv6OptionSequence sets Ipv6OptionSequence field to given value.

### HasIpv6OptionSequence

`func (o *GetFingerprintResponse) HasIpv6OptionSequence() bool`

HasIpv6OptionSequence returns a boolean if a field has been set.

### GetName

`func (o *GetFingerprintResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFingerprintResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFingerprintResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFingerprintResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionSequence

`func (o *GetFingerprintResponse) GetOptionSequence() []string`

GetOptionSequence returns the OptionSequence field if non-nil, zero value otherwise.

### GetOptionSequenceOk

`func (o *GetFingerprintResponse) GetOptionSequenceOk() (*[]string, bool)`

GetOptionSequenceOk returns a tuple with the OptionSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSequence

`func (o *GetFingerprintResponse) SetOptionSequence(v []string)`

SetOptionSequence sets OptionSequence field to given value.

### HasOptionSequence

`func (o *GetFingerprintResponse) HasOptionSequence() bool`

HasOptionSequence returns a boolean if a field has been set.

### GetType

`func (o *GetFingerprintResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetFingerprintResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetFingerprintResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetFingerprintResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorId

`func (o *GetFingerprintResponse) GetVendorId() []string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GetFingerprintResponse) GetVendorIdOk() (*[]string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GetFingerprintResponse) SetVendorId(v []string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GetFingerprintResponse) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetResult

`func (o *GetFingerprintResponse) GetResult() Fingerprint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFingerprintResponse) GetResultOk() (*Fingerprint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFingerprintResponse) SetResult(v Fingerprint)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFingerprintResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


