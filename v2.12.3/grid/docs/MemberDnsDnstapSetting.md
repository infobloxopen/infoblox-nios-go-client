# MemberDnsDnstapSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnstapReceiverPort** | Pointer to **int64** | DNSTAP receiver port number. | [optional] 
**DnstapIdentity** | Pointer to **string** | DNSTAP id string. | [optional] [readonly] 
**DnstapVersion** | Pointer to **string** | DNSTAP version. | [optional] [readonly] 
**DnstapReceiverAddress** | Pointer to **string** | Address or FQDN of DNSTAP receiver. | [optional] 

## Methods

### NewMemberDnsDnstapSetting

`func NewMemberDnsDnstapSetting() *MemberDnsDnstapSetting`

NewMemberDnsDnstapSetting instantiates a new MemberDnsDnstapSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsDnstapSettingWithDefaults

`func NewMemberDnsDnstapSettingWithDefaults() *MemberDnsDnstapSetting`

NewMemberDnsDnstapSettingWithDefaults instantiates a new MemberDnsDnstapSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnstapReceiverPort

`func (o *MemberDnsDnstapSetting) GetDnstapReceiverPort() int64`

GetDnstapReceiverPort returns the DnstapReceiverPort field if non-nil, zero value otherwise.

### GetDnstapReceiverPortOk

`func (o *MemberDnsDnstapSetting) GetDnstapReceiverPortOk() (*int64, bool)`

GetDnstapReceiverPortOk returns a tuple with the DnstapReceiverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapReceiverPort

`func (o *MemberDnsDnstapSetting) SetDnstapReceiverPort(v int64)`

SetDnstapReceiverPort sets DnstapReceiverPort field to given value.

### HasDnstapReceiverPort

`func (o *MemberDnsDnstapSetting) HasDnstapReceiverPort() bool`

HasDnstapReceiverPort returns a boolean if a field has been set.

### GetDnstapIdentity

`func (o *MemberDnsDnstapSetting) GetDnstapIdentity() string`

GetDnstapIdentity returns the DnstapIdentity field if non-nil, zero value otherwise.

### GetDnstapIdentityOk

`func (o *MemberDnsDnstapSetting) GetDnstapIdentityOk() (*string, bool)`

GetDnstapIdentityOk returns a tuple with the DnstapIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapIdentity

`func (o *MemberDnsDnstapSetting) SetDnstapIdentity(v string)`

SetDnstapIdentity sets DnstapIdentity field to given value.

### HasDnstapIdentity

`func (o *MemberDnsDnstapSetting) HasDnstapIdentity() bool`

HasDnstapIdentity returns a boolean if a field has been set.

### GetDnstapVersion

`func (o *MemberDnsDnstapSetting) GetDnstapVersion() string`

GetDnstapVersion returns the DnstapVersion field if non-nil, zero value otherwise.

### GetDnstapVersionOk

`func (o *MemberDnsDnstapSetting) GetDnstapVersionOk() (*string, bool)`

GetDnstapVersionOk returns a tuple with the DnstapVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapVersion

`func (o *MemberDnsDnstapSetting) SetDnstapVersion(v string)`

SetDnstapVersion sets DnstapVersion field to given value.

### HasDnstapVersion

`func (o *MemberDnsDnstapSetting) HasDnstapVersion() bool`

HasDnstapVersion returns a boolean if a field has been set.

### GetDnstapReceiverAddress

`func (o *MemberDnsDnstapSetting) GetDnstapReceiverAddress() string`

GetDnstapReceiverAddress returns the DnstapReceiverAddress field if non-nil, zero value otherwise.

### GetDnstapReceiverAddressOk

`func (o *MemberDnsDnstapSetting) GetDnstapReceiverAddressOk() (*string, bool)`

GetDnstapReceiverAddressOk returns a tuple with the DnstapReceiverAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnstapReceiverAddress

`func (o *MemberDnsDnstapSetting) SetDnstapReceiverAddress(v string)`

SetDnstapReceiverAddress sets DnstapReceiverAddress field to given value.

### HasDnstapReceiverAddress

`func (o *MemberDnsDnstapSetting) HasDnstapReceiverAddress() bool`

HasDnstapReceiverAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


