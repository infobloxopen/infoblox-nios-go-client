# GridDnsDnssecTrustedKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | The FQDN of the domain for which the member validates responses to recursive queries. | [optional] 
**Algorithm** | Pointer to **string** | The DNSSEC algorithm used to generate the key. | [optional] 
**Key** | Pointer to **string** | The DNSSEC key. | [optional] 
**SecureEntryPoint** | Pointer to **bool** | The secure entry point flag, if set it means this is a KSK configuration. | [optional] 
**DnssecMustBeSecure** | Pointer to **bool** | Responses must be DNSSEC secure for this hierarchy/domain. | [optional] 

## Methods

### NewGridDnsDnssecTrustedKeys

`func NewGridDnsDnssecTrustedKeys() *GridDnsDnssecTrustedKeys`

NewGridDnsDnssecTrustedKeys instantiates a new GridDnsDnssecTrustedKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsDnssecTrustedKeysWithDefaults

`func NewGridDnsDnssecTrustedKeysWithDefaults() *GridDnsDnssecTrustedKeys`

NewGridDnsDnssecTrustedKeysWithDefaults instantiates a new GridDnsDnssecTrustedKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *GridDnsDnssecTrustedKeys) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GridDnsDnssecTrustedKeys) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GridDnsDnssecTrustedKeys) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GridDnsDnssecTrustedKeys) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetAlgorithm

`func (o *GridDnsDnssecTrustedKeys) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GridDnsDnssecTrustedKeys) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GridDnsDnssecTrustedKeys) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *GridDnsDnssecTrustedKeys) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKey

`func (o *GridDnsDnssecTrustedKeys) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GridDnsDnssecTrustedKeys) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GridDnsDnssecTrustedKeys) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GridDnsDnssecTrustedKeys) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecureEntryPoint

`func (o *GridDnsDnssecTrustedKeys) GetSecureEntryPoint() bool`

GetSecureEntryPoint returns the SecureEntryPoint field if non-nil, zero value otherwise.

### GetSecureEntryPointOk

`func (o *GridDnsDnssecTrustedKeys) GetSecureEntryPointOk() (*bool, bool)`

GetSecureEntryPointOk returns a tuple with the SecureEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureEntryPoint

`func (o *GridDnsDnssecTrustedKeys) SetSecureEntryPoint(v bool)`

SetSecureEntryPoint sets SecureEntryPoint field to given value.

### HasSecureEntryPoint

`func (o *GridDnsDnssecTrustedKeys) HasSecureEntryPoint() bool`

HasSecureEntryPoint returns a boolean if a field has been set.

### GetDnssecMustBeSecure

`func (o *GridDnsDnssecTrustedKeys) GetDnssecMustBeSecure() bool`

GetDnssecMustBeSecure returns the DnssecMustBeSecure field if non-nil, zero value otherwise.

### GetDnssecMustBeSecureOk

`func (o *GridDnsDnssecTrustedKeys) GetDnssecMustBeSecureOk() (*bool, bool)`

GetDnssecMustBeSecureOk returns a tuple with the DnssecMustBeSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecMustBeSecure

`func (o *GridDnsDnssecTrustedKeys) SetDnssecMustBeSecure(v bool)`

SetDnssecMustBeSecure sets DnssecMustBeSecure field to given value.

### HasDnssecMustBeSecure

`func (o *GridDnsDnssecTrustedKeys) HasDnssecMustBeSecure() bool`

HasDnssecMustBeSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


