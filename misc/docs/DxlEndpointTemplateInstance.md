# DxlEndpointTemplateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** | The name of the REST API template parameter. | [optional] 
**Parameters** | Pointer to [**[]DxlendpointtemplateinstanceParameters**](DxlendpointtemplateinstanceParameters.md) | The notification REST template parameters. | [optional] 

## Methods

### NewDxlEndpointTemplateInstance

`func NewDxlEndpointTemplateInstance() *DxlEndpointTemplateInstance`

NewDxlEndpointTemplateInstance instantiates a new DxlEndpointTemplateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDxlEndpointTemplateInstanceWithDefaults

`func NewDxlEndpointTemplateInstanceWithDefaults() *DxlEndpointTemplateInstance`

NewDxlEndpointTemplateInstanceWithDefaults instantiates a new DxlEndpointTemplateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *DxlEndpointTemplateInstance) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *DxlEndpointTemplateInstance) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *DxlEndpointTemplateInstance) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *DxlEndpointTemplateInstance) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetParameters

`func (o *DxlEndpointTemplateInstance) GetParameters() []DxlendpointtemplateinstanceParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DxlEndpointTemplateInstance) GetParametersOk() (*[]DxlendpointtemplateinstanceParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DxlEndpointTemplateInstance) SetParameters(v []DxlendpointtemplateinstanceParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DxlEndpointTemplateInstance) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


