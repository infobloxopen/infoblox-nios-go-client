# MemberSyslogServersCertificate

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

### NewMemberSyslogServersCertificate

`func NewMemberSyslogServersCertificate() *MemberSyslogServersCertificate`

NewMemberSyslogServersCertificate instantiates a new MemberSyslogServersCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSyslogServersCertificateWithDefaults

`func NewMemberSyslogServersCertificateWithDefaults() *MemberSyslogServersCertificate`

NewMemberSyslogServersCertificateWithDefaults instantiates a new MemberSyslogServersCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberSyslogServersCertificate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberSyslogServersCertificate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberSyslogServersCertificate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberSyslogServersCertificate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIssuer

`func (o *MemberSyslogServersCertificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MemberSyslogServersCertificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MemberSyslogServersCertificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *MemberSyslogServersCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerial

`func (o *MemberSyslogServersCertificate) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *MemberSyslogServersCertificate) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *MemberSyslogServersCertificate) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *MemberSyslogServersCertificate) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSubject

`func (o *MemberSyslogServersCertificate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MemberSyslogServersCertificate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MemberSyslogServersCertificate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MemberSyslogServersCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidNotAfter

`func (o *MemberSyslogServersCertificate) GetValidNotAfter() int32`

GetValidNotAfter returns the ValidNotAfter field if non-nil, zero value otherwise.

### GetValidNotAfterOk

`func (o *MemberSyslogServersCertificate) GetValidNotAfterOk() (*int32, bool)`

GetValidNotAfterOk returns a tuple with the ValidNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotAfter

`func (o *MemberSyslogServersCertificate) SetValidNotAfter(v int32)`

SetValidNotAfter sets ValidNotAfter field to given value.

### HasValidNotAfter

`func (o *MemberSyslogServersCertificate) HasValidNotAfter() bool`

HasValidNotAfter returns a boolean if a field has been set.

### GetValidNotBefore

`func (o *MemberSyslogServersCertificate) GetValidNotBefore() int32`

GetValidNotBefore returns the ValidNotBefore field if non-nil, zero value otherwise.

### GetValidNotBeforeOk

`func (o *MemberSyslogServersCertificate) GetValidNotBeforeOk() (*int32, bool)`

GetValidNotBeforeOk returns a tuple with the ValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotBefore

`func (o *MemberSyslogServersCertificate) SetValidNotBefore(v int32)`

SetValidNotBefore sets ValidNotBefore field to given value.

### HasValidNotBefore

`func (o *MemberSyslogServersCertificate) HasValidNotBefore() bool`

HasValidNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


