# NetworkviewRemoteForwardZones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | The FQDN of the remote server. | [optional] 
**ServerAddress** | Pointer to **string** | The remote server IP address. | [optional] 
**GssTsigDnsPrincipal** | Pointer to **string** | The principal name in which GSS-TSIG for dynamic updates is enabled. | [optional] 
**GssTsigDomain** | Pointer to **string** | The domain in which GSS-TSIG for dynamic updates is enabled. | [optional] 
**TsigKey** | Pointer to **string** | The TSIG key value. | [optional] 
**TsigKeyAlg** | Pointer to **string** | The TSIG key alorithm name. | [optional] 
**TsigKeyName** | Pointer to **string** | The name of the TSIG key. The key name entered here must match the TSIG key name on the external name server. | [optional] 
**KeyType** | Pointer to **string** | The key type to be used. | [optional] 

## Methods

### NewNetworkviewRemoteForwardZones

`func NewNetworkviewRemoteForwardZones() *NetworkviewRemoteForwardZones`

NewNetworkviewRemoteForwardZones instantiates a new NetworkviewRemoteForwardZones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkviewRemoteForwardZonesWithDefaults

`func NewNetworkviewRemoteForwardZonesWithDefaults() *NetworkviewRemoteForwardZones`

NewNetworkviewRemoteForwardZonesWithDefaults instantiates a new NetworkviewRemoteForwardZones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *NetworkviewRemoteForwardZones) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NetworkviewRemoteForwardZones) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NetworkviewRemoteForwardZones) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NetworkviewRemoteForwardZones) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetServerAddress

`func (o *NetworkviewRemoteForwardZones) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *NetworkviewRemoteForwardZones) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *NetworkviewRemoteForwardZones) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *NetworkviewRemoteForwardZones) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetGssTsigDnsPrincipal

`func (o *NetworkviewRemoteForwardZones) GetGssTsigDnsPrincipal() string`

GetGssTsigDnsPrincipal returns the GssTsigDnsPrincipal field if non-nil, zero value otherwise.

### GetGssTsigDnsPrincipalOk

`func (o *NetworkviewRemoteForwardZones) GetGssTsigDnsPrincipalOk() (*string, bool)`

GetGssTsigDnsPrincipalOk returns a tuple with the GssTsigDnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigDnsPrincipal

`func (o *NetworkviewRemoteForwardZones) SetGssTsigDnsPrincipal(v string)`

SetGssTsigDnsPrincipal sets GssTsigDnsPrincipal field to given value.

### HasGssTsigDnsPrincipal

`func (o *NetworkviewRemoteForwardZones) HasGssTsigDnsPrincipal() bool`

HasGssTsigDnsPrincipal returns a boolean if a field has been set.

### GetGssTsigDomain

`func (o *NetworkviewRemoteForwardZones) GetGssTsigDomain() string`

GetGssTsigDomain returns the GssTsigDomain field if non-nil, zero value otherwise.

### GetGssTsigDomainOk

`func (o *NetworkviewRemoteForwardZones) GetGssTsigDomainOk() (*string, bool)`

GetGssTsigDomainOk returns a tuple with the GssTsigDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigDomain

`func (o *NetworkviewRemoteForwardZones) SetGssTsigDomain(v string)`

SetGssTsigDomain sets GssTsigDomain field to given value.

### HasGssTsigDomain

`func (o *NetworkviewRemoteForwardZones) HasGssTsigDomain() bool`

HasGssTsigDomain returns a boolean if a field has been set.

### GetTsigKey

`func (o *NetworkviewRemoteForwardZones) GetTsigKey() string`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *NetworkviewRemoteForwardZones) GetTsigKeyOk() (*string, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *NetworkviewRemoteForwardZones) SetTsigKey(v string)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *NetworkviewRemoteForwardZones) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetTsigKeyAlg

`func (o *NetworkviewRemoteForwardZones) GetTsigKeyAlg() string`

GetTsigKeyAlg returns the TsigKeyAlg field if non-nil, zero value otherwise.

### GetTsigKeyAlgOk

`func (o *NetworkviewRemoteForwardZones) GetTsigKeyAlgOk() (*string, bool)`

GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyAlg

`func (o *NetworkviewRemoteForwardZones) SetTsigKeyAlg(v string)`

SetTsigKeyAlg sets TsigKeyAlg field to given value.

### HasTsigKeyAlg

`func (o *NetworkviewRemoteForwardZones) HasTsigKeyAlg() bool`

HasTsigKeyAlg returns a boolean if a field has been set.

### GetTsigKeyName

`func (o *NetworkviewRemoteForwardZones) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *NetworkviewRemoteForwardZones) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *NetworkviewRemoteForwardZones) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.

### HasTsigKeyName

`func (o *NetworkviewRemoteForwardZones) HasTsigKeyName() bool`

HasTsigKeyName returns a boolean if a field has been set.

### GetKeyType

`func (o *NetworkviewRemoteForwardZones) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *NetworkviewRemoteForwardZones) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *NetworkviewRemoteForwardZones) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *NetworkviewRemoteForwardZones) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


