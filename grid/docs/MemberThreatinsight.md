# MemberThreatinsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The Grid member descriptive comment. | [optional] [readonly] 
**EnableService** | Pointer to **bool** | Determines whether the threat insight service is enabled. | [optional] 
**HostName** | Pointer to **string** | The Grid member host name. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | The IPv4 Address address of the Grid member. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | The IPv6 Address address of the Grid member. | [optional] [readonly] 
**Status** | Pointer to **string** | The Grid member threat insight status. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewMemberThreatinsight

`func NewMemberThreatinsight() *MemberThreatinsight`

NewMemberThreatinsight instantiates a new MemberThreatinsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberThreatinsightWithDefaults

`func NewMemberThreatinsightWithDefaults() *MemberThreatinsight`

NewMemberThreatinsightWithDefaults instantiates a new MemberThreatinsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberThreatinsight) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberThreatinsight) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberThreatinsight) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberThreatinsight) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *MemberThreatinsight) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MemberThreatinsight) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MemberThreatinsight) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MemberThreatinsight) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnableService

`func (o *MemberThreatinsight) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *MemberThreatinsight) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *MemberThreatinsight) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *MemberThreatinsight) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetHostName

`func (o *MemberThreatinsight) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *MemberThreatinsight) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *MemberThreatinsight) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *MemberThreatinsight) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpv4Address

`func (o *MemberThreatinsight) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *MemberThreatinsight) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *MemberThreatinsight) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *MemberThreatinsight) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *MemberThreatinsight) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *MemberThreatinsight) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *MemberThreatinsight) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *MemberThreatinsight) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetStatus

`func (o *MemberThreatinsight) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemberThreatinsight) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemberThreatinsight) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MemberThreatinsight) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *MemberThreatinsight) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MemberThreatinsight) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MemberThreatinsight) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MemberThreatinsight) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


