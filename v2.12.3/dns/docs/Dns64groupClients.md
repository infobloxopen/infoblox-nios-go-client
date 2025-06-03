# Dns64groupClients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 

## Methods

### NewDns64groupClients

`func NewDns64groupClients() *Dns64groupClients`

NewDns64groupClients instantiates a new Dns64groupClients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDns64groupClientsWithDefaults

`func NewDns64groupClientsWithDefaults() *Dns64groupClients`

NewDns64groupClientsWithDefaults instantiates a new Dns64groupClients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Dns64groupClients) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Dns64groupClients) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Dns64groupClients) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Dns64groupClients) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *Dns64groupClients) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Dns64groupClients) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Dns64groupClients) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Dns64groupClients) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


