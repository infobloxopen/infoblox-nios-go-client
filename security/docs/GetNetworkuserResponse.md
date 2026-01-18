# GetNetworkuserResponse

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
**Result** | Pointer to [**Networkuser**](Networkuser.md) |  | [optional] 

## Methods

### NewGetNetworkuserResponse

`func NewGetNetworkuserResponse() *GetNetworkuserResponse`

NewGetNetworkuserResponse instantiates a new GetNetworkuserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkuserResponseWithDefaults

`func NewGetNetworkuserResponseWithDefaults() *GetNetworkuserResponse`

NewGetNetworkuserResponseWithDefaults instantiates a new GetNetworkuserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNetworkuserResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNetworkuserResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNetworkuserResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNetworkuserResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetNetworkuserResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetNetworkuserResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetNetworkuserResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetNetworkuserResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressObject

`func (o *GetNetworkuserResponse) GetAddressObject() string`

GetAddressObject returns the AddressObject field if non-nil, zero value otherwise.

### GetAddressObjectOk

`func (o *GetNetworkuserResponse) GetAddressObjectOk() (*string, bool)`

GetAddressObjectOk returns a tuple with the AddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressObject

`func (o *GetNetworkuserResponse) SetAddressObject(v string)`

SetAddressObject sets AddressObject field to given value.

### HasAddressObject

`func (o *GetNetworkuserResponse) HasAddressObject() bool`

HasAddressObject returns a boolean if a field has been set.

### GetDataSource

`func (o *GetNetworkuserResponse) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *GetNetworkuserResponse) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *GetNetworkuserResponse) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *GetNetworkuserResponse) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceIp

`func (o *GetNetworkuserResponse) GetDataSourceIp() string`

GetDataSourceIp returns the DataSourceIp field if non-nil, zero value otherwise.

### GetDataSourceIpOk

`func (o *GetNetworkuserResponse) GetDataSourceIpOk() (*string, bool)`

GetDataSourceIpOk returns a tuple with the DataSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceIp

`func (o *GetNetworkuserResponse) SetDataSourceIp(v string)`

SetDataSourceIp sets DataSourceIp field to given value.

### HasDataSourceIp

`func (o *GetNetworkuserResponse) HasDataSourceIp() bool`

HasDataSourceIp returns a boolean if a field has been set.

### GetDomainname

`func (o *GetNetworkuserResponse) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *GetNetworkuserResponse) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *GetNetworkuserResponse) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *GetNetworkuserResponse) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetFirstSeenTime

`func (o *GetNetworkuserResponse) GetFirstSeenTime() int64`

GetFirstSeenTime returns the FirstSeenTime field if non-nil, zero value otherwise.

### GetFirstSeenTimeOk

`func (o *GetNetworkuserResponse) GetFirstSeenTimeOk() (*int64, bool)`

GetFirstSeenTimeOk returns a tuple with the FirstSeenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTime

`func (o *GetNetworkuserResponse) SetFirstSeenTime(v int64)`

SetFirstSeenTime sets FirstSeenTime field to given value.

### HasFirstSeenTime

`func (o *GetNetworkuserResponse) HasFirstSeenTime() bool`

HasFirstSeenTime returns a boolean if a field has been set.

### GetGuid

`func (o *GetNetworkuserResponse) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *GetNetworkuserResponse) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *GetNetworkuserResponse) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *GetNetworkuserResponse) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLastSeenTime

`func (o *GetNetworkuserResponse) GetLastSeenTime() int64`

GetLastSeenTime returns the LastSeenTime field if non-nil, zero value otherwise.

### GetLastSeenTimeOk

`func (o *GetNetworkuserResponse) GetLastSeenTimeOk() (*int64, bool)`

GetLastSeenTimeOk returns a tuple with the LastSeenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTime

`func (o *GetNetworkuserResponse) SetLastSeenTime(v int64)`

SetLastSeenTime sets LastSeenTime field to given value.

### HasLastSeenTime

`func (o *GetNetworkuserResponse) HasLastSeenTime() bool`

HasLastSeenTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *GetNetworkuserResponse) GetLastUpdatedTime() int64`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GetNetworkuserResponse) GetLastUpdatedTimeOk() (*int64, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GetNetworkuserResponse) SetLastUpdatedTime(v int64)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *GetNetworkuserResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLogonId

`func (o *GetNetworkuserResponse) GetLogonId() string`

GetLogonId returns the LogonId field if non-nil, zero value otherwise.

### GetLogonIdOk

`func (o *GetNetworkuserResponse) GetLogonIdOk() (*string, bool)`

GetLogonIdOk returns a tuple with the LogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonId

`func (o *GetNetworkuserResponse) SetLogonId(v string)`

SetLogonId sets LogonId field to given value.

### HasLogonId

`func (o *GetNetworkuserResponse) HasLogonId() bool`

HasLogonId returns a boolean if a field has been set.

### GetLogoutTime

`func (o *GetNetworkuserResponse) GetLogoutTime() int64`

GetLogoutTime returns the LogoutTime field if non-nil, zero value otherwise.

### GetLogoutTimeOk

`func (o *GetNetworkuserResponse) GetLogoutTimeOk() (*int64, bool)`

GetLogoutTimeOk returns a tuple with the LogoutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutTime

`func (o *GetNetworkuserResponse) SetLogoutTime(v int64)`

SetLogoutTime sets LogoutTime field to given value.

### HasLogoutTime

`func (o *GetNetworkuserResponse) HasLogoutTime() bool`

HasLogoutTime returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkuserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkuserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkuserResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkuserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *GetNetworkuserResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetNetworkuserResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetNetworkuserResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetNetworkuserResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetNetworkuserResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetNetworkuserResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetNetworkuserResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetNetworkuserResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetUserStatus

`func (o *GetNetworkuserResponse) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *GetNetworkuserResponse) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *GetNetworkuserResponse) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *GetNetworkuserResponse) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetNetworkuserResponse) GetResult() Networkuser`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNetworkuserResponse) GetResultOk() (*Networkuser, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNetworkuserResponse) SetResult(v Networkuser)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNetworkuserResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


