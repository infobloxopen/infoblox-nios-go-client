# GridThreatprotectionNatRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | Pointer to **string** | The rule type for the threat protection NAT mapping rule. | [optional] 
**Address** | Pointer to **string** | The IP address for the threat protection NAT mapping rule. | [optional] 
**Network** | Pointer to **string** | The network address for the threat protection NAT mapping rule. | [optional] 
**Cidr** | Pointer to **int64** | The network CIDR for the threat protection NAT mapping rule. | [optional] 
**StartAddress** | Pointer to **string** | The start address for the range of the threat protection NAT mapping rule. | [optional] 
**EndAddress** | Pointer to **string** | The end address for the range of the threat protection NAT mapping rule. | [optional] 
**NatPorts** | Pointer to [**GridthreatprotectionnatrulesNatPorts**](GridthreatprotectionnatrulesNatPorts.md) |  | [optional] 

## Methods

### NewGridThreatprotectionNatRules

`func NewGridThreatprotectionNatRules() *GridThreatprotectionNatRules`

NewGridThreatprotectionNatRules instantiates a new GridThreatprotectionNatRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridThreatprotectionNatRulesWithDefaults

`func NewGridThreatprotectionNatRulesWithDefaults() *GridThreatprotectionNatRules`

NewGridThreatprotectionNatRulesWithDefaults instantiates a new GridThreatprotectionNatRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *GridThreatprotectionNatRules) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GridThreatprotectionNatRules) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GridThreatprotectionNatRules) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GridThreatprotectionNatRules) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetAddress

`func (o *GridThreatprotectionNatRules) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GridThreatprotectionNatRules) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GridThreatprotectionNatRules) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GridThreatprotectionNatRules) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNetwork

`func (o *GridThreatprotectionNatRules) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GridThreatprotectionNatRules) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GridThreatprotectionNatRules) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GridThreatprotectionNatRules) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCidr

`func (o *GridThreatprotectionNatRules) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GridThreatprotectionNatRules) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GridThreatprotectionNatRules) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GridThreatprotectionNatRules) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetStartAddress

`func (o *GridThreatprotectionNatRules) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *GridThreatprotectionNatRules) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *GridThreatprotectionNatRules) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *GridThreatprotectionNatRules) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *GridThreatprotectionNatRules) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *GridThreatprotectionNatRules) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *GridThreatprotectionNatRules) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *GridThreatprotectionNatRules) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetNatPorts

`func (o *GridThreatprotectionNatRules) GetNatPorts() GridthreatprotectionnatrulesNatPorts`

GetNatPorts returns the NatPorts field if non-nil, zero value otherwise.

### GetNatPortsOk

`func (o *GridThreatprotectionNatRules) GetNatPortsOk() (*GridthreatprotectionnatrulesNatPorts, bool)`

GetNatPortsOk returns a tuple with the NatPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPorts

`func (o *GridThreatprotectionNatRules) SetNatPorts(v GridthreatprotectionnatrulesNatPorts)`

SetNatPorts sets NatPorts field to given value.

### HasNatPorts

`func (o *GridThreatprotectionNatRules) HasNatPorts() bool`

HasNatPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


