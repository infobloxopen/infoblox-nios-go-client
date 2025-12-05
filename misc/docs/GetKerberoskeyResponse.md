# GetKerberoskeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Domain** | Pointer to **string** | The Kerberos domain name. | [optional] [readonly] 
**Enctype** | Pointer to **string** | The Kerberos key encryption type. | [optional] [readonly] 
**InUse** | Pointer to **bool** | Determines whether the Kerberos key is assigned to the Grid or Grid member. | [optional] [readonly] 
**Members** | Pointer to **[]string** | The list of hostnames and services of Grid members where the key is assigned or Grid/DHCP4 or Grid/DHCP6 or Grid/DNS. | [optional] [readonly] 
**Principal** | Pointer to **string** | The principal of the Kerberos key object. | [optional] [readonly] 
**UploadTimestamp** | Pointer to **int64** | The timestamp of the Kerberos key upload operation. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Version** | Pointer to **int64** | The Kerberos key version number (KVNO). | [optional] [readonly] 
**Result** | Pointer to [**Kerberoskey**](Kerberoskey.md) |  | [optional] 

## Methods

### NewGetKerberoskeyResponse

`func NewGetKerberoskeyResponse() *GetKerberoskeyResponse`

NewGetKerberoskeyResponse instantiates a new GetKerberoskeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKerberoskeyResponseWithDefaults

`func NewGetKerberoskeyResponseWithDefaults() *GetKerberoskeyResponse`

NewGetKerberoskeyResponseWithDefaults instantiates a new GetKerberoskeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetKerberoskeyResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetKerberoskeyResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetKerberoskeyResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetKerberoskeyResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDomain

`func (o *GetKerberoskeyResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetKerberoskeyResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetKerberoskeyResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetKerberoskeyResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEnctype

`func (o *GetKerberoskeyResponse) GetEnctype() string`

GetEnctype returns the Enctype field if non-nil, zero value otherwise.

### GetEnctypeOk

`func (o *GetKerberoskeyResponse) GetEnctypeOk() (*string, bool)`

GetEnctypeOk returns a tuple with the Enctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnctype

`func (o *GetKerberoskeyResponse) SetEnctype(v string)`

SetEnctype sets Enctype field to given value.

### HasEnctype

`func (o *GetKerberoskeyResponse) HasEnctype() bool`

HasEnctype returns a boolean if a field has been set.

### GetInUse

`func (o *GetKerberoskeyResponse) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *GetKerberoskeyResponse) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *GetKerberoskeyResponse) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *GetKerberoskeyResponse) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetMembers

`func (o *GetKerberoskeyResponse) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetKerberoskeyResponse) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetKerberoskeyResponse) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetKerberoskeyResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPrincipal

`func (o *GetKerberoskeyResponse) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *GetKerberoskeyResponse) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *GetKerberoskeyResponse) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *GetKerberoskeyResponse) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetUploadTimestamp

`func (o *GetKerberoskeyResponse) GetUploadTimestamp() int64`

GetUploadTimestamp returns the UploadTimestamp field if non-nil, zero value otherwise.

### GetUploadTimestampOk

`func (o *GetKerberoskeyResponse) GetUploadTimestampOk() (*int64, bool)`

GetUploadTimestampOk returns a tuple with the UploadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTimestamp

`func (o *GetKerberoskeyResponse) SetUploadTimestamp(v int64)`

SetUploadTimestamp sets UploadTimestamp field to given value.

### HasUploadTimestamp

`func (o *GetKerberoskeyResponse) HasUploadTimestamp() bool`

HasUploadTimestamp returns a boolean if a field has been set.

### GetUuid

`func (o *GetKerberoskeyResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetKerberoskeyResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetKerberoskeyResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetKerberoskeyResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *GetKerberoskeyResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetKerberoskeyResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetKerberoskeyResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetKerberoskeyResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetKerberoskeyResponse) GetResult() Kerberoskey`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetKerberoskeyResponse) GetResultOk() (*Kerberoskey, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetKerberoskeyResponse) SetResult(v Kerberoskey)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetKerberoskeyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


