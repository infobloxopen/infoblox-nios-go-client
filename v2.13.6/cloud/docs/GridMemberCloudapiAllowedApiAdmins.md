# GridMemberCloudapiAllowedApiAdmins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRemote** | Pointer to **bool** | Determines whether this is a remote admin user. | [optional] 
**RemoteAdmin** | Pointer to **string** | Username that matches a remote administrator who can perform cloud API requests on the Cloud Platform Appliance. | [optional] 
**LocalAdmin** | Pointer to **string** | Local administrator who can perform cloud API requests on the Cloud Platform Appliance. | [optional] 

## Methods

### NewGridMemberCloudapiAllowedApiAdmins

`func NewGridMemberCloudapiAllowedApiAdmins() *GridMemberCloudapiAllowedApiAdmins`

NewGridMemberCloudapiAllowedApiAdmins instantiates a new GridMemberCloudapiAllowedApiAdmins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridMemberCloudapiAllowedApiAdminsWithDefaults

`func NewGridMemberCloudapiAllowedApiAdminsWithDefaults() *GridMemberCloudapiAllowedApiAdmins`

NewGridMemberCloudapiAllowedApiAdminsWithDefaults instantiates a new GridMemberCloudapiAllowedApiAdmins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRemote

`func (o *GridMemberCloudapiAllowedApiAdmins) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *GridMemberCloudapiAllowedApiAdmins) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *GridMemberCloudapiAllowedApiAdmins) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *GridMemberCloudapiAllowedApiAdmins) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetRemoteAdmin

`func (o *GridMemberCloudapiAllowedApiAdmins) GetRemoteAdmin() string`

GetRemoteAdmin returns the RemoteAdmin field if non-nil, zero value otherwise.

### GetRemoteAdminOk

`func (o *GridMemberCloudapiAllowedApiAdmins) GetRemoteAdminOk() (*string, bool)`

GetRemoteAdminOk returns a tuple with the RemoteAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAdmin

`func (o *GridMemberCloudapiAllowedApiAdmins) SetRemoteAdmin(v string)`

SetRemoteAdmin sets RemoteAdmin field to given value.

### HasRemoteAdmin

`func (o *GridMemberCloudapiAllowedApiAdmins) HasRemoteAdmin() bool`

HasRemoteAdmin returns a boolean if a field has been set.

### GetLocalAdmin

`func (o *GridMemberCloudapiAllowedApiAdmins) GetLocalAdmin() string`

GetLocalAdmin returns the LocalAdmin field if non-nil, zero value otherwise.

### GetLocalAdminOk

`func (o *GridMemberCloudapiAllowedApiAdmins) GetLocalAdminOk() (*string, bool)`

GetLocalAdminOk returns a tuple with the LocalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAdmin

`func (o *GridMemberCloudapiAllowedApiAdmins) SetLocalAdmin(v string)`

SetLocalAdmin sets LocalAdmin field to given value.

### HasLocalAdmin

`func (o *GridMemberCloudapiAllowedApiAdmins) HasLocalAdmin() bool`

HasLocalAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


