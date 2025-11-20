# GetGridX509certificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Issuer** | Pointer to **string** | Certificate issuer. | [optional] [readonly] 
**Serial** | Pointer to **string** | X509Certificate serial number. | [optional] [readonly] 
**Subject** | Pointer to **string** | A Distinguished Name that is made of multiple relative distinguished names (RDNs). | [optional] [readonly] 
**ValidNotAfter** | Pointer to **int64** | Certificate expiry date. | [optional] [readonly] 
**ValidNotBefore** | Pointer to **int64** | Certificate validity start date. | [optional] [readonly] 
**Result** | Pointer to [**GridX509certificate**](GridX509certificate.md) |  | [optional] 

## Methods

### NewGetGridX509certificateResponse

`func NewGetGridX509certificateResponse() *GetGridX509certificateResponse`

NewGetGridX509certificateResponse instantiates a new GetGridX509certificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridX509certificateResponseWithDefaults

`func NewGetGridX509certificateResponseWithDefaults() *GetGridX509certificateResponse`

NewGetGridX509certificateResponseWithDefaults instantiates a new GetGridX509certificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridX509certificateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridX509certificateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridX509certificateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridX509certificateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIssuer

`func (o *GetGridX509certificateResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetGridX509certificateResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetGridX509certificateResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *GetGridX509certificateResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerial

`func (o *GetGridX509certificateResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetGridX509certificateResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetGridX509certificateResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetGridX509certificateResponse) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSubject

`func (o *GetGridX509certificateResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetGridX509certificateResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetGridX509certificateResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetGridX509certificateResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidNotAfter

`func (o *GetGridX509certificateResponse) GetValidNotAfter() int64`

GetValidNotAfter returns the ValidNotAfter field if non-nil, zero value otherwise.

### GetValidNotAfterOk

`func (o *GetGridX509certificateResponse) GetValidNotAfterOk() (*int64, bool)`

GetValidNotAfterOk returns a tuple with the ValidNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotAfter

`func (o *GetGridX509certificateResponse) SetValidNotAfter(v int64)`

SetValidNotAfter sets ValidNotAfter field to given value.

### HasValidNotAfter

`func (o *GetGridX509certificateResponse) HasValidNotAfter() bool`

HasValidNotAfter returns a boolean if a field has been set.

### GetValidNotBefore

`func (o *GetGridX509certificateResponse) GetValidNotBefore() int64`

GetValidNotBefore returns the ValidNotBefore field if non-nil, zero value otherwise.

### GetValidNotBeforeOk

`func (o *GetGridX509certificateResponse) GetValidNotBeforeOk() (*int64, bool)`

GetValidNotBeforeOk returns a tuple with the ValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotBefore

`func (o *GetGridX509certificateResponse) SetValidNotBefore(v int64)`

SetValidNotBefore sets ValidNotBefore field to given value.

### HasValidNotBefore

`func (o *GetGridX509certificateResponse) HasValidNotBefore() bool`

HasValidNotBefore returns a boolean if a field has been set.

### GetResult

`func (o *GetGridX509certificateResponse) GetResult() GridX509certificate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridX509certificateResponse) GetResultOk() (*GridX509certificate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridX509certificateResponse) SetResult(v GridX509certificate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridX509certificateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


