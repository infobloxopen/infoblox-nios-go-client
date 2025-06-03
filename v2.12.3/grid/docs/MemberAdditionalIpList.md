# MemberAdditionalIpList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anycast** | Pointer to **bool** | Determines if anycast for the Interface object is enabled or not. | [optional] 
**Ipv4NetworkSetting** | Pointer to [**MemberadditionaliplistIpv4NetworkSetting**](MemberadditionaliplistIpv4NetworkSetting.md) |  | [optional] 
**Ipv6NetworkSetting** | Pointer to [**MemberadditionaliplistIpv6NetworkSetting**](MemberadditionaliplistIpv6NetworkSetting.md) |  | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of this structure. | [optional] 
**EnableBgp** | Pointer to **bool** | Determines if the BGP advertisement setting is enabled for this interface or not. | [optional] 
**EnableOspf** | Pointer to **bool** | Determines if the OSPF advertisement setting is enabled for this interface or not. | [optional] 
**Interface** | Pointer to **string** | The interface type for the Interface object. | [optional] 

## Methods

### NewMemberAdditionalIpList

`func NewMemberAdditionalIpList() *MemberAdditionalIpList`

NewMemberAdditionalIpList instantiates a new MemberAdditionalIpList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberAdditionalIpListWithDefaults

`func NewMemberAdditionalIpListWithDefaults() *MemberAdditionalIpList`

NewMemberAdditionalIpListWithDefaults instantiates a new MemberAdditionalIpList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnycast

`func (o *MemberAdditionalIpList) GetAnycast() bool`

GetAnycast returns the Anycast field if non-nil, zero value otherwise.

### GetAnycastOk

`func (o *MemberAdditionalIpList) GetAnycastOk() (*bool, bool)`

GetAnycastOk returns a tuple with the Anycast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycast

`func (o *MemberAdditionalIpList) SetAnycast(v bool)`

SetAnycast sets Anycast field to given value.

### HasAnycast

`func (o *MemberAdditionalIpList) HasAnycast() bool`

HasAnycast returns a boolean if a field has been set.

### GetIpv4NetworkSetting

`func (o *MemberAdditionalIpList) GetIpv4NetworkSetting() MemberadditionaliplistIpv4NetworkSetting`

GetIpv4NetworkSetting returns the Ipv4NetworkSetting field if non-nil, zero value otherwise.

### GetIpv4NetworkSettingOk

`func (o *MemberAdditionalIpList) GetIpv4NetworkSettingOk() (*MemberadditionaliplistIpv4NetworkSetting, bool)`

GetIpv4NetworkSettingOk returns a tuple with the Ipv4NetworkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4NetworkSetting

`func (o *MemberAdditionalIpList) SetIpv4NetworkSetting(v MemberadditionaliplistIpv4NetworkSetting)`

SetIpv4NetworkSetting sets Ipv4NetworkSetting field to given value.

### HasIpv4NetworkSetting

`func (o *MemberAdditionalIpList) HasIpv4NetworkSetting() bool`

HasIpv4NetworkSetting returns a boolean if a field has been set.

### GetIpv6NetworkSetting

`func (o *MemberAdditionalIpList) GetIpv6NetworkSetting() MemberadditionaliplistIpv6NetworkSetting`

GetIpv6NetworkSetting returns the Ipv6NetworkSetting field if non-nil, zero value otherwise.

### GetIpv6NetworkSettingOk

`func (o *MemberAdditionalIpList) GetIpv6NetworkSettingOk() (*MemberadditionaliplistIpv6NetworkSetting, bool)`

GetIpv6NetworkSettingOk returns a tuple with the Ipv6NetworkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6NetworkSetting

`func (o *MemberAdditionalIpList) SetIpv6NetworkSetting(v MemberadditionaliplistIpv6NetworkSetting)`

SetIpv6NetworkSetting sets Ipv6NetworkSetting field to given value.

### HasIpv6NetworkSetting

`func (o *MemberAdditionalIpList) HasIpv6NetworkSetting() bool`

HasIpv6NetworkSetting returns a boolean if a field has been set.

### GetComment

`func (o *MemberAdditionalIpList) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MemberAdditionalIpList) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MemberAdditionalIpList) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MemberAdditionalIpList) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnableBgp

`func (o *MemberAdditionalIpList) GetEnableBgp() bool`

GetEnableBgp returns the EnableBgp field if non-nil, zero value otherwise.

### GetEnableBgpOk

`func (o *MemberAdditionalIpList) GetEnableBgpOk() (*bool, bool)`

GetEnableBgpOk returns a tuple with the EnableBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBgp

`func (o *MemberAdditionalIpList) SetEnableBgp(v bool)`

SetEnableBgp sets EnableBgp field to given value.

### HasEnableBgp

`func (o *MemberAdditionalIpList) HasEnableBgp() bool`

HasEnableBgp returns a boolean if a field has been set.

### GetEnableOspf

`func (o *MemberAdditionalIpList) GetEnableOspf() bool`

GetEnableOspf returns the EnableOspf field if non-nil, zero value otherwise.

### GetEnableOspfOk

`func (o *MemberAdditionalIpList) GetEnableOspfOk() (*bool, bool)`

GetEnableOspfOk returns a tuple with the EnableOspf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOspf

`func (o *MemberAdditionalIpList) SetEnableOspf(v bool)`

SetEnableOspf sets EnableOspf field to given value.

### HasEnableOspf

`func (o *MemberAdditionalIpList) HasEnableOspf() bool`

HasEnableOspf returns a boolean if a field has been set.

### GetInterface

`func (o *MemberAdditionalIpList) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *MemberAdditionalIpList) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *MemberAdditionalIpList) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *MemberAdditionalIpList) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


