# ViewMatchDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Struct** | Pointer to **string** | The struct type of the object. The value must be one of &#39;addressac&#39; and &#39;tsigac&#39;. | [optional] 
**Ref** | Pointer to **string** | The reference to the Named ACL object. | [optional] 
**Address** | Pointer to **string** | The address this rule applies to or \&quot;Any\&quot;. | [optional] 
**Permission** | Pointer to **string** | The permission to use for this address. | [optional] 
**TsigKey** | Pointer to **string** | A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT. | [optional] 
**TsigKeyAlg** | Pointer to **string** | The TSIG key algorithm. | [optional] 
**TsigKeyName** | Pointer to **string** | The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to &#39;tsig_xfer&#39; on retrieval, and ignored on insert or update. | [optional] 
**UseTsigKeyName** | Pointer to **bool** | Use flag for: tsig_key_name | [optional] 

## Methods

### NewViewMatchDestinations

`func NewViewMatchDestinations() *ViewMatchDestinations`

NewViewMatchDestinations instantiates a new ViewMatchDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewMatchDestinationsWithDefaults

`func NewViewMatchDestinationsWithDefaults() *ViewMatchDestinations`

NewViewMatchDestinationsWithDefaults instantiates a new ViewMatchDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStruct

`func (o *ViewMatchDestinations) GetStruct() string`

GetStruct returns the Struct field if non-nil, zero value otherwise.

### GetStructOk

`func (o *ViewMatchDestinations) GetStructOk() (*string, bool)`

GetStructOk returns a tuple with the Struct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStruct

`func (o *ViewMatchDestinations) SetStruct(v string)`

SetStruct sets Struct field to given value.

### HasStruct

`func (o *ViewMatchDestinations) HasStruct() bool`

HasStruct returns a boolean if a field has been set.

### GetRef

`func (o *ViewMatchDestinations) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ViewMatchDestinations) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ViewMatchDestinations) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ViewMatchDestinations) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *ViewMatchDestinations) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ViewMatchDestinations) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ViewMatchDestinations) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ViewMatchDestinations) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPermission

`func (o *ViewMatchDestinations) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ViewMatchDestinations) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ViewMatchDestinations) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ViewMatchDestinations) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetTsigKey

`func (o *ViewMatchDestinations) GetTsigKey() string`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *ViewMatchDestinations) GetTsigKeyOk() (*string, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *ViewMatchDestinations) SetTsigKey(v string)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *ViewMatchDestinations) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetTsigKeyAlg

`func (o *ViewMatchDestinations) GetTsigKeyAlg() string`

GetTsigKeyAlg returns the TsigKeyAlg field if non-nil, zero value otherwise.

### GetTsigKeyAlgOk

`func (o *ViewMatchDestinations) GetTsigKeyAlgOk() (*string, bool)`

GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyAlg

`func (o *ViewMatchDestinations) SetTsigKeyAlg(v string)`

SetTsigKeyAlg sets TsigKeyAlg field to given value.

### HasTsigKeyAlg

`func (o *ViewMatchDestinations) HasTsigKeyAlg() bool`

HasTsigKeyAlg returns a boolean if a field has been set.

### GetTsigKeyName

`func (o *ViewMatchDestinations) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *ViewMatchDestinations) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *ViewMatchDestinations) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.

### HasTsigKeyName

`func (o *ViewMatchDestinations) HasTsigKeyName() bool`

HasTsigKeyName returns a boolean if a field has been set.

### GetUseTsigKeyName

`func (o *ViewMatchDestinations) GetUseTsigKeyName() bool`

GetUseTsigKeyName returns the UseTsigKeyName field if non-nil, zero value otherwise.

### GetUseTsigKeyNameOk

`func (o *ViewMatchDestinations) GetUseTsigKeyNameOk() (*bool, bool)`

GetUseTsigKeyNameOk returns a tuple with the UseTsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTsigKeyName

`func (o *ViewMatchDestinations) SetUseTsigKeyName(v bool)`

SetUseTsigKeyName sets UseTsigKeyName field to given value.

### HasUseTsigKeyName

`func (o *ViewMatchDestinations) HasUseTsigKeyName() bool`

HasUseTsigKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


