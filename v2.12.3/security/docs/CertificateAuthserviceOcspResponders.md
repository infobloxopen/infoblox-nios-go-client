# CertificateAuthserviceOcspResponders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FqdnOrIp** | Pointer to **string** | The FQDN (Fully Qualified Domain Name) or IP address of the server. | [optional] 
**Port** | Pointer to **int64** | The port used for connecting. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the OCSP authentication responder. | [optional] 
**Disabled** | Pointer to **bool** | Determines if this OCSP authentication responder is disabled. | [optional] 
**Certificate** | Pointer to **string** | The reference to the OCSP responder certificate. | [optional] [readonly] 
**CertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop. | [optional] 

## Methods

### NewCertificateAuthserviceOcspResponders

`func NewCertificateAuthserviceOcspResponders() *CertificateAuthserviceOcspResponders`

NewCertificateAuthserviceOcspResponders instantiates a new CertificateAuthserviceOcspResponders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthserviceOcspRespondersWithDefaults

`func NewCertificateAuthserviceOcspRespondersWithDefaults() *CertificateAuthserviceOcspResponders`

NewCertificateAuthserviceOcspRespondersWithDefaults instantiates a new CertificateAuthserviceOcspResponders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdnOrIp

`func (o *CertificateAuthserviceOcspResponders) GetFqdnOrIp() string`

GetFqdnOrIp returns the FqdnOrIp field if non-nil, zero value otherwise.

### GetFqdnOrIpOk

`func (o *CertificateAuthserviceOcspResponders) GetFqdnOrIpOk() (*string, bool)`

GetFqdnOrIpOk returns a tuple with the FqdnOrIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnOrIp

`func (o *CertificateAuthserviceOcspResponders) SetFqdnOrIp(v string)`

SetFqdnOrIp sets FqdnOrIp field to given value.

### HasFqdnOrIp

`func (o *CertificateAuthserviceOcspResponders) HasFqdnOrIp() bool`

HasFqdnOrIp returns a boolean if a field has been set.

### GetPort

`func (o *CertificateAuthserviceOcspResponders) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CertificateAuthserviceOcspResponders) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CertificateAuthserviceOcspResponders) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CertificateAuthserviceOcspResponders) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetComment

`func (o *CertificateAuthserviceOcspResponders) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CertificateAuthserviceOcspResponders) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CertificateAuthserviceOcspResponders) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CertificateAuthserviceOcspResponders) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *CertificateAuthserviceOcspResponders) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CertificateAuthserviceOcspResponders) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CertificateAuthserviceOcspResponders) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CertificateAuthserviceOcspResponders) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateAuthserviceOcspResponders) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateAuthserviceOcspResponders) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateAuthserviceOcspResponders) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateAuthserviceOcspResponders) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateToken

`func (o *CertificateAuthserviceOcspResponders) GetCertificateToken() string`

GetCertificateToken returns the CertificateToken field if non-nil, zero value otherwise.

### GetCertificateTokenOk

`func (o *CertificateAuthserviceOcspResponders) GetCertificateTokenOk() (*string, bool)`

GetCertificateTokenOk returns a tuple with the CertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateToken

`func (o *CertificateAuthserviceOcspResponders) SetCertificateToken(v string)`

SetCertificateToken sets CertificateToken field to given value.

### HasCertificateToken

`func (o *CertificateAuthserviceOcspResponders) HasCertificateToken() bool`

HasCertificateToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


