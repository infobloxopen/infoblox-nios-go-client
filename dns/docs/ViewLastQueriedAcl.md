# ViewLastQueriedAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 

## Methods

### NewViewLastQueriedAcl

`func NewViewLastQueriedAcl() *ViewLastQueriedAcl`

NewViewLastQueriedAcl instantiates a new ViewLastQueriedAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewLastQueriedAclWithDefaults

`func NewViewLastQueriedAclWithDefaults() *ViewLastQueriedAcl`

NewViewLastQueriedAclWithDefaults instantiates a new ViewLastQueriedAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ViewLastQueriedAcl) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ViewLastQueriedAcl) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ViewLastQueriedAcl) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ViewLastQueriedAcl) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *ViewLastQueriedAcl) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ViewLastQueriedAcl) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ViewLastQueriedAcl) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ViewLastQueriedAcl) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


