# GetCertificateAuthserviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**RemoteLookupService** | Pointer to **string** | The service that will be used for remote lookup. | [optional] 
**RemoteLookupUsername** | Pointer to **string** | The username for the service account. | [optional] 
**ResponseTimeout** | Pointer to **int64** | The validation timeout period in milliseconds. | [optional] 
**TrustModel** | Pointer to **string** | The OCSP trust model. | [optional] 
**UserMatchType** | Pointer to **string** | Specifies how to search for a user. | [optional] 
**Result** | Pointer to [**CertificateAuthservice**](CertificateAuthservice.md) |  | [optional] 

## Methods

### NewGetCertificateAuthserviceResponse

`func NewGetCertificateAuthserviceResponse() *GetCertificateAuthserviceResponse`

NewGetCertificateAuthserviceResponse instantiates a new GetCertificateAuthserviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertificateAuthserviceResponseWithDefaults

`func NewGetCertificateAuthserviceResponseWithDefaults() *GetCertificateAuthserviceResponse`

NewGetCertificateAuthserviceResponseWithDefaults instantiates a new GetCertificateAuthserviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetCertificateAuthserviceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetCertificateAuthserviceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetCertificateAuthserviceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetCertificateAuthserviceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoPopulateLogin

`func (o *GetCertificateAuthserviceResponse) GetAutoPopulateLogin() string`

GetAutoPopulateLogin returns the AutoPopulateLogin field if non-nil, zero value otherwise.

### GetAutoPopulateLoginOk

`func (o *GetCertificateAuthserviceResponse) GetAutoPopulateLoginOk() (*string, bool)`

GetAutoPopulateLoginOk returns a tuple with the AutoPopulateLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPopulateLogin

`func (o *GetCertificateAuthserviceResponse) SetAutoPopulateLogin(v string)`

SetAutoPopulateLogin sets AutoPopulateLogin field to given value.

### HasAutoPopulateLogin

`func (o *GetCertificateAuthserviceResponse) HasAutoPopulateLogin() bool`

HasAutoPopulateLogin returns a boolean if a field has been set.

### GetCaCertificates

`func (o *GetCertificateAuthserviceResponse) GetCaCertificates() []string`

GetCaCertificates returns the CaCertificates field if non-nil, zero value otherwise.

### GetCaCertificatesOk

`func (o *GetCertificateAuthserviceResponse) GetCaCertificatesOk() (*[]string, bool)`

GetCaCertificatesOk returns a tuple with the CaCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificates

`func (o *GetCertificateAuthserviceResponse) SetCaCertificates(v []string)`

SetCaCertificates sets CaCertificates field to given value.

### HasCaCertificates

`func (o *GetCertificateAuthserviceResponse) HasCaCertificates() bool`

HasCaCertificates returns a boolean if a field has been set.

### GetComment

`func (o *GetCertificateAuthserviceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetCertificateAuthserviceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetCertificateAuthserviceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetCertificateAuthserviceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *GetCertificateAuthserviceResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetCertificateAuthserviceResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetCertificateAuthserviceResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetCertificateAuthserviceResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEnablePasswordRequest

`func (o *GetCertificateAuthserviceResponse) GetEnablePasswordRequest() bool`

GetEnablePasswordRequest returns the EnablePasswordRequest field if non-nil, zero value otherwise.

### GetEnablePasswordRequestOk

`func (o *GetCertificateAuthserviceResponse) GetEnablePasswordRequestOk() (*bool, bool)`

GetEnablePasswordRequestOk returns a tuple with the EnablePasswordRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordRequest

`func (o *GetCertificateAuthserviceResponse) SetEnablePasswordRequest(v bool)`

SetEnablePasswordRequest sets EnablePasswordRequest field to given value.

### HasEnablePasswordRequest

`func (o *GetCertificateAuthserviceResponse) HasEnablePasswordRequest() bool`

HasEnablePasswordRequest returns a boolean if a field has been set.

### GetEnableRemoteLookup

`func (o *GetCertificateAuthserviceResponse) GetEnableRemoteLookup() bool`

GetEnableRemoteLookup returns the EnableRemoteLookup field if non-nil, zero value otherwise.

### GetEnableRemoteLookupOk

`func (o *GetCertificateAuthserviceResponse) GetEnableRemoteLookupOk() (*bool, bool)`

GetEnableRemoteLookupOk returns a tuple with the EnableRemoteLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteLookup

`func (o *GetCertificateAuthserviceResponse) SetEnableRemoteLookup(v bool)`

SetEnableRemoteLookup sets EnableRemoteLookup field to given value.

### HasEnableRemoteLookup

`func (o *GetCertificateAuthserviceResponse) HasEnableRemoteLookup() bool`

HasEnableRemoteLookup returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetCertificateAuthserviceResponse) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetCertificateAuthserviceResponse) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetCertificateAuthserviceResponse) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetCertificateAuthserviceResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetName

`func (o *GetCertificateAuthserviceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCertificateAuthserviceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCertificateAuthserviceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCertificateAuthserviceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOcspCheck

`func (o *GetCertificateAuthserviceResponse) GetOcspCheck() string`

GetOcspCheck returns the OcspCheck field if non-nil, zero value otherwise.

### GetOcspCheckOk

`func (o *GetCertificateAuthserviceResponse) GetOcspCheckOk() (*string, bool)`

GetOcspCheckOk returns a tuple with the OcspCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspCheck

`func (o *GetCertificateAuthserviceResponse) SetOcspCheck(v string)`

SetOcspCheck sets OcspCheck field to given value.

### HasOcspCheck

`func (o *GetCertificateAuthserviceResponse) HasOcspCheck() bool`

HasOcspCheck returns a boolean if a field has been set.

### GetOcspResponders

`func (o *GetCertificateAuthserviceResponse) GetOcspResponders() []CertificateAuthserviceOcspResponders`

GetOcspResponders returns the OcspResponders field if non-nil, zero value otherwise.

### GetOcspRespondersOk

`func (o *GetCertificateAuthserviceResponse) GetOcspRespondersOk() (*[]CertificateAuthserviceOcspResponders, bool)`

GetOcspRespondersOk returns a tuple with the OcspResponders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspResponders

`func (o *GetCertificateAuthserviceResponse) SetOcspResponders(v []CertificateAuthserviceOcspResponders)`

SetOcspResponders sets OcspResponders field to given value.

### HasOcspResponders

`func (o *GetCertificateAuthserviceResponse) HasOcspResponders() bool`

HasOcspResponders returns a boolean if a field has been set.

### GetRecoveryInterval

`func (o *GetCertificateAuthserviceResponse) GetRecoveryInterval() int64`

GetRecoveryInterval returns the RecoveryInterval field if non-nil, zero value otherwise.

### GetRecoveryIntervalOk

`func (o *GetCertificateAuthserviceResponse) GetRecoveryIntervalOk() (*int64, bool)`

GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInterval

`func (o *GetCertificateAuthserviceResponse) SetRecoveryInterval(v int64)`

SetRecoveryInterval sets RecoveryInterval field to given value.

### HasRecoveryInterval

`func (o *GetCertificateAuthserviceResponse) HasRecoveryInterval() bool`

HasRecoveryInterval returns a boolean if a field has been set.

### GetRemoteLookupPassword

`func (o *GetCertificateAuthserviceResponse) GetRemoteLookupPassword() string`

GetRemoteLookupPassword returns the RemoteLookupPassword field if non-nil, zero value otherwise.

### GetRemoteLookupPasswordOk

`func (o *GetCertificateAuthserviceResponse) GetRemoteLookupPasswordOk() (*string, bool)`

GetRemoteLookupPasswordOk returns a tuple with the RemoteLookupPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLookupPassword

`func (o *GetCertificateAuthserviceResponse) SetRemoteLookupPassword(v string)`

SetRemoteLookupPassword sets RemoteLookupPassword field to given value.

### HasRemoteLookupPassword

`func (o *GetCertificateAuthserviceResponse) HasRemoteLookupPassword() bool`

HasRemoteLookupPassword returns a boolean if a field has been set.

### GetRemoteLookupService

`func (o *GetCertificateAuthserviceResponse) GetRemoteLookupService() string`

GetRemoteLookupService returns the RemoteLookupService field if non-nil, zero value otherwise.

### GetRemoteLookupServiceOk

`func (o *GetCertificateAuthserviceResponse) GetRemoteLookupServiceOk() (*string, bool)`

GetRemoteLookupServiceOk returns a tuple with the RemoteLookupService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLookupService

`func (o *GetCertificateAuthserviceResponse) SetRemoteLookupService(v string)`

SetRemoteLookupService sets RemoteLookupService field to given value.

### HasRemoteLookupService

`func (o *GetCertificateAuthserviceResponse) HasRemoteLookupService() bool`

HasRemoteLookupService returns a boolean if a field has been set.

### GetRemoteLookupUsername

`func (o *GetCertificateAuthserviceResponse) GetRemoteLookupUsername() string`

GetRemoteLookupUsername returns the RemoteLookupUsername field if non-nil, zero value otherwise.

### GetRemoteLookupUsernameOk

`func (o *GetCertificateAuthserviceResponse) GetRemoteLookupUsernameOk() (*string, bool)`

GetRemoteLookupUsernameOk returns a tuple with the RemoteLookupUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLookupUsername

`func (o *GetCertificateAuthserviceResponse) SetRemoteLookupUsername(v string)`

SetRemoteLookupUsername sets RemoteLookupUsername field to given value.

### HasRemoteLookupUsername

`func (o *GetCertificateAuthserviceResponse) HasRemoteLookupUsername() bool`

HasRemoteLookupUsername returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *GetCertificateAuthserviceResponse) GetResponseTimeout() int64`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *GetCertificateAuthserviceResponse) GetResponseTimeoutOk() (*int64, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *GetCertificateAuthserviceResponse) SetResponseTimeout(v int64)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *GetCertificateAuthserviceResponse) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetTrustModel

`func (o *GetCertificateAuthserviceResponse) GetTrustModel() string`

GetTrustModel returns the TrustModel field if non-nil, zero value otherwise.

### GetTrustModelOk

`func (o *GetCertificateAuthserviceResponse) GetTrustModelOk() (*string, bool)`

GetTrustModelOk returns a tuple with the TrustModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustModel

`func (o *GetCertificateAuthserviceResponse) SetTrustModel(v string)`

SetTrustModel sets TrustModel field to given value.

### HasTrustModel

`func (o *GetCertificateAuthserviceResponse) HasTrustModel() bool`

HasTrustModel returns a boolean if a field has been set.

### GetUserMatchType

`func (o *GetCertificateAuthserviceResponse) GetUserMatchType() string`

GetUserMatchType returns the UserMatchType field if non-nil, zero value otherwise.

### GetUserMatchTypeOk

`func (o *GetCertificateAuthserviceResponse) GetUserMatchTypeOk() (*string, bool)`

GetUserMatchTypeOk returns a tuple with the UserMatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchType

`func (o *GetCertificateAuthserviceResponse) SetUserMatchType(v string)`

SetUserMatchType sets UserMatchType field to given value.

### HasUserMatchType

`func (o *GetCertificateAuthserviceResponse) HasUserMatchType() bool`

HasUserMatchType returns a boolean if a field has been set.

### GetResult

`func (o *GetCertificateAuthserviceResponse) GetResult() CertificateAuthservice`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetCertificateAuthserviceResponse) GetResultOk() (*CertificateAuthservice, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetCertificateAuthserviceResponse) SetResult(v CertificateAuthservice)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetCertificateAuthserviceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


