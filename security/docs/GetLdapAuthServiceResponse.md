# GetLdapAuthServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The LDAP descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines if the LDAP authentication service is disabled. | [optional] 
**EaMapping** | Pointer to [**[]LdapAuthServiceEaMapping**](LdapAuthServiceEaMapping.md) | The mapping LDAP fields to extensible attributes. | [optional] 
**LdapGroupAttribute** | Pointer to **string** | The name of the LDAP attribute that defines group membership. | [optional] 
**LdapGroupAuthenticationType** | Pointer to **string** | The LDAP group authentication type. | [optional] 
**LdapUserAttribute** | Pointer to **string** | The LDAP userid attribute that is used for search. | [optional] 
**Mode** | Pointer to **string** | The LDAP authentication mode. | [optional] 
**Name** | Pointer to **string** | The LDAP authentication service name. | [optional] 
**RecoveryInterval** | Pointer to **int64** | The period of time in seconds to wait before trying to contact a LDAP server that has been marked as &#39;DOWN&#39;. | [optional] 
**Retries** | Pointer to **int64** | The maximum number of LDAP authentication attempts. | [optional] 
**SearchScope** | Pointer to **string** | The starting point of the LDAP search. | [optional] 
**Servers** | Pointer to [**[]LdapAuthServiceServers**](LdapAuthServiceServers.md) | The list of LDAP servers used for authentication. | [optional] 
**Timeout** | Pointer to **int64** | The LDAP authentication timeout in seconds. | [optional] 
**Result** | Pointer to [**LdapAuthService**](LdapAuthService.md) |  | [optional] 

## Methods

### NewGetLdapAuthServiceResponse

`func NewGetLdapAuthServiceResponse() *GetLdapAuthServiceResponse`

NewGetLdapAuthServiceResponse instantiates a new GetLdapAuthServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLdapAuthServiceResponseWithDefaults

`func NewGetLdapAuthServiceResponseWithDefaults() *GetLdapAuthServiceResponse`

NewGetLdapAuthServiceResponseWithDefaults instantiates a new GetLdapAuthServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetLdapAuthServiceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetLdapAuthServiceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetLdapAuthServiceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetLdapAuthServiceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetLdapAuthServiceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetLdapAuthServiceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetLdapAuthServiceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetLdapAuthServiceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetLdapAuthServiceResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetLdapAuthServiceResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetLdapAuthServiceResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetLdapAuthServiceResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEaMapping

`func (o *GetLdapAuthServiceResponse) GetEaMapping() []LdapAuthServiceEaMapping`

GetEaMapping returns the EaMapping field if non-nil, zero value otherwise.

### GetEaMappingOk

`func (o *GetLdapAuthServiceResponse) GetEaMappingOk() (*[]LdapAuthServiceEaMapping, bool)`

GetEaMappingOk returns a tuple with the EaMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaMapping

`func (o *GetLdapAuthServiceResponse) SetEaMapping(v []LdapAuthServiceEaMapping)`

SetEaMapping sets EaMapping field to given value.

### HasEaMapping

`func (o *GetLdapAuthServiceResponse) HasEaMapping() bool`

HasEaMapping returns a boolean if a field has been set.

### GetLdapGroupAttribute

`func (o *GetLdapAuthServiceResponse) GetLdapGroupAttribute() string`

GetLdapGroupAttribute returns the LdapGroupAttribute field if non-nil, zero value otherwise.

### GetLdapGroupAttributeOk

`func (o *GetLdapAuthServiceResponse) GetLdapGroupAttributeOk() (*string, bool)`

GetLdapGroupAttributeOk returns a tuple with the LdapGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAttribute

`func (o *GetLdapAuthServiceResponse) SetLdapGroupAttribute(v string)`

SetLdapGroupAttribute sets LdapGroupAttribute field to given value.

### HasLdapGroupAttribute

`func (o *GetLdapAuthServiceResponse) HasLdapGroupAttribute() bool`

HasLdapGroupAttribute returns a boolean if a field has been set.

### GetLdapGroupAuthenticationType

`func (o *GetLdapAuthServiceResponse) GetLdapGroupAuthenticationType() string`

GetLdapGroupAuthenticationType returns the LdapGroupAuthenticationType field if non-nil, zero value otherwise.

### GetLdapGroupAuthenticationTypeOk

`func (o *GetLdapAuthServiceResponse) GetLdapGroupAuthenticationTypeOk() (*string, bool)`

GetLdapGroupAuthenticationTypeOk returns a tuple with the LdapGroupAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAuthenticationType

`func (o *GetLdapAuthServiceResponse) SetLdapGroupAuthenticationType(v string)`

SetLdapGroupAuthenticationType sets LdapGroupAuthenticationType field to given value.

### HasLdapGroupAuthenticationType

`func (o *GetLdapAuthServiceResponse) HasLdapGroupAuthenticationType() bool`

HasLdapGroupAuthenticationType returns a boolean if a field has been set.

### GetLdapUserAttribute

`func (o *GetLdapAuthServiceResponse) GetLdapUserAttribute() string`

GetLdapUserAttribute returns the LdapUserAttribute field if non-nil, zero value otherwise.

### GetLdapUserAttributeOk

`func (o *GetLdapAuthServiceResponse) GetLdapUserAttributeOk() (*string, bool)`

GetLdapUserAttributeOk returns a tuple with the LdapUserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserAttribute

`func (o *GetLdapAuthServiceResponse) SetLdapUserAttribute(v string)`

SetLdapUserAttribute sets LdapUserAttribute field to given value.

### HasLdapUserAttribute

`func (o *GetLdapAuthServiceResponse) HasLdapUserAttribute() bool`

HasLdapUserAttribute returns a boolean if a field has been set.

### GetMode

`func (o *GetLdapAuthServiceResponse) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetLdapAuthServiceResponse) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetLdapAuthServiceResponse) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetLdapAuthServiceResponse) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GetLdapAuthServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLdapAuthServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLdapAuthServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLdapAuthServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecoveryInterval

`func (o *GetLdapAuthServiceResponse) GetRecoveryInterval() int64`

GetRecoveryInterval returns the RecoveryInterval field if non-nil, zero value otherwise.

### GetRecoveryIntervalOk

`func (o *GetLdapAuthServiceResponse) GetRecoveryIntervalOk() (*int64, bool)`

GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInterval

`func (o *GetLdapAuthServiceResponse) SetRecoveryInterval(v int64)`

SetRecoveryInterval sets RecoveryInterval field to given value.

### HasRecoveryInterval

`func (o *GetLdapAuthServiceResponse) HasRecoveryInterval() bool`

HasRecoveryInterval returns a boolean if a field has been set.

### GetRetries

`func (o *GetLdapAuthServiceResponse) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *GetLdapAuthServiceResponse) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *GetLdapAuthServiceResponse) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *GetLdapAuthServiceResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetSearchScope

`func (o *GetLdapAuthServiceResponse) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *GetLdapAuthServiceResponse) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *GetLdapAuthServiceResponse) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.

### HasSearchScope

`func (o *GetLdapAuthServiceResponse) HasSearchScope() bool`

HasSearchScope returns a boolean if a field has been set.

### GetServers

`func (o *GetLdapAuthServiceResponse) GetServers() []LdapAuthServiceServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetLdapAuthServiceResponse) GetServersOk() (*[]LdapAuthServiceServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetLdapAuthServiceResponse) SetServers(v []LdapAuthServiceServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetLdapAuthServiceResponse) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTimeout

`func (o *GetLdapAuthServiceResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetLdapAuthServiceResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetLdapAuthServiceResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetLdapAuthServiceResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetResult

`func (o *GetLdapAuthServiceResponse) GetResult() LdapAuthService`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetLdapAuthServiceResponse) GetResultOk() (*LdapAuthService, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetLdapAuthServiceResponse) SetResult(v LdapAuthService)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetLdapAuthServiceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


