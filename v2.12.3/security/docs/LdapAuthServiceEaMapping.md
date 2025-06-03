# LdapAuthServiceEaMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The LDAP attribute name. | [optional] 
**MappedEa** | Pointer to **string** | The name of the extensible attribute definition object to which the LDAP attribute is mapped. | [optional] 

## Methods

### NewLdapAuthServiceEaMapping

`func NewLdapAuthServiceEaMapping() *LdapAuthServiceEaMapping`

NewLdapAuthServiceEaMapping instantiates a new LdapAuthServiceEaMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAuthServiceEaMappingWithDefaults

`func NewLdapAuthServiceEaMappingWithDefaults() *LdapAuthServiceEaMapping`

NewLdapAuthServiceEaMappingWithDefaults instantiates a new LdapAuthServiceEaMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LdapAuthServiceEaMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapAuthServiceEaMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapAuthServiceEaMapping) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapAuthServiceEaMapping) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMappedEa

`func (o *LdapAuthServiceEaMapping) GetMappedEa() string`

GetMappedEa returns the MappedEa field if non-nil, zero value otherwise.

### GetMappedEaOk

`func (o *LdapAuthServiceEaMapping) GetMappedEaOk() (*string, bool)`

GetMappedEaOk returns a tuple with the MappedEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedEa

`func (o *LdapAuthServiceEaMapping) SetMappedEa(v string)`

SetMappedEa sets MappedEa field to given value.

### HasMappedEa

`func (o *LdapAuthServiceEaMapping) HasMappedEa() bool`

HasMappedEa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


