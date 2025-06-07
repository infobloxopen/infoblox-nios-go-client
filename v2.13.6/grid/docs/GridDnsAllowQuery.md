# GridDnsAllowQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 
**TsigKey** | Pointer to **string** | A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT. | [optional] 
**TsigKeyAlg** | Pointer to **string** | The TSIG key algorithm. | [optional] 
**TsigKeyName** | Pointer to **string** | The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to &#39;tsig_xfer&#39; on retrieval, and ignored on insert or update. | [optional] 
**UseTsigKeyName** | Pointer to **bool** | Use flag for: tsig_key_name | [optional] 

## Methods

### NewGridDnsAllowQuery

`func NewGridDnsAllowQuery() *GridDnsAllowQuery`

NewGridDnsAllowQuery instantiates a new GridDnsAllowQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsAllowQueryWithDefaults

`func NewGridDnsAllowQueryWithDefaults() *GridDnsAllowQuery`

NewGridDnsAllowQueryWithDefaults instantiates a new GridDnsAllowQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GridDnsAllowQuery) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GridDnsAllowQuery) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GridDnsAllowQuery) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GridDnsAllowQuery) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *GridDnsAllowQuery) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *GridDnsAllowQuery) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *GridDnsAllowQuery) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *GridDnsAllowQuery) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTsigKey

`func (o *GridDnsAllowQuery) GetTsigKey() string`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *GridDnsAllowQuery) GetTsigKeyOk() (*string, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *GridDnsAllowQuery) SetTsigKey(v string)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *GridDnsAllowQuery) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetTsigKeyAlg

`func (o *GridDnsAllowQuery) GetTsigKeyAlg() string`

GetTsigKeyAlg returns the TsigKeyAlg field if non-nil, zero value otherwise.

### GetTsigKeyAlgOk

`func (o *GridDnsAllowQuery) GetTsigKeyAlgOk() (*string, bool)`

GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyAlg

`func (o *GridDnsAllowQuery) SetTsigKeyAlg(v string)`

SetTsigKeyAlg sets TsigKeyAlg field to given value.

### HasTsigKeyAlg

`func (o *GridDnsAllowQuery) HasTsigKeyAlg() bool`

HasTsigKeyAlg returns a boolean if a field has been set.

### GetTsigKeyName

`func (o *GridDnsAllowQuery) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *GridDnsAllowQuery) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *GridDnsAllowQuery) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.

### HasTsigKeyName

`func (o *GridDnsAllowQuery) HasTsigKeyName() bool`

HasTsigKeyName returns a boolean if a field has been set.

### GetUseTsigKeyName

`func (o *GridDnsAllowQuery) GetUseTsigKeyName() bool`

GetUseTsigKeyName returns the UseTsigKeyName field if non-nil, zero value otherwise.

### GetUseTsigKeyNameOk

`func (o *GridDnsAllowQuery) GetUseTsigKeyNameOk() (*bool, bool)`

GetUseTsigKeyNameOk returns a tuple with the UseTsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTsigKeyName

`func (o *GridDnsAllowQuery) SetUseTsigKeyName(v bool)`

SetUseTsigKeyName sets UseTsigKeyName field to given value.

### HasUseTsigKeyName

`func (o *GridDnsAllowQuery) HasUseTsigKeyName() bool`

HasUseTsigKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


