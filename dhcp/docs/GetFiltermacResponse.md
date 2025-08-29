# GetFiltermacResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP MAC Filter object. | [optional] 
**DefaultMacAddressExpiration** | Pointer to **int64** | The default MAC expiration time of the DHCP MAC Address Filter object. By default, the MAC address filter never expires; otherwise, it is the absolute interval when the MAC address filter expires. The maximum value can extend up to 4294967295 secs. The minimum value is 60 secs (1 min). | [optional] 
**Disable** | Pointer to **bool** | Determines if the DHCP Fingerprint object is disabled or not. | [optional] 
**EnforceExpirationTimes** | Pointer to **bool** | The flag to enforce MAC address expiration of the DHCP MAC Address Filter object. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs+:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrsMinus**](ExtAttrsMinus.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs-:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | The length of time the DHCP server leases an IP address to a client. The lease time applies to hosts that meet the filter criteria. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP MAC Filter object. | [optional] 
**NeverExpires** | Pointer to **bool** | Determines if DHCP MAC Filter never expires or automatically expires. | [optional] 
**Options** | Pointer to [**[]FiltermacOptions**](FiltermacOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**ReservedForInfoblox** | Pointer to **string** | This is reserved for writing comments related to the particular MAC address filter. The length of comment cannot exceed 1024 bytes. | [optional] 
**Result** | Pointer to [**Filtermac**](Filtermac.md) |  | [optional] 

## Methods

### NewGetFiltermacResponse

`func NewGetFiltermacResponse() *GetFiltermacResponse`

NewGetFiltermacResponse instantiates a new GetFiltermacResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiltermacResponseWithDefaults

`func NewGetFiltermacResponseWithDefaults() *GetFiltermacResponse`

NewGetFiltermacResponseWithDefaults instantiates a new GetFiltermacResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFiltermacResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFiltermacResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFiltermacResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFiltermacResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetFiltermacResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetFiltermacResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetFiltermacResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetFiltermacResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultMacAddressExpiration

`func (o *GetFiltermacResponse) GetDefaultMacAddressExpiration() int64`

GetDefaultMacAddressExpiration returns the DefaultMacAddressExpiration field if non-nil, zero value otherwise.

### GetDefaultMacAddressExpirationOk

`func (o *GetFiltermacResponse) GetDefaultMacAddressExpirationOk() (*int64, bool)`

GetDefaultMacAddressExpirationOk returns a tuple with the DefaultMacAddressExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMacAddressExpiration

`func (o *GetFiltermacResponse) SetDefaultMacAddressExpiration(v int64)`

SetDefaultMacAddressExpiration sets DefaultMacAddressExpiration field to given value.

### HasDefaultMacAddressExpiration

`func (o *GetFiltermacResponse) HasDefaultMacAddressExpiration() bool`

HasDefaultMacAddressExpiration returns a boolean if a field has been set.

### GetDisable

`func (o *GetFiltermacResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetFiltermacResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetFiltermacResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetFiltermacResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnforceExpirationTimes

`func (o *GetFiltermacResponse) GetEnforceExpirationTimes() bool`

GetEnforceExpirationTimes returns the EnforceExpirationTimes field if non-nil, zero value otherwise.

### GetEnforceExpirationTimesOk

`func (o *GetFiltermacResponse) GetEnforceExpirationTimesOk() (*bool, bool)`

GetEnforceExpirationTimesOk returns a tuple with the EnforceExpirationTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceExpirationTimes

`func (o *GetFiltermacResponse) SetEnforceExpirationTimes(v bool)`

SetEnforceExpirationTimes sets EnforceExpirationTimes field to given value.

### HasEnforceExpirationTimes

`func (o *GetFiltermacResponse) HasEnforceExpirationTimes() bool`

HasEnforceExpirationTimes returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetFiltermacResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetFiltermacResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetFiltermacResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetFiltermacResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetFiltermacResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetFiltermacResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetFiltermacResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetFiltermacResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetFiltermacResponse) GetExtAttrsMinus() map[string]ExtAttrsMinus`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetFiltermacResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrsMinus, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetFiltermacResponse) SetExtAttrsMinus(v map[string]ExtAttrsMinus)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetFiltermacResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetLeaseTime

`func (o *GetFiltermacResponse) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *GetFiltermacResponse) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *GetFiltermacResponse) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *GetFiltermacResponse) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *GetFiltermacResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFiltermacResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFiltermacResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFiltermacResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeverExpires

`func (o *GetFiltermacResponse) GetNeverExpires() bool`

GetNeverExpires returns the NeverExpires field if non-nil, zero value otherwise.

### GetNeverExpiresOk

`func (o *GetFiltermacResponse) GetNeverExpiresOk() (*bool, bool)`

GetNeverExpiresOk returns a tuple with the NeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverExpires

`func (o *GetFiltermacResponse) SetNeverExpires(v bool)`

SetNeverExpires sets NeverExpires field to given value.

### HasNeverExpires

`func (o *GetFiltermacResponse) HasNeverExpires() bool`

HasNeverExpires returns a boolean if a field has been set.

### GetOptions

`func (o *GetFiltermacResponse) GetOptions() []FiltermacOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetFiltermacResponse) GetOptionsOk() (*[]FiltermacOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetFiltermacResponse) SetOptions(v []FiltermacOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetFiltermacResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReservedForInfoblox

`func (o *GetFiltermacResponse) GetReservedForInfoblox() string`

GetReservedForInfoblox returns the ReservedForInfoblox field if non-nil, zero value otherwise.

### GetReservedForInfobloxOk

`func (o *GetFiltermacResponse) GetReservedForInfobloxOk() (*string, bool)`

GetReservedForInfobloxOk returns a tuple with the ReservedForInfoblox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedForInfoblox

`func (o *GetFiltermacResponse) SetReservedForInfoblox(v string)`

SetReservedForInfoblox sets ReservedForInfoblox field to given value.

### HasReservedForInfoblox

`func (o *GetFiltermacResponse) HasReservedForInfoblox() bool`

HasReservedForInfoblox returns a boolean if a field has been set.

### GetResult

`func (o *GetFiltermacResponse) GetResult() Filtermac`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFiltermacResponse) GetResultOk() (*Filtermac, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFiltermacResponse) SetResult(v Filtermac)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFiltermacResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


