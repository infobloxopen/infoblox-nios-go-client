# DiscoveryDeviceinterfaceAdminStatusTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The configured admin status value. | [optional] [readonly] 
**Details** | Pointer to [**DiscoverydeviceinterfaceadminstatustaskinfoDetails**](DiscoverydeviceinterfaceadminstatustaskinfoDetails.md) |  | [optional] 

## Methods

### NewDiscoveryDeviceinterfaceAdminStatusTaskInfo

`func NewDiscoveryDeviceinterfaceAdminStatusTaskInfo() *DiscoveryDeviceinterfaceAdminStatusTaskInfo`

NewDiscoveryDeviceinterfaceAdminStatusTaskInfo instantiates a new DiscoveryDeviceinterfaceAdminStatusTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDeviceinterfaceAdminStatusTaskInfoWithDefaults

`func NewDiscoveryDeviceinterfaceAdminStatusTaskInfoWithDefaults() *DiscoveryDeviceinterfaceAdminStatusTaskInfo`

NewDiscoveryDeviceinterfaceAdminStatusTaskInfoWithDefaults instantiates a new DiscoveryDeviceinterfaceAdminStatusTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) GetDetails() DiscoverydeviceinterfaceadminstatustaskinfoDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) GetDetailsOk() (*DiscoverydeviceinterfaceadminstatustaskinfoDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) SetDetails(v DiscoverydeviceinterfaceadminstatustaskinfoDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DiscoveryDeviceinterfaceAdminStatusTaskInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


