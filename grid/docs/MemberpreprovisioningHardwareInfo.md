# MemberpreprovisioningHardwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hwtype** | Pointer to **string** | Hardware type. | [optional] 
**Hwmodel** | Pointer to **string** | Hardware model - for IB-4010 are Rev1, Rev2; for IB-4030 are Rev1, Rev2; for PT-4000 is Rev2; for IB-VNIOS are IB-VM-100, IB-VM-810, IB-VM-820, IB-VM-RSP, IB-VM-1410, IB-VM-1420, IB-VM-2210, IB-VM-2220, IB-VM-4010, CP-V800, CP-V1400, CP-V2200. Note that you cannot specify hwmodel for following hardware types: IB-FLEX, IB-V2215, IB-V1425, IB-V4025, IB-V4015, IB-V1415, IB-V815, IB-V825, IB-V2225, CP-V805, CP-V1405, CP-V2205, IB-2215, IB-2225. | [optional] 

## Methods

### NewMemberpreprovisioningHardwareInfo

`func NewMemberpreprovisioningHardwareInfo() *MemberpreprovisioningHardwareInfo`

NewMemberpreprovisioningHardwareInfo instantiates a new MemberpreprovisioningHardwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberpreprovisioningHardwareInfoWithDefaults

`func NewMemberpreprovisioningHardwareInfoWithDefaults() *MemberpreprovisioningHardwareInfo`

NewMemberpreprovisioningHardwareInfoWithDefaults instantiates a new MemberpreprovisioningHardwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHwtype

`func (o *MemberpreprovisioningHardwareInfo) GetHwtype() string`

GetHwtype returns the Hwtype field if non-nil, zero value otherwise.

### GetHwtypeOk

`func (o *MemberpreprovisioningHardwareInfo) GetHwtypeOk() (*string, bool)`

GetHwtypeOk returns a tuple with the Hwtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwtype

`func (o *MemberpreprovisioningHardwareInfo) SetHwtype(v string)`

SetHwtype sets Hwtype field to given value.

### HasHwtype

`func (o *MemberpreprovisioningHardwareInfo) HasHwtype() bool`

HasHwtype returns a boolean if a field has been set.

### GetHwmodel

`func (o *MemberpreprovisioningHardwareInfo) GetHwmodel() string`

GetHwmodel returns the Hwmodel field if non-nil, zero value otherwise.

### GetHwmodelOk

`func (o *MemberpreprovisioningHardwareInfo) GetHwmodelOk() (*string, bool)`

GetHwmodelOk returns a tuple with the Hwmodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwmodel

`func (o *MemberpreprovisioningHardwareInfo) SetHwmodel(v string)`

SetHwmodel sets Hwmodel field to given value.

### HasHwmodel

`func (o *MemberpreprovisioningHardwareInfo) HasHwmodel() bool`

HasHwmodel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


