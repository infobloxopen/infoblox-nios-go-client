# GridDnsClientSubnetDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The FQDN that represents the ECS zone domain name. | [optional] 
**Permission** | Pointer to **string** | The ECS domain name permission. | [optional] 

## Methods

### NewGridDnsClientSubnetDomains

`func NewGridDnsClientSubnetDomains() *GridDnsClientSubnetDomains`

NewGridDnsClientSubnetDomains instantiates a new GridDnsClientSubnetDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsClientSubnetDomainsWithDefaults

`func NewGridDnsClientSubnetDomainsWithDefaults() *GridDnsClientSubnetDomains`

NewGridDnsClientSubnetDomainsWithDefaults instantiates a new GridDnsClientSubnetDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *GridDnsClientSubnetDomains) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GridDnsClientSubnetDomains) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GridDnsClientSubnetDomains) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GridDnsClientSubnetDomains) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPermission

`func (o *GridDnsClientSubnetDomains) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *GridDnsClientSubnetDomains) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *GridDnsClientSubnetDomains) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *GridDnsClientSubnetDomains) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


