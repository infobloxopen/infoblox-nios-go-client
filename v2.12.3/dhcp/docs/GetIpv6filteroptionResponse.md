# GetIpv6filteroptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ApplyAsClass** | Pointer to **bool** | Determines if apply as class is enabled or not. If this flag is set to \&quot;true\&quot; the filter is treated as global DHCP class, e.g it is written to DHCPv6 configuration file even if it is not present in any DHCP range. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP IPv6 filter option object. | [optional] 
**Expression** | Pointer to **string** | The conditional expression of a DHCP IPv6 filter option object. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | Determines the lease time of a DHCP IPv6 filter option object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP IPv6 option filter object. | [optional] 
**OptionList** | Pointer to [**[]Ipv6filteroptionOptionList**](Ipv6filteroptionOptionList.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**OptionSpace** | Pointer to **string** | The option space of a DHCP IPv6 filter option object. | [optional] 
**Result** | Pointer to [**Ipv6filteroption**](Ipv6filteroption.md) |  | [optional] 

## Methods

### NewGetIpv6filteroptionResponse

`func NewGetIpv6filteroptionResponse() *GetIpv6filteroptionResponse`

NewGetIpv6filteroptionResponse instantiates a new GetIpv6filteroptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6filteroptionResponseWithDefaults

`func NewGetIpv6filteroptionResponseWithDefaults() *GetIpv6filteroptionResponse`

NewGetIpv6filteroptionResponseWithDefaults instantiates a new GetIpv6filteroptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6filteroptionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6filteroptionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6filteroptionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6filteroptionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetApplyAsClass

`func (o *GetIpv6filteroptionResponse) GetApplyAsClass() bool`

GetApplyAsClass returns the ApplyAsClass field if non-nil, zero value otherwise.

### GetApplyAsClassOk

`func (o *GetIpv6filteroptionResponse) GetApplyAsClassOk() (*bool, bool)`

GetApplyAsClassOk returns a tuple with the ApplyAsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAsClass

`func (o *GetIpv6filteroptionResponse) SetApplyAsClass(v bool)`

SetApplyAsClass sets ApplyAsClass field to given value.

### HasApplyAsClass

`func (o *GetIpv6filteroptionResponse) HasApplyAsClass() bool`

HasApplyAsClass returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6filteroptionResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6filteroptionResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6filteroptionResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6filteroptionResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpression

`func (o *GetIpv6filteroptionResponse) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *GetIpv6filteroptionResponse) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *GetIpv6filteroptionResponse) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *GetIpv6filteroptionResponse) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetIpv6filteroptionResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetIpv6filteroptionResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetIpv6filteroptionResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetIpv6filteroptionResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLeaseTime

`func (o *GetIpv6filteroptionResponse) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *GetIpv6filteroptionResponse) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *GetIpv6filteroptionResponse) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *GetIpv6filteroptionResponse) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6filteroptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6filteroptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6filteroptionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6filteroptionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionList

`func (o *GetIpv6filteroptionResponse) GetOptionList() []Ipv6filteroptionOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *GetIpv6filteroptionResponse) GetOptionListOk() (*[]Ipv6filteroptionOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *GetIpv6filteroptionResponse) SetOptionList(v []Ipv6filteroptionOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *GetIpv6filteroptionResponse) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetOptionSpace

`func (o *GetIpv6filteroptionResponse) GetOptionSpace() string`

GetOptionSpace returns the OptionSpace field if non-nil, zero value otherwise.

### GetOptionSpaceOk

`func (o *GetIpv6filteroptionResponse) GetOptionSpaceOk() (*string, bool)`

GetOptionSpaceOk returns a tuple with the OptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSpace

`func (o *GetIpv6filteroptionResponse) SetOptionSpace(v string)`

SetOptionSpace sets OptionSpace field to given value.

### HasOptionSpace

`func (o *GetIpv6filteroptionResponse) HasOptionSpace() bool`

HasOptionSpace returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6filteroptionResponse) GetResult() Ipv6filteroption`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6filteroptionResponse) GetResultOk() (*Ipv6filteroption, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6filteroptionResponse) SetResult(v Ipv6filteroption)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6filteroptionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


