# Networkuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the Network User. | [optional] 
**AddressObject** | Pointer to **string** | The reference of the IPAM IPv4Address or IPv6Address object describing the address of the Network User. | [optional] [readonly] 
**DataSource** | Pointer to **string** | The Network User data source. | [optional] [readonly] 
**DataSourceIp** | Pointer to **string** | The Network User data source IPv4 Address or IPv6 Address or FQDN address. | [optional] [readonly] 
**Domainname** | Pointer to **string** | The domain name of the Network User. | [optional] 
**FirstSeenTime** | Pointer to **int64** | The first seen timestamp of the Network User. | [optional] 
**Guid** | Pointer to **string** | The group identifier of the Network User. | [optional] 
**LastSeenTime** | Pointer to **int64** | The last seen timestamp of the Network User. | [optional] 
**LastUpdatedTime** | Pointer to **int64** | The last updated timestamp of the Network User. | [optional] 
**LogonId** | Pointer to **string** | The logon identifier of the Network User. | [optional] 
**LogoutTime** | Pointer to **int64** | The logout timestamp of the Network User. | [optional] 
**Name** | Pointer to **string** | The name of the Network User. | [optional] 
**Network** | Pointer to **string** | The reference to the network to which the Network User belongs. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this Network User resides. | [optional] 
**UserStatus** | Pointer to **string** | The status of the Network User. | [optional] [readonly] 

## Methods

### NewNetworkuser

`func NewNetworkuser() *Networkuser`

NewNetworkuser instantiates a new Networkuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkuserWithDefaults

`func NewNetworkuserWithDefaults() *Networkuser`

NewNetworkuserWithDefaults instantiates a new Networkuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Networkuser) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Networkuser) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Networkuser) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Networkuser) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *Networkuser) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Networkuser) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Networkuser) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Networkuser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressObject

`func (o *Networkuser) GetAddressObject() string`

GetAddressObject returns the AddressObject field if non-nil, zero value otherwise.

### GetAddressObjectOk

`func (o *Networkuser) GetAddressObjectOk() (*string, bool)`

GetAddressObjectOk returns a tuple with the AddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressObject

`func (o *Networkuser) SetAddressObject(v string)`

SetAddressObject sets AddressObject field to given value.

### HasAddressObject

`func (o *Networkuser) HasAddressObject() bool`

HasAddressObject returns a boolean if a field has been set.

### GetDataSource

`func (o *Networkuser) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *Networkuser) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *Networkuser) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *Networkuser) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceIp

`func (o *Networkuser) GetDataSourceIp() string`

GetDataSourceIp returns the DataSourceIp field if non-nil, zero value otherwise.

### GetDataSourceIpOk

`func (o *Networkuser) GetDataSourceIpOk() (*string, bool)`

GetDataSourceIpOk returns a tuple with the DataSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceIp

`func (o *Networkuser) SetDataSourceIp(v string)`

SetDataSourceIp sets DataSourceIp field to given value.

### HasDataSourceIp

`func (o *Networkuser) HasDataSourceIp() bool`

HasDataSourceIp returns a boolean if a field has been set.

### GetDomainname

`func (o *Networkuser) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *Networkuser) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *Networkuser) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *Networkuser) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetFirstSeenTime

`func (o *Networkuser) GetFirstSeenTime() int64`

GetFirstSeenTime returns the FirstSeenTime field if non-nil, zero value otherwise.

### GetFirstSeenTimeOk

`func (o *Networkuser) GetFirstSeenTimeOk() (*int64, bool)`

GetFirstSeenTimeOk returns a tuple with the FirstSeenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTime

`func (o *Networkuser) SetFirstSeenTime(v int64)`

SetFirstSeenTime sets FirstSeenTime field to given value.

### HasFirstSeenTime

`func (o *Networkuser) HasFirstSeenTime() bool`

HasFirstSeenTime returns a boolean if a field has been set.

### GetGuid

`func (o *Networkuser) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Networkuser) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Networkuser) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Networkuser) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLastSeenTime

`func (o *Networkuser) GetLastSeenTime() int64`

GetLastSeenTime returns the LastSeenTime field if non-nil, zero value otherwise.

### GetLastSeenTimeOk

`func (o *Networkuser) GetLastSeenTimeOk() (*int64, bool)`

GetLastSeenTimeOk returns a tuple with the LastSeenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTime

`func (o *Networkuser) SetLastSeenTime(v int64)`

SetLastSeenTime sets LastSeenTime field to given value.

### HasLastSeenTime

`func (o *Networkuser) HasLastSeenTime() bool`

HasLastSeenTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Networkuser) GetLastUpdatedTime() int64`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Networkuser) GetLastUpdatedTimeOk() (*int64, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Networkuser) SetLastUpdatedTime(v int64)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Networkuser) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLogonId

`func (o *Networkuser) GetLogonId() string`

GetLogonId returns the LogonId field if non-nil, zero value otherwise.

### GetLogonIdOk

`func (o *Networkuser) GetLogonIdOk() (*string, bool)`

GetLogonIdOk returns a tuple with the LogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonId

`func (o *Networkuser) SetLogonId(v string)`

SetLogonId sets LogonId field to given value.

### HasLogonId

`func (o *Networkuser) HasLogonId() bool`

HasLogonId returns a boolean if a field has been set.

### GetLogoutTime

`func (o *Networkuser) GetLogoutTime() int64`

GetLogoutTime returns the LogoutTime field if non-nil, zero value otherwise.

### GetLogoutTimeOk

`func (o *Networkuser) GetLogoutTimeOk() (*int64, bool)`

GetLogoutTimeOk returns a tuple with the LogoutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutTime

`func (o *Networkuser) SetLogoutTime(v int64)`

SetLogoutTime sets LogoutTime field to given value.

### HasLogoutTime

`func (o *Networkuser) HasLogoutTime() bool`

HasLogoutTime returns a boolean if a field has been set.

### GetName

`func (o *Networkuser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Networkuser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Networkuser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Networkuser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Networkuser) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Networkuser) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Networkuser) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Networkuser) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Networkuser) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Networkuser) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Networkuser) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Networkuser) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetUserStatus

`func (o *Networkuser) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *Networkuser) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *Networkuser) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *Networkuser) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


