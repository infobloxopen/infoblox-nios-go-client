# MemberTrapNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrapType** | Pointer to **string** | Determines the type of a given trap. | [optional] 
**EnableEmail** | Pointer to **bool** | Determines if the email notifications for the given trap are enabled or not. | [optional] 
**EnableTrap** | Pointer to **bool** | Determines if the trap is enabled or not. | [optional] 

## Methods

### NewMemberTrapNotifications

`func NewMemberTrapNotifications() *MemberTrapNotifications`

NewMemberTrapNotifications instantiates a new MemberTrapNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberTrapNotificationsWithDefaults

`func NewMemberTrapNotificationsWithDefaults() *MemberTrapNotifications`

NewMemberTrapNotificationsWithDefaults instantiates a new MemberTrapNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrapType

`func (o *MemberTrapNotifications) GetTrapType() string`

GetTrapType returns the TrapType field if non-nil, zero value otherwise.

### GetTrapTypeOk

`func (o *MemberTrapNotifications) GetTrapTypeOk() (*string, bool)`

GetTrapTypeOk returns a tuple with the TrapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapType

`func (o *MemberTrapNotifications) SetTrapType(v string)`

SetTrapType sets TrapType field to given value.

### HasTrapType

`func (o *MemberTrapNotifications) HasTrapType() bool`

HasTrapType returns a boolean if a field has been set.

### GetEnableEmail

`func (o *MemberTrapNotifications) GetEnableEmail() bool`

GetEnableEmail returns the EnableEmail field if non-nil, zero value otherwise.

### GetEnableEmailOk

`func (o *MemberTrapNotifications) GetEnableEmailOk() (*bool, bool)`

GetEnableEmailOk returns a tuple with the EnableEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmail

`func (o *MemberTrapNotifications) SetEnableEmail(v bool)`

SetEnableEmail sets EnableEmail field to given value.

### HasEnableEmail

`func (o *MemberTrapNotifications) HasEnableEmail() bool`

HasEnableEmail returns a boolean if a field has been set.

### GetEnableTrap

`func (o *MemberTrapNotifications) GetEnableTrap() bool`

GetEnableTrap returns the EnableTrap field if non-nil, zero value otherwise.

### GetEnableTrapOk

`func (o *MemberTrapNotifications) GetEnableTrapOk() (*bool, bool)`

GetEnableTrapOk returns a tuple with the EnableTrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrap

`func (o *MemberTrapNotifications) SetEnableTrap(v bool)`

SetEnableTrap sets EnableTrap field to given value.

### HasEnableTrap

`func (o *MemberTrapNotifications) HasEnableTrap() bool`

HasEnableTrap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


