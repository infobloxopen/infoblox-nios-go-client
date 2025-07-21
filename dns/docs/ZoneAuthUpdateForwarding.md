# ZoneAuthUpdateForwarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Struct** | Pointer to **string** | The struct type of the object. The value must be one of &#39;addressac&#39; and &#39;tsigac&#39;. | [optional] 
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 
**TsigKey** | Pointer to **string** | A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT. | [optional] 
**TsigKeyAlg** | Pointer to **string** | The TSIG key algorithm. | [optional] 
**TsigKeyName** | Pointer to **string** | The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to &#39;tsig_xfer&#39; on retrieval, and ignored on insert or update. | [optional] 
**UseTsigKeyName** | Pointer to **bool** | Use flag for: tsig_key_name | [optional] 

## Methods

### NewZoneAuthUpdateForwarding

`func NewZoneAuthUpdateForwarding() *ZoneAuthUpdateForwarding`

NewZoneAuthUpdateForwarding instantiates a new ZoneAuthUpdateForwarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthUpdateForwardingWithDefaults

`func NewZoneAuthUpdateForwardingWithDefaults() *ZoneAuthUpdateForwarding`

NewZoneAuthUpdateForwardingWithDefaults instantiates a new ZoneAuthUpdateForwarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStruct

`func (o *ZoneAuthUpdateForwarding) GetStruct() string`

GetStruct returns the Struct field if non-nil, zero value otherwise.

### GetStructOk

`func (o *ZoneAuthUpdateForwarding) GetStructOk() (*string, bool)`

GetStructOk returns a tuple with the Struct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStruct

`func (o *ZoneAuthUpdateForwarding) SetStruct(v string)`

SetStruct sets Struct field to given value.

### HasStruct

`func (o *ZoneAuthUpdateForwarding) HasStruct() bool`

HasStruct returns a boolean if a field has been set.

### GetAddress

`func (o *ZoneAuthUpdateForwarding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ZoneAuthUpdateForwarding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ZoneAuthUpdateForwarding) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ZoneAuthUpdateForwarding) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *ZoneAuthUpdateForwarding) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ZoneAuthUpdateForwarding) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ZoneAuthUpdateForwarding) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ZoneAuthUpdateForwarding) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTsigKey

`func (o *ZoneAuthUpdateForwarding) GetTsigKey() string`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *ZoneAuthUpdateForwarding) GetTsigKeyOk() (*string, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *ZoneAuthUpdateForwarding) SetTsigKey(v string)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *ZoneAuthUpdateForwarding) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetTsigKeyAlg

`func (o *ZoneAuthUpdateForwarding) GetTsigKeyAlg() string`

GetTsigKeyAlg returns the TsigKeyAlg field if non-nil, zero value otherwise.

### GetTsigKeyAlgOk

`func (o *ZoneAuthUpdateForwarding) GetTsigKeyAlgOk() (*string, bool)`

GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyAlg

`func (o *ZoneAuthUpdateForwarding) SetTsigKeyAlg(v string)`

SetTsigKeyAlg sets TsigKeyAlg field to given value.

### HasTsigKeyAlg

`func (o *ZoneAuthUpdateForwarding) HasTsigKeyAlg() bool`

HasTsigKeyAlg returns a boolean if a field has been set.

### GetTsigKeyName

`func (o *ZoneAuthUpdateForwarding) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *ZoneAuthUpdateForwarding) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *ZoneAuthUpdateForwarding) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.

### HasTsigKeyName

`func (o *ZoneAuthUpdateForwarding) HasTsigKeyName() bool`

HasTsigKeyName returns a boolean if a field has been set.

### GetUseTsigKeyName

`func (o *ZoneAuthUpdateForwarding) GetUseTsigKeyName() bool`

GetUseTsigKeyName returns the UseTsigKeyName field if non-nil, zero value otherwise.

### GetUseTsigKeyNameOk

`func (o *ZoneAuthUpdateForwarding) GetUseTsigKeyNameOk() (*bool, bool)`

GetUseTsigKeyNameOk returns a tuple with the UseTsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTsigKeyName

`func (o *ZoneAuthUpdateForwarding) SetUseTsigKeyName(v bool)`

SetUseTsigKeyName sets UseTsigKeyName field to given value.

### HasUseTsigKeyName

`func (o *ZoneAuthUpdateForwarding) HasUseTsigKeyName() bool`

HasUseTsigKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


