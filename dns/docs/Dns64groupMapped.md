# Dns64groupMapped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 

## Methods

### NewDns64groupMapped

`func NewDns64groupMapped() *Dns64groupMapped`

NewDns64groupMapped instantiates a new Dns64groupMapped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDns64groupMappedWithDefaults

`func NewDns64groupMappedWithDefaults() *Dns64groupMapped`

NewDns64groupMappedWithDefaults instantiates a new Dns64groupMapped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Dns64groupMapped) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Dns64groupMapped) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Dns64groupMapped) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Dns64groupMapped) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *Dns64groupMapped) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Dns64groupMapped) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Dns64groupMapped) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Dns64groupMapped) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


