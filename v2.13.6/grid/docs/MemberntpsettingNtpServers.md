# MemberntpsettingNtpServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The NTP server IP address or FQDN. | [optional] 
**EnableAuthentication** | Pointer to **bool** | Determines whether the NTP authentication is enabled. | [optional] 
**NtpKeyNumber** | Pointer to **int64** | The NTP authentication key number. | [optional] 
**Preferred** | Pointer to **bool** | Determines whether the NTP server is a preferred one or not. | [optional] 
**Burst** | Pointer to **bool** | Determines whether the BURST operation mode is enabled. In BURST operating mode, when the external server is reachable and a valid source of synchronization is available, NTP sends a burst of 8 packets with a 2 second interval between packets. | [optional] 
**Iburst** | Pointer to **bool** | Determines whether the IBURST operation mode is enabled. In IBURST operating mode, when the external server is unreachable, NTP server sends a burst of 8 packets with a 2 second interval between packets. | [optional] 

## Methods

### NewMemberntpsettingNtpServers

`func NewMemberntpsettingNtpServers() *MemberntpsettingNtpServers`

NewMemberntpsettingNtpServers instantiates a new MemberntpsettingNtpServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberntpsettingNtpServersWithDefaults

`func NewMemberntpsettingNtpServersWithDefaults() *MemberntpsettingNtpServers`

NewMemberntpsettingNtpServersWithDefaults instantiates a new MemberntpsettingNtpServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberntpsettingNtpServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberntpsettingNtpServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberntpsettingNtpServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberntpsettingNtpServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnableAuthentication

`func (o *MemberntpsettingNtpServers) GetEnableAuthentication() bool`

GetEnableAuthentication returns the EnableAuthentication field if non-nil, zero value otherwise.

### GetEnableAuthenticationOk

`func (o *MemberntpsettingNtpServers) GetEnableAuthenticationOk() (*bool, bool)`

GetEnableAuthenticationOk returns a tuple with the EnableAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuthentication

`func (o *MemberntpsettingNtpServers) SetEnableAuthentication(v bool)`

SetEnableAuthentication sets EnableAuthentication field to given value.

### HasEnableAuthentication

`func (o *MemberntpsettingNtpServers) HasEnableAuthentication() bool`

HasEnableAuthentication returns a boolean if a field has been set.

### GetNtpKeyNumber

`func (o *MemberntpsettingNtpServers) GetNtpKeyNumber() int64`

GetNtpKeyNumber returns the NtpKeyNumber field if non-nil, zero value otherwise.

### GetNtpKeyNumberOk

`func (o *MemberntpsettingNtpServers) GetNtpKeyNumberOk() (*int64, bool)`

GetNtpKeyNumberOk returns a tuple with the NtpKeyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpKeyNumber

`func (o *MemberntpsettingNtpServers) SetNtpKeyNumber(v int64)`

SetNtpKeyNumber sets NtpKeyNumber field to given value.

### HasNtpKeyNumber

`func (o *MemberntpsettingNtpServers) HasNtpKeyNumber() bool`

HasNtpKeyNumber returns a boolean if a field has been set.

### GetPreferred

`func (o *MemberntpsettingNtpServers) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *MemberntpsettingNtpServers) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *MemberntpsettingNtpServers) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *MemberntpsettingNtpServers) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### GetBurst

`func (o *MemberntpsettingNtpServers) GetBurst() bool`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *MemberntpsettingNtpServers) GetBurstOk() (*bool, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *MemberntpsettingNtpServers) SetBurst(v bool)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *MemberntpsettingNtpServers) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetIburst

`func (o *MemberntpsettingNtpServers) GetIburst() bool`

GetIburst returns the Iburst field if non-nil, zero value otherwise.

### GetIburstOk

`func (o *MemberntpsettingNtpServers) GetIburstOk() (*bool, bool)`

GetIburstOk returns a tuple with the Iburst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIburst

`func (o *MemberntpsettingNtpServers) SetIburst(v bool)`

SetIburst sets Iburst field to given value.

### HasIburst

`func (o *MemberntpsettingNtpServers) HasIburst() bool`

HasIburst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


