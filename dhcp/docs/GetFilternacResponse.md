# GetFilternacResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP NAC Filter object. | [optional] 
**Expression** | Pointer to **string** | The conditional expression of a DHCP NAC Filter object. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | The length of time the DHCP server leases an IP address to a client. The lease time applies to hosts that meet the filter criteria. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP NAC Filter object. | [optional] 
**Options** | Pointer to [**[]FilternacOptions**](FilternacOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**Result** | Pointer to [**Filternac**](Filternac.md) |  | [optional] 

## Methods

### NewGetFilternacResponse

`func NewGetFilternacResponse() *GetFilternacResponse`

NewGetFilternacResponse instantiates a new GetFilternacResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFilternacResponseWithDefaults

`func NewGetFilternacResponseWithDefaults() *GetFilternacResponse`

NewGetFilternacResponseWithDefaults instantiates a new GetFilternacResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFilternacResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFilternacResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFilternacResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFilternacResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetFilternacResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFilternacResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFilternacResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFilternacResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpression

`func (o *GetFilternacResponse) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *GetFilternacResponse) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *GetFilternacResponse) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *GetFilternacResponse) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetFilternacResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetFilternacResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetFilternacResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetFilternacResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLeaseTime

`func (o *GetFilternacResponse) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *GetFilternacResponse) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *GetFilternacResponse) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *GetFilternacResponse) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *GetFilternacResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFilternacResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFilternacResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFilternacResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *GetFilternacResponse) GetOptions() []FilternacOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetFilternacResponse) GetOptionsOk() (*[]FilternacOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetFilternacResponse) SetOptions(v []FilternacOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetFilternacResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetResult

`func (o *GetFilternacResponse) GetResult() Filternac`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFilternacResponse) GetResultOk() (*Filternac, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFilternacResponse) SetResult(v Filternac)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFilternacResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


