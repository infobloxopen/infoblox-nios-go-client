# SamlAuthserviceIdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpType** | Pointer to **string** | SAML Identity Provider type. | [optional] 
**Comment** | Pointer to **string** | The SAML Identity Provider descriptive comment. | [optional] 
**MetadataUrl** | Pointer to **string** | Identity Provider Metadata URL. | [optional] 
**MetadataToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop. | [optional] 
**Groupname** | Pointer to **string** | The SAML groupname optional user group attribute. | [optional] 
**SsoRedirectUrl** | Pointer to **string** | host name or IP address of the GM | [optional] 

## Methods

### NewSamlAuthserviceIdp

`func NewSamlAuthserviceIdp() *SamlAuthserviceIdp`

NewSamlAuthserviceIdp instantiates a new SamlAuthserviceIdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAuthserviceIdpWithDefaults

`func NewSamlAuthserviceIdpWithDefaults() *SamlAuthserviceIdp`

NewSamlAuthserviceIdpWithDefaults instantiates a new SamlAuthserviceIdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpType

`func (o *SamlAuthserviceIdp) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *SamlAuthserviceIdp) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *SamlAuthserviceIdp) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *SamlAuthserviceIdp) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.

### GetComment

`func (o *SamlAuthserviceIdp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SamlAuthserviceIdp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SamlAuthserviceIdp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SamlAuthserviceIdp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *SamlAuthserviceIdp) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *SamlAuthserviceIdp) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *SamlAuthserviceIdp) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *SamlAuthserviceIdp) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetMetadataToken

`func (o *SamlAuthserviceIdp) GetMetadataToken() string`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *SamlAuthserviceIdp) GetMetadataTokenOk() (*string, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *SamlAuthserviceIdp) SetMetadataToken(v string)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *SamlAuthserviceIdp) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetGroupname

`func (o *SamlAuthserviceIdp) GetGroupname() string`

GetGroupname returns the Groupname field if non-nil, zero value otherwise.

### GetGroupnameOk

`func (o *SamlAuthserviceIdp) GetGroupnameOk() (*string, bool)`

GetGroupnameOk returns a tuple with the Groupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupname

`func (o *SamlAuthserviceIdp) SetGroupname(v string)`

SetGroupname sets Groupname field to given value.

### HasGroupname

`func (o *SamlAuthserviceIdp) HasGroupname() bool`

HasGroupname returns a boolean if a field has been set.

### GetSsoRedirectUrl

`func (o *SamlAuthserviceIdp) GetSsoRedirectUrl() string`

GetSsoRedirectUrl returns the SsoRedirectUrl field if non-nil, zero value otherwise.

### GetSsoRedirectUrlOk

`func (o *SamlAuthserviceIdp) GetSsoRedirectUrlOk() (*string, bool)`

GetSsoRedirectUrlOk returns a tuple with the SsoRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoRedirectUrl

`func (o *SamlAuthserviceIdp) SetSsoRedirectUrl(v string)`

SetSsoRedirectUrl sets SsoRedirectUrl field to given value.

### HasSsoRedirectUrl

`func (o *SamlAuthserviceIdp) HasSsoRedirectUrl() bool`

HasSsoRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


