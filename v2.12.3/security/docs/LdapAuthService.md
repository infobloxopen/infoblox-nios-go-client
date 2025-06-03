# LdapAuthService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CheckLdapServerSettings** | Pointer to **map[string]interface{}** |  | [optional] 
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

## Methods

### NewLdapAuthService

`func NewLdapAuthService() *LdapAuthService`

NewLdapAuthService instantiates a new LdapAuthService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAuthServiceWithDefaults

`func NewLdapAuthServiceWithDefaults() *LdapAuthService`

NewLdapAuthServiceWithDefaults instantiates a new LdapAuthService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *LdapAuthService) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *LdapAuthService) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *LdapAuthService) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *LdapAuthService) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCheckLdapServerSettings

`func (o *LdapAuthService) GetCheckLdapServerSettings() map[string]interface{}`

GetCheckLdapServerSettings returns the CheckLdapServerSettings field if non-nil, zero value otherwise.

### GetCheckLdapServerSettingsOk

`func (o *LdapAuthService) GetCheckLdapServerSettingsOk() (*map[string]interface{}, bool)`

GetCheckLdapServerSettingsOk returns a tuple with the CheckLdapServerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckLdapServerSettings

`func (o *LdapAuthService) SetCheckLdapServerSettings(v map[string]interface{})`

SetCheckLdapServerSettings sets CheckLdapServerSettings field to given value.

### HasCheckLdapServerSettings

`func (o *LdapAuthService) HasCheckLdapServerSettings() bool`

HasCheckLdapServerSettings returns a boolean if a field has been set.

### GetComment

`func (o *LdapAuthService) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LdapAuthService) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LdapAuthService) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LdapAuthService) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *LdapAuthService) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *LdapAuthService) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *LdapAuthService) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *LdapAuthService) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEaMapping

`func (o *LdapAuthService) GetEaMapping() []LdapAuthServiceEaMapping`

GetEaMapping returns the EaMapping field if non-nil, zero value otherwise.

### GetEaMappingOk

`func (o *LdapAuthService) GetEaMappingOk() (*[]LdapAuthServiceEaMapping, bool)`

GetEaMappingOk returns a tuple with the EaMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaMapping

`func (o *LdapAuthService) SetEaMapping(v []LdapAuthServiceEaMapping)`

SetEaMapping sets EaMapping field to given value.

### HasEaMapping

`func (o *LdapAuthService) HasEaMapping() bool`

HasEaMapping returns a boolean if a field has been set.

### GetLdapGroupAttribute

`func (o *LdapAuthService) GetLdapGroupAttribute() string`

GetLdapGroupAttribute returns the LdapGroupAttribute field if non-nil, zero value otherwise.

### GetLdapGroupAttributeOk

`func (o *LdapAuthService) GetLdapGroupAttributeOk() (*string, bool)`

GetLdapGroupAttributeOk returns a tuple with the LdapGroupAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAttribute

`func (o *LdapAuthService) SetLdapGroupAttribute(v string)`

SetLdapGroupAttribute sets LdapGroupAttribute field to given value.

### HasLdapGroupAttribute

`func (o *LdapAuthService) HasLdapGroupAttribute() bool`

HasLdapGroupAttribute returns a boolean if a field has been set.

### GetLdapGroupAuthenticationType

`func (o *LdapAuthService) GetLdapGroupAuthenticationType() string`

GetLdapGroupAuthenticationType returns the LdapGroupAuthenticationType field if non-nil, zero value otherwise.

### GetLdapGroupAuthenticationTypeOk

`func (o *LdapAuthService) GetLdapGroupAuthenticationTypeOk() (*string, bool)`

GetLdapGroupAuthenticationTypeOk returns a tuple with the LdapGroupAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAuthenticationType

`func (o *LdapAuthService) SetLdapGroupAuthenticationType(v string)`

SetLdapGroupAuthenticationType sets LdapGroupAuthenticationType field to given value.

### HasLdapGroupAuthenticationType

`func (o *LdapAuthService) HasLdapGroupAuthenticationType() bool`

HasLdapGroupAuthenticationType returns a boolean if a field has been set.

### GetLdapUserAttribute

`func (o *LdapAuthService) GetLdapUserAttribute() string`

GetLdapUserAttribute returns the LdapUserAttribute field if non-nil, zero value otherwise.

### GetLdapUserAttributeOk

`func (o *LdapAuthService) GetLdapUserAttributeOk() (*string, bool)`

GetLdapUserAttributeOk returns a tuple with the LdapUserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserAttribute

`func (o *LdapAuthService) SetLdapUserAttribute(v string)`

SetLdapUserAttribute sets LdapUserAttribute field to given value.

### HasLdapUserAttribute

`func (o *LdapAuthService) HasLdapUserAttribute() bool`

HasLdapUserAttribute returns a boolean if a field has been set.

### GetMode

`func (o *LdapAuthService) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LdapAuthService) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LdapAuthService) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LdapAuthService) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *LdapAuthService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapAuthService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapAuthService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapAuthService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecoveryInterval

`func (o *LdapAuthService) GetRecoveryInterval() int64`

GetRecoveryInterval returns the RecoveryInterval field if non-nil, zero value otherwise.

### GetRecoveryIntervalOk

`func (o *LdapAuthService) GetRecoveryIntervalOk() (*int64, bool)`

GetRecoveryIntervalOk returns a tuple with the RecoveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInterval

`func (o *LdapAuthService) SetRecoveryInterval(v int64)`

SetRecoveryInterval sets RecoveryInterval field to given value.

### HasRecoveryInterval

`func (o *LdapAuthService) HasRecoveryInterval() bool`

HasRecoveryInterval returns a boolean if a field has been set.

### GetRetries

`func (o *LdapAuthService) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *LdapAuthService) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *LdapAuthService) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *LdapAuthService) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetSearchScope

`func (o *LdapAuthService) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *LdapAuthService) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *LdapAuthService) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.

### HasSearchScope

`func (o *LdapAuthService) HasSearchScope() bool`

HasSearchScope returns a boolean if a field has been set.

### GetServers

`func (o *LdapAuthService) GetServers() []LdapAuthServiceServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *LdapAuthService) GetServersOk() (*[]LdapAuthServiceServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *LdapAuthService) SetServers(v []LdapAuthServiceServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *LdapAuthService) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTimeout

`func (o *LdapAuthService) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LdapAuthService) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LdapAuthService) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *LdapAuthService) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


