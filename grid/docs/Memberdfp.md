# Memberdfp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**DfpForwardFirst** | Pointer to **bool** | Option to resolve DNS query if resolution over Active Trust Cloud failed. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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

### GetExtattrs

`func (o *Memberdfp) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Memberdfp) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Memberdfp) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Memberdfp) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

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


