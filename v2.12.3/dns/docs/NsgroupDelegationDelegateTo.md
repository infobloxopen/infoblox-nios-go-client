# NsgroupDelegationDelegateTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the server. | [optional] 
**Name** | Pointer to **string** | A resolvable domain name for the external DNS server. | [optional] 
**SharedWithMsParentDelegation** | Pointer to **bool** | This flag represents whether the name server is shared with the parent Microsoft primary zone&#39;s delegation server. | [optional] [readonly] 
**Stealth** | Pointer to **bool** | Set this flag to hide the NS record for the primary name server from DNS queries. | [optional] 
**TsigKey** | Pointer to **string** | A generated TSIG key. | [optional] 
**TsigKeyAlg** | Pointer to **string** | The TSIG key algorithm. | [optional] 
**TsigKeyName** | Pointer to **string** | The TSIG key name. | [optional] 
**UseTsigKeyName** | Pointer to **bool** | Use flag for: tsig_key_name | [optional] 

## Methods

### NewNsgroupDelegationDelegateTo

`func NewNsgroupDelegationDelegateTo() *NsgroupDelegationDelegateTo`

NewNsgroupDelegationDelegateTo instantiates a new NsgroupDelegationDelegateTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupDelegationDelegateToWithDefaults

`func NewNsgroupDelegationDelegateToWithDefaults() *NsgroupDelegationDelegateTo`

NewNsgroupDelegationDelegateToWithDefaults instantiates a new NsgroupDelegationDelegateTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NsgroupDelegationDelegateTo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NsgroupDelegationDelegateTo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NsgroupDelegationDelegateTo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NsgroupDelegationDelegateTo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *NsgroupDelegationDelegateTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NsgroupDelegationDelegateTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NsgroupDelegationDelegateTo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NsgroupDelegationDelegateTo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedWithMsParentDelegation

`func (o *NsgroupDelegationDelegateTo) GetSharedWithMsParentDelegation() bool`

GetSharedWithMsParentDelegation returns the SharedWithMsParentDelegation field if non-nil, zero value otherwise.

### GetSharedWithMsParentDelegationOk

`func (o *NsgroupDelegationDelegateTo) GetSharedWithMsParentDelegationOk() (*bool, bool)`

GetSharedWithMsParentDelegationOk returns a tuple with the SharedWithMsParentDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithMsParentDelegation

`func (o *NsgroupDelegationDelegateTo) SetSharedWithMsParentDelegation(v bool)`

SetSharedWithMsParentDelegation sets SharedWithMsParentDelegation field to given value.

### HasSharedWithMsParentDelegation

`func (o *NsgroupDelegationDelegateTo) HasSharedWithMsParentDelegation() bool`

HasSharedWithMsParentDelegation returns a boolean if a field has been set.

### GetStealth

`func (o *NsgroupDelegationDelegateTo) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *NsgroupDelegationDelegateTo) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *NsgroupDelegationDelegateTo) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *NsgroupDelegationDelegateTo) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetTsigKey

`func (o *NsgroupDelegationDelegateTo) GetTsigKey() string`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *NsgroupDelegationDelegateTo) GetTsigKeyOk() (*string, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *NsgroupDelegationDelegateTo) SetTsigKey(v string)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *NsgroupDelegationDelegateTo) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetTsigKeyAlg

`func (o *NsgroupDelegationDelegateTo) GetTsigKeyAlg() string`

GetTsigKeyAlg returns the TsigKeyAlg field if non-nil, zero value otherwise.

### GetTsigKeyAlgOk

`func (o *NsgroupDelegationDelegateTo) GetTsigKeyAlgOk() (*string, bool)`

GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyAlg

`func (o *NsgroupDelegationDelegateTo) SetTsigKeyAlg(v string)`

SetTsigKeyAlg sets TsigKeyAlg field to given value.

### HasTsigKeyAlg

`func (o *NsgroupDelegationDelegateTo) HasTsigKeyAlg() bool`

HasTsigKeyAlg returns a boolean if a field has been set.

### GetTsigKeyName

`func (o *NsgroupDelegationDelegateTo) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *NsgroupDelegationDelegateTo) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *NsgroupDelegationDelegateTo) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.

### HasTsigKeyName

`func (o *NsgroupDelegationDelegateTo) HasTsigKeyName() bool`

HasTsigKeyName returns a boolean if a field has been set.

### GetUseTsigKeyName

`func (o *NsgroupDelegationDelegateTo) GetUseTsigKeyName() bool`

GetUseTsigKeyName returns the UseTsigKeyName field if non-nil, zero value otherwise.

### GetUseTsigKeyNameOk

`func (o *NsgroupDelegationDelegateTo) GetUseTsigKeyNameOk() (*bool, bool)`

GetUseTsigKeyNameOk returns a tuple with the UseTsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTsigKeyName

`func (o *NsgroupDelegationDelegateTo) SetUseTsigKeyName(v bool)`

SetUseTsigKeyName sets UseTsigKeyName field to given value.

### HasUseTsigKeyName

`func (o *NsgroupDelegationDelegateTo) HasUseTsigKeyName() bool`

HasUseTsigKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


