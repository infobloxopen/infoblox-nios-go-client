# GetDtcCertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Certificate** | Pointer to **string** | Reference to underlying X509Certificate. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Determines whether the certificate is in use or not. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DtcCertificate**](DtcCertificate.md) |  | [optional] 

## Methods

### NewGetDtcCertificateResponse

`func NewGetDtcCertificateResponse() *GetDtcCertificateResponse`

NewGetDtcCertificateResponse instantiates a new GetDtcCertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcCertificateResponseWithDefaults

`func NewGetDtcCertificateResponseWithDefaults() *GetDtcCertificateResponse`

NewGetDtcCertificateResponseWithDefaults instantiates a new GetDtcCertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcCertificateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcCertificateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcCertificateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcCertificateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCertificate

`func (o *GetDtcCertificateResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *GetDtcCertificateResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *GetDtcCertificateResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *GetDtcCertificateResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetInUse

`func (o *GetDtcCertificateResponse) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *GetDtcCertificateResponse) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *GetDtcCertificateResponse) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *GetDtcCertificateResponse) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetUuid

`func (o *GetDtcCertificateResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDtcCertificateResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDtcCertificateResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDtcCertificateResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcCertificateResponse) GetResult() DtcCertificate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcCertificateResponse) GetResultOk() (*DtcCertificate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcCertificateResponse) SetResult(v DtcCertificate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcCertificateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


