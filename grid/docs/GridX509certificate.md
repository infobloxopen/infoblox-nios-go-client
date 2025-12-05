# GridX509certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Issuer** | Pointer to **string** | Certificate issuer. | [optional] [readonly] 
**Serial** | Pointer to **string** | X509Certificate serial number. | [optional] [readonly] 
**Subject** | Pointer to **string** | A Distinguished Name that is made of multiple relative distinguished names (RDNs). | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**ValidNotAfter** | Pointer to **int64** | Certificate expiry date. | [optional] [readonly] 
**ValidNotBefore** | Pointer to **int64** | Certificate validity start date. | [optional] [readonly] 

## Methods

### NewGridX509certificate

`func NewGridX509certificate() *GridX509certificate`

NewGridX509certificate instantiates a new GridX509certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridX509certificateWithDefaults

`func NewGridX509certificateWithDefaults() *GridX509certificate`

NewGridX509certificateWithDefaults instantiates a new GridX509certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridX509certificate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridX509certificate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridX509certificate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridX509certificate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIssuer

`func (o *GridX509certificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GridX509certificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GridX509certificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *GridX509certificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerial

`func (o *GridX509certificate) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GridX509certificate) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GridX509certificate) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GridX509certificate) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSubject

`func (o *GridX509certificate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GridX509certificate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GridX509certificate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GridX509certificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetUuid

`func (o *GridX509certificate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GridX509certificate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GridX509certificate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GridX509certificate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValidNotAfter

`func (o *GridX509certificate) GetValidNotAfter() int64`

GetValidNotAfter returns the ValidNotAfter field if non-nil, zero value otherwise.

### GetValidNotAfterOk

`func (o *GridX509certificate) GetValidNotAfterOk() (*int64, bool)`

GetValidNotAfterOk returns a tuple with the ValidNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotAfter

`func (o *GridX509certificate) SetValidNotAfter(v int64)`

SetValidNotAfter sets ValidNotAfter field to given value.

### HasValidNotAfter

`func (o *GridX509certificate) HasValidNotAfter() bool`

HasValidNotAfter returns a boolean if a field has been set.

### GetValidNotBefore

`func (o *GridX509certificate) GetValidNotBefore() int64`

GetValidNotBefore returns the ValidNotBefore field if non-nil, zero value otherwise.

### GetValidNotBeforeOk

`func (o *GridX509certificate) GetValidNotBeforeOk() (*int64, bool)`

GetValidNotBeforeOk returns a tuple with the ValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotBefore

`func (o *GridX509certificate) SetValidNotBefore(v int64)`

SetValidNotBefore sets ValidNotBefore field to given value.

### HasValidNotBefore

`func (o *GridX509certificate) HasValidNotBefore() bool`

HasValidNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


