# NsgroupforwardingmemberforwardingserversForwardTo

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

### NewNsgroupforwardingmemberforwardingserversForwardTo

`func NewNsgroupforwardingmemberforwardingserversForwardTo() *NsgroupforwardingmemberforwardingserversForwardTo`

NewNsgroupforwardingmemberforwardingserversForwardTo instantiates a new NsgroupforwardingmemberforwardingserversForwardTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupforwardingmemberforwardingserversForwardToWithDefaults

`func NewNsgroupforwardingmemberforwardingserversForwardToWithDefaults() *NsgroupforwardingmemberforwardingserversForwardTo`

NewNsgroupforwardingmemberforwardingserversForwardToWithDefaults instantiates a new NsgroupforwardingmemberforwardingserversForwardTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedWithMsParentDelegation

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetSharedWithMsParentDelegation() bool`

GetSharedWithMsParentDelegation returns the SharedWithMsParentDelegation field if non-nil, zero value otherwise.

### GetSharedWithMsParentDelegationOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetSharedWithMsParentDelegationOk() (*bool, bool)`

GetSharedWithMsParentDelegationOk returns a tuple with the SharedWithMsParentDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithMsParentDelegation

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetSharedWithMsParentDelegation(v bool)`

SetSharedWithMsParentDelegation sets SharedWithMsParentDelegation field to given value.

### HasSharedWithMsParentDelegation

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasSharedWithMsParentDelegation() bool`

HasSharedWithMsParentDelegation returns a boolean if a field has been set.

### GetStealth

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetTsigKey

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetTsigKey() string`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetTsigKeyOk() (*string, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetTsigKey(v string)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetTsigKeyAlg

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetTsigKeyAlg() string`

GetTsigKeyAlg returns the TsigKeyAlg field if non-nil, zero value otherwise.

### GetTsigKeyAlgOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetTsigKeyAlgOk() (*string, bool)`

GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyAlg

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetTsigKeyAlg(v string)`

SetTsigKeyAlg sets TsigKeyAlg field to given value.

### HasTsigKeyAlg

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasTsigKeyAlg() bool`

HasTsigKeyAlg returns a boolean if a field has been set.

### GetTsigKeyName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetTsigKeyName() string`

GetTsigKeyName returns the TsigKeyName field if non-nil, zero value otherwise.

### GetTsigKeyNameOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetTsigKeyNameOk() (*string, bool)`

GetTsigKeyNameOk returns a tuple with the TsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKeyName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetTsigKeyName(v string)`

SetTsigKeyName sets TsigKeyName field to given value.

### HasTsigKeyName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasTsigKeyName() bool`

HasTsigKeyName returns a boolean if a field has been set.

### GetUseTsigKeyName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetUseTsigKeyName() bool`

GetUseTsigKeyName returns the UseTsigKeyName field if non-nil, zero value otherwise.

### GetUseTsigKeyNameOk

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) GetUseTsigKeyNameOk() (*bool, bool)`

GetUseTsigKeyNameOk returns a tuple with the UseTsigKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTsigKeyName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) SetUseTsigKeyName(v bool)`

SetUseTsigKeyName sets UseTsigKeyName field to given value.

### HasUseTsigKeyName

`func (o *NsgroupforwardingmemberforwardingserversForwardTo) HasUseTsigKeyName() bool`

HasUseTsigKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


