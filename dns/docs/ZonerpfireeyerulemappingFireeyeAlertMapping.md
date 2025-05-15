# ZonerpfireeyerulemappingFireeyeAlertMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | Pointer to **string** | The type of Fireeye Alert. | [optional] 
**RpzRule** | Pointer to **string** | The RPZ rule for the alert. | [optional] 
**Lifetime** | Pointer to **int64** | The expiration Lifetime of alert type. The 32-bit unsigned integer represents the amount of seconds this alert type will live for. 0 means the alert will never expire. | [optional] 

## Methods

### NewZonerpfireeyerulemappingFireeyeAlertMapping

`func NewZonerpfireeyerulemappingFireeyeAlertMapping() *ZonerpfireeyerulemappingFireeyeAlertMapping`

NewZonerpfireeyerulemappingFireeyeAlertMapping instantiates a new ZonerpfireeyerulemappingFireeyeAlertMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonerpfireeyerulemappingFireeyeAlertMappingWithDefaults

`func NewZonerpfireeyerulemappingFireeyeAlertMappingWithDefaults() *ZonerpfireeyerulemappingFireeyeAlertMapping`

NewZonerpfireeyerulemappingFireeyeAlertMappingWithDefaults instantiates a new ZonerpfireeyerulemappingFireeyeAlertMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetRpzRule

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) GetRpzRule() string`

GetRpzRule returns the RpzRule field if non-nil, zero value otherwise.

### GetRpzRuleOk

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) GetRpzRuleOk() (*string, bool)`

GetRpzRuleOk returns a tuple with the RpzRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzRule

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) SetRpzRule(v string)`

SetRpzRule sets RpzRule field to given value.

### HasRpzRule

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) HasRpzRule() bool`

HasRpzRule returns a boolean if a field has been set.

### GetLifetime

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) GetLifetime() int64`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) GetLifetimeOk() (*int64, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) SetLifetime(v int64)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *ZonerpfireeyerulemappingFireeyeAlertMapping) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


