# NotificationRuleTemplateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** | The name of the REST API template parameter. | [optional] 
**Parameters** | Pointer to [**[]NotificationruletemplateinstanceParameters**](NotificationruletemplateinstanceParameters.md) | The notification REST template parameters. | [optional] 

## Methods

### NewNotificationRuleTemplateInstance

`func NewNotificationRuleTemplateInstance() *NotificationRuleTemplateInstance`

NewNotificationRuleTemplateInstance instantiates a new NotificationRuleTemplateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleTemplateInstanceWithDefaults

`func NewNotificationRuleTemplateInstanceWithDefaults() *NotificationRuleTemplateInstance`

NewNotificationRuleTemplateInstanceWithDefaults instantiates a new NotificationRuleTemplateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *NotificationRuleTemplateInstance) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *NotificationRuleTemplateInstance) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *NotificationRuleTemplateInstance) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *NotificationRuleTemplateInstance) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetParameters

`func (o *NotificationRuleTemplateInstance) GetParameters() []NotificationruletemplateinstanceParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *NotificationRuleTemplateInstance) GetParametersOk() (*[]NotificationruletemplateinstanceParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *NotificationRuleTemplateInstance) SetParameters(v []NotificationruletemplateinstanceParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *NotificationRuleTemplateInstance) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


