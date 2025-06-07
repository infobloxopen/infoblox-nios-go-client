# GetDiscoveryDevicecomponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ComponentName** | Pointer to **string** | The component name. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the device component. | [optional] [readonly] 
**Device** | Pointer to **string** | A reference to a device, to which this component belongs to. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the device component. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the device component. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of device component. | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryDevicecomponent**](DiscoveryDevicecomponent.md) |  | [optional] 

## Methods

### NewGetDiscoveryDevicecomponentResponse

`func NewGetDiscoveryDevicecomponentResponse() *GetDiscoveryDevicecomponentResponse`

NewGetDiscoveryDevicecomponentResponse instantiates a new GetDiscoveryDevicecomponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryDevicecomponentResponseWithDefaults

`func NewGetDiscoveryDevicecomponentResponseWithDefaults() *GetDiscoveryDevicecomponentResponse`

NewGetDiscoveryDevicecomponentResponseWithDefaults instantiates a new GetDiscoveryDevicecomponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryDevicecomponentResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryDevicecomponentResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryDevicecomponentResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryDevicecomponentResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComponentName

`func (o *GetDiscoveryDevicecomponentResponse) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *GetDiscoveryDevicecomponentResponse) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *GetDiscoveryDevicecomponentResponse) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *GetDiscoveryDevicecomponentResponse) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetDescription

`func (o *GetDiscoveryDevicecomponentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetDiscoveryDevicecomponentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetDiscoveryDevicecomponentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetDiscoveryDevicecomponentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *GetDiscoveryDevicecomponentResponse) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetDiscoveryDevicecomponentResponse) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetDiscoveryDevicecomponentResponse) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetDiscoveryDevicecomponentResponse) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModel

`func (o *GetDiscoveryDevicecomponentResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetDiscoveryDevicecomponentResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetDiscoveryDevicecomponentResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetDiscoveryDevicecomponentResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *GetDiscoveryDevicecomponentResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetDiscoveryDevicecomponentResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetDiscoveryDevicecomponentResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetDiscoveryDevicecomponentResponse) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetType

`func (o *GetDiscoveryDevicecomponentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDiscoveryDevicecomponentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDiscoveryDevicecomponentResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDiscoveryDevicecomponentResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryDevicecomponentResponse) GetResult() DiscoveryDevicecomponent`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryDevicecomponentResponse) GetResultOk() (*DiscoveryDevicecomponent, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryDevicecomponentResponse) SetResult(v DiscoveryDevicecomponent)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryDevicecomponentResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


