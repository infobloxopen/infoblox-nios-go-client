# DtcCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Certificate** | Pointer to **string** | Reference to underlying X509Certificate. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Determines whether the certificate is in use or not. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewDtcCertificate

`func NewDtcCertificate() *DtcCertificate`

NewDtcCertificate instantiates a new DtcCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcCertificateWithDefaults

`func NewDtcCertificateWithDefaults() *DtcCertificate`

NewDtcCertificateWithDefaults instantiates a new DtcCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcCertificate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcCertificate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcCertificate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcCertificate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCertificate

`func (o *DtcCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *DtcCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *DtcCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *DtcCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetInUse

`func (o *DtcCertificate) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *DtcCertificate) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *DtcCertificate) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *DtcCertificate) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetUuid

`func (o *DtcCertificate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DtcCertificate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DtcCertificate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DtcCertificate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


