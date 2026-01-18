# GetGcpuserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AuthProviderX509CertUrl** | Pointer to **string** | The URL where the public key certificates provided by the authentication provider can be retrieved.. Maximum 255 characters. | [optional] [readonly] 
**AuthUri** | Pointer to **string** | The URI where authentication requests should be directed.. Maximum 255 characters. | [optional] [readonly] 
**ClientEmail** | Pointer to **string** | The email address associated with the service account. Maximum 255 characters. | [optional] [readonly] 
**ClientId** | Pointer to **string** | The unique identifier for the service account. Maximum 64 characts.er | [optional] [readonly] 
**ClientX509CertUrl** | Pointer to **string** | The URL where the public key certificate for the service account can be retrieved. Maximum 255 characters. | [optional] [readonly] 
**FileName** | Pointer to **string** | GCP client credentials file name.. Maximum 255 characters. | [optional] [readonly] 
**LastUsed** | Pointer to **int64** | The timestamp when this Azure user credentials was last used. | [optional] [readonly] 
**PrivateKey** | Pointer to **string** | The private key used for authentication. Maximum 255 characters. | [optional] [readonly] 
**PrivateKeyId** | Pointer to **string** | The identifier for the private key associated with the service account. Maximum 64 characters. | [optional] [readonly] 
**ProjectId** | Pointer to **string** | The ID of the GCP project associated with the service account. Maximum 64 characters. | [optional] [readonly] 
**Status** | Pointer to **string** | Indicate the validity status of this GCP user. | [optional] [readonly] 
**TokenUri** | Pointer to **string** | The URI where token requests should be directed.. Maximum 255 characters. | [optional] [readonly] 
**Type** | Pointer to **string** | Specifies the type of the credential. Maximum 255 characters. | [optional] [readonly] 
**UserName** | Pointer to **string** | The GCP client&#39;s user name. Maximum 64 characters. | [optional] 
**Result** | Pointer to [**Gcpuser**](Gcpuser.md) |  | [optional] 

## Methods

### NewGetGcpuserResponse

`func NewGetGcpuserResponse() *GetGcpuserResponse`

NewGetGcpuserResponse instantiates a new GetGcpuserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGcpuserResponseWithDefaults

`func NewGetGcpuserResponseWithDefaults() *GetGcpuserResponse`

NewGetGcpuserResponseWithDefaults instantiates a new GetGcpuserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGcpuserResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGcpuserResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGcpuserResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGcpuserResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthProviderX509CertUrl

`func (o *GetGcpuserResponse) GetAuthProviderX509CertUrl() string`

GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field if non-nil, zero value otherwise.

### GetAuthProviderX509CertUrlOk

`func (o *GetGcpuserResponse) GetAuthProviderX509CertUrlOk() (*string, bool)`

GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviderX509CertUrl

`func (o *GetGcpuserResponse) SetAuthProviderX509CertUrl(v string)`

SetAuthProviderX509CertUrl sets AuthProviderX509CertUrl field to given value.

### HasAuthProviderX509CertUrl

`func (o *GetGcpuserResponse) HasAuthProviderX509CertUrl() bool`

HasAuthProviderX509CertUrl returns a boolean if a field has been set.

### GetAuthUri

`func (o *GetGcpuserResponse) GetAuthUri() string`

GetAuthUri returns the AuthUri field if non-nil, zero value otherwise.

### GetAuthUriOk

`func (o *GetGcpuserResponse) GetAuthUriOk() (*string, bool)`

GetAuthUriOk returns a tuple with the AuthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUri

`func (o *GetGcpuserResponse) SetAuthUri(v string)`

SetAuthUri sets AuthUri field to given value.

### HasAuthUri

`func (o *GetGcpuserResponse) HasAuthUri() bool`

HasAuthUri returns a boolean if a field has been set.

### GetClientEmail

`func (o *GetGcpuserResponse) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *GetGcpuserResponse) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *GetGcpuserResponse) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *GetGcpuserResponse) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetClientId

`func (o *GetGcpuserResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetGcpuserResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetGcpuserResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetGcpuserResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientX509CertUrl

`func (o *GetGcpuserResponse) GetClientX509CertUrl() string`

GetClientX509CertUrl returns the ClientX509CertUrl field if non-nil, zero value otherwise.

### GetClientX509CertUrlOk

`func (o *GetGcpuserResponse) GetClientX509CertUrlOk() (*string, bool)`

GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientX509CertUrl

`func (o *GetGcpuserResponse) SetClientX509CertUrl(v string)`

SetClientX509CertUrl sets ClientX509CertUrl field to given value.

### HasClientX509CertUrl

`func (o *GetGcpuserResponse) HasClientX509CertUrl() bool`

HasClientX509CertUrl returns a boolean if a field has been set.

### GetFileName

`func (o *GetGcpuserResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetGcpuserResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetGcpuserResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *GetGcpuserResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetLastUsed

`func (o *GetGcpuserResponse) GetLastUsed() int64`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *GetGcpuserResponse) GetLastUsedOk() (*int64, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *GetGcpuserResponse) SetLastUsed(v int64)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *GetGcpuserResponse) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GetGcpuserResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GetGcpuserResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GetGcpuserResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GetGcpuserResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyId

`func (o *GetGcpuserResponse) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *GetGcpuserResponse) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *GetGcpuserResponse) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.

### HasPrivateKeyId

`func (o *GetGcpuserResponse) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### GetProjectId

`func (o *GetGcpuserResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetGcpuserResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetGcpuserResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetGcpuserResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *GetGcpuserResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGcpuserResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGcpuserResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGcpuserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenUri

`func (o *GetGcpuserResponse) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *GetGcpuserResponse) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *GetGcpuserResponse) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *GetGcpuserResponse) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetType

`func (o *GetGcpuserResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGcpuserResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGcpuserResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetGcpuserResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserName

`func (o *GetGcpuserResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetGcpuserResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetGcpuserResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetGcpuserResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetResult

`func (o *GetGcpuserResponse) GetResult() Gcpuser`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGcpuserResponse) GetResultOk() (*Gcpuser, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGcpuserResponse) SetResult(v Gcpuser)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGcpuserResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


