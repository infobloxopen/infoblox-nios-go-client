# Ipv6rangeExclude

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAddress** | Pointer to **string** | The IPv4 Address starting address of the exclusion range. | [optional] 
**EndAddress** | Pointer to **string** | The IPv4 Address ending address of the exclusion range. | [optional] 
**Comment** | Pointer to **string** | Comment for the exclusion range; maximum 256 characters. | [optional] 

## Methods

### NewIpv6rangeExclude

`func NewIpv6rangeExclude() *Ipv6rangeExclude`

NewIpv6rangeExclude instantiates a new Ipv6rangeExclude object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangeExcludeWithDefaults

`func NewIpv6rangeExcludeWithDefaults() *Ipv6rangeExclude`

NewIpv6rangeExcludeWithDefaults instantiates a new Ipv6rangeExclude object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAddress

`func (o *Ipv6rangeExclude) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *Ipv6rangeExclude) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *Ipv6rangeExclude) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *Ipv6rangeExclude) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *Ipv6rangeExclude) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *Ipv6rangeExclude) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *Ipv6rangeExclude) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *Ipv6rangeExclude) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6rangeExclude) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6rangeExclude) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6rangeExclude) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6rangeExclude) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


