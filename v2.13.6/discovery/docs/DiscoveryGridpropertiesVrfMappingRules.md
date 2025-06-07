# DiscoveryGridpropertiesVrfMappingRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkView** | Pointer to **string** | The name of the network view associated with the VRF mapping rule. | [optional] 
**Criteria** | Pointer to **string** | Extended POSIX regular expression matching the VRF name. The regular expression string should contain the leading caret anchor ^ and the trailing dollar anchor $ symbols. | [optional] 
**Comment** | Pointer to **string** | The comment. | [optional] 

## Methods

### NewDiscoveryGridpropertiesVrfMappingRules

`func NewDiscoveryGridpropertiesVrfMappingRules() *DiscoveryGridpropertiesVrfMappingRules`

NewDiscoveryGridpropertiesVrfMappingRules instantiates a new DiscoveryGridpropertiesVrfMappingRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesVrfMappingRulesWithDefaults

`func NewDiscoveryGridpropertiesVrfMappingRulesWithDefaults() *DiscoveryGridpropertiesVrfMappingRules`

NewDiscoveryGridpropertiesVrfMappingRulesWithDefaults instantiates a new DiscoveryGridpropertiesVrfMappingRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkView

`func (o *DiscoveryGridpropertiesVrfMappingRules) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *DiscoveryGridpropertiesVrfMappingRules) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *DiscoveryGridpropertiesVrfMappingRules) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *DiscoveryGridpropertiesVrfMappingRules) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetCriteria

`func (o *DiscoveryGridpropertiesVrfMappingRules) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *DiscoveryGridpropertiesVrfMappingRules) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *DiscoveryGridpropertiesVrfMappingRules) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *DiscoveryGridpropertiesVrfMappingRules) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetComment

`func (o *DiscoveryGridpropertiesVrfMappingRules) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DiscoveryGridpropertiesVrfMappingRules) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DiscoveryGridpropertiesVrfMappingRules) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DiscoveryGridpropertiesVrfMappingRules) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


