# Authpolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AdminGroups** | Pointer to **[]string** | List of names of local administration groups that are mapped to remote administration groups. | [optional] 
**AuthServices** | Pointer to **[]string** | The array that contains an ordered list of refs to :doc:&#x60;localuser:authservice object &lt;/objects/localuser.authservice&gt;&#x60;, ldap_auth_service object ldap_auth_service, :doc:&#x60;radius:authservice object &lt;/objects/radius.authservice&gt;&#x60;, :doc:&#x60;tacacsplus:authservice object &lt;/objects/tacacsplus.authservice&gt;&#x60;, ad_auth_service object ad_auth_service, :doc:&#x60;certificate:authservice object &lt;/objects/certificate.authservice&gt;&#x60;. :doc:&#x60;saml:authservice object &lt;/objects/saml.authservice&gt;&#x60;, | [optional] 
**DefaultGroup** | Pointer to **string** | The default admin group that provides authentication in case no valid group is found. | [optional] 
**UsageType** | Pointer to **string** | Remote policies usage. | [optional] 

## Methods

### NewAuthpolicy

`func NewAuthpolicy() *Authpolicy`

NewAuthpolicy instantiates a new Authpolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthpolicyWithDefaults

`func NewAuthpolicyWithDefaults() *Authpolicy`

NewAuthpolicyWithDefaults instantiates a new Authpolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Authpolicy) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Authpolicy) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Authpolicy) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Authpolicy) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdminGroups

`func (o *Authpolicy) GetAdminGroups() []string`

GetAdminGroups returns the AdminGroups field if non-nil, zero value otherwise.

### GetAdminGroupsOk

`func (o *Authpolicy) GetAdminGroupsOk() (*[]string, bool)`

GetAdminGroupsOk returns a tuple with the AdminGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroups

`func (o *Authpolicy) SetAdminGroups(v []string)`

SetAdminGroups sets AdminGroups field to given value.

### HasAdminGroups

`func (o *Authpolicy) HasAdminGroups() bool`

HasAdminGroups returns a boolean if a field has been set.

### GetAuthServices

`func (o *Authpolicy) GetAuthServices() []string`

GetAuthServices returns the AuthServices field if non-nil, zero value otherwise.

### GetAuthServicesOk

`func (o *Authpolicy) GetAuthServicesOk() (*[]string, bool)`

GetAuthServicesOk returns a tuple with the AuthServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServices

`func (o *Authpolicy) SetAuthServices(v []string)`

SetAuthServices sets AuthServices field to given value.

### HasAuthServices

`func (o *Authpolicy) HasAuthServices() bool`

HasAuthServices returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *Authpolicy) GetDefaultGroup() string`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *Authpolicy) GetDefaultGroupOk() (*string, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *Authpolicy) SetDefaultGroup(v string)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *Authpolicy) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetUsageType

`func (o *Authpolicy) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Authpolicy) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Authpolicy) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Authpolicy) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


