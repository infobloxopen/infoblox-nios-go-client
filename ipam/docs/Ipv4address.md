# Ipv4address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the address; maximum 256 characters. | [optional] [readonly] 
**ConflictTypes** | Pointer to **[]string** | Types of the conflict. | [optional] [readonly] 
**DhcpClientIdentifier** | Pointer to **string** | The client unique identifier. | [optional] [readonly] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**Ipv4addressDiscoveredData**](Ipv4addressDiscoveredData.md) |  | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Fingerprint** | Pointer to **string** | DHCP fingerprint for the address. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | The IP address. | [optional] [readonly] 
**IsConflict** | Pointer to **bool** | If set to True, the IP address has either a MAC address conflict or a DHCP lease conflict detected through a network discovery. | [optional] [readonly] 
**IsInvalidMac** | Pointer to **bool** | This flag reflects whether the MAC address for this address is invalid. | [optional] [readonly] 
**LeaseState** | Pointer to **string** | The lease state of the address. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**Ipv4addressMsAdUserData**](Ipv4addressMsAdUserData.md) |  | [optional] 
**Names** | Pointer to **[]string** | The DNS names. For example, if the IP address belongs to a host record, this field contains the hostname. This field supports both single and array search. | [optional] [readonly] 
**Network** | Pointer to **string** | The network to which this address belongs, in FQDN/CIDR format. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view. | [optional] [readonly] 
**Objects** | Pointer to **string** | The objects associated with the IP address. | [optional] [readonly] 
**ReservedPort** | Pointer to **string** | The reserved port for the address. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the address. | [optional] [readonly] 
**Types** | Pointer to **[]string** | The types of associated objects. This field supports both single and array search. | [optional] [readonly] 
**Usage** | Pointer to **[]string** | Indicates whether the IP address is configured for DNS or DHCP. This field supports both single and array search. | [optional] [readonly] 
**Username** | Pointer to **string** | The name of the user who created or modified the record. | [optional] [readonly] 

## Methods

### NewIpv4address

`func NewIpv4address() *Ipv4address`

NewIpv4address instantiates a new Ipv4address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4addressWithDefaults

`func NewIpv4addressWithDefaults() *Ipv4address`

NewIpv4addressWithDefaults instantiates a new Ipv4address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv4address) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv4address) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv4address) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv4address) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Ipv4address) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv4address) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv4address) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv4address) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConflictTypes

`func (o *Ipv4address) GetConflictTypes() []string`

GetConflictTypes returns the ConflictTypes field if non-nil, zero value otherwise.

### GetConflictTypesOk

`func (o *Ipv4address) GetConflictTypesOk() (*[]string, bool)`

GetConflictTypesOk returns a tuple with the ConflictTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictTypes

`func (o *Ipv4address) SetConflictTypes(v []string)`

SetConflictTypes sets ConflictTypes field to given value.

### HasConflictTypes

`func (o *Ipv4address) HasConflictTypes() bool`

HasConflictTypes returns a boolean if a field has been set.

### GetDhcpClientIdentifier

`func (o *Ipv4address) GetDhcpClientIdentifier() string`

GetDhcpClientIdentifier returns the DhcpClientIdentifier field if non-nil, zero value otherwise.

### GetDhcpClientIdentifierOk

`func (o *Ipv4address) GetDhcpClientIdentifierOk() (*string, bool)`

GetDhcpClientIdentifierOk returns a tuple with the DhcpClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpClientIdentifier

`func (o *Ipv4address) SetDhcpClientIdentifier(v string)`

SetDhcpClientIdentifier sets DhcpClientIdentifier field to given value.

### HasDhcpClientIdentifier

`func (o *Ipv4address) HasDhcpClientIdentifier() bool`

HasDhcpClientIdentifier returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Ipv4address) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Ipv4address) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Ipv4address) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Ipv4address) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *Ipv4address) GetDiscoveredData() Ipv4addressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *Ipv4address) GetDiscoveredDataOk() (*Ipv4addressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *Ipv4address) SetDiscoveredData(v Ipv4addressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *Ipv4address) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Ipv4address) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Ipv4address) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Ipv4address) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Ipv4address) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Ipv4address) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Ipv4address) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Ipv4address) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Ipv4address) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Ipv4address) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Ipv4address) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Ipv4address) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Ipv4address) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFingerprint

`func (o *Ipv4address) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Ipv4address) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Ipv4address) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Ipv4address) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIpAddress

`func (o *Ipv4address) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Ipv4address) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Ipv4address) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Ipv4address) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsConflict

`func (o *Ipv4address) GetIsConflict() bool`

GetIsConflict returns the IsConflict field if non-nil, zero value otherwise.

### GetIsConflictOk

`func (o *Ipv4address) GetIsConflictOk() (*bool, bool)`

GetIsConflictOk returns a tuple with the IsConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConflict

`func (o *Ipv4address) SetIsConflict(v bool)`

SetIsConflict sets IsConflict field to given value.

### HasIsConflict

`func (o *Ipv4address) HasIsConflict() bool`

HasIsConflict returns a boolean if a field has been set.

### GetIsInvalidMac

`func (o *Ipv4address) GetIsInvalidMac() bool`

GetIsInvalidMac returns the IsInvalidMac field if non-nil, zero value otherwise.

### GetIsInvalidMacOk

`func (o *Ipv4address) GetIsInvalidMacOk() (*bool, bool)`

GetIsInvalidMacOk returns a tuple with the IsInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalidMac

`func (o *Ipv4address) SetIsInvalidMac(v bool)`

SetIsInvalidMac sets IsInvalidMac field to given value.

### HasIsInvalidMac

`func (o *Ipv4address) HasIsInvalidMac() bool`

HasIsInvalidMac returns a boolean if a field has been set.

### GetLeaseState

`func (o *Ipv4address) GetLeaseState() string`

GetLeaseState returns the LeaseState field if non-nil, zero value otherwise.

### GetLeaseStateOk

`func (o *Ipv4address) GetLeaseStateOk() (*string, bool)`

GetLeaseStateOk returns a tuple with the LeaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseState

`func (o *Ipv4address) SetLeaseState(v string)`

SetLeaseState sets LeaseState field to given value.

### HasLeaseState

`func (o *Ipv4address) HasLeaseState() bool`

HasLeaseState returns a boolean if a field has been set.

### GetMacAddress

`func (o *Ipv4address) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Ipv4address) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Ipv4address) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Ipv4address) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Ipv4address) GetMsAdUserData() Ipv4addressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Ipv4address) GetMsAdUserDataOk() (*Ipv4addressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Ipv4address) SetMsAdUserData(v Ipv4addressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Ipv4address) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNames

`func (o *Ipv4address) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Ipv4address) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Ipv4address) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *Ipv4address) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetNetwork

`func (o *Ipv4address) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ipv4address) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ipv4address) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ipv4address) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Ipv4address) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Ipv4address) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Ipv4address) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Ipv4address) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetObjects

`func (o *Ipv4address) GetObjects() string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Ipv4address) GetObjectsOk() (*string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Ipv4address) SetObjects(v string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *Ipv4address) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetReservedPort

`func (o *Ipv4address) GetReservedPort() string`

GetReservedPort returns the ReservedPort field if non-nil, zero value otherwise.

### GetReservedPortOk

`func (o *Ipv4address) GetReservedPortOk() (*string, bool)`

GetReservedPortOk returns a tuple with the ReservedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPort

`func (o *Ipv4address) SetReservedPort(v string)`

SetReservedPort sets ReservedPort field to given value.

### HasReservedPort

`func (o *Ipv4address) HasReservedPort() bool`

HasReservedPort returns a boolean if a field has been set.

### GetStatus

`func (o *Ipv4address) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ipv4address) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ipv4address) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ipv4address) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *Ipv4address) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Ipv4address) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Ipv4address) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *Ipv4address) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUsage

`func (o *Ipv4address) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Ipv4address) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Ipv4address) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Ipv4address) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsername

`func (o *Ipv4address) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Ipv4address) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Ipv4address) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Ipv4address) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


