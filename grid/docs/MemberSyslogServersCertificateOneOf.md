# MemberSyslogServersCertificateOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the Syslog endpoint server certificate. | [optional] [readonly] 
**Issuer** | Pointer to **string** | The certificate issuer subject name. | [optional] [readonly] 
**Serial** | Pointer to **string** | The certificate serial number in hex format. | [optional] [readonly] 
**Subject** | Pointer to **string** | The certificate subject name. | [optional] [readonly] 
**ValidNotAfter** | Pointer to **int32** | The date after which the certificate becomes invalid. | [optional] [readonly] 
**ValidNotBefore** | Pointer to **int32** | The date before which the certificate is not valid. | [optional] [readonly] 

## Methods

### NewMemberSyslogServersCertificateOneOf

`func NewMemberSyslogServersCertificateOneOf() *MemberSyslogServersCertificateOneOf`

NewMemberSyslogServersCertificateOneOf instantiates a new MemberSyslogServersCertificateOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSyslogServersCertificateOneOfWithDefaults

`func NewMemberSyslogServersCertificateOneOfWithDefaults() *MemberSyslogServersCertificateOneOf`

NewMemberSyslogServersCertificateOneOfWithDefaults instantiates a new MemberSyslogServersCertificateOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberSyslogServersCertificateOneOf) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberSyslogServersCertificateOneOf) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberSyslogServersCertificateOneOf) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberSyslogServersCertificateOneOf) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIssuer

`func (o *MemberSyslogServersCertificateOneOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MemberSyslogServersCertificateOneOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MemberSyslogServersCertificateOneOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *MemberSyslogServersCertificateOneOf) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerial

`func (o *MemberSyslogServersCertificateOneOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *MemberSyslogServersCertificateOneOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *MemberSyslogServersCertificateOneOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *MemberSyslogServersCertificateOneOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSubject

`func (o *MemberSyslogServersCertificateOneOf) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MemberSyslogServersCertificateOneOf) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MemberSyslogServersCertificateOneOf) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MemberSyslogServersCertificateOneOf) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidNotAfter

`func (o *MemberSyslogServersCertificateOneOf) GetValidNotAfter() int32`

GetValidNotAfter returns the ValidNotAfter field if non-nil, zero value otherwise.

### GetValidNotAfterOk

`func (o *MemberSyslogServersCertificateOneOf) GetValidNotAfterOk() (*int32, bool)`

GetValidNotAfterOk returns a tuple with the ValidNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotAfter

`func (o *MemberSyslogServersCertificateOneOf) SetValidNotAfter(v int32)`

SetValidNotAfter sets ValidNotAfter field to given value.

### HasValidNotAfter

`func (o *MemberSyslogServersCertificateOneOf) HasValidNotAfter() bool`

HasValidNotAfter returns a boolean if a field has been set.

### GetValidNotBefore

`func (o *MemberSyslogServersCertificateOneOf) GetValidNotBefore() int32`

GetValidNotBefore returns the ValidNotBefore field if non-nil, zero value otherwise.

### GetValidNotBeforeOk

`func (o *MemberSyslogServersCertificateOneOf) GetValidNotBeforeOk() (*int32, bool)`

GetValidNotBeforeOk returns a tuple with the ValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotBefore

`func (o *MemberSyslogServersCertificateOneOf) SetValidNotBefore(v int32)`

SetValidNotBefore sets ValidNotBefore field to given value.

### HasValidNotBefore

`func (o *MemberSyslogServersCertificateOneOf) HasValidNotBefore() bool`

HasValidNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


