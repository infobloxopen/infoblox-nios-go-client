# MemberntpsettingNtpAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclType** | Pointer to **string** | The NTP access control list type. | [optional] 
**AcList** | Pointer to [**[]MemberntpsettingntpaclAcList**](MemberntpsettingntpaclAcList.md) | The list of NTP access control items. | [optional] 
**NamedAcl** | Pointer to **string** | The NTP access named ACL. | [optional] 
**Service** | Pointer to **string** | The type of service with access control for the assigned named ACL. | [optional] 

## Methods

### NewMemberntpsettingNtpAcl

`func NewMemberntpsettingNtpAcl() *MemberntpsettingNtpAcl`

NewMemberntpsettingNtpAcl instantiates a new MemberntpsettingNtpAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberntpsettingNtpAclWithDefaults

`func NewMemberntpsettingNtpAclWithDefaults() *MemberntpsettingNtpAcl`

NewMemberntpsettingNtpAclWithDefaults instantiates a new MemberntpsettingNtpAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclType

`func (o *MemberntpsettingNtpAcl) GetAclType() string`

GetAclType returns the AclType field if non-nil, zero value otherwise.

### GetAclTypeOk

`func (o *MemberntpsettingNtpAcl) GetAclTypeOk() (*string, bool)`

GetAclTypeOk returns a tuple with the AclType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclType

`func (o *MemberntpsettingNtpAcl) SetAclType(v string)`

SetAclType sets AclType field to given value.

### HasAclType

`func (o *MemberntpsettingNtpAcl) HasAclType() bool`

HasAclType returns a boolean if a field has been set.

### GetAcList

`func (o *MemberntpsettingNtpAcl) GetAcList() []MemberntpsettingntpaclAcList`

GetAcList returns the AcList field if non-nil, zero value otherwise.

### GetAcListOk

`func (o *MemberntpsettingNtpAcl) GetAcListOk() (*[]MemberntpsettingntpaclAcList, bool)`

GetAcListOk returns a tuple with the AcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcList

`func (o *MemberntpsettingNtpAcl) SetAcList(v []MemberntpsettingntpaclAcList)`

SetAcList sets AcList field to given value.

### HasAcList

`func (o *MemberntpsettingNtpAcl) HasAcList() bool`

HasAcList returns a boolean if a field has been set.

### GetNamedAcl

`func (o *MemberntpsettingNtpAcl) GetNamedAcl() string`

GetNamedAcl returns the NamedAcl field if non-nil, zero value otherwise.

### GetNamedAclOk

`func (o *MemberntpsettingNtpAcl) GetNamedAclOk() (*string, bool)`

GetNamedAclOk returns a tuple with the NamedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedAcl

`func (o *MemberntpsettingNtpAcl) SetNamedAcl(v string)`

SetNamedAcl sets NamedAcl field to given value.

### HasNamedAcl

`func (o *MemberntpsettingNtpAcl) HasNamedAcl() bool`

HasNamedAcl returns a boolean if a field has been set.

### GetService

`func (o *MemberntpsettingNtpAcl) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MemberntpsettingNtpAcl) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MemberntpsettingNtpAcl) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MemberntpsettingNtpAcl) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


