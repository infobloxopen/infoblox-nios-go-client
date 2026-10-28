# RestartServiceActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** | Service Restart Group names to restart. Conflicts with &#x60;members&#x60;. If both are omitted, NIOS uses the default restart group. | [optional] 
**Members** | Pointer to **[]string** | Grid Member names to restart. Conflicts with &#x60;groups&#x60;. If both are omitted, NIOS uses the default restart group. | [optional] 
**Mode** | Pointer to **string** | Grid restart method. Valid values are &#x60;GROUPED&#x60;, &#x60;SEQUENTIAL&#x60;, and &#x60;SIMULTANEOUS&#x60;. | [optional] 
**RestartOption** | Pointer to **string** | Controls whether NIOS restarts services only when needed or unconditionally. Valid values are &#x60;RESTART_IF_NEEDED&#x60; and &#x60;FORCE_RESTART&#x60;. Defaults to &#x60;RESTART_IF_NEEDED&#x60;. | [optional] 
**Services** | Pointer to **[]string** | Services to restart. Valid values are &#x60;ALL&#x60;, &#x60;DNS&#x60;, &#x60;DHCP&#x60;, &#x60;DHCPV4&#x60;, and &#x60;DHCPV6&#x60;. Defaults to &#x60;ALL&#x60;. | [optional] 
**UserName** | Pointer to **string** | Name of the user requesting the restart, recorded by NIOS. | [optional] 

## Methods

### NewRestartServiceActionRequest

`func NewRestartServiceActionRequest() *RestartServiceActionRequest`

NewRestartServiceActionRequest instantiates a new RestartServiceActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartServiceActionRequestWithDefaults

`func NewRestartServiceActionRequestWithDefaults() *RestartServiceActionRequest`

NewRestartServiceActionRequestWithDefaults instantiates a new RestartServiceActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *RestartServiceActionRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RestartServiceActionRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RestartServiceActionRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RestartServiceActionRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMembers

`func (o *RestartServiceActionRequest) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RestartServiceActionRequest) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *RestartServiceActionRequest) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *RestartServiceActionRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMode

`func (o *RestartServiceActionRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RestartServiceActionRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RestartServiceActionRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RestartServiceActionRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRestartOption

`func (o *RestartServiceActionRequest) GetRestartOption() string`

GetRestartOption returns the RestartOption field if non-nil, zero value otherwise.

### GetRestartOptionOk

`func (o *RestartServiceActionRequest) GetRestartOptionOk() (*string, bool)`

GetRestartOptionOk returns a tuple with the RestartOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOption

`func (o *RestartServiceActionRequest) SetRestartOption(v string)`

SetRestartOption sets RestartOption field to given value.

### HasRestartOption

`func (o *RestartServiceActionRequest) HasRestartOption() bool`

HasRestartOption returns a boolean if a field has been set.

### GetServices

`func (o *RestartServiceActionRequest) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *RestartServiceActionRequest) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *RestartServiceActionRequest) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *RestartServiceActionRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetUserName

`func (o *RestartServiceActionRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RestartServiceActionRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RestartServiceActionRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RestartServiceActionRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


