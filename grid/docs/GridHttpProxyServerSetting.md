# GridHttpProxyServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address of the HTTP proxy server. | [optional] 
**Port** | Pointer to **int64** | The port on which the HTTP proxy server listens. | [optional] 
**EnableProxy** | Pointer to **bool** | Determines if the HTTP proxy server is enabled or not. | [optional] 
**EnableContentInspection** | Pointer to **bool** | Determines if HTTPS content inspection by the HTTP proxy server is enabled or not. | [optional] 
**VerifyCname** | Pointer to **bool** | Determines if the CNAME record query verification is enabled or not. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the HTTP proxy server configuration. | [optional] 
**Username** | Pointer to **string** | The user name for the HTTP proxy server. | [optional] 
**Password** | Pointer to **string** | The password for the HTTP proxy server. | [optional] 
**Certificate** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for the CA certificate file used in the content inspection by an HTTP proxy server. | [optional] 
**EnableUsernameAndPassword** | Pointer to **bool** | Determines if username and password for HTTP Proxy Server connectivity is used or not. | [optional] 

## Methods

### NewGridHttpProxyServerSetting

`func NewGridHttpProxyServerSetting() *GridHttpProxyServerSetting`

NewGridHttpProxyServerSetting instantiates a new GridHttpProxyServerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridHttpProxyServerSettingWithDefaults

`func NewGridHttpProxyServerSettingWithDefaults() *GridHttpProxyServerSetting`

NewGridHttpProxyServerSettingWithDefaults instantiates a new GridHttpProxyServerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GridHttpProxyServerSetting) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GridHttpProxyServerSetting) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GridHttpProxyServerSetting) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GridHttpProxyServerSetting) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *GridHttpProxyServerSetting) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GridHttpProxyServerSetting) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GridHttpProxyServerSetting) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GridHttpProxyServerSetting) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEnableProxy

`func (o *GridHttpProxyServerSetting) GetEnableProxy() bool`

GetEnableProxy returns the EnableProxy field if non-nil, zero value otherwise.

### GetEnableProxyOk

`func (o *GridHttpProxyServerSetting) GetEnableProxyOk() (*bool, bool)`

GetEnableProxyOk returns a tuple with the EnableProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxy

`func (o *GridHttpProxyServerSetting) SetEnableProxy(v bool)`

SetEnableProxy sets EnableProxy field to given value.

### HasEnableProxy

`func (o *GridHttpProxyServerSetting) HasEnableProxy() bool`

HasEnableProxy returns a boolean if a field has been set.

### GetEnableContentInspection

`func (o *GridHttpProxyServerSetting) GetEnableContentInspection() bool`

GetEnableContentInspection returns the EnableContentInspection field if non-nil, zero value otherwise.

### GetEnableContentInspectionOk

`func (o *GridHttpProxyServerSetting) GetEnableContentInspectionOk() (*bool, bool)`

GetEnableContentInspectionOk returns a tuple with the EnableContentInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentInspection

`func (o *GridHttpProxyServerSetting) SetEnableContentInspection(v bool)`

SetEnableContentInspection sets EnableContentInspection field to given value.

### HasEnableContentInspection

`func (o *GridHttpProxyServerSetting) HasEnableContentInspection() bool`

HasEnableContentInspection returns a boolean if a field has been set.

### GetVerifyCname

`func (o *GridHttpProxyServerSetting) GetVerifyCname() bool`

GetVerifyCname returns the VerifyCname field if non-nil, zero value otherwise.

### GetVerifyCnameOk

`func (o *GridHttpProxyServerSetting) GetVerifyCnameOk() (*bool, bool)`

GetVerifyCnameOk returns a tuple with the VerifyCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCname

`func (o *GridHttpProxyServerSetting) SetVerifyCname(v bool)`

SetVerifyCname sets VerifyCname field to given value.

### HasVerifyCname

`func (o *GridHttpProxyServerSetting) HasVerifyCname() bool`

HasVerifyCname returns a boolean if a field has been set.

### GetComment

`func (o *GridHttpProxyServerSetting) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GridHttpProxyServerSetting) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GridHttpProxyServerSetting) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GridHttpProxyServerSetting) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetUsername

`func (o *GridHttpProxyServerSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GridHttpProxyServerSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GridHttpProxyServerSetting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GridHttpProxyServerSetting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GridHttpProxyServerSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GridHttpProxyServerSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GridHttpProxyServerSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GridHttpProxyServerSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCertificate

`func (o *GridHttpProxyServerSetting) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *GridHttpProxyServerSetting) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *GridHttpProxyServerSetting) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *GridHttpProxyServerSetting) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetEnableUsernameAndPassword

`func (o *GridHttpProxyServerSetting) GetEnableUsernameAndPassword() bool`

GetEnableUsernameAndPassword returns the EnableUsernameAndPassword field if non-nil, zero value otherwise.

### GetEnableUsernameAndPasswordOk

`func (o *GridHttpProxyServerSetting) GetEnableUsernameAndPasswordOk() (*bool, bool)`

GetEnableUsernameAndPasswordOk returns a tuple with the EnableUsernameAndPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUsernameAndPassword

`func (o *GridHttpProxyServerSetting) SetEnableUsernameAndPassword(v bool)`

SetEnableUsernameAndPassword sets EnableUsernameAndPassword field to given value.

### HasEnableUsernameAndPassword

`func (o *GridHttpProxyServerSetting) HasEnableUsernameAndPassword() bool`

HasEnableUsernameAndPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


