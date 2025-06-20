# GetMemberdfpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**DfpForwardFirst** | Pointer to **bool** | Option to resolve DNS query if resolution over Active Trust Cloud failed. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**HostName** | Pointer to **string** | Host name of the parent Member | [optional] [readonly] 
**IsDfpOverride** | Pointer to **bool** | DFP override lock&#39;. | [optional] 
**Result** | Pointer to [**Memberdfp**](Memberdfp.md) |  | [optional] 

## Methods

### NewGetMemberdfpResponse

`func NewGetMemberdfpResponse() *GetMemberdfpResponse`

NewGetMemberdfpResponse instantiates a new GetMemberdfpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMemberdfpResponseWithDefaults

`func NewGetMemberdfpResponseWithDefaults() *GetMemberdfpResponse`

NewGetMemberdfpResponseWithDefaults instantiates a new GetMemberdfpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMemberdfpResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMemberdfpResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMemberdfpResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMemberdfpResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDfpForwardFirst

`func (o *GetMemberdfpResponse) GetDfpForwardFirst() bool`

GetDfpForwardFirst returns the DfpForwardFirst field if non-nil, zero value otherwise.

### GetDfpForwardFirstOk

`func (o *GetMemberdfpResponse) GetDfpForwardFirstOk() (*bool, bool)`

GetDfpForwardFirstOk returns a tuple with the DfpForwardFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpForwardFirst

`func (o *GetMemberdfpResponse) SetDfpForwardFirst(v bool)`

SetDfpForwardFirst sets DfpForwardFirst field to given value.

### HasDfpForwardFirst

`func (o *GetMemberdfpResponse) HasDfpForwardFirst() bool`

HasDfpForwardFirst returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetMemberdfpResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetMemberdfpResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetMemberdfpResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetMemberdfpResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHostName

`func (o *GetMemberdfpResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetMemberdfpResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetMemberdfpResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetMemberdfpResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIsDfpOverride

`func (o *GetMemberdfpResponse) GetIsDfpOverride() bool`

GetIsDfpOverride returns the IsDfpOverride field if non-nil, zero value otherwise.

### GetIsDfpOverrideOk

`func (o *GetMemberdfpResponse) GetIsDfpOverrideOk() (*bool, bool)`

GetIsDfpOverrideOk returns a tuple with the IsDfpOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDfpOverride

`func (o *GetMemberdfpResponse) SetIsDfpOverride(v bool)`

SetIsDfpOverride sets IsDfpOverride field to given value.

### HasIsDfpOverride

`func (o *GetMemberdfpResponse) HasIsDfpOverride() bool`

HasIsDfpOverride returns a boolean if a field has been set.

### GetResult

`func (o *GetMemberdfpResponse) GetResult() Memberdfp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMemberdfpResponse) GetResultOk() (*Memberdfp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMemberdfpResponse) SetResult(v Memberdfp)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMemberdfpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


