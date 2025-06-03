# GetIpv4addressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | Comment for the address; maximum 256 characters. | [optional] [readonly] 
**ConflictTypes** | Pointer to **[]string** | Types of the conflict. | [optional] [readonly] 
**DhcpClientIdentifier** | Pointer to **string** | The client unique identifier. | [optional] [readonly] 
**DiscoverNowStatus** | Pointer to **string** | Discover now status for this address. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**Ipv4addressDiscoveredData**](Ipv4addressDiscoveredData.md) |  | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**Ipv4address**](Ipv4address.md) |  | [optional] 

## Methods

### NewGetIpv4addressResponse

`func NewGetIpv4addressResponse() *GetIpv4addressResponse`

NewGetIpv4addressResponse instantiates a new GetIpv4addressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv4addressResponseWithDefaults

`func NewGetIpv4addressResponseWithDefaults() *GetIpv4addressResponse`

NewGetIpv4addressResponseWithDefaults instantiates a new GetIpv4addressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv4addressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv4addressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv4addressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv4addressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv4addressResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv4addressResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv4addressResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv4addressResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConflictTypes

`func (o *GetIpv4addressResponse) GetConflictTypes() []string`

GetConflictTypes returns the ConflictTypes field if non-nil, zero value otherwise.

### GetConflictTypesOk

`func (o *GetIpv4addressResponse) GetConflictTypesOk() (*[]string, bool)`

GetConflictTypesOk returns a tuple with the ConflictTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictTypes

`func (o *GetIpv4addressResponse) SetConflictTypes(v []string)`

SetConflictTypes sets ConflictTypes field to given value.

### HasConflictTypes

`func (o *GetIpv4addressResponse) HasConflictTypes() bool`

HasConflictTypes returns a boolean if a field has been set.

### GetDhcpClientIdentifier

`func (o *GetIpv4addressResponse) GetDhcpClientIdentifier() string`

GetDhcpClientIdentifier returns the DhcpClientIdentifier field if non-nil, zero value otherwise.

### GetDhcpClientIdentifierOk

`func (o *GetIpv4addressResponse) GetDhcpClientIdentifierOk() (*string, bool)`

GetDhcpClientIdentifierOk returns a tuple with the DhcpClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpClientIdentifier

`func (o *GetIpv4addressResponse) SetDhcpClientIdentifier(v string)`

SetDhcpClientIdentifier sets DhcpClientIdentifier field to given value.

### HasDhcpClientIdentifier

`func (o *GetIpv4addressResponse) HasDhcpClientIdentifier() bool`

HasDhcpClientIdentifier returns a boolean if a field has been set.

### GetDiscoverNowStatus

`func (o *GetIpv4addressResponse) GetDiscoverNowStatus() string`

GetDiscoverNowStatus returns the DiscoverNowStatus field if non-nil, zero value otherwise.

### GetDiscoverNowStatusOk

`func (o *GetIpv4addressResponse) GetDiscoverNowStatusOk() (*string, bool)`

GetDiscoverNowStatusOk returns a tuple with the DiscoverNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNowStatus

`func (o *GetIpv4addressResponse) SetDiscoverNowStatus(v string)`

SetDiscoverNowStatus sets DiscoverNowStatus field to given value.

### HasDiscoverNowStatus

`func (o *GetIpv4addressResponse) HasDiscoverNowStatus() bool`

HasDiscoverNowStatus returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *GetIpv4addressResponse) GetDiscoveredData() Ipv4addressDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *GetIpv4addressResponse) GetDiscoveredDataOk() (*Ipv4addressDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *GetIpv4addressResponse) SetDiscoveredData(v Ipv4addressDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *GetIpv4addressResponse) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetIpv4addressResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetIpv4addressResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetIpv4addressResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetIpv4addressResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFingerprint

`func (o *GetIpv4addressResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GetIpv4addressResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GetIpv4addressResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GetIpv4addressResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetIpv4addressResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetIpv4addressResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetIpv4addressResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetIpv4addressResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsConflict

`func (o *GetIpv4addressResponse) GetIsConflict() bool`

GetIsConflict returns the IsConflict field if non-nil, zero value otherwise.

### GetIsConflictOk

`func (o *GetIpv4addressResponse) GetIsConflictOk() (*bool, bool)`

GetIsConflictOk returns a tuple with the IsConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConflict

`func (o *GetIpv4addressResponse) SetIsConflict(v bool)`

SetIsConflict sets IsConflict field to given value.

### HasIsConflict

`func (o *GetIpv4addressResponse) HasIsConflict() bool`

HasIsConflict returns a boolean if a field has been set.

### GetIsInvalidMac

`func (o *GetIpv4addressResponse) GetIsInvalidMac() bool`

GetIsInvalidMac returns the IsInvalidMac field if non-nil, zero value otherwise.

### GetIsInvalidMacOk

`func (o *GetIpv4addressResponse) GetIsInvalidMacOk() (*bool, bool)`

GetIsInvalidMacOk returns a tuple with the IsInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalidMac

`func (o *GetIpv4addressResponse) SetIsInvalidMac(v bool)`

SetIsInvalidMac sets IsInvalidMac field to given value.

### HasIsInvalidMac

`func (o *GetIpv4addressResponse) HasIsInvalidMac() bool`

HasIsInvalidMac returns a boolean if a field has been set.

### GetLeaseState

`func (o *GetIpv4addressResponse) GetLeaseState() string`

GetLeaseState returns the LeaseState field if non-nil, zero value otherwise.

### GetLeaseStateOk

`func (o *GetIpv4addressResponse) GetLeaseStateOk() (*string, bool)`

GetLeaseStateOk returns a tuple with the LeaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseState

`func (o *GetIpv4addressResponse) SetLeaseState(v string)`

SetLeaseState sets LeaseState field to given value.

### HasLeaseState

`func (o *GetIpv4addressResponse) HasLeaseState() bool`

HasLeaseState returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetIpv4addressResponse) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetIpv4addressResponse) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetIpv4addressResponse) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetIpv4addressResponse) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetIpv4addressResponse) GetMsAdUserData() Ipv4addressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetIpv4addressResponse) GetMsAdUserDataOk() (*Ipv4addressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetIpv4addressResponse) SetMsAdUserData(v Ipv4addressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetIpv4addressResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNames

`func (o *GetIpv4addressResponse) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GetIpv4addressResponse) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GetIpv4addressResponse) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *GetIpv4addressResponse) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetNetwork

`func (o *GetIpv4addressResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetIpv4addressResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetIpv4addressResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetIpv4addressResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetIpv4addressResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetIpv4addressResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetIpv4addressResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetIpv4addressResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetObjects

`func (o *GetIpv4addressResponse) GetObjects() string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetIpv4addressResponse) GetObjectsOk() (*string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetIpv4addressResponse) SetObjects(v string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *GetIpv4addressResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetReservedPort

`func (o *GetIpv4addressResponse) GetReservedPort() string`

GetReservedPort returns the ReservedPort field if non-nil, zero value otherwise.

### GetReservedPortOk

`func (o *GetIpv4addressResponse) GetReservedPortOk() (*string, bool)`

GetReservedPortOk returns a tuple with the ReservedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPort

`func (o *GetIpv4addressResponse) SetReservedPort(v string)`

SetReservedPort sets ReservedPort field to given value.

### HasReservedPort

`func (o *GetIpv4addressResponse) HasReservedPort() bool`

HasReservedPort returns a boolean if a field has been set.

### GetStatus

`func (o *GetIpv4addressResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIpv4addressResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIpv4addressResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIpv4addressResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTypes

`func (o *GetIpv4addressResponse) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GetIpv4addressResponse) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GetIpv4addressResponse) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *GetIpv4addressResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUsage

`func (o *GetIpv4addressResponse) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetIpv4addressResponse) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetIpv4addressResponse) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetIpv4addressResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsername

`func (o *GetIpv4addressResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetIpv4addressResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetIpv4addressResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetIpv4addressResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv4addressResponse) GetResult() Ipv4address`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv4addressResponse) GetResultOk() (*Ipv4address, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv4addressResponse) SetResult(v Ipv4address)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv4addressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


