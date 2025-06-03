# MemberPreProvisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareInfo** | Pointer to [**MemberpreprovisioningHardwareInfo**](MemberpreprovisioningHardwareInfo.md) |  | [optional] 
**Licenses** | Pointer to **[]string** | An array of license types the pre-provisioned member should have in order to join the Grid, or the licenses that must be allocated to the member when it joins the Grid using the token-based authentication. | [optional] 

## Methods

### NewMemberPreProvisioning

`func NewMemberPreProvisioning() *MemberPreProvisioning`

NewMemberPreProvisioning instantiates a new MemberPreProvisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberPreProvisioningWithDefaults

`func NewMemberPreProvisioningWithDefaults() *MemberPreProvisioning`

NewMemberPreProvisioningWithDefaults instantiates a new MemberPreProvisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareInfo

`func (o *MemberPreProvisioning) GetHardwareInfo() MemberpreprovisioningHardwareInfo`

GetHardwareInfo returns the HardwareInfo field if non-nil, zero value otherwise.

### GetHardwareInfoOk

`func (o *MemberPreProvisioning) GetHardwareInfoOk() (*MemberpreprovisioningHardwareInfo, bool)`

GetHardwareInfoOk returns a tuple with the HardwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInfo

`func (o *MemberPreProvisioning) SetHardwareInfo(v MemberpreprovisioningHardwareInfo)`

SetHardwareInfo sets HardwareInfo field to given value.

### HasHardwareInfo

`func (o *MemberPreProvisioning) HasHardwareInfo() bool`

HasHardwareInfo returns a boolean if a field has been set.

### GetLicenses

`func (o *MemberPreProvisioning) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *MemberPreProvisioning) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *MemberPreProvisioning) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *MemberPreProvisioning) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


