# MemberFiledistributionHttpAcls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 

## Methods

### NewMemberFiledistributionHttpAcls

`func NewMemberFiledistributionHttpAcls() *MemberFiledistributionHttpAcls`

NewMemberFiledistributionHttpAcls instantiates a new MemberFiledistributionHttpAcls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberFiledistributionHttpAclsWithDefaults

`func NewMemberFiledistributionHttpAclsWithDefaults() *MemberFiledistributionHttpAcls`

NewMemberFiledistributionHttpAclsWithDefaults instantiates a new MemberFiledistributionHttpAcls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberFiledistributionHttpAcls) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberFiledistributionHttpAcls) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberFiledistributionHttpAcls) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberFiledistributionHttpAcls) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *MemberFiledistributionHttpAcls) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *MemberFiledistributionHttpAcls) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *MemberFiledistributionHttpAcls) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *MemberFiledistributionHttpAcls) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


