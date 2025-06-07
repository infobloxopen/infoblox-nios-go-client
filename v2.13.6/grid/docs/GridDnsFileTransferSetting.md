# GridDnsFileTransferSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to **string** | The directory to save the captured DNS queries and responses. | [optional] 
**ServerAddressOrFqdn** | Pointer to **string** | The server address or a FQDN name of the destination server for DNS capture transfer. | [optional] 
**Password** | Pointer to **string** | The password to access the destination server directory. | [optional] 
**Type** | Pointer to **string** | The transfer protocol for the captured DNS queries and responses. | [optional] 
**Username** | Pointer to **string** | The username to access the destination server directory. | [optional] 
**Port** | Pointer to **int64** | Transfer scp port. | [optional] 

## Methods

### NewGridDnsFileTransferSetting

`func NewGridDnsFileTransferSetting() *GridDnsFileTransferSetting`

NewGridDnsFileTransferSetting instantiates a new GridDnsFileTransferSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsFileTransferSettingWithDefaults

`func NewGridDnsFileTransferSettingWithDefaults() *GridDnsFileTransferSetting`

NewGridDnsFileTransferSettingWithDefaults instantiates a new GridDnsFileTransferSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *GridDnsFileTransferSetting) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *GridDnsFileTransferSetting) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *GridDnsFileTransferSetting) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *GridDnsFileTransferSetting) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetServerAddressOrFqdn

`func (o *GridDnsFileTransferSetting) GetServerAddressOrFqdn() string`

GetServerAddressOrFqdn returns the ServerAddressOrFqdn field if non-nil, zero value otherwise.

### GetServerAddressOrFqdnOk

`func (o *GridDnsFileTransferSetting) GetServerAddressOrFqdnOk() (*string, bool)`

GetServerAddressOrFqdnOk returns a tuple with the ServerAddressOrFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddressOrFqdn

`func (o *GridDnsFileTransferSetting) SetServerAddressOrFqdn(v string)`

SetServerAddressOrFqdn sets ServerAddressOrFqdn field to given value.

### HasServerAddressOrFqdn

`func (o *GridDnsFileTransferSetting) HasServerAddressOrFqdn() bool`

HasServerAddressOrFqdn returns a boolean if a field has been set.

### GetPassword

`func (o *GridDnsFileTransferSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GridDnsFileTransferSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GridDnsFileTransferSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GridDnsFileTransferSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *GridDnsFileTransferSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridDnsFileTransferSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridDnsFileTransferSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GridDnsFileTransferSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *GridDnsFileTransferSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GridDnsFileTransferSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GridDnsFileTransferSetting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GridDnsFileTransferSetting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPort

`func (o *GridDnsFileTransferSetting) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GridDnsFileTransferSetting) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GridDnsFileTransferSetting) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GridDnsFileTransferSetting) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


