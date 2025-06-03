# GridExternalSyslogBackupServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryPath** | Pointer to **string** | The directory path for the replication of the rotated syslog files. | [optional] 
**Enable** | Pointer to **bool** | If set to True, the syslog backup server is enabled. | [optional] 
**Password** | Pointer to **string** | The password of the backup syslog server. | [optional] 
**Port** | Pointer to **int64** | The port used to connect to the backup syslog server. | [optional] 
**Protocol** | Pointer to **string** | The transport protocol used to connect to the backup syslog server. | [optional] 
**Username** | Pointer to **string** | The username of the backup syslog server. | [optional] 
**Address** | Pointer to **string** | The IPv4/IPv6 address or FQDN of the backup syslog server. | [optional] 

## Methods

### NewGridExternalSyslogBackupServers

`func NewGridExternalSyslogBackupServers() *GridExternalSyslogBackupServers`

NewGridExternalSyslogBackupServers instantiates a new GridExternalSyslogBackupServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridExternalSyslogBackupServersWithDefaults

`func NewGridExternalSyslogBackupServersWithDefaults() *GridExternalSyslogBackupServers`

NewGridExternalSyslogBackupServersWithDefaults instantiates a new GridExternalSyslogBackupServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryPath

`func (o *GridExternalSyslogBackupServers) GetDirectoryPath() string`

GetDirectoryPath returns the DirectoryPath field if non-nil, zero value otherwise.

### GetDirectoryPathOk

`func (o *GridExternalSyslogBackupServers) GetDirectoryPathOk() (*string, bool)`

GetDirectoryPathOk returns a tuple with the DirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPath

`func (o *GridExternalSyslogBackupServers) SetDirectoryPath(v string)`

SetDirectoryPath sets DirectoryPath field to given value.

### HasDirectoryPath

`func (o *GridExternalSyslogBackupServers) HasDirectoryPath() bool`

HasDirectoryPath returns a boolean if a field has been set.

### GetEnable

`func (o *GridExternalSyslogBackupServers) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GridExternalSyslogBackupServers) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GridExternalSyslogBackupServers) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GridExternalSyslogBackupServers) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPassword

`func (o *GridExternalSyslogBackupServers) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GridExternalSyslogBackupServers) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GridExternalSyslogBackupServers) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GridExternalSyslogBackupServers) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *GridExternalSyslogBackupServers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GridExternalSyslogBackupServers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GridExternalSyslogBackupServers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GridExternalSyslogBackupServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *GridExternalSyslogBackupServers) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GridExternalSyslogBackupServers) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GridExternalSyslogBackupServers) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GridExternalSyslogBackupServers) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *GridExternalSyslogBackupServers) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GridExternalSyslogBackupServers) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GridExternalSyslogBackupServers) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GridExternalSyslogBackupServers) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAddress

`func (o *GridExternalSyslogBackupServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GridExternalSyslogBackupServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GridExternalSyslogBackupServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GridExternalSyslogBackupServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


