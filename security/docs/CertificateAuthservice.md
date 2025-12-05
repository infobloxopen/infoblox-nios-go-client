# CertificateAuthservice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AutoPopulateLogin** | Pointer to **string** | Specifies the value of the client certificate for automatically populating the NIOS login name. | [optional] 
**CaCertificates** | Pointer to **[]string** | The list of CA certificates. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the certificate authentication service. | [optional] 
**Disabled** | Pointer to **bool** | Determines if this certificate authentication service is enabled or disabled. | [optional] 
**EnablePasswordRequest** | Pointer to **bool** | Determines if username/password authentication together with client certificate authentication is enabled or disabled. | [optional] 
**EnableRemoteLookup** | Pointer to **bool** | Determines if the lookup for user group membership information on remote services is enabled or disabled. | [optional] 
**MaxRetries** | Pointer to **int64** | The number of validation attempts before the appliance contacts the next responder. | [optional] 
**Name** | Pointer to **string** | The name of the certificate authentication service. | [optional] 
**OcspCheck** | Pointer to **string** | Specifies the source of OCSP settings. | [optional] 
**OcspResponders** | Pointer to [**[]CertificateAuthserviceOcspResponders**](CertificateAuthserviceOcspResponders.md) | An ordered list of OCSP responders that are part of the certificate authentication service. | [optional] 
**RecoveryInterval** | Pointer to **int64** | The period of time the appliance waits before it attempts to contact a responder that is out of service again. The value must be between 1 and 600 seconds. | [optional] 
**RemoteLookupPassword** | Pointer to **string** | The password for the service account. | [optional] 
**RemoteLookupService** | Pointer to [**CertificateAuthserviceRemoteLookupService**](CertificateAuthserviceRemoteLookupService.md) |  | [optional] 
**RemoteLookupUsername** | Pointer to **string** | The username for the service account. | [optional] 
**ResponseTimeout** | Pointer to **int64** | The validation timeout period in milliseconds. | [optional] 
**TrustModel** | Pointer to **string** | The OCSP trust model. | [optional] 
**UserMatchType** | Pointer to **string** | Specifies how to search for a user. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewCertificateAuthservice

`func NewCertificateAuthservice() *CertificateAuthservice`

NewCertificateAuthservice instantiates a new CertificateAuthservice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthserviceWithDefaults

`func NewCertificateAuthserviceWithDefaults() *CertificateAuthservice`

NewCertificateAuthserviceWithDefaults instantiates a new CertificateAuthservice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *CertificateAuthservice) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CertificateAuthservice) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CertificateAuthservice) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CertificateAuthservice) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoPopulateLogin

`func (o *CertificateAuthservice) GetAutoPopulateLogin() string`

GetAutoPopulateLogin returns the AutoPopulateLogin field if non-nil, zero value otherwise.

### GetAutoPopulateLoginOk

`func (o *CertificateAuthservice) GetAutoPopulateLoginOk() (*string, bool)`

GetAutoPopulateLoginOk returns a tuple with the AutoPopulateLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPopulateLogin

`func (o *CertificateAuthservice) SetAutoPopulateLogin(v string)`

SetAutoPopulateLogin sets AutoPopulateLogin field to given value.

### HasAutoPopulateLogin

`func (o *CertificateAuthservice) HasAutoPopulateLogin() bool`

HasAutoPopulateLogin returns a boolean if a field has been set.

### GetCaCertificates

`func (o *CertificateAuthservice) GetCaCertificates() []string`

GetCaCertificates returns the CaCertificates field if non-nil, zero value otherwise.

### GetCaCertificatesOk

`func (o *CertificateAuthservice) GetCaCertificatesOk() (*[]string, bool)`

GetCaCertificatesOk returns a tuple with the CaCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificates

`func (o *CertificateAuthservice) SetCaCertificates(v []string)`

SetCaCertificates sets CaCertificates field to given value.

### HasCaCertificates

`func (o *CertificateAuthservice) HasCaCertificates() bool`

HasCaCertificates returns a boolean if a field has been set.

### GetComment

`func (o *CertificateAuthservice) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CertificateAuthservice) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CertificateAuthservice) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CertificateAuthservice) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *CertificateAuthservice) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CertificateAuthservice) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CertificateAuthservice) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CertificateAuthservice) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEnablePasswordRequest

`func (o *CertificateAuthservice) GetEnablePasswordRequest() bool`

GetEnablePasswordRequest returns the EnablePasswordRequest field if non-nil, zero value otherwise.

### GetEnablePasswordRequestOk

`func (o *CertificateAuthservice) GetEnablePasswordRequestOk() (*bool, bool)`

GetEnablePasswordRequestOk returns a tuple with the EnablePasswordRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordRequest

`func (o *CertificateAuthservice) SetEnablePasswordRequest(v bool)`

SetEnablePasswordRequest sets EnablePasswordRequest field to given value.

### HasEnablePasswordRequest

`func (o *CertificateAuthservice) HasEnablePasswordRequest() bool`

HasEnablePasswordRequest returns a boolean if a field has been set.

### GetEnableRemoteLookup

`func (o *CertificateAuthservice) GetEnableRemoteLookup() bool`

GetEnableRemoteLookup returns the EnableRemoteLookup field if non-nil, zero value otherwise.

### GetEnableRemoteLookupOk

`func (o *CertificateAuthservice) GetEnableRemoteLookupOk() (*bool, bool)`

GetEnableRemoteLookupOk returns a tuple with the EnableRemoteLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteLookup

`func (o *CertificateAuthservice) SetEnableRemoteLookup(v bool)`

SetEnableRemoteLookup sets EnableRemoteLookup field to given value.

### HasEnableRemoteLookup

`func (o *CertificateAuthservice) HasEnableRemoteLookup() bool`

HasEnableRemoteLookup returns a boolean if a field has been set.

### GetMaxRetries

`func (o *CertificateAuthservice) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *CertificateAuthservice) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *CertificateAuthservice) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *CertificateAuthservice) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetName

`func (o *CertificateAuthservice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateAuthservice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateAuthservice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateAuthservice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOcspCheck

`func (o *CertificateAuthservice) GetOcspCheck() string`

GetOcspCheck returns the OcspCheck field if non-nil, zero value otherwise.

### GetOcspCheckOk

`func (o *CertificateAuthservice) GetOcspCheckOk() (*string, bool)`

GetOcspCheckOk returns a tuple with the OcspCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspCheck

`func (o *CertificateAuthservice) SetOcspCheck(v string)`

SetOcspCheck sets OcspCheck field to given value.

### HasOcspCheck

`func (o *CertificateAuthservice) HasOcspCheck() bool`

HasOcspCheck returns a boolean if a field has been set.

### GetOcspResponders

`func (o *CertificateAuthservice) GetOcspResponders() []CertificateAuthserviceOcspResponders`

GetOcspResponders returns the OcspResponders field if non-nil, zero value otherwise.

### GetOcspRespondersOk

`func (o *CertificateAuthservice) GetOcspRespondersOk() (*[]CertificateAuthserviceOcspResponders, bool)`

GetOcspRespondersOk returns a tuple with the OcspResponders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspResponders

`func (o *CertificateAuthservice) SetOcspResponders(v []CertificateAuthserviceOcspResponders)`

SetOcspResponders sets OcspResponders field to given value.

### HasOcspResponders

`func (o *CertificateAuthservice) HasOcspResponders() bool`

HasOcspResponders returns a boolean if a field has been set.

### GetRecoveryInterval

`func (o *CertificateAuthservice) GetRecoveryInterval() int64`

GetRecoveryInterval returns the RecoveryInterval field if non-nil, zero value otherwise.

### GetRecoveryIntervalOk

`func (o *CertificateAuthservice) GetRecoveryIntervalOk() (*int64, bool)`

GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInterval

`func (o *CertificateAuthservice) SetRecoveryInterval(v int64)`

SetRecoveryInterval sets RecoveryInterval field to given value.

### HasRecoveryInterval

`func (o *CertificateAuthservice) HasRecoveryInterval() bool`

HasRecoveryInterval returns a boolean if a field has been set.

### GetRemoteLookupPassword

`func (o *CertificateAuthservice) GetRemoteLookupPassword() string`

GetRemoteLookupPassword returns the RemoteLookupPassword field if non-nil, zero value otherwise.

### GetRemoteLookupPasswordOk

`func (o *CertificateAuthservice) GetRemoteLookupPasswordOk() (*string, bool)`

GetRemoteLookupPasswordOk returns a tuple with the RemoteLookupPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLookupPassword

`func (o *CertificateAuthservice) SetRemoteLookupPassword(v string)`

SetRemoteLookupPassword sets RemoteLookupPassword field to given value.

### HasRemoteLookupPassword

`func (o *CertificateAuthservice) HasRemoteLookupPassword() bool`

HasRemoteLookupPassword returns a boolean if a field has been set.

### GetRemoteLookupService

`func (o *CertificateAuthservice) GetRemoteLookupService() CertificateAuthserviceRemoteLookupService`

GetRemoteLookupService returns the RemoteLookupService field if non-nil, zero value otherwise.

### GetRemoteLookupServiceOk

`func (o *CertificateAuthservice) GetRemoteLookupServiceOk() (*CertificateAuthserviceRemoteLookupService, bool)`

GetRemoteLookupServiceOk returns a tuple with the RemoteLookupService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLookupService

`func (o *CertificateAuthservice) SetRemoteLookupService(v CertificateAuthserviceRemoteLookupService)`

SetRemoteLookupService sets RemoteLookupService field to given value.

### HasRemoteLookupService

`func (o *CertificateAuthservice) HasRemoteLookupService() bool`

HasRemoteLookupService returns a boolean if a field has been set.

### GetRemoteLookupUsername

`func (o *CertificateAuthservice) GetRemoteLookupUsername() string`

GetRemoteLookupUsername returns the RemoteLookupUsername field if non-nil, zero value otherwise.

### GetRemoteLookupUsernameOk

`func (o *CertificateAuthservice) GetRemoteLookupUsernameOk() (*string, bool)`

GetRemoteLookupUsernameOk returns a tuple with the RemoteLookupUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLookupUsername

`func (o *CertificateAuthservice) SetRemoteLookupUsername(v string)`

SetRemoteLookupUsername sets RemoteLookupUsername field to given value.

### HasRemoteLookupUsername

`func (o *CertificateAuthservice) HasRemoteLookupUsername() bool`

HasRemoteLookupUsername returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *CertificateAuthservice) GetResponseTimeout() int64`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *CertificateAuthservice) GetResponseTimeoutOk() (*int64, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *CertificateAuthservice) SetResponseTimeout(v int64)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *CertificateAuthservice) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetTrustModel

`func (o *CertificateAuthservice) GetTrustModel() string`

GetTrustModel returns the TrustModel field if non-nil, zero value otherwise.

### GetTrustModelOk

`func (o *CertificateAuthservice) GetTrustModelOk() (*string, bool)`

GetTrustModelOk returns a tuple with the TrustModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustModel

`func (o *CertificateAuthservice) SetTrustModel(v string)`

SetTrustModel sets TrustModel field to given value.

### HasTrustModel

`func (o *CertificateAuthservice) HasTrustModel() bool`

HasTrustModel returns a boolean if a field has been set.

### GetUserMatchType

`func (o *CertificateAuthservice) GetUserMatchType() string`

GetUserMatchType returns the UserMatchType field if non-nil, zero value otherwise.

### GetUserMatchTypeOk

`func (o *CertificateAuthservice) GetUserMatchTypeOk() (*string, bool)`

GetUserMatchTypeOk returns a tuple with the UserMatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchType

`func (o *CertificateAuthservice) SetUserMatchType(v string)`

SetUserMatchType sets UserMatchType field to given value.

### HasUserMatchType

`func (o *CertificateAuthservice) HasUserMatchType() bool`

HasUserMatchType returns a boolean if a field has been set.

### GetUuid

`func (o *CertificateAuthservice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CertificateAuthservice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CertificateAuthservice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CertificateAuthservice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


