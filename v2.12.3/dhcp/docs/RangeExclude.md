# RangeExclude

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAddress** | Pointer to **string** | The IPv4 Address starting address of the exclusion range. | [optional] 
**EndAddress** | Pointer to **string** | The IPv4 Address ending address of the exclusion range. | [optional] 
**Comment** | Pointer to **string** | Comment for the exclusion range; maximum 256 characters. | [optional] 

## Methods

### NewRangeExclude

`func NewRangeExclude() *RangeExclude`

NewRangeExclude instantiates a new RangeExclude object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeExcludeWithDefaults

`func NewRangeExcludeWithDefaults() *RangeExclude`

NewRangeExcludeWithDefaults instantiates a new RangeExclude object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAddress

`func (o *RangeExclude) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *RangeExclude) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *RangeExclude) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *RangeExclude) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *RangeExclude) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *RangeExclude) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *RangeExclude) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *RangeExclude) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetComment

`func (o *RangeExclude) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RangeExclude) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RangeExclude) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RangeExclude) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


