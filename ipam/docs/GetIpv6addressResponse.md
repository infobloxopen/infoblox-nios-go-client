# GetIpv6addressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the address; maximum 256 characters. | [optional] [readonly] 
**ConflictTypes** | Pointer to **[]string** | Types of the conflict. | [optional] [readonly] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**Ipv6addressDiscoveredData**](Ipv6addressDiscoveredData.md) |  | [optional] 
**Duid** | Pointer to **string** | DHCPv6 Unique Identifier (DUID) of the address object. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**Ipv6address**](Ipv6address.md) |  | [optional] 

## Methods

### NewGetIpv6addressResponse

`func NewGetIpv6addressResponse() *GetIpv6addressResponse`

NewGetIpv6addressResponse instantiates a new GetIpv6addressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6addressResponseWithDefaults

`func NewGetIpv6addressResponseWithDefaults() *GetIpv6addressResponse`

NewGetIpv6addressResponseWithDefaults instantiates a new GetIpv6addressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6addressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6addressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6addressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6addressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6addressResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6addressResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6addressResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6addressResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConflictTypes

`func (o *GetIpv6addressResponse) GetConflictTypes() []string`

GetConflictTypes returns the ConflictTypes field if non-nil, zero value otherwise.

### GetConflictTypesOk

`func (o *GetIpv6addressResponse) GetConflictTypesOk() (*[]string, bool)`

GetConflictTypesOk returns a tuple with the ConflictTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictTypes

`func (o *GetIpv6addressResponse) SetConflictTypes(v []string)`

SetConflictTypes sets ConflictTypes field to given value.

### HasConflictTypes

`func (o *GetIpv6addressResponse) HasConflictTypes() bool`

HasConflictTypes returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetIpv6addressResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetIpv6addressResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetIpv6addressResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetIpv6addressResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *GetIpv6addressResponse) GetDiscoveredData() Ipv6addressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *GetIpv6addressResponse) GetDiscoveredDataOk() (*Ipv6addressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *GetIpv6addressResponse) SetDiscoveredData(v Ipv6addressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *GetIpv6addressResponse) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDuid

`func (o *GetIpv6addressResponse) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *GetIpv6addressResponse) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *GetIpv6addressResponse) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *GetIpv6addressResponse) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetIpv6addressResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetIpv6addressResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetIpv6addressResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetIpv6addressResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetIpv6addressResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetIpv6addressResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetIpv6addressResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetIpv6addressResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetIpv6addressResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetIpv6addressResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetIpv6addressResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetIpv6addressResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFingerprint

`func (o *GetIpv6addressResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GetIpv6addressResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GetIpv6addressResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GetIpv6addressResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetIpv6addressResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetIpv6addressResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetIpv6addressResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetIpv6addressResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsConflict

`func (o *GetIpv6addressResponse) GetIsConflict() bool`

GetIsConflict returns the IsConflict field if non-nil, zero value otherwise.

### GetIsConflictOk

`func (o *GetIpv6addressResponse) GetIsConflictOk() (*bool, bool)`

GetIsConflictOk returns a tuple with the IsConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConflict

`func (o *GetIpv6addressResponse) SetIsConflict(v bool)`

SetIsConflict sets IsConflict field to given value.

### HasIsConflict

`func (o *GetIpv6addressResponse) HasIsConflict() bool`

HasIsConflict returns a boolean if a field has been set.

### GetLeaseState

`func (o *GetIpv6addressResponse) GetLeaseState() string`

GetLeaseState returns the LeaseState field if non-nil, zero value otherwise.

### GetLeaseStateOk

`func (o *GetIpv6addressResponse) GetLeaseStateOk() (*string, bool)`

GetLeaseStateOk returns a tuple with the LeaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseState

`func (o *GetIpv6addressResponse) SetLeaseState(v string)`

SetLeaseState sets LeaseState field to given value.

### HasLeaseState

`func (o *GetIpv6addressResponse) HasLeaseState() bool`

HasLeaseState returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetIpv6addressResponse) GetMsAdUserData() Ipv6addressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetIpv6addressResponse) GetMsAdUserDataOk() (*Ipv6addressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetIpv6addressResponse) SetMsAdUserData(v Ipv6addressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetIpv6addressResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNames

`func (o *GetIpv6addressResponse) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GetIpv6addressResponse) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GetIpv6addressResponse) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *GetIpv6addressResponse) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetNetwork

`func (o *GetIpv6addressResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetIpv6addressResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetIpv6addressResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetIpv6addressResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetIpv6addressResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetIpv6addressResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetIpv6addressResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetIpv6addressResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetObjects

`func (o *GetIpv6addressResponse) GetObjects() string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetIpv6addressResponse) GetObjectsOk() (*string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetIpv6addressResponse) SetObjects(v string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *GetIpv6addressResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetReservedPort

`func (o *GetIpv6addressResponse) GetReservedPort() string`

GetReservedPort returns the ReservedPort field if non-nil, zero value otherwise.

### GetReservedPortOk

`func (o *GetIpv6addressResponse) GetReservedPortOk() (*string, bool)`

GetReservedPortOk returns a tuple with the ReservedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPort

`func (o *GetIpv6addressResponse) SetReservedPort(v string)`

SetReservedPort sets ReservedPort field to given value.

### HasReservedPort

`func (o *GetIpv6addressResponse) HasReservedPort() bool`

HasReservedPort returns a boolean if a field has been set.

### GetStatus

`func (o *GetIpv6addressResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIpv6addressResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIpv6addressResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIpv6addressResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *GetIpv6addressResponse) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GetIpv6addressResponse) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GetIpv6addressResponse) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *GetIpv6addressResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUsage

`func (o *GetIpv6addressResponse) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetIpv6addressResponse) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetIpv6addressResponse) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetIpv6addressResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6addressResponse) GetResult() Ipv6address`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6addressResponse) GetResultOk() (*Ipv6address, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6addressResponse) SetResult(v Ipv6address)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6addressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


