# MemberDnsFileTransferSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to **string** | The directory to save the captured DNS queries and responses. | [optional] 
**Password** | Pointer to **string** | The password to access the destination server directory. | [optional] 
**Type** | Pointer to **string** | The transfer protocol for the captured DNS queries and responses. | [optional] 
**Username** | Pointer to **string** | The username to access the destination server directory. | [optional] 
**Port** | Pointer to **int64** | Transfer scp port. | [optional] 
**Host** | Pointer to **string** | The host name of the destination server for DNS capture transfer. | [optional] 

## Methods

### NewMemberDnsFileTransferSetting

`func NewMemberDnsFileTransferSetting() *MemberDnsFileTransferSetting`

NewMemberDnsFileTransferSetting instantiates a new MemberDnsFileTransferSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsFileTransferSettingWithDefaults

`func NewMemberDnsFileTransferSettingWithDefaults() *MemberDnsFileTransferSetting`

NewMemberDnsFileTransferSettingWithDefaults instantiates a new MemberDnsFileTransferSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *MemberDnsFileTransferSetting) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *MemberDnsFileTransferSetting) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *MemberDnsFileTransferSetting) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *MemberDnsFileTransferSetting) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetPassword

`func (o *MemberDnsFileTransferSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MemberDnsFileTransferSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MemberDnsFileTransferSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MemberDnsFileTransferSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *MemberDnsFileTransferSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberDnsFileTransferSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberDnsFileTransferSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemberDnsFileTransferSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *MemberDnsFileTransferSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MemberDnsFileTransferSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MemberDnsFileTransferSetting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MemberDnsFileTransferSetting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPort

`func (o *MemberDnsFileTransferSetting) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MemberDnsFileTransferSetting) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MemberDnsFileTransferSetting) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *MemberDnsFileTransferSetting) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHost

`func (o *MemberDnsFileTransferSetting) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MemberDnsFileTransferSetting) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MemberDnsFileTransferSetting) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MemberDnsFileTransferSetting) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


