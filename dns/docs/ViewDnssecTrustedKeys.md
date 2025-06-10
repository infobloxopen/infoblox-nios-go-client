# ViewDnssecTrustedKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | The FQDN of the domain for which the member validates responses to recursive queries. | [optional] 
**Algorithm** | Pointer to **string** | The DNSSEC algorithm used to generate the key. | [optional] 
**Key** | Pointer to **string** | The DNSSEC key. | [optional] 
**SecureEntryPoint** | Pointer to **bool** | The secure entry point flag, if set it means this is a KSK configuration. | [optional] 
**DnssecMustBeSecure** | Pointer to **bool** | Responses must be DNSSEC secure for this hierarchy/domain. | [optional] 

## Methods

### NewViewDnssecTrustedKeys

`func NewViewDnssecTrustedKeys() *ViewDnssecTrustedKeys`

NewViewDnssecTrustedKeys instantiates a new ViewDnssecTrustedKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewDnssecTrustedKeysWithDefaults

`func NewViewDnssecTrustedKeysWithDefaults() *ViewDnssecTrustedKeys`

NewViewDnssecTrustedKeysWithDefaults instantiates a new ViewDnssecTrustedKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *ViewDnssecTrustedKeys) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ViewDnssecTrustedKeys) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ViewDnssecTrustedKeys) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ViewDnssecTrustedKeys) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetAlgorithm

`func (o *ViewDnssecTrustedKeys) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ViewDnssecTrustedKeys) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ViewDnssecTrustedKeys) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ViewDnssecTrustedKeys) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKey

`func (o *ViewDnssecTrustedKeys) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ViewDnssecTrustedKeys) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ViewDnssecTrustedKeys) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ViewDnssecTrustedKeys) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecureEntryPoint

`func (o *ViewDnssecTrustedKeys) GetSecureEntryPoint() bool`

GetSecureEntryPoint returns the SecureEntryPoint field if non-nil, zero value otherwise.

### GetSecureEntryPointOk

`func (o *ViewDnssecTrustedKeys) GetSecureEntryPointOk() (*bool, bool)`

GetSecureEntryPointOk returns a tuple with the SecureEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureEntryPoint

`func (o *ViewDnssecTrustedKeys) SetSecureEntryPoint(v bool)`

SetSecureEntryPoint sets SecureEntryPoint field to given value.

### HasSecureEntryPoint

`func (o *ViewDnssecTrustedKeys) HasSecureEntryPoint() bool`

HasSecureEntryPoint returns a boolean if a field has been set.

### GetDnssecMustBeSecure

`func (o *ViewDnssecTrustedKeys) GetDnssecMustBeSecure() bool`

GetDnssecMustBeSecure returns the DnssecMustBeSecure field if non-nil, zero value otherwise.

### GetDnssecMustBeSecureOk

`func (o *ViewDnssecTrustedKeys) GetDnssecMustBeSecureOk() (*bool, bool)`

GetDnssecMustBeSecureOk returns a tuple with the DnssecMustBeSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecMustBeSecure

`func (o *ViewDnssecTrustedKeys) SetDnssecMustBeSecure(v bool)`

SetDnssecMustBeSecure sets DnssecMustBeSecure field to given value.

### HasDnssecMustBeSecure

`func (o *ViewDnssecTrustedKeys) HasDnssecMustBeSecure() bool`

HasDnssecMustBeSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


