# GetDiscoveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ClearNetworkPortAssignment** | Pointer to **map[string]interface{}** |  | [optional] 
**ControlSwitchPort** | Pointer to **map[string]interface{}** |  | [optional] 
**DiscoveryDataConversion** | Pointer to **map[string]interface{}** |  | [optional] 
**GetDeviceSupportInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**GetJobDevices** | Pointer to **map[string]interface{}** |  | [optional] 
**GetJobProcessDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**ImportDeviceSupportBundle** | Pointer to **map[string]interface{}** |  | [optional] 
**ModifySdnAssignment** | Pointer to **map[string]interface{}** |  | [optional] 
**ModifyVrfAssignment** | Pointer to **map[string]interface{}** |  | [optional] 
**ProvisionNetworkDhcpRelay** | Pointer to **map[string]interface{}** |  | [optional] 
**ProvisionNetworkPort** | Pointer to **map[string]interface{}** |  | [optional] 
**Result** | Pointer to [**Discovery**](Discovery.md) |  | [optional] 

## Methods

### NewGetDiscoveryResponse

`func NewGetDiscoveryResponse() *GetDiscoveryResponse`

NewGetDiscoveryResponse instantiates a new GetDiscoveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryResponseWithDefaults

`func NewGetDiscoveryResponseWithDefaults() *GetDiscoveryResponse`

NewGetDiscoveryResponseWithDefaults instantiates a new GetDiscoveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClearNetworkPortAssignment

`func (o *GetDiscoveryResponse) GetClearNetworkPortAssignment() map[string]interface{}`

GetClearNetworkPortAssignment returns the ClearNetworkPortAssignment field if non-nil, zero value otherwise.

### GetClearNetworkPortAssignmentOk

`func (o *GetDiscoveryResponse) GetClearNetworkPortAssignmentOk() (*map[string]interface{}, bool)`

GetClearNetworkPortAssignmentOk returns a tuple with the ClearNetworkPortAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearNetworkPortAssignment

`func (o *GetDiscoveryResponse) SetClearNetworkPortAssignment(v map[string]interface{})`

SetClearNetworkPortAssignment sets ClearNetworkPortAssignment field to given value.

### HasClearNetworkPortAssignment

`func (o *GetDiscoveryResponse) HasClearNetworkPortAssignment() bool`

HasClearNetworkPortAssignment returns a boolean if a field has been set.

### GetControlSwitchPort

`func (o *GetDiscoveryResponse) GetControlSwitchPort() map[string]interface{}`

GetControlSwitchPort returns the ControlSwitchPort field if non-nil, zero value otherwise.

### GetControlSwitchPortOk

`func (o *GetDiscoveryResponse) GetControlSwitchPortOk() (*map[string]interface{}, bool)`

GetControlSwitchPortOk returns a tuple with the ControlSwitchPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSwitchPort

`func (o *GetDiscoveryResponse) SetControlSwitchPort(v map[string]interface{})`

SetControlSwitchPort sets ControlSwitchPort field to given value.

### HasControlSwitchPort

`func (o *GetDiscoveryResponse) HasControlSwitchPort() bool`

HasControlSwitchPort returns a boolean if a field has been set.

### GetDiscoveryDataConversion

`func (o *GetDiscoveryResponse) GetDiscoveryDataConversion() map[string]interface{}`

GetDiscoveryDataConversion returns the DiscoveryDataConversion field if non-nil, zero value otherwise.

### GetDiscoveryDataConversionOk

`func (o *GetDiscoveryResponse) GetDiscoveryDataConversionOk() (*map[string]interface{}, bool)`

GetDiscoveryDataConversionOk returns a tuple with the DiscoveryDataConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryDataConversion

`func (o *GetDiscoveryResponse) SetDiscoveryDataConversion(v map[string]interface{})`

SetDiscoveryDataConversion sets DiscoveryDataConversion field to given value.

### HasDiscoveryDataConversion

`func (o *GetDiscoveryResponse) HasDiscoveryDataConversion() bool`

HasDiscoveryDataConversion returns a boolean if a field has been set.

### GetGetDeviceSupportInfo

`func (o *GetDiscoveryResponse) GetGetDeviceSupportInfo() map[string]interface{}`

GetGetDeviceSupportInfo returns the GetDeviceSupportInfo field if non-nil, zero value otherwise.

### GetGetDeviceSupportInfoOk

`func (o *GetDiscoveryResponse) GetGetDeviceSupportInfoOk() (*map[string]interface{}, bool)`

GetGetDeviceSupportInfoOk returns a tuple with the GetDeviceSupportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetDeviceSupportInfo

`func (o *GetDiscoveryResponse) SetGetDeviceSupportInfo(v map[string]interface{})`

SetGetDeviceSupportInfo sets GetDeviceSupportInfo field to given value.

### HasGetDeviceSupportInfo

`func (o *GetDiscoveryResponse) HasGetDeviceSupportInfo() bool`

HasGetDeviceSupportInfo returns a boolean if a field has been set.

### GetGetJobDevices

`func (o *GetDiscoveryResponse) GetGetJobDevices() map[string]interface{}`

GetGetJobDevices returns the GetJobDevices field if non-nil, zero value otherwise.

### GetGetJobDevicesOk

`func (o *GetDiscoveryResponse) GetGetJobDevicesOk() (*map[string]interface{}, bool)`

GetGetJobDevicesOk returns a tuple with the GetJobDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetJobDevices

`func (o *GetDiscoveryResponse) SetGetJobDevices(v map[string]interface{})`

SetGetJobDevices sets GetJobDevices field to given value.

### HasGetJobDevices

`func (o *GetDiscoveryResponse) HasGetJobDevices() bool`

HasGetJobDevices returns a boolean if a field has been set.

### GetGetJobProcessDetails

`func (o *GetDiscoveryResponse) GetGetJobProcessDetails() map[string]interface{}`

GetGetJobProcessDetails returns the GetJobProcessDetails field if non-nil, zero value otherwise.

### GetGetJobProcessDetailsOk

`func (o *GetDiscoveryResponse) GetGetJobProcessDetailsOk() (*map[string]interface{}, bool)`

GetGetJobProcessDetailsOk returns a tuple with the GetJobProcessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetJobProcessDetails

`func (o *GetDiscoveryResponse) SetGetJobProcessDetails(v map[string]interface{})`

SetGetJobProcessDetails sets GetJobProcessDetails field to given value.

### HasGetJobProcessDetails

`func (o *GetDiscoveryResponse) HasGetJobProcessDetails() bool`

HasGetJobProcessDetails returns a boolean if a field has been set.

### GetImportDeviceSupportBundle

`func (o *GetDiscoveryResponse) GetImportDeviceSupportBundle() map[string]interface{}`

GetImportDeviceSupportBundle returns the ImportDeviceSupportBundle field if non-nil, zero value otherwise.

### GetImportDeviceSupportBundleOk

`func (o *GetDiscoveryResponse) GetImportDeviceSupportBundleOk() (*map[string]interface{}, bool)`

GetImportDeviceSupportBundleOk returns a tuple with the ImportDeviceSupportBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDeviceSupportBundle

`func (o *GetDiscoveryResponse) SetImportDeviceSupportBundle(v map[string]interface{})`

SetImportDeviceSupportBundle sets ImportDeviceSupportBundle field to given value.

### HasImportDeviceSupportBundle

`func (o *GetDiscoveryResponse) HasImportDeviceSupportBundle() bool`

HasImportDeviceSupportBundle returns a boolean if a field has been set.

### GetModifySdnAssignment

`func (o *GetDiscoveryResponse) GetModifySdnAssignment() map[string]interface{}`

GetModifySdnAssignment returns the ModifySdnAssignment field if non-nil, zero value otherwise.

### GetModifySdnAssignmentOk

`func (o *GetDiscoveryResponse) GetModifySdnAssignmentOk() (*map[string]interface{}, bool)`

GetModifySdnAssignmentOk returns a tuple with the ModifySdnAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifySdnAssignment

`func (o *GetDiscoveryResponse) SetModifySdnAssignment(v map[string]interface{})`

SetModifySdnAssignment sets ModifySdnAssignment field to given value.

### HasModifySdnAssignment

`func (o *GetDiscoveryResponse) HasModifySdnAssignment() bool`

HasModifySdnAssignment returns a boolean if a field has been set.

### GetModifyVrfAssignment

`func (o *GetDiscoveryResponse) GetModifyVrfAssignment() map[string]interface{}`

GetModifyVrfAssignment returns the ModifyVrfAssignment field if non-nil, zero value otherwise.

### GetModifyVrfAssignmentOk

`func (o *GetDiscoveryResponse) GetModifyVrfAssignmentOk() (*map[string]interface{}, bool)`

GetModifyVrfAssignmentOk returns a tuple with the ModifyVrfAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyVrfAssignment

`func (o *GetDiscoveryResponse) SetModifyVrfAssignment(v map[string]interface{})`

SetModifyVrfAssignment sets ModifyVrfAssignment field to given value.

### HasModifyVrfAssignment

`func (o *GetDiscoveryResponse) HasModifyVrfAssignment() bool`

HasModifyVrfAssignment returns a boolean if a field has been set.

### GetProvisionNetworkDhcpRelay

`func (o *GetDiscoveryResponse) GetProvisionNetworkDhcpRelay() map[string]interface{}`

GetProvisionNetworkDhcpRelay returns the ProvisionNetworkDhcpRelay field if non-nil, zero value otherwise.

### GetProvisionNetworkDhcpRelayOk

`func (o *GetDiscoveryResponse) GetProvisionNetworkDhcpRelayOk() (*map[string]interface{}, bool)`

GetProvisionNetworkDhcpRelayOk returns a tuple with the ProvisionNetworkDhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionNetworkDhcpRelay

`func (o *GetDiscoveryResponse) SetProvisionNetworkDhcpRelay(v map[string]interface{})`

SetProvisionNetworkDhcpRelay sets ProvisionNetworkDhcpRelay field to given value.

### HasProvisionNetworkDhcpRelay

`func (o *GetDiscoveryResponse) HasProvisionNetworkDhcpRelay() bool`

HasProvisionNetworkDhcpRelay returns a boolean if a field has been set.

### GetProvisionNetworkPort

`func (o *GetDiscoveryResponse) GetProvisionNetworkPort() map[string]interface{}`

GetProvisionNetworkPort returns the ProvisionNetworkPort field if non-nil, zero value otherwise.

### GetProvisionNetworkPortOk

`func (o *GetDiscoveryResponse) GetProvisionNetworkPortOk() (*map[string]interface{}, bool)`

GetProvisionNetworkPortOk returns a tuple with the ProvisionNetworkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionNetworkPort

`func (o *GetDiscoveryResponse) SetProvisionNetworkPort(v map[string]interface{})`

SetProvisionNetworkPort sets ProvisionNetworkPort field to given value.

### HasProvisionNetworkPort

`func (o *GetDiscoveryResponse) HasProvisionNetworkPort() bool`

HasProvisionNetworkPort returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryResponse) GetResult() Discovery`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryResponse) GetResultOk() (*Discovery, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryResponse) SetResult(v Discovery)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


