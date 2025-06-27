# Ipv6filteroption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ApplyAsClass** | Pointer to **bool** | Determines if apply as class is enabled or not. If this flag is set to \&quot;true\&quot; the filter is treated as global DHCP class, e.g it is written to DHCPv6 configuration file even if it is not present in any DHCP range. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP IPv6 filter option object. | [optional] 
**Expression** | Pointer to **string** | The conditional expression of a DHCP IPv6 filter option object. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | Determines the lease time of a DHCP IPv6 filter option object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP IPv6 option filter object. | [optional] 
**OptionList** | Pointer to [**[]Ipv6filteroptionOptionList**](Ipv6filteroptionOptionList.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**OptionSpace** | Pointer to **string** | The option space of a DHCP IPv6 filter option object. | [optional] 

## Methods

### NewIpv6filteroption

`func NewIpv6filteroption() *Ipv6filteroption`

NewIpv6filteroption instantiates a new Ipv6filteroption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6filteroptionWithDefaults

`func NewIpv6filteroptionWithDefaults() *Ipv6filteroption`

NewIpv6filteroptionWithDefaults instantiates a new Ipv6filteroption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6filteroption) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6filteroption) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6filteroption) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6filteroption) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetApplyAsClass

`func (o *Ipv6filteroption) GetApplyAsClass() bool`

GetApplyAsClass returns the ApplyAsClass field if non-nil, zero value otherwise.

### GetApplyAsClassOk

`func (o *Ipv6filteroption) GetApplyAsClassOk() (*bool, bool)`

GetApplyAsClassOk returns a tuple with the ApplyAsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAsClass

`func (o *Ipv6filteroption) SetApplyAsClass(v bool)`

SetApplyAsClass sets ApplyAsClass field to given value.

### HasApplyAsClass

`func (o *Ipv6filteroption) HasApplyAsClass() bool`

HasApplyAsClass returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6filteroption) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6filteroption) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6filteroption) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6filteroption) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpression

`func (o *Ipv6filteroption) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Ipv6filteroption) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Ipv6filteroption) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Ipv6filteroption) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Ipv6filteroption) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Ipv6filteroption) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Ipv6filteroption) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Ipv6filteroption) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLeaseTime

`func (o *Ipv6filteroption) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *Ipv6filteroption) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *Ipv6filteroption) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *Ipv6filteroption) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *Ipv6filteroption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6filteroption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6filteroption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6filteroption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionList

`func (o *Ipv6filteroption) GetOptionList() []Ipv6filteroptionOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *Ipv6filteroption) GetOptionListOk() (*[]Ipv6filteroptionOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *Ipv6filteroption) SetOptionList(v []Ipv6filteroptionOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *Ipv6filteroption) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetOptionSpace

`func (o *Ipv6filteroption) GetOptionSpace() string`

GetOptionSpace returns the OptionSpace field if non-nil, zero value otherwise.

### GetOptionSpaceOk

`func (o *Ipv6filteroption) GetOptionSpaceOk() (*string, bool)`

GetOptionSpaceOk returns a tuple with the OptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSpace

`func (o *Ipv6filteroption) SetOptionSpace(v string)`

SetOptionSpace sets OptionSpace field to given value.

### HasOptionSpace

`func (o *Ipv6filteroption) HasOptionSpace() bool`

HasOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


