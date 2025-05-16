# Ipv6address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the address; maximum 256 characters. | [optional] [readonly] 
**ConflictTypes** | Pointer to **[]string** | Types of the conflict. | [optional] [readonly] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**Ipv6addressDiscoveredData**](Ipv6addressDiscoveredData.md) |  | [optional] 
**Duid** | Pointer to **string** | DHCPv6 Unique Identifier (DUID) of the address object. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Fingerprint** | Pointer to **string** | DHCP fingerprint for the address. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IPv6 addresses of the address object. | [optional] [readonly] 
**IsConflict** | Pointer to **bool** | IP address has either a duid conflict or a DHCP lease conflict detected through a network discovery. | [optional] [readonly] 
**LeaseState** | Pointer to **string** | The lease state of the address. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**Ipv6addressMsAdUserData**](Ipv6addressMsAdUserData.md) |  | [optional] 
**Names** | Pointer to **[]string** | The DNS names. For example, if the IP address belongs to a host record, this field contains the hostname. This field supports both single and array search. | [optional] [readonly] 
**Network** | Pointer to **string** | The network to which this address belongs, in FQDN/CIDR format. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view. | [optional] [readonly] 
**Objects** | Pointer to **string** | The objects associated with the IP address. | [optional] [readonly] 
**ReservedPort** | Pointer to **string** | The reserved port for the address. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the address. | [optional] [readonly] 
**Types** | Pointer to **[]string** | The types of associated objects. This field supports both single and array search. | [optional] [readonly] 
**Usage** | Pointer to **[]string** | Indicates whether the IP address is configured for DNS or DHCP. This field supports both single and array search. | [optional] [readonly] 

## Methods

### NewIpv6address

`func NewIpv6address() *Ipv6address`

NewIpv6address instantiates a new Ipv6address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6addressWithDefaults

`func NewIpv6addressWithDefaults() *Ipv6address`

NewIpv6addressWithDefaults instantiates a new Ipv6address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6address) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6address) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6address) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6address) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6address) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6address) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6address) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6address) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConflictTypes

`func (o *Ipv6address) GetConflictTypes() []string`

GetConflictTypes returns the ConflictTypes field if non-nil, zero value otherwise.

### GetConflictTypesOk

`func (o *Ipv6address) GetConflictTypesOk() (*[]string, bool)`

GetConflictTypesOk returns a tuple with the ConflictTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictTypes

`func (o *Ipv6address) SetConflictTypes(v []string)`

SetConflictTypes sets ConflictTypes field to given value.

### HasConflictTypes

`func (o *Ipv6address) HasConflictTypes() bool`

HasConflictTypes returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *Ipv6address) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *Ipv6address) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *Ipv6address) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *Ipv6address) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *Ipv6address) GetDiscoveredData() Ipv6addressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *Ipv6address) GetDiscoveredDataOk() (*Ipv6addressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *Ipv6address) SetDiscoveredData(v Ipv6addressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *Ipv6address) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDuid

`func (o *Ipv6address) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *Ipv6address) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *Ipv6address) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *Ipv6address) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetExtattrs

`func (o *Ipv6address) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Ipv6address) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Ipv6address) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Ipv6address) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFingerprint

`func (o *Ipv6address) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Ipv6address) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Ipv6address) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Ipv6address) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIpAddress

`func (o *Ipv6address) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Ipv6address) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Ipv6address) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Ipv6address) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsConflict

`func (o *Ipv6address) GetIsConflict() bool`

GetIsConflict returns the IsConflict field if non-nil, zero value otherwise.

### GetIsConflictOk

`func (o *Ipv6address) GetIsConflictOk() (*bool, bool)`

GetIsConflictOk returns a tuple with the IsConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConflict

`func (o *Ipv6address) SetIsConflict(v bool)`

SetIsConflict sets IsConflict field to given value.

### HasIsConflict

`func (o *Ipv6address) HasIsConflict() bool`

HasIsConflict returns a boolean if a field has been set.

### GetLeaseState

`func (o *Ipv6address) GetLeaseState() string`

GetLeaseState returns the LeaseState field if non-nil, zero value otherwise.

### GetLeaseStateOk

`func (o *Ipv6address) GetLeaseStateOk() (*string, bool)`

GetLeaseStateOk returns a tuple with the LeaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseState

`func (o *Ipv6address) SetLeaseState(v string)`

SetLeaseState sets LeaseState field to given value.

### HasLeaseState

`func (o *Ipv6address) HasLeaseState() bool`

HasLeaseState returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Ipv6address) GetMsAdUserData() Ipv6addressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Ipv6address) GetMsAdUserDataOk() (*Ipv6addressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Ipv6address) SetMsAdUserData(v Ipv6addressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Ipv6address) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNames

`func (o *Ipv6address) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Ipv6address) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Ipv6address) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *Ipv6address) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetNetwork

`func (o *Ipv6address) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Ipv6address) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Ipv6address) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Ipv6address) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Ipv6address) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Ipv6address) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Ipv6address) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Ipv6address) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetObjects

`func (o *Ipv6address) GetObjects() string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Ipv6address) GetObjectsOk() (*string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Ipv6address) SetObjects(v string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *Ipv6address) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetReservedPort

`func (o *Ipv6address) GetReservedPort() string`

GetReservedPort returns the ReservedPort field if non-nil, zero value otherwise.

### GetReservedPortOk

`func (o *Ipv6address) GetReservedPortOk() (*string, bool)`

GetReservedPortOk returns a tuple with the ReservedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPort

`func (o *Ipv6address) SetReservedPort(v string)`

SetReservedPort sets ReservedPort field to given value.

### HasReservedPort

`func (o *Ipv6address) HasReservedPort() bool`

HasReservedPort returns a boolean if a field has been set.

### GetStatus

`func (o *Ipv6address) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ipv6address) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ipv6address) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ipv6address) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *Ipv6address) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Ipv6address) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Ipv6address) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *Ipv6address) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUsage

`func (o *Ipv6address) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Ipv6address) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Ipv6address) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Ipv6address) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


