# Filtermac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of a DHCP MAC Filter object. | [optional] 
**DefaultMacAddressExpiration** | Pointer to **int64** | The default MAC expiration time of the DHCP MAC Address Filter object. By default, the MAC address filter never expires; otherwise, it is the absolute interval when the MAC address filter expires. The maximum value can extend up to 4294967295 secs. The minimum value is 60 secs (1 min). | [optional] 
**Disable** | Pointer to **bool** | Determines if the DHCP Fingerprint object is disabled or not. | [optional] 
**EnforceExpirationTimes** | Pointer to **bool** | The flag to enforce MAC address expiration of the DHCP MAC Address Filter object. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LeaseTime** | Pointer to **int64** | The length of time the DHCP server leases an IP address to a client. The lease time applies to hosts that meet the filter criteria. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP MAC Filter object. | [optional] 
**NeverExpires** | Pointer to **bool** | Determines if DHCP MAC Filter never expires or automatically expires. | [optional] 
**Options** | Pointer to [**[]FiltermacOptions**](FiltermacOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**ReservedForInfoblox** | Pointer to **string** | This is reserved for writing comments related to the particular MAC address filter. The length of comment cannot exceed 1024 bytes. | [optional] 

## Methods

### NewFiltermac

`func NewFiltermac() *Filtermac`

NewFiltermac instantiates a new Filtermac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltermacWithDefaults

`func NewFiltermacWithDefaults() *Filtermac`

NewFiltermacWithDefaults instantiates a new Filtermac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Filtermac) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Filtermac) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Filtermac) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Filtermac) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Filtermac) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Filtermac) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Filtermac) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Filtermac) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultMacAddressExpiration

`func (o *Filtermac) GetDefaultMacAddressExpiration() int64`

GetDefaultMacAddressExpiration returns the DefaultMacAddressExpiration field if non-nil, zero value otherwise.

### GetDefaultMacAddressExpirationOk

`func (o *Filtermac) GetDefaultMacAddressExpirationOk() (*int64, bool)`

GetDefaultMacAddressExpirationOk returns a tuple with the DefaultMacAddressExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMacAddressExpiration

`func (o *Filtermac) SetDefaultMacAddressExpiration(v int64)`

SetDefaultMacAddressExpiration sets DefaultMacAddressExpiration field to given value.

### HasDefaultMacAddressExpiration

`func (o *Filtermac) HasDefaultMacAddressExpiration() bool`

HasDefaultMacAddressExpiration returns a boolean if a field has been set.

### GetDisable

`func (o *Filtermac) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Filtermac) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Filtermac) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Filtermac) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnforceExpirationTimes

`func (o *Filtermac) GetEnforceExpirationTimes() bool`

GetEnforceExpirationTimes returns the EnforceExpirationTimes field if non-nil, zero value otherwise.

### GetEnforceExpirationTimesOk

`func (o *Filtermac) GetEnforceExpirationTimesOk() (*bool, bool)`

GetEnforceExpirationTimesOk returns a tuple with the EnforceExpirationTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceExpirationTimes

`func (o *Filtermac) SetEnforceExpirationTimes(v bool)`

SetEnforceExpirationTimes sets EnforceExpirationTimes field to given value.

### HasEnforceExpirationTimes

`func (o *Filtermac) HasEnforceExpirationTimes() bool`

HasEnforceExpirationTimes returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Filtermac) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Filtermac) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Filtermac) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Filtermac) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLeaseTime

`func (o *Filtermac) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *Filtermac) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *Filtermac) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *Filtermac) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *Filtermac) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filtermac) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filtermac) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Filtermac) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeverExpires

`func (o *Filtermac) GetNeverExpires() bool`

GetNeverExpires returns the NeverExpires field if non-nil, zero value otherwise.

### GetNeverExpiresOk

`func (o *Filtermac) GetNeverExpiresOk() (*bool, bool)`

GetNeverExpiresOk returns a tuple with the NeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverExpires

`func (o *Filtermac) SetNeverExpires(v bool)`

SetNeverExpires sets NeverExpires field to given value.

### HasNeverExpires

`func (o *Filtermac) HasNeverExpires() bool`

HasNeverExpires returns a boolean if a field has been set.

### GetOptions

`func (o *Filtermac) GetOptions() []FiltermacOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Filtermac) GetOptionsOk() (*[]FiltermacOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Filtermac) SetOptions(v []FiltermacOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Filtermac) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReservedForInfoblox

`func (o *Filtermac) GetReservedForInfoblox() string`

GetReservedForInfoblox returns the ReservedForInfoblox field if non-nil, zero value otherwise.

### GetReservedForInfobloxOk

`func (o *Filtermac) GetReservedForInfobloxOk() (*string, bool)`

GetReservedForInfobloxOk returns a tuple with the ReservedForInfoblox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedForInfoblox

`func (o *Filtermac) SetReservedForInfoblox(v string)`

SetReservedForInfoblox sets ReservedForInfoblox field to given value.

### HasReservedForInfoblox

`func (o *Filtermac) HasReservedForInfoblox() bool`

HasReservedForInfoblox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


