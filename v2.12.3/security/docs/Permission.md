# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Group** | Pointer to **string** | The name of the admin group this permission applies to. | [optional] 
**Object** | Pointer to **string** | A reference to a WAPI object, which will be the object this permission applies to. | [optional] 
**Permission** | Pointer to **string** | The type of permission. | [optional] 
**ResourceType** | Pointer to **string** | The type of resource this permission applies to. If &#39;object&#39; is set, the permission is going to apply to child objects of the specified type, for example if &#39;object&#39; was set to an authoritative zone reference and &#39;resource_type&#39; was set to &#39;A&#39;, the permission would apply to A Resource Records within the specified zone. | [optional] 
**Role** | Pointer to **string** | The name of the role this permission applies to. | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Permission) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Permission) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Permission) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Permission) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetGroup

`func (o *Permission) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Permission) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Permission) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Permission) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetObject

`func (o *Permission) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Permission) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Permission) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Permission) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPermission

`func (o *Permission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Permission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Permission) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Permission) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetResourceType

`func (o *Permission) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Permission) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Permission) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Permission) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRole

`func (o *Permission) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Permission) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Permission) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Permission) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


