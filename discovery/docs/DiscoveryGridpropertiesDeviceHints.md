# DiscoveryGridpropertiesDeviceHints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | The type of Device | [optional] 
**Criteria** | Pointer to **string** | Pattern or extended POSIX regular expression matching the Device Hint. The regular expression string should contain the leading caret anchor ^ and the trailing dollar anchor $ symbols. | [optional] 
**Comment** | Pointer to **string** | The comment. | [optional] 

## Methods

### NewDiscoveryGridpropertiesDeviceHints

`func NewDiscoveryGridpropertiesDeviceHints() *DiscoveryGridpropertiesDeviceHints`

NewDiscoveryGridpropertiesDeviceHints instantiates a new DiscoveryGridpropertiesDeviceHints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesDeviceHintsWithDefaults

`func NewDiscoveryGridpropertiesDeviceHintsWithDefaults() *DiscoveryGridpropertiesDeviceHints`

NewDiscoveryGridpropertiesDeviceHintsWithDefaults instantiates a new DiscoveryGridpropertiesDeviceHints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *DiscoveryGridpropertiesDeviceHints) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DiscoveryGridpropertiesDeviceHints) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DiscoveryGridpropertiesDeviceHints) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DiscoveryGridpropertiesDeviceHints) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCriteria

`func (o *DiscoveryGridpropertiesDeviceHints) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *DiscoveryGridpropertiesDeviceHints) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *DiscoveryGridpropertiesDeviceHints) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *DiscoveryGridpropertiesDeviceHints) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetComment

`func (o *DiscoveryGridpropertiesDeviceHints) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DiscoveryGridpropertiesDeviceHints) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DiscoveryGridpropertiesDeviceHints) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DiscoveryGridpropertiesDeviceHints) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


