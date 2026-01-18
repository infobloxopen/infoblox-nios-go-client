# Memberdfp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**DfpForwardFirst** | Pointer to **bool** | Option to resolve DNS query if resolution over Active Trust Cloud failed. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**HostName** | Pointer to **string** | Host name of the parent Member | [optional] [readonly] 
**IsDfpOverride** | Pointer to **bool** | DFP override lock&#39;. | [optional] 

## Methods

### NewMemberdfp

`func NewMemberdfp() *Memberdfp`

NewMemberdfp instantiates a new Memberdfp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberdfpWithDefaults

`func NewMemberdfpWithDefaults() *Memberdfp`

NewMemberdfpWithDefaults instantiates a new Memberdfp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Memberdfp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Memberdfp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Memberdfp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Memberdfp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDfpForwardFirst

`func (o *Memberdfp) GetDfpForwardFirst() bool`

GetDfpForwardFirst returns the DfpForwardFirst field if non-nil, zero value otherwise.

### GetDfpForwardFirstOk

`func (o *Memberdfp) GetDfpForwardFirstOk() (*bool, bool)`

GetDfpForwardFirstOk returns a tuple with the DfpForwardFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpForwardFirst

`func (o *Memberdfp) SetDfpForwardFirst(v bool)`

SetDfpForwardFirst sets DfpForwardFirst field to given value.

### HasDfpForwardFirst

`func (o *Memberdfp) HasDfpForwardFirst() bool`

HasDfpForwardFirst returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Memberdfp) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Memberdfp) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Memberdfp) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Memberdfp) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Memberdfp) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Memberdfp) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Memberdfp) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Memberdfp) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Memberdfp) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Memberdfp) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Memberdfp) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Memberdfp) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHostName

`func (o *Memberdfp) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Memberdfp) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Memberdfp) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *Memberdfp) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIsDfpOverride

`func (o *Memberdfp) GetIsDfpOverride() bool`

GetIsDfpOverride returns the IsDfpOverride field if non-nil, zero value otherwise.

### GetIsDfpOverrideOk

`func (o *Memberdfp) GetIsDfpOverrideOk() (*bool, bool)`

GetIsDfpOverrideOk returns a tuple with the IsDfpOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDfpOverride

`func (o *Memberdfp) SetIsDfpOverride(v bool)`

SetIsDfpOverride sets IsDfpOverride field to given value.

### HasIsDfpOverride

`func (o *Memberdfp) HasIsDfpOverride() bool`

HasIsDfpOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


