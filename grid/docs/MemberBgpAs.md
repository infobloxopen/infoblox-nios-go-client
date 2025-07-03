# MemberBgpAs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**As** | Pointer to **int64** | The number of this autonomous system. | [optional] 
**Keepalive** | Pointer to **int64** | The AS keepalive timer (in seconds). The valid value is from 1 to 21845. | [optional] 
**Holddown** | Pointer to **int64** | The AS holddown timer (in seconds). The valid value is from 3 to 65535. | [optional] 
**Neighbors** | Pointer to [**[]MemberbgpasNeighbors**](MemberbgpasNeighbors.md) | The BGP neighbors for this AS. | [optional] 
**LinkDetect** | Pointer to **bool** | Determines if link detection on the interface is enabled or not. | [optional] 

## Methods

### NewMemberBgpAs

`func NewMemberBgpAs() *MemberBgpAs`

NewMemberBgpAs instantiates a new MemberBgpAs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberBgpAsWithDefaults

`func NewMemberBgpAsWithDefaults() *MemberBgpAs`

NewMemberBgpAsWithDefaults instantiates a new MemberBgpAs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAs

`func (o *MemberBgpAs) GetAs() int64`

GetAs returns the As field if non-nil, zero value otherwise.

### GetAsOk

`func (o *MemberBgpAs) GetAsOk() (*int64, bool)`

GetAsOk returns a tuple with the As field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAs

`func (o *MemberBgpAs) SetAs(v int64)`

SetAs sets As field to given value.

### HasAs

`func (o *MemberBgpAs) HasAs() bool`

HasAs returns a boolean if a field has been set.

### GetKeepalive

`func (o *MemberBgpAs) GetKeepalive() int64`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *MemberBgpAs) GetKeepaliveOk() (*int64, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *MemberBgpAs) SetKeepalive(v int64)`

SetKeepalive sets Keepalive field to given value.

### HasKeepalive

`func (o *MemberBgpAs) HasKeepalive() bool`

HasKeepalive returns a boolean if a field has been set.

### GetHolddown

`func (o *MemberBgpAs) GetHolddown() int64`

GetHolddown returns the Holddown field if non-nil, zero value otherwise.

### GetHolddownOk

`func (o *MemberBgpAs) GetHolddownOk() (*int64, bool)`

GetHolddownOk returns a tuple with the Holddown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolddown

`func (o *MemberBgpAs) SetHolddown(v int64)`

SetHolddown sets Holddown field to given value.

### HasHolddown

`func (o *MemberBgpAs) HasHolddown() bool`

HasHolddown returns a boolean if a field has been set.

### GetNeighbors

`func (o *MemberBgpAs) GetNeighbors() []MemberbgpasNeighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *MemberBgpAs) GetNeighborsOk() (*[]MemberbgpasNeighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *MemberBgpAs) SetNeighbors(v []MemberbgpasNeighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *MemberBgpAs) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetLinkDetect

`func (o *MemberBgpAs) GetLinkDetect() bool`

GetLinkDetect returns the LinkDetect field if non-nil, zero value otherwise.

### GetLinkDetectOk

`func (o *MemberBgpAs) GetLinkDetectOk() (*bool, bool)`

GetLinkDetectOk returns a tuple with the LinkDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDetect

`func (o *MemberBgpAs) SetLinkDetect(v bool)`

SetLinkDetect sets LinkDetect field to given value.

### HasLinkDetect

`func (o *MemberBgpAs) HasLinkDetect() bool`

HasLinkDetect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


