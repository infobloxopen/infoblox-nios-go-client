# MemberExternalSyslogBackupServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressOrFqdn** | Pointer to **string** | The IPv4 or IPv6 address or FQDN of the backup syslog server. | [optional] 
**DirectoryPath** | Pointer to **string** | The directory path for the replication of the rotated syslog files. | [optional] 
**Enable** | Pointer to **bool** | If set to True, the syslog backup server is enabled. | [optional] 
**Password** | Pointer to **string** | The password of the backup syslog server. | [optional] 
**Port** | Pointer to **int64** | The port used to connect to the backup syslog server. | [optional] 
**Protocol** | Pointer to **string** | The transport protocol used to connect to the backup syslog server. | [optional] 
**Username** | Pointer to **string** | The username of the backup syslog server. | [optional] 

## Methods

### NewMemberExternalSyslogBackupServers

`func NewMemberExternalSyslogBackupServers() *MemberExternalSyslogBackupServers`

NewMemberExternalSyslogBackupServers instantiates a new MemberExternalSyslogBackupServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberExternalSyslogBackupServersWithDefaults

`func NewMemberExternalSyslogBackupServersWithDefaults() *MemberExternalSyslogBackupServers`

NewMemberExternalSyslogBackupServersWithDefaults instantiates a new MemberExternalSyslogBackupServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressOrFqdn

`func (o *MemberExternalSyslogBackupServers) GetAddressOrFqdn() string`

GetAddressOrFqdn returns the AddressOrFqdn field if non-nil, zero value otherwise.

### GetAddressOrFqdnOk

`func (o *MemberExternalSyslogBackupServers) GetAddressOrFqdnOk() (*string, bool)`

GetAddressOrFqdnOk returns a tuple with the AddressOrFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOrFqdn

`func (o *MemberExternalSyslogBackupServers) SetAddressOrFqdn(v string)`

SetAddressOrFqdn sets AddressOrFqdn field to given value.

### HasAddressOrFqdn

`func (o *MemberExternalSyslogBackupServers) HasAddressOrFqdn() bool`

HasAddressOrFqdn returns a boolean if a field has been set.

### GetDirectoryPath

`func (o *MemberExternalSyslogBackupServers) GetDirectoryPath() string`

GetDirectoryPath returns the DirectoryPath field if non-nil, zero value otherwise.

### GetDirectoryPathOk

`func (o *MemberExternalSyslogBackupServers) GetDirectoryPathOk() (*string, bool)`

GetDirectoryPathOk returns a tuple with the DirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPath

`func (o *MemberExternalSyslogBackupServers) SetDirectoryPath(v string)`

SetDirectoryPath sets DirectoryPath field to given value.

### HasDirectoryPath

`func (o *MemberExternalSyslogBackupServers) HasDirectoryPath() bool`

HasDirectoryPath returns a boolean if a field has been set.

### GetEnable

`func (o *MemberExternalSyslogBackupServers) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MemberExternalSyslogBackupServers) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MemberExternalSyslogBackupServers) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MemberExternalSyslogBackupServers) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPassword

`func (o *MemberExternalSyslogBackupServers) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MemberExternalSyslogBackupServers) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MemberExternalSyslogBackupServers) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MemberExternalSyslogBackupServers) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *MemberExternalSyslogBackupServers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MemberExternalSyslogBackupServers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MemberExternalSyslogBackupServers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *MemberExternalSyslogBackupServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *MemberExternalSyslogBackupServers) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MemberExternalSyslogBackupServers) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *MemberExternalSyslogBackupServers) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *MemberExternalSyslogBackupServers) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *MemberExternalSyslogBackupServers) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MemberExternalSyslogBackupServers) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MemberExternalSyslogBackupServers) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MemberExternalSyslogBackupServers) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


