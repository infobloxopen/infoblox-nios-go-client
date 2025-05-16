# GridCloudapiAllowedApiAdmins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRemote** | Pointer to **bool** | Determines whether this is a remote admin user. | [optional] 
**RemoteAdmin** | Pointer to **string** | Username that matches a remote administrator who can perform cloud API requests on the Cloud Platform Appliance. | [optional] 
**LocalAdmin** | Pointer to **string** | Local administrator who can perform cloud API requests on the Cloud Platform Appliance. | [optional] 

## Methods

### NewGridCloudapiAllowedApiAdmins

`func NewGridCloudapiAllowedApiAdmins() *GridCloudapiAllowedApiAdmins`

NewGridCloudapiAllowedApiAdmins instantiates a new GridCloudapiAllowedApiAdmins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCloudapiAllowedApiAdminsWithDefaults

`func NewGridCloudapiAllowedApiAdminsWithDefaults() *GridCloudapiAllowedApiAdmins`

NewGridCloudapiAllowedApiAdminsWithDefaults instantiates a new GridCloudapiAllowedApiAdmins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRemote

`func (o *GridCloudapiAllowedApiAdmins) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *GridCloudapiAllowedApiAdmins) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *GridCloudapiAllowedApiAdmins) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *GridCloudapiAllowedApiAdmins) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetRemoteAdmin

`func (o *GridCloudapiAllowedApiAdmins) GetRemoteAdmin() string`

GetRemoteAdmin returns the RemoteAdmin field if non-nil, zero value otherwise.

### GetRemoteAdminOk

`func (o *GridCloudapiAllowedApiAdmins) GetRemoteAdminOk() (*string, bool)`

GetRemoteAdminOk returns a tuple with the RemoteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAdmin

`func (o *GridCloudapiAllowedApiAdmins) SetRemoteAdmin(v string)`

SetRemoteAdmin sets RemoteAdmin field to given value.

### HasRemoteAdmin

`func (o *GridCloudapiAllowedApiAdmins) HasRemoteAdmin() bool`

HasRemoteAdmin returns a boolean if a field has been set.

### GetLocalAdmin

`func (o *GridCloudapiAllowedApiAdmins) GetLocalAdmin() string`

GetLocalAdmin returns the LocalAdmin field if non-nil, zero value otherwise.

### GetLocalAdminOk

`func (o *GridCloudapiAllowedApiAdmins) GetLocalAdminOk() (*string, bool)`

GetLocalAdminOk returns a tuple with the LocalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdmin

`func (o *GridCloudapiAllowedApiAdmins) SetLocalAdmin(v string)`

SetLocalAdmin sets LocalAdmin field to given value.

### HasLocalAdmin

`func (o *GridCloudapiAllowedApiAdmins) HasLocalAdmin() bool`

HasLocalAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


