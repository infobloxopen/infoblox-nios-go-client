# HsmEntrustnshieldgroupEntrustnshieldHsm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteIp** | Pointer to **string** | The IPv4 Address of the Entrust nShield HSM device. | [optional] 
**RemotePort** | Pointer to **int64** | The Entrust nShield HSM device destination port. | [optional] 
**Status** | Pointer to **string** | The Entrust nShield HSM device status. | [optional] [readonly] 
**RemoteEsn** | Pointer to **string** | The Entrust nShield HSM device electronic serial number. | [optional] [readonly] 
**Keyhash** | Pointer to **string** | The Entrust nShield HSM device public key digest. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the Entrust nShield HSM device is disabled. | [optional] 

## Methods

### NewHsmEntrustnshieldgroupEntrustnshieldHsm

`func NewHsmEntrustnshieldgroupEntrustnshieldHsm() *HsmEntrustnshieldgroupEntrustnshieldHsm`

NewHsmEntrustnshieldgroupEntrustnshieldHsm instantiates a new HsmEntrustnshieldgroupEntrustnshieldHsm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmEntrustnshieldgroupEntrustnshieldHsmWithDefaults

`func NewHsmEntrustnshieldgroupEntrustnshieldHsmWithDefaults() *HsmEntrustnshieldgroupEntrustnshieldHsm`

NewHsmEntrustnshieldgroupEntrustnshieldHsmWithDefaults instantiates a new HsmEntrustnshieldgroupEntrustnshieldHsm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteIp

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemotePort

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetStatus

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRemoteEsn

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetRemoteEsn() string`

GetRemoteEsn returns the RemoteEsn field if non-nil, zero value otherwise.

### GetRemoteEsnOk

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetRemoteEsnOk() (*string, bool)`

GetRemoteEsnOk returns a tuple with the RemoteEsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEsn

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) SetRemoteEsn(v string)`

SetRemoteEsn sets RemoteEsn field to given value.

### HasRemoteEsn

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) HasRemoteEsn() bool`

HasRemoteEsn returns a boolean if a field has been set.

### GetKeyhash

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetKeyhash() string`

GetKeyhash returns the Keyhash field if non-nil, zero value otherwise.

### GetKeyhashOk

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetKeyhashOk() (*string, bool)`

GetKeyhashOk returns a tuple with the Keyhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyhash

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) SetKeyhash(v string)`

SetKeyhash sets Keyhash field to given value.

### HasKeyhash

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) HasKeyhash() bool`

HasKeyhash returns a boolean if a field has been set.

### GetDisable

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *HsmEntrustnshieldgroupEntrustnshieldHsm) HasDisable() bool`

HasDisable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


