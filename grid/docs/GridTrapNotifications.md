# GridTrapNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrapType** | Pointer to **string** | Determines the type of a given trap. | [optional] 
**EnableEmail** | Pointer to **bool** | Determines if the email notifications for the given trap are enabled or not. | [optional] 
**EnableTrap** | Pointer to **bool** | Determines if the trap is enabled or not. | [optional] 

## Methods

### NewGridTrapNotifications

`func NewGridTrapNotifications() *GridTrapNotifications`

NewGridTrapNotifications instantiates a new GridTrapNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridTrapNotificationsWithDefaults

`func NewGridTrapNotificationsWithDefaults() *GridTrapNotifications`

NewGridTrapNotificationsWithDefaults instantiates a new GridTrapNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrapType

`func (o *GridTrapNotifications) GetTrapType() string`

GetTrapType returns the TrapType field if non-nil, zero value otherwise.

### GetTrapTypeOk

`func (o *GridTrapNotifications) GetTrapTypeOk() (*string, bool)`

GetTrapTypeOk returns a tuple with the TrapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapType

`func (o *GridTrapNotifications) SetTrapType(v string)`

SetTrapType sets TrapType field to given value.

### HasTrapType

`func (o *GridTrapNotifications) HasTrapType() bool`

HasTrapType returns a boolean if a field has been set.

### GetEnableEmail

`func (o *GridTrapNotifications) GetEnableEmail() bool`

GetEnableEmail returns the EnableEmail field if non-nil, zero value otherwise.

### GetEnableEmailOk

`func (o *GridTrapNotifications) GetEnableEmailOk() (*bool, bool)`

GetEnableEmailOk returns a tuple with the EnableEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmail

`func (o *GridTrapNotifications) SetEnableEmail(v bool)`

SetEnableEmail sets EnableEmail field to given value.

### HasEnableEmail

`func (o *GridTrapNotifications) HasEnableEmail() bool`

HasEnableEmail returns a boolean if a field has been set.

### GetEnableTrap

`func (o *GridTrapNotifications) GetEnableTrap() bool`

GetEnableTrap returns the EnableTrap field if non-nil, zero value otherwise.

### GetEnableTrapOk

`func (o *GridTrapNotifications) GetEnableTrapOk() (*bool, bool)`

GetEnableTrapOk returns a tuple with the EnableTrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrap

`func (o *GridTrapNotifications) SetEnableTrap(v bool)`

SetEnableTrap sets EnableTrap field to given value.

### HasEnableTrap

`func (o *GridTrapNotifications) HasEnableTrap() bool`

HasEnableTrap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


